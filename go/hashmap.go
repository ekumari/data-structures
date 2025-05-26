package main

import "fmt"

func main() {
    m := make(map[string]int)

    // Insert
    m["apple"] = 10
    m["banana"] = 20
    m["orange"] = 30

    fmt.Println("Map:", m)

    // Access
    fmt.Println("Apple price:", m["apple"])

    // Delete
    delete(m, "banana")
    fmt.Println("After deletion:", m)

    // Check if key exists
    val, ok := m["banana"]
    if !ok {
        fmt.Println("Banana not found")
    } else {
        fmt.Println("Banana price:", val)
    }
}

