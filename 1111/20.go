package  main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
    r io.Reader
}


func main() {
    s := strings.NewReader("LLLLL");
	r:= rot13Reader{s}
	io.Copy(os.Stdout, &r)
}