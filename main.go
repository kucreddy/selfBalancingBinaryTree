package main

import (
	"fmt"
	"github.com/YashwanthReddy098/selfBalancingBinaryTree/sbt"
)

func main() {
	myTree := sbt.TreeA{}
	a := make([]int, 0)
	for i:= 0; i < 10; i++ {
		a = append(a, i)
	}
	err := myTree.InsertFromSlice(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("tilt your head left to see the tree")
	myTree.PrintInorder()
}
