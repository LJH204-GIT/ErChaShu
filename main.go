package main

import "fmt"

type Node struct {
	left, right *Node
	value       int
}

func LevelOrder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.value)

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return result
}

func buildNode(this *Node, num int, a *Node) {
	if this.value == 0 {
		this.value = num
		return
	}
	a = this
	for {
		if num > a.value {
			if a.right == nil {
				a.right = &Node{value: num}
				return
			}
			a = a.right
			continue
		}
		if num < a.value {
			if a.left == nil {
				a.left = &Node{value: num}
				return
			}
			a = a.left
			continue
		}
	}
}
func BuildTree(nums []int, root *Node) {
	for i := 0; i < len(nums); i++ {
		buildNode(root, nums[i], &Node{})
	}
}

func DeleteNode(node *Node, root *Node) {
	if node == root.left && node.left == nil && node.right == nil {
		root.left = nil
		return
	}
	if node == root.right && node.left == nil && node.right == nil {
		root.right = nil
		return
	}
	deleteNode(node)
}

func deleteNode(n *Node) *Node {
	if n.left == nil && n.right == nil {
		n.value = -1
		return n
	}
	if n.left != nil {
		n.value = n.left.value
		temp := deleteNode(n.left)
		if temp == n.left {
			n.left = nil
		}
		return nil
	}

	if n.right != nil {
		n.value = n.right.value
		temp := deleteNode(n.right)
		if temp == n.right {
			n.right = nil
		}
		return nil
	}
	return nil
}

func FoundNode(num int, root *Node, a *Node) *Node {
	a = root
	for {
		if a == nil {
			return nil
		}
		if a.value == num {
			return a
		}
		if num < a.value {
			a = a.left
			continue
		}
		if num > a.value {
			a = a.right
			continue
		}

	}
}

func main() {
	arr := []int{100, 50, 150, 125, 175, 115, 135, 165, 200, 160, 170, 180, 210}
	root := &Node{value: 0}
	BuildTree(arr, root)
	fmt.Println(LevelOrder(root))
	DeleteNode(root.right, root)
	fmt.Println(LevelOrder(root))
	a := FoundNode(165, root, &Node{})
	println(a.value)
}
