package main

import "fmt"

func get11()  int{
	defer func() {
		fmt.Println(2222)
	}()
	return 222
}

func main()  {
	get11()
}
