package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.

var bufPool = sync.Pool{New: func() interface{} {
	fmt.Println("Creating new bytes.Buffer instance")
	return new(bytes.Buffer)
}}

func log(w io.Writer, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(b)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
