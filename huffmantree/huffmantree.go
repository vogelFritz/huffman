package huffmantree

import (
	"fmt"
	"strings"
)

type Node struct {
	Words    []byte
	Freq     int
	Children []Node
}

type NodeHeap []Node

func (b *NodeHeap) Push(x any) {
	*b = append(*b, x.(Node))
}

func (b *NodeHeap) Pop() any {
	old := *b
	n := len(old)
	x := old[n-1]
	*b = old[0 : n-1]
	return x
}

func (b NodeHeap) Len() int {
	return len(b)
}
func (b NodeHeap) Less(i, j int) bool {
	return b[i].Freq < b[j].Freq
}

func (b NodeHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b NodeHeap) GetCodes() map[byte][]byte {
	arr := []Node(b)
	if len(arr) != 2 {
		return nil
	}
	r := make(map[byte][]byte)
	for i, elem := range arr {
		for _, symbol := range elem.Words {
			r[symbol] = []byte{byte(i)}
		}
		for k, v := range elem.GetCodes() {
			r[k] = append(r[k], v...)
		}
	}
	return r
}

func (n Node) GetCodes() map[byte][]byte {
	r := make(map[byte][]byte)
	for i, child := range n.Children {
		for _, symbol := range child.Words {
			r[symbol] = []byte{byte(i)}
		}
		for k, v := range child.GetCodes() {
			r[k] = append(r[k], v...)
		}
	}
	return r
}

func (b NodeHeap) String() string {
	if len(b) == 0 {
		return ""
	}
	r := b[0].String()
	for _, v := range b[1:] {
		r = fmt.Sprintf("%v%v", r, v.String())
	}
	return r
}

func (n Node) String() string {
	r := fmt.Sprintf("words: %v\nfreq: %v\n", n.Words, n.Freq)
	for _, v := range n.Children {
		vstr := indent(v.String())
		r = fmt.Sprintf("%v%v", r, vstr)
	}
	return r
}

func indent(s string) string {
	lines := strings.Split(s, "\n")
	var indented []string
	for _, line := range lines {
		indented = append(indented, fmt.Sprintf("  %v", line))
	}
	return fmt.Sprintf("%v\n", strings.Join(indented, "\n"))
}
