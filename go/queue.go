package main

import "fmt"

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
    if len(q.items) == 0 {
        fmt.Println("Queue is empty")
        return -1
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func main() {
    q := Queue{}
    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)
    fmt.Println(q.Dequeue()) // 10
    fmt.Println(q.Dequeue()) // 20
    fmt.Println(q.IsEmpty()) // false
}

