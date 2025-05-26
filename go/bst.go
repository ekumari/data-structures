package main

import "fmt"

type Node struct {
    data  int
    left  *Node
    right *Node
}

type BST struct {
    root *Node
}

func (bst *BST) Insert(data int) {
    bst.root = insertRec(bst.root, data)
}

func insertRec(root *Node, data int) *Node {
    if root == nil {
        return &Node{data: data}
    }
    if data < root.data {
        root.left = insertRec(root.left, data)
    } else if data > root.data {
        root.right = insertRec(root.right, data)
    }
    return root
}

func (bst *BST) Inorder() {
    inorderRec(bst.root)
    fmt.Println()
}

func inorderRec(root *Node) {
    if root != nil {
        inorderRec(root.left)
        fmt.Print(root.data, " ")
        inorderRec(root.right)
    }
}

func main() {
    bst := BST{}
    bst.Insert(50)
    bst.Insert(30)
    bst.Insert(70)
    bst.Insert(20)
    bst.Insert(40)
    bst.Insert(60)
    bst.Insert(80)

    bst.Inorder() // Output: 20 30 40 50 60 70 80
}

