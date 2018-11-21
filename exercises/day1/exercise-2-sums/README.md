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