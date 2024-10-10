package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) InOrder(w io.Writer) {
	if n != nil {
		n.Left.InOrder(w)
		_, _ = fmt.Fprintf(w, "value: %d, address: %p\n", n.Data, n)
		n.Right.InOrder(w)
	}
}

func insert(root *Node, data int) *Node {
	/*
		From the book "The art of immutable architecture" by  Michael L. Perry

		"No matter where you insert a new number, you will always end up creating a new root node,
		This new root node represents the shape of the tree after the insertion, the previous
		root node still exists."
	*/
	if root == nil {
		return &Node{data, nil, nil}
	}

	newNode := &Node{root.Data, nil, nil}

	if data < root.Data {
		(*newNode).Right = root.Right // take all on the right, we "insert" on left
	} else {
		(*newNode).Left = root.Left // take all on left, we "insert" on right
	}

	if data < root.Data {
		newNode.Left = insert(root.Left, data)
	} else if data > root.Data {
		newNode.Right = insert(root.Right, data)
	}
	return newNode
}

func loadTree() *Node {
	/*
		Load the tree from the DB
	*/
	var root = Node{12, nil, nil}
	root.Left = &Node{7, nil, nil}
	root.Left.Left = &Node{3, nil, nil}
	root.Left.Left.Left = &Node{1, nil, nil}
	root.Left.Right = &Node{9, nil, nil}
	root.Left.Left.Right = &Node{5, nil, nil}

	root.Right = &Node{27, nil, nil}
	root.Right.Right = &Node{32, nil, nil}
	root.Right.Left = &Node{17, nil, nil}
	root.Right.Left.Right = &Node{25, nil, nil}
	return &root
}

func main() {
	root := loadTree()

	root.InOrder(os.Stdout)

	fmt.Println("S-------------Original tree--------------")
	root.InOrder(os.Stdout)
	fmt.Println("E-------------Original tree--------------")

	fmt.Println("--------------Inserting new value of 22----------")
	newRoot := insert(root, 22)
	//newRoot = insert(newRoot, 36)
	fmt.Println("S-------------'Modified' tree--------------")
	newRoot.InOrder(os.Stdout)
	fmt.Println("E-------------'Modified' tree--------------")

	fmt.Println("S-------------Original tree not changed--------------")
	root.InOrder(os.Stdout)
	fmt.Println("E-------------Original tree--------------")
}


func getPath(root *Node, lst []int, data int) {
	if root == nil {
		fmt.Printf("path is: %+v\n", lst)
		return
	}
	lst = append(lst, root.Data)

	if data < root.Data {
		getPath(root.Left, lst, data)
	} else {
		getPath(root.Right, lst, data)
	}
}

