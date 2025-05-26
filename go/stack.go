package main

import "fmt"

// Stack defines a stack of integers
type Stack struct {
    items []int
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack
// Returns -1 if the stack is empty
func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return -1
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top
}

// Peek returns the top element without removing it
func (s *Stack) Peek() int {
    if len(s.items) == 0 {
        fmt.Println("Stack is empty")
        return -1
    }
    return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
    return len(s.items)
}

func main() {
    stack := Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println("Stack size:", stack.Size())      // Output: 3
    fmt.Println("Top element:", stack.Peek())     // Output: 30

    fmt.Println("Popped:", stack.Pop())           // Output: 30
    fmt.Println("Popped:", stack.Pop())           // Output: 20
    fmt.Println("Stack empty?", stack.IsEmpty())  // Output: false

    fmt.Println("Popped:", stack.Pop())           // Output: 10
    fmt.Println("Stack empty?", stack.IsEmpty())  // Output: true
}

