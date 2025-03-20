package fileutils

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ReadChunk reads and returns a chunk from an open file
func ReadChunk(f *os.File, chunkSizeInBytes int) ([]byte, error) {
	buff := make([]byte, chunkSizeInBytes)
	n, err := f.Read(buff)
	buff = buff[:n]
	if err != nil {
		return []byte{}, err
	}
	return buff, nil
}

func PrintHexa(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Got error: %v\n", err)
		return err
	}
	for _, v := range content {
		fmt.Printf("%x ", v)
	}
	return nil
}

func ProcessFileInChunks(filename string, chunkSizeInMB int, handleChunk func(chunk []byte)) {
	var bytes []byte
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Could not open file: %v\n", filename)
	}
	for err != io.EOF {
		bytes, err = ReadChunk(file, chunkSizeInMB*1000000)
		if err != nil && err != io.EOF {
			log.Fatalf("Error while reading chunk")
		}
		if len(bytes) > 0 {
			handleChunk(bytes)
		}
	}
}
