package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    ll := LinkedList{}
    ll.Insert(10)
    ll.Insert(20)
    ll.Insert(30)
    ll.Display() // Output: 10 -> 20 -> 30 -> nil
}

