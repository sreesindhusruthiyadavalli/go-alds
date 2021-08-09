package main

import (
	//"github.com/sreesindhusruthiyadavalli/go-alds/src/ll"
	"github.com/sreesindhusruthiyadavalli/go-alds/src/stack"
)

func main(){
	//ll.CheckLl()
	stack.CheckStack()
}

/*
-- go to go-alds/src/app
-- go build
 */

/*
RUN TESTS

Gist source: https://gist.github.com/arsham/bbc93990d8e5c9b54128a3d88901ab90

go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out ./package | tee bench.txt
go tool pprof -http :8080 cpu.out
go tool pprof -http :8081 mem.out
go tool trace trace.out

go tool pprof $FILENAME.test cpu.out
# (pprof) list <func name>

# go get -u golang.org/x/perf/cmd/benchstat
benchstat bench.txt
rm cpu.out mem.out trace.out *.test

go tool pprof -sample_index=inuse_space stack.test mem1.out
>pprof : pdf, svg
 */

/*
curl http://localhost:8080/debug/pprof/heap > heap.0.pprof
sleep 30
curl http://localhost:8080/debug/pprof/heap > heap.1.pprof
sleep 30
curl http://localhost:8080/debug/pprof/heap > heap.2.pprof
sleep 30
curl http://localhost:8080/debug/pprof/heap > heap.3.pprof
 */
