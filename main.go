package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Key int
	Height int
	Parent *TreeNode
	Left *TreeNode
	Right *TreeNode
}

func getMaxHeight (root *TreeNode) int  {
	var leftHeight int
	var rightHeight int

	if root.Left != nil {
		leftHeight = root.Left.Height
	} else {
		leftHeight = 0
	}

	if root.Right != nil {
		rightHeight = root.Right.Height
	} else {
		rightHeight = 0
	}

	return int(math.Max(float64(rightHeight), float64(leftHeight)))
}

func balanceTreeClockWise(root *TreeNode) *TreeNode {
	leftRoot := root.Left
	root.Left = leftRoot.Right
	leftRoot.Right = root
	leftRoot.Parent = root.Parent
	root.Parent = leftRoot

	root.Height = getMaxHeight(root)+1
	leftRoot.Height = getMaxHeight(leftRoot)+1

	return leftRoot
}

func balanceTreeCounterClockWise(root *TreeNode) *TreeNode {
	rightRoot := root.Right
	root.Right = rightRoot.Left
	rightRoot.Left = root
	rightRoot = root.Parent
	root.Parent = rightRoot

	root.Height = getMaxHeight(root)+1
	rightRoot.Height = getMaxHeight(rightRoot)+1
	return rightRoot
}

func balance(root *TreeNode) *TreeNode {
	var rightHeight int
	var leftHeight int

	if root.Left != nil {
		leftHeight = root.Left.Height
	} else {
		leftHeight = 0
	}

	if root.Right != nil {
		rightHeight = root.Right.Height
	} else {
		rightHeight = 0
	}

	if math.Abs(float64(rightHeight - leftHeight)) >= 2 {
		if rightHeight > leftHeight {
			return balanceTreeCounterClockWise(root)
		} else {
			return balanceTreeClockWise(root)
		}
	} else {
		root.Height = int(math.Max(float64(leftHeight), float64(rightHeight)))+1
		return root
	}
}

func insert (root *TreeNode, node *TreeNode) *TreeNode {

	if root.Key == node.Key {
		root.Val = node.Val
		return root
	}

	if root.Key > node.Key {
		if root.Left == nil  {
			root.Left = node
			node.Parent = root
		} else {
			insert(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			insert(root.Right, node)
		}
	}

	return balance(root)
}

func removeRight (parent *TreeNode, node *TreeNode) *TreeNode {

	if node.Left == nil {
		if node.Right == nil {
			parent.Right = nil
			return nil
		}

		node.Right.Parent = parent
		if parent != nil {
			parent.Right = node.Right
		}
		return balance(node.Right)
	} else {
		rightMost := node.Left

		if rightMost == nil {
			parent.Right = nil
			return nil
		}

		for rightMost.Right != nil {
			rightMost = rightMost.Right
		}

		if rightMost.Key != node.Left.Key {
			rightMost.Left = node.Left
			node.Left.Parent = rightMost
		}

		if parent != nil {
			parent.Right = rightMost
		}
			rightMost.Parent = parent
			rightMost.Right = node.Right

			return balance(rightMost)
	}

}

func removeLeft (parent *TreeNode, node *TreeNode) *TreeNode {
	if node.Right == nil {
		if node.Left == nil {
			parent.Left = nil
			return nil
		}

		node.Left.Parent = parent
		if parent != nil {
			parent.Left = node.Left
		}
		return node.Left
	} else {
		leftMost := node.Right

		if leftMost == nil {
			parent.Left = nil
			return nil
		}

		for leftMost.Left != nil {
			leftMost = leftMost.Left
		}

		if leftMost.Key != node.Right.Key {
			leftMost.Right = node.Right
			node.Right.Parent = leftMost
		}

		if parent != nil {
			parent.Left = leftMost
		}
		leftMost.Parent = parent
		leftMost.Left = node.Left
		return balance(leftMost)
	}
}

func remove(root *TreeNode) *TreeNode {
	parent := root.Parent

	if parent != nil &&  parent.Right != nil && parent.Right.Key == root.Key {
		removeRight(parent, root)
		return balance(parent)
	}

	removeLeft(parent, root)
	return balance(parent)
}

func search(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}

	if root.Key == key {
		return root.Val
	}

	if root.Key > key {
		return search(root.Left, key)
	} else {
		return search(root.Right, key)
	}
}

func searchNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Key == key {
		return root
	}

	if root.Key > key {
		return searchNode(root.Left, key)
	} else {
		return searchNode(root.Right, key)
	}
}

type MyHashMap struct {
	root *TreeNode
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}


/** value will always be non-negative. */
func (hm *MyHashMap) Put(key int, value int)  {
	if hm.root == nil {
		hm.root = &TreeNode{value, key, 1, nil, nil, nil}
	}

	hm.root = insert(hm.root, &TreeNode{value, key, 1, nil, nil, nil})
}


/** Returns the value to which the specified key is mapped, or -1 if hm map contains no mapping for the key */
func (hm *MyHashMap) Get(key int) int {
	return search(hm.root, key)
}


/** Removes the mapping of the specified value key if hm map contains a mapping for the key */
func (hm *MyHashMap) Remove(key int)  {
	node := searchNode(hm.root, key)

	if node == nil {
		return
	}

	if node.Key == hm.root.Key {
		hm.root = remove(node)
	} else {
		remove(node)
	}

}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	hm := Constructor()

	hm.Put(1, 1)
	hm.Put(2, 2)

	fmt.Printf("value of %d is %d\n", 1, hm.Get(1))
	fmt.Printf("value of %d is %d\n", 3, hm.Get(3))

	hm.Put(2, 1)

	fmt.Printf("value of %d is %d\n", 2, hm.Get(2))
	hm.Remove(2)
	fmt.Printf("value of %d is %d\n", 2, hm.Get(2))

}
