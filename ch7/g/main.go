package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	fmt.Println(os.Args[1:])
	str:=`<?xml version="1.0" encoding="UTF-8"?><note><to>Tove</to><from>Jani</from><heading>Reminder</heading><body>Don't forget me this weekend!</body></note>`
	buf:=bytes.NewBuffer([]byte(str))
	dec:=xml.NewDecoder(buf)
//	dec:=xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok,err:=dec.Token()
		if err==io.EOF {
			break
		}else if(err!=nil){
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x,y[]string)  bool{
	for len(y)<=len(x){
		if(len(y)==0){
			return true
		}
		if(x[0]==y[0]){
			y=y[1:]
		}
		x=x[1:]
	}
	return  false
}
