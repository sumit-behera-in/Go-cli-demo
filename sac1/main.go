package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	greeting(os.Stdout, "sumit")
}

func greeting(out io.Writer, name string) {
	fmt.Fprintf(out, "Hello %v", name)
}

// io.Writer used for specify where to do io as writing
// os.Stdout is the standard output of the os
