
.PHONY: bench
bench:
	go test -bench . -benchtime 10s -count 5

# sudo apt install linux-tools-common linux-tools-6.6.10-76060610-generic
.PHONY: perfbench
perfbench:
	perf stat -e cache-misses ./cache-perf.test -test.bench . -test.benchtime=10s -test.count=5

build:
	go test -c
