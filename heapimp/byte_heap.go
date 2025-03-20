package heapimp

type nodo struct {
	total float32
	symbols []byte
	elems []nodo
}

type ByteHeap []byte

func (b *ByteHeap) Push(x any) {
	*b = append(*b, x.(byte))
}

func (b *ByteHeap) Pop() any {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

func (b ByteHeap) Len() int {
	return len(b)
}
func (b ByteHeap) Less(i, j int) bool {
	return int(b[i]) < int(b[j])
}

func (b ByteHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
