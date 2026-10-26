package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}

func BFS(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.Value)
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func SearchBST(root *Node, target int) *Node {
	current := root
	for current != nil {
		if current.Value == target {
			return current
		} else if current.Value < target {
			current = current.Right
		} else {
			current = current.Left
		}

	}
	return nil
}

func main() {
	node35 := &Node{Value: 35}
	node40 := &Node{Value: 40, Left: node35}
	node20 := &Node{Value: 20}
	node30 := &Node{Value: 30, Left: node20, Right: node40}
	node60 := &Node{Value: 60}
	node80 := &Node{Value: 80}
	node70 := &Node{Value: 70, Left: node60, Right: node80}

	root := &Node{
		Value: 50,
		Left:  node30,
		Right: node70,
	}

	found := SearchBST(root, 35)
	if found != nil {
		fmt.Println("found", found.Value)
	} else {
		fmt.Println("not found")
	}
	found = SearchBST(root, 50)
	if found != nil {
		fmt.Println("found", found.Value)
	} else {
		fmt.Println("not found")
	}
	found = SearchBST(root, 99)
	if found != nil {
		fmt.Println("found", found.Value)
	} else {
		fmt.Println("not found")
	}
}
