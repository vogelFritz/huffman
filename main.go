package main

import (
	"container/heap"
	"fmt"
	"os"
)

type byteHeap []byte

func (b *byteHeap) Push(x any) {
	*b = append(*b, x.(byte))
}

func (b *byteHeap) Pop() any {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

func (b byteHeap) Len() int {
	return len(b)
}
func (b byteHeap) Less(i, j int) bool {
	return int(b[i]) < int(b[j])
}

func (b byteHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	args := os.Args
	fmt.Println(args)
	var b byteHeap = make(byteHeap, 10)

	doWithHeap(&b)

	s := []byte{
		'a',
		'c',
		'e',
		'd',
		'b',
	}
	h := byteHeap(s)
	heap.Init(&h)
	heap.Push(&h, byte(3))

	fmt.Println(h)
	fmt.Println(h.Pop())

	m := &byteHeap{2, 1, 5}
	heap.Init(m)
	heap.Push(m, byte(3))
	fmt.Printf("minimum: %d\n", (*m)[0])
	for m.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(m))
	}
	fmt.Println()
}

func doWithHeap(h heap.Interface) {
	heap.Init(h)
	h.Push(uint8(2))
	fmt.Println("Pushed 2")
	x := h.Pop()
	fmt.Println("pop:", x)
	x = h.Pop()
	fmt.Println("pop:", x)
}
