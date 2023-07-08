package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func Walk(tree *Tree, ch chan int) {
	if tree == nil {
		return
	}
	Walk(tree.left, ch)
	ch <- tree.value
	Walk(tree.right, ch)
}

func Walker(tree *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(tree, ch)
		close(ch)
	}()
	return ch
}

func Same(tree1, tree2 *Tree) bool {
	ch1 := Walker(tree1)
	ch2 := Walker(tree2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func randBst(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = addBst(t, (1+v)*k)
	}
	return t
}

func addBst(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{v, nil, nil}
	}
	if v < t.value {
		t.left = addBst(t.left, v)
		return t
	}
	t.right = addBst(t.right, v)
	return t
}

func main() {
	fmt.Println(Same(randBst(100, 2), randBst(100, 2)), "Same trees")
	fmt.Println(Same(randBst(100, 4), randBst(100, 2)), "Different values")
	fmt.Println(Same(randBst(101, 1), randBst(100, 1)), "Different sizes")
}
