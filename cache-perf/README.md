Code from this article: https://www.ardanlabs.com/blog/2023/07/getting-friendly-with-cpu-caches.html

## Before: ~10B cache misses

```
$ make perfbench
perf stat -e cache-misses ./cache-perf.test -test.bench . -test.benchtime=10s -test.count=5
goos: linux
goarch: amd64
pkg: github.com/rtfb/sketchbook/cache-perf
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkCountryCount-20    	    4596	   2544228 ns/op
BenchmarkCountryCount-20    	    4446	   2656406 ns/op
BenchmarkCountryCount-20    	    4573	   2648627 ns/op
BenchmarkCountryCount-20    	    4290	   2617001 ns/op
BenchmarkCountryCount-20    	    4488	   2569501 ns/op
PASS

 Performance counter stats for './cache-perf.test -test.bench . -test.benchtime=10s -test.count=5':

     2,119,480,424      cpu_atom/cache-misses:u/                                                (0.25%)
     7,722,507,847      cpu_core/cache-misses:u/                                                (99.77%)

      59.802650990 seconds time elapsed

      59.908770000 seconds user
       0.188140000 seconds sys
```

## After: ~215M cache misses

```
$ make perfbench
perf stat -e cache-misses ./cache-perf.test -test.bench . -test.benchtime=10s -test.count=5
goos: linux
goarch: amd64
pkg: github.com/rtfb/sketchbook/cache-perf
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkCountryCount-20    	  179596	     67544 ns/op
BenchmarkCountryCount-20    	  179594	     67160 ns/op
BenchmarkCountryCount-20    	  180772	     67128 ns/op
BenchmarkCountryCount-20    	  180088	     67165 ns/op
BenchmarkCountryCount-20    	  179932	     67119 ns/op
PASS

 Performance counter stats for './cache-perf.test -test.bench . -test.benchtime=10s -test.count=5':

       213,313,235      cpu_atom/cache-misses:u/                                                (0.41%)
         2,073,997      cpu_core/cache-misses:u/                                                (99.72%)

      63.950004378 seconds time elapsed

      64.084403000 seconds user
       0.120210000 seconds sys
```
