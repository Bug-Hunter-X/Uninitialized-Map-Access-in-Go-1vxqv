```go
func main() {
    var m map[string]int
    // Check if the map is initialized before accessing elements
    if m == nil {
        fmt.Println("Map is nil") //Handle the nil map case
    } else {
        fmt.Println(m["a"])
    }

    // Alternatively, initialize the map before use
    m = make(map[string]int)
    fmt.Println(m["a"]) // Prints 0, as expected for an empty map
    m["a"] = 10
    fmt.Println(m["a"]) // Prints 10
}
```