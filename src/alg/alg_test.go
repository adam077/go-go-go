package alg

import (
	"fmt"
	"testing"
)

type Node struct {
	Value int
	Next  *Node
}

// 记住map可以降低时间复杂度
func TestFunc(t *testing.T) {
	a := AddTwoNumber(&Node{5, &Node{6, nil}}, &Node{5, &Node{6, nil}})
	result := &Node{0, nil}
	for a != nil {
		temp := &Node{a.Value, result}
		result.Next = temp
		a = a.Next
		// result.Next = result // 这个是绝对不行的
	}
	for result != nil {
		fmt.Println(result.Value)
		result = result.Next
	}
}

func AddTwoNumber(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	var dummy = &Node{0, nil}
	var current = dummy
	var carry = 0 // 进位
	for a != nil && b != nil {
		dig := a.Value + b.Value + carry
		val := dig % 10
		carry = dig / 10
		newNode := &Node{val, nil}
		current.Next = newNode
		current = current.Next
		a = a.Next
		b = b.Next
	}
	for a != nil {
		val := (a.Value + carry) % 10
		carry = (a.Value + carry) / 10
		current.Next = &Node{val, nil}
		current = current.Next
		a = a.Next
	}
	for b != nil {
		val := (b.Value + carry) % 10
		carry = (b.Value + carry) / 10
		current.Next = &Node{val, nil}
		current = current.Next
		b = b.Next
	}
	if carry != 0 {
		current.Next = &Node{carry, nil}
	}
	return dummy.Next
}
