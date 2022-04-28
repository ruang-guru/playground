cpuprof-basic-row:
	go test -run=BenchmarkSumRow \
	-bench=BenchmarkSumRow \
	-benchtime=1x -benchmem \
	-cpuprofile profile.out

memprof-basic-row:
	go test -run=BenchmarkSumRow \
	-bench=BenchmarkSumRow \
	-benchtime=1x -benchmem \
	-memprofile memprofile.out

cpuprof-basic-column:
	go test -run=BenchmarkSumColumn \
	-bench=BenchmarkSumColumn \
	-benchtime=1x -benchmem \
	-cpuprofile profile.out


memprof-basic-column:
	go test -run=BenchmarkSumColumn \
	-bench=BenchmarkSumColumn \
	-benchtime=1x -benchmem \
	-memprofile memprofile.out

gen-profiling-mainOne:
	go test -run TestRunMainOne