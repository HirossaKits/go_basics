package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	// PathSeparator()
	// ShortDeclaration()
	// TypeConversion()
	CommandlineArgs()

}

// PathSeparator separate path
func PathSeparator() {
	var dir, file string
	dir, file = path.Split("css/main.css")

	fmt.Println("dir:", dir)
	fmt.Println("file:", file)

	_, file = path.Split("css/main.css")
	fmt.Println("file:", file)

	ndir, nfile := path.Split("css/main.css")
	fmt.Println("file:", ndir, nfile)
}

// ShortDeclaration about declaration
func ShortDeclaration() {

	scoreA := 0    // ❌ DONT!
	var scoreB int // 👍  already score = 0
	fmt.Println(scoreA, scoreB)

	var (
		// related:
		video    string
		duration int
		current  int
	)
	fmt.Println(video, duration, current)

	var widthA, heightA = 100, 50 // ❌ DONT!
	widthB, heightB := 100, 50    // 👍
	fmt.Println(widthA, heightA, widthB, heightB)

	// DONT!
	// width = 50 // assigns 50 to width
	// color := "red" // new variable: color
	width, color := 50, "red" // 👍 same as above
	fmt.Println(width, color)
}

// TypeConversion type conversion
func TypeConversion() {
	speed := 100 // speed is int
	force := 2.5 // force is float64

	speed = int(float64(speed) * force)
	fmt.Println(speed)
}

// CommandlineArgs commandline args
func CommandlineArgs() {
	fmt.Printf("%#v\n", os.Args)
	// "go run" tool creates temporary executable files.
	// "go build" creates an executable file permanently.
}

// Read ❌ NOT IDIOMATIC
// func Read(buffer *Buffer, inBuffer []byte) (size int, err error) {
// 	if buffer.empty() {
// 		buffer.Reset()
// 	}
// 	size = copy(
// 		inBuffer,
// 		buffer.buffer[buffer.offset:])

// 	buffer.offset += size
// 	return size, nil
// }

// ⭕ IDIOMATIC
func Read(b *Buffer, p []byte) (n int, err error) {
	if b.empty() {
		b.Reset()
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return n, nil
}
