package main

import (
	"errors"
	"fmt"
)

// Node represents a node in a binary search tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// PreOrder visits Root -> Left -> Right.
func PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

// InOrder visits Left -> Root -> Right.
// For a valid BST, this prints values in ascending order.
func InOrder(node *Node) {
	if node == nil {
		return
	}

	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}

// PostOrder visits Left -> Right -> Root.
func PostOrder(node *Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}

// BFS visits nodes level by level using a FIFO queue.
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

// SearchBST iteratively searches for target in a binary search tree.
// It returns the matching node, or nil if the value is not present.
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

var ErrDuplicateValues = errors.New("value already exists in BST")

func InsertBST(root *Node, value int) (*Node, error) {
	if root == nil {
		return &Node{Value: value}, nil
	}
	current := root
	for current != nil {
		if value == current.Value {
			return root, ErrDuplicateValues
		} else if value < current.Value {
			if current.Left == nil {
				current.Left = &Node{Value: value}
				return root, nil
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &Node{Value: value}
				return root, nil
			}
			current = current.Right
		}
	}
	return root, nil
}

func DeleteBST(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if root.Value < value {
		root.Right = DeleteBST(root.Right, value)
	} else if root.Value > value {
		root.Left = DeleteBST(root.Left, value)
	} else if root.Value == value {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := FindMin(root.Right)
		root.Value = successor.Value
		root.Right = DeleteBST(root.Right, successor.Value)
	}
	return root
}

func FindMin(root *Node) *Node {
	if root == nil {
		return nil
	}
	current := root

	for current.Left != nil {
		current = current.Left
	}
	return current
}

func main() {
	node37 := Node{Value: 37}
	node35 := Node{Value: 35, Right: &node37}
	node40 := Node{Value: 40, Left: &node35}
	node20 := Node{Value: 20}
	node30 := Node{Value: 30, Left: &node20, Right: &node40}
	node60 := Node{Value: 60}
	node80 := Node{Value: 80}
	node70 := Node{Value: 70, Left: &node60, Right: &node80}
	root := &Node{Value: 50, Left: &node30, Right: &node70}

	root = DeleteBST(root, 30)
	root, err := InsertBST(root, 45)
	if err != nil {
		fmt.Println(err)
		return
	}
	root = DeleteBST(root, 70)
	InOrder(root)
	PostOrder(root)
	PreOrder(root)
	BFS(root)
	fmt.Println(SearchBST(root, 30))
	fmt.Println(SearchBST(root, 70))
	fmt.Println(SearchBST(root, 37))
	fmt.Println(SearchBST(root, 45))
	fmt.Println("Hello world!")

}
