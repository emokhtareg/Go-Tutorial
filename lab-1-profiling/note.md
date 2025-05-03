

https://github.com/google/pprof
go install github.com/google/pprof@latest


- You will have to install the graphviz library in ubuntu you can do it with
```
sudo apt-get install graphviz
```


```
$ go test -bench=.
```

```
$ go test -bench='BenchmarkSquare2'
```

```
$ go test -bench='BenchmarkSquare2' -count=3
```

```
$ go test -bench='BenchmarkSquare2' -count=3 -benchtime='1s'
```

```
$ go test -bench=. -cpuprofile='cpu.prof' -memprofile='mem.prof'
$ go tool pprof cpu.prof
```



```
$ go run . -cpu=cpu.prof -mem=mem.prof
$ go tool pprof cpu.prof 

```

```
$ go tool pprof -http localhost:8080 cpu.prof 
```