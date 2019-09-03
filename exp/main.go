package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func findLinks(url string) ([]string, error) {
	var links []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%s", url, err)
	}
	links = visit(nil, doc)
	return links, nil
}

func waitForServer(url string) error {
	const timeout = 1 * time.Second
	deadline := time.Now().Add(timeout)
	fmt.Println(deadline)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		interval := time.Second << uint(tries)
		fmt.Println("睡眠%s秒", interval)
		time.Sleep(interval)
	}
	return fmt.Errorf("response errror:%s after %s minutes", url, timeout)
}

func NetWorkStatus() error {
	cmd := exec.Command("ping", "baidu.com")
	fmt.Println("NetWorkStatus Start:", time.Now().Unix())
	err := cmd.Run()
	fmt.Println("NetWorkStatus End  :", time.Now().Unix())
	if err != nil {
		log.Printf("network error:%v", err)
		return err
	}
	return nil
}

func createTempDir(prefix string) (string, error) {
	dir, err := ioutil.TempDir("", prefix)
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir:%v", err)
	}
	return dir, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if r == '\r' {
			fmt.Println("回车符")
		}
		if r == '\n' {
			fmt.Println("换行符")
		}
		fmt.Println(string(r))
	}
	//log.SetPrefix("wait:")
	//log.SetFlags(0)
	//
	////第四种策略
	//if err:=NetWorkStatus();err!=nil {
	//	log.Printf("ping failed:%v",err)
	//}
	//
	//dir,err:=createTempDir("scratch")
	//if err!=nil {
	//	log.Fatal(err)
	//}
	////println(dir)
	//os.RemoveAll(dir)
	//os.RemoveAll("D:/test11")
	//
	//
	////fmt.Println(ht1.UnescapeString("&lt;&gt;&#39;"))
	////log
	//fmt.Println(time.Second<<2)
	//err=waitForServer("http://www.baidu.com")
	//if err!=nil{
	//	log.Fatalf("site is error:%s",err)
	//}
	//urls:=os.Args[1:]
	//fmt.Println(urls)
	//for _,link:=range urls {
	//	links,err:=findLinks(link)
	//	if err!=nil{
	//		fmt.Fprintf(os.Stderr,"%s:%v",link,err)
	//	//doc, err := html.Parse(os.Stdin)
	//	//if err != nil {
	//	//	fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	//	//	os.Exit(1)
	//	//}
	//	//outline(nil, doc)
	//	//var links[]string
	//	//for _,v:=range visit(links,doc){
	//	//	fmt.Println(v)
	//	}
	//	for _,v:=range links {
	//		fmt.Println(v)
	//	}
	//}

}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			fmt.Println(a)
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
