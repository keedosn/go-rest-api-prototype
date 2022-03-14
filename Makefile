.PHONY: test

run:
	go run cmd/server/main.go

race:
	go run --race cmd/server/main.go

build:
	go build cmd/server/main.go

test:
	go test -v -run=. test/**/*.go

bench: bench-clean
	go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile test/data/cpu.out -memprofile test/data/mem.out -trace test/data/trace.out ./test/handler/create_category_test.go | tee test/data/bench.log

bench-cpu:
	go tool pprof -http :19800 test/data/cpu.out

bench-mem:
	go tool pprof -http :19801 test/data/mem.out

bench-trace:
	go tool trace test/data/trace.out

bench-clean:
	rm -f test/data/*.out test/data/bench.log

# run-profiler:
# 	go run cmd/server/main.go -cpuprofile cpu.prof -memprofile mem.prof
#
# profile-cpu-web:
# 	go tool pprof -http=127.0.0.1:9990 cpu.prof
#
# profile-mem-web:
# 	go tool pprof -http=127.0.0.1:9991 mem.prof
