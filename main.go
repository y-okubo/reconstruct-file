package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Chunk is a pseudo chunk of split download
type Chunk struct {
	Data   []byte
	Offset int64
}

// BUFSIZE is divided size
const BUFSIZE = 2

func main() {
	fi, err := os.Stat("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(fi.Size())

	// chunks := make([]Chunk, fi.Size()/2+fi.Size()%2)

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf []byte
	var offset int64
	var chunks []Chunk

	for {

		buf = make([]byte, BUFSIZE)

		n, err := f.ReadAt(buf, offset)
		if n == 0 {
			fmt.Println("n == 0")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		chunk := Chunk{}
		chunk.Offset = offset
		chunk.Data = buf
		chunks = append(chunks, chunk)

		offset, err = f.Seek(2, 1)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	// 斑にダウンロードされることを想定してシャッフル
	shuffle(chunks)

	// fmt.Println(chunks)

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, chunk := range chunks {
		// fmt.Println(chunk.Offset)
		file.WriteAt(chunk.Data, chunk.Offset) // オフセットに従い書き込み
	}
}

func shuffle(data []Chunk) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}
