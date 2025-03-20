package main

import (
	"container/heap"
	"fmt"
	"os"
	"sync"

	"github.com/vogelFritz/huffman/fileutils"
	"github.com/vogelFritz/huffman/huffmantree"
)

func getFrequencyTable(filename string) map[byte]int {
	// using fanout concurrency pattern
	freq := make(map[byte]int)
	var wg sync.WaitGroup
	results := make(chan map[byte]int)
	processChunk := func(chunk []byte) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			freq := make(map[byte]int)
			for _, r := range chunk {
				freq[r]++
			}
			results <- freq
		}()
	}
	fileutils.ProcessFileInChunks(filename, 1, processChunk)
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		for r, count := range result {
			freq[r] += count
		}
	}
	return freq
}

func createHuffmanTree(freqTable map[byte]int) huffmantree.NodeHeap {
	var tree huffmantree.NodeHeap
	heap.Init(&tree)
	for k, v := range freqTable {
		heap.Push(&tree, huffmantree.Node{
			Words: []byte{k},
			Freq:  v,
		})
	}
	for tree.Len() > 2 {
		c1 := heap.Pop(&tree).(huffmantree.Node)
		c2 := heap.Pop(&tree).(huffmantree.Node)
		n := huffmantree.Node{
			Words:    append(c1.Words, c2.Words...),
			Freq:     c1.Freq + c2.Freq,
			Children: []huffmantree.Node{c1, c2},
		}
		heap.Push(&tree, n)
	}
	return tree
}

func main() {
	encode := os.Args[1] == "-c"
	filename := os.Args[2]
	// outputPath := os.Args[3]

	if encode {
		freq := getFrequencyTable(filename)
		fmt.Println(freq)
		tree := createHuffmanTree(freq)
		fmt.Println("huffman tree:")
		fmt.Println(tree.String())
		codes := tree.GetCodes()
		fmt.Println(codes)
		// Use codes to create encrypted file with codes table at the start
	}
}
