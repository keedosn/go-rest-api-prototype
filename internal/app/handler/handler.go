package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Env struct {
	Addr string
	DB   *pgxpool.Pool
}

type HandlerFunc func(e *Env, req interface{}, w http.ResponseWriter) (interface{}, error)

type Handler struct {
	*Env
	Handle   HandlerFunc
	Request  interface{}
	Response interface{}
}

func New(env *Env, h *Handler) *Handler {
	return &Handler{env, h.Handle, h.Request, h.Response}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := decodeRequestData(r, &h.Request); err != nil {
		log.Println(err.Error())
	}

	res, err := h.Handle(h.Env, h.Request, w)

	encodeResponse(w, res, err)
}

func decodeRequestData(r *http.Request, v interface{}) error {
	buf, _ := ioutil.ReadAll(r.Body)
	rdr := ioutil.NopCloser(bytes.NewReader(buf))
	r.Body = ioutil.NopCloser(bytes.NewReader(buf))

	if len(buf) == 0 {
		// log.Println("empty request body1") //FIXME #1

		return nil
	}

	if err := json.NewDecoder(rdr).Decode(&v); err != nil {
		return err
	}

	return nil
}

func encodeResponse(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		encodeError(w, err)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func encodeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": e.Error(),
	})
}
