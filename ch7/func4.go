package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	fmt.Printf("%#v", out)
	if out != nil {
		fmt.Println("33")
		out.Write([]byte("done!\n"))
	}
}

//func main()  {
//	var p *bytes.Buffer
//	//p = new(bytes.Buffer)
//	fmt.Printf("%#v\n",p)
//	if p!=nil {
//		p.Write([]byte("hello"))
//	}else{
//		fmt.Println(22)
//	}
//}
