run:
	go run cmd/server/main.go

race:
	go run --race cmd/server/main.go

build:
	go build cmd/server/main.go

test:
	go test cmd/server/main.go

# run-profiler:
# 	go run cmd/server/main.go -cpuprofile cpu.prof -memprofile mem.prof
#
# profile-cpu-web:
# 	go tool pprof -http=127.0.0.1:9990 cpu.prof
#
# profile-mem-web:
# 	go tool pprof -http=127.0.0.1:9991 mem.prof
