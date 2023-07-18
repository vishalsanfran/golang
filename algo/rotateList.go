/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	data int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func (l *List) add(val int) {
	node := &ListNode{data: val}
	if l.head == nil {
		l.head = node
		return
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		l += 1
	}

	rots := k % l
	if rots == 0 {
		return head
	}
	cur := head
	for i := 0; i < l-rots-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	tail.Next = head
	return newHead
}

func main() {
	l := &List{head: &ListNode{data: 1}}
	for i := 2; i < 6; i++ {
		l.add(i)
	}
	fmt.Println(l.head.data)
	l.head = rotateRight(l.head, 3)
	fmt.Println(l.head.data)
}
