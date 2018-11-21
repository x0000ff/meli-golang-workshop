# Exercise 2

> Texto original:
> https://talks.godoc.org/github.com/ifreddyrondon/go-workshop/santiago-nov2018/go-day1.slide#39

Hacer un programa que reciba un parámetro `p` por consola y que calcule la suma de los valores `v` tales que 0 `<= v <= p` , siendo `v` un número divisible por `3` o por `5`

- Ayuda para leer argumentos desde la consola

```go
for i := 1; i < len(os.Args); i++ {
    fmt.Println(os.Args[i])
}
```

## Result

```termanal
$ go run main.go

Usage:
$ go run main <MAX>
$ go run main 31

Only integer values allowed
exit status 1
```

```termanal
$ go run main.go x

Incorrect input: "x", expected integer
exit status 1
```

```termanal
$ go run main.go -2
0%
```

```termanal
$ go run main.go 3
3%
```

```termanal
$ go run main.go 5
8%
```

```termanal
$ go run main.go 10
33%
```