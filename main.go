package main

import (
	"fmt"
)

const size = 5

type Node struct {

	Val string
	Left *Node
	Right *Node

}

type Queue struct {

	Head *Node
	Tail *Node
	Length int

}

type Cache struct {

	Queue Queue
	Hash  Hash

}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func NewQueue() Queue {

	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{
		Head : head,
		Tail : tail,
	}

}

type Hash map[string]*Node

func (c *Cache) Check(str string){
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {

	fmt.Printf("Removing %s\n", n.Val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length--
	delete(c.Hash, n.Val)
	return n
}

func (c *Cache) Add(n *Node) {

	fmt.Printf("Adding %s\n", n.Val)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n
	c.Queue.Length ++
	if c.Queue.Length > size {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display(){
	c.Queue.Display()
}

func (q *Queue) Display(){
	node := q.Head.Right

	fmt.Printf("%d - [", q.Length)

	for i:=0; i<q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i<q.Length - 1 {
			fmt.Println("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Cache server")

	cache := NewCache()

	for _, word := range []string{"one", "two", "three", "four", "five"} {

		cache.Check(word)
		cache.Display()

	}
}
