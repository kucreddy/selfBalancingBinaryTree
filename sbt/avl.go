package sbt

import (
	"errors"
	"fmt"
	"strconv"
)

type node struct {
	val, leftLen, rightLen int
	left, right, parent *node
}

type TreeA struct {
	root *node
}

func (a *TreeA) InsertFromSlice(input []int) error {
	if a.root != nil {
		return errors.New("tree not empty, cannot initialize")
	}
	for _, val := range input {
		a.Insert(val)
	}
	return nil
}

func (a *TreeA) Insert(val int) {
	if a.root == nil {
		root := &node{val, 0, 0, nil, nil, nil}
		a.root = root
		return
	}
	curr := a.root
	newNode := &node{val, 0, 0, nil, nil, nil}
	for {
		// ignore in case of duplication
		if val == curr.val {
			return
		}
		if val < curr.val {
			if curr.left == nil {
				curr.left = newNode
				newNode.parent = curr
				break
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = newNode
				newNode.parent = curr
				break
			}
			curr = curr.right
		}
	}
	updateLengths(newNode)
	a.balance(newNode)
}

func (a *TreeA)rotateLeft(parent, child *node) {
	elder := parent.parent
	if elder == nil {
		// parent is the root of the tree
		a.root = child
	}else {
		if elder.left == parent {
			elder.left = child
		} else {
			elder.right = child
		}
	}
	parent.right = child.left
	child.left = parent

	child.parent = elder
	parent.parent = child

	// updating lengths
	if parent.right != nil {
		parent.rightLen = parent.right.leftLen + 1
		if parent.right.rightLen > parent.right.leftLen {
			parent.rightLen = parent.right.rightLen + 1
		}
	}else{
		parent.rightLen = 0
	}

	child.leftLen = parent.leftLen + 1
	if parent.rightLen > parent.leftLen {
		child.leftLen = parent.rightLen + 1
	}
	updateLengths(child)
}

func (a *TreeA)rotateRight(parent, child *node) {
	elder := parent.parent
	if elder == nil {
		// parent is the root of the tree
		a.root = child
	}else {
		if elder.left == parent {
			elder.left = child
		} else {
			elder.right = child
		}
	}
	parent.left = child.right
	child.right = parent

	child.parent = elder
	parent.parent = child

	//updating lengths
	if parent.left != nil {
		parent.leftLen = parent.left.leftLen + 1
		if parent.left.rightLen > parent.left.leftLen {
			parent.leftLen = parent.left.rightLen + 1
		}
	}else{
		parent.leftLen = 0
	}

	child.rightLen = parent.leftLen + 1
	if parent.rightLen > parent.leftLen {
		child.rightLen = parent.rightLen + 1
	}

	updateLengths(child)
}

func updateLengths(newNode *node) {
	curr := newNode
	parent := newNode.parent
	for parent != nil {
		if parent.left == curr {
			parent.leftLen = curr.leftLen + 1
			if curr.rightLen > curr.leftLen {
				parent.leftLen = curr.rightLen + 1
			}
			if parent.leftLen <= parent.rightLen {
				return
			}
		}else {
			parent.rightLen = curr.leftLen + 1
			if curr.rightLen > curr.leftLen {
				parent.rightLen = curr.rightLen + 1
			}
			if parent.rightLen <= parent.leftLen {
				return
			}
		}
		curr = parent
		parent = parent.parent
	}
}

func (a *TreeA) balance(newNode *node) {
	// take care of the case where root changes

	// Z = first unbalanced node encountered while traversing
	// from newNode towards root
	// Y, X are the nodes on the path from newNode to Z
	// spatial positions below represent parent-child strictly, but not left-right
	//	 	     /
	//	 		Z
	//	 	   / \
	// 		  Y
	//	 	 / \
	//	 	X
	//     / \
	X := newNode
	if X.parent == nil {
		return
	}
	Y := X.parent
	if Y.parent == nil {
		return
	}
	Z := Y.parent

	// no way parent of newNode will be unbalanced
	// start checking from Z
	for Z != nil {
		diff := Z.leftLen - Z.rightLen
		if diff < -1 || diff > 1 {
			if Z.left == Y {
				if Y.left == X {
					a.rotateRight(Z, Y)
				} else {
					a.rotateLeft(Y, X)
					a.rotateRight(Z, X)
				}
			} else {
				if Y.right == X {
					a.rotateLeft(Z, Y)
				} else {
					a.rotateRight(Y, X)
					a.rotateLeft(Z, X)
				}
			}
			return
		}
		X = Y
		Y = Z
		Z = Z.parent
	}
}

func (a *TreeA)PrintInorder() {
	printInOrder(a.root, 0)
}

func printInOrder (a *node, nodeHeight int) {
	if a != nil {
		printInOrder(a.right, nodeHeight + 1)
	}
	for i := 0; i < nodeHeight * 3; i++ {
		fmt.Print("  ")
	}
	if a == nil {
		fmt.Println("[" + strconv.Itoa(nodeHeight) + "]")
	}else {
		fmt.Println(a.val)
	}
	if a != nil {
		printInOrder(a.left, nodeHeight+1)
	}
}
