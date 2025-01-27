```go
func main() {
    var x int
    fmt.Println(x) // This will print 0
    defer fmt.Println(x)
    x = 10
}
```