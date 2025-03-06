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

func deleteNode1(root *Node, node *Node) *Node {
	if root.left == node || root.right == node {
		return root
	}
	if root.left == nil && root.right == nil {
		return nil
	}

	temp := deleteNode1(root.left, node)

	if temp == nil {
		temp = deleteNode1(root.right, node)
	}
	return temp
}

func DDeleteNode(root *Node, node *Node) {
	a := deleteNode1(root, node) // 找父节点
	if a.left == node {
		a.left = node.left
		add2(node.left, node.right)

	}
	if a.right == node {
		a.right = node.right
		add1(node.right, node.left)
	}
}
func add1(root *Node, add *Node) {
	if root.left == nil && root.right == nil {
		root.left = add
		return
	}
	if root.left != nil {
		add1(root.left, add)
	}
}
func add2(root *Node, add *Node) {
	if root.left == nil && root.right == nil {
		root.right = add
		return
	}
	if root.right != nil {
		add1(root.right, add)
	}
}
func add3(father *Node, node *Node) {
	if node == father.left {

	}
	if node.left == nil && node.right == nil {

	}
}

//func deleteNode21(des *Node) *Node {
//	if des.right == nil && des.left == nil {
//		return des
//	}
//	temp := deleteNode21(des.right)
//	temp.left = des.left
//	return des
//}
//func deleteNode22(des *Node) *Node {
//	if des.right == nil && des.left == nil {
//		return des
//	}
//	temp := deleteNode22(des.left)
//	temp.right = des.right
//	return des
//}

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

// PRF ==============================
func PRF(root *Node, node *Node) {
	if root.left != nil {
		PRF(root.left, node)
	}
	if root.left == nil {
		root.left = node
	}
}
func DNM(root *Node, node *Node) {
	if root.right == node {
		if node.left == nil {
			root.right = node.right
			return
		}
		if node.left == nil && node.right == nil {
			root.right = nil
			return
		}
		root.right = node.right
		PRF(root.right, node.left)
		return
	}
	if root.left == node {
		if node.right == nil {
			root.left = node.left
			return
		}
		if node.left == nil && node.right == nil {
			root.left = nil
			return
		}
		root.left = node.right
		PRF(root.left, node.left)
		return
	}
	if root.left != nil {
		DNM(root.left, node)
		DNM(root.right, node)
	}
}

// =====================================
func main() {
	arr := []int{100, 50, 25, 150 /*, 125*/, 175 /* 115, 135,*/, 165, 200, 160, 170, 180, 210}
	root := &Node{value: 0}
	BuildTree(arr, root)
	fmt.Println(LevelOrder(root))
	//DeleteNode(root.right, root)
	DNM(root, root.right) // 最佳递归去节点

	fmt.Println(LevelOrder(root))
	//a := FoundNode(165, root, &Node{})
	//println(a.value)
}
