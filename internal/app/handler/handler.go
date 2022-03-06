package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Env struct {
	Addr string
	DB   *pgxpool.Pool
}

type Handler struct {
	*Env
	Handle   HandlerFunc
	Request  interface{}
	Response interface{}
	Params   Set
}

type HandlerFunc func(h *Handler, w http.ResponseWriter) (interface{}, int, error)

type Set map[string]string

type response struct {
	Status int
	Data   interface{}
}

func New(e *Env, h *Handler) *Handler {
	return &Handler{e, h.Handle, h.Request, h.Response, Set{}}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := decodeRequestData(r, h.Request); err != nil {
		log.Println("err_ServeHTTP:", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
	}

	h.Params = mux.Vars(r)
	res, code, err := h.Handle(h, w)

	encodeResponse(w, &response{code, res}, err)
}

func decodeRequestData(r *http.Request, v interface{}) error {
	buf, _ := ioutil.ReadAll(r.Body)
	rdr := ioutil.NopCloser(bytes.NewReader(buf))
	r.Body = ioutil.NopCloser(bytes.NewReader(buf))

	json.NewDecoder(rdr).Decode(&v)

	return nil
}

func encodeResponse(w http.ResponseWriter, res *response, err error) {
	if err != nil {
		encodeError(w, res.Status, err)
		return
	}

	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res.Data)
}

func encodeError(w http.ResponseWriter, status int, e error) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Error(e.Error()))
}
