package main

import (
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

const file1 string = "/tmp/file1.txt"
const file2 string = "/tmp/file2.txt"

func write() {
	data := []byte("hello\ngo\n")
	err := os.WriteFile(file1, data, 0644)
	checkError(err)

	f, err := os.Create(file2)
	checkError(err)

	defer f.Close()

	n, err := f.Write(data)
	fmt.Printf("Wrote %d bytes \n", n)

	n, err = f.WriteString("test\n")
	fmt.Printf("Wrote %d bytes\n", n)

	f.Sync()
}

func read() {
	data, err := os.ReadFile(file1)
	checkError(err)
	fmt.Println("Read data ", string(data))

	f, err := os.Open(file2)
	checkError(err)

	defer f.Close()

	dataRead := make([]byte, 5)
	n, err := f.Read(dataRead)
	fmt.Printf("Read %d bytes \n", n)
}

func main() {

	write()
	read()
}
