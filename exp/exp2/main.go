package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"time"
)

type Slo struct {
	Title string `json:"title"`
	Cont string `json:"cont"`
}

var sloMap =[]Slo{}

var visited = map[string]bool{}

func main() {
	c:=colly.NewCollector(colly.AllowedDomains("www.jbr.net.cn"),
		colly.MaxDepth(1),
	)

	detailRegex,_:=regexp.Compile(`company/info_\d+.html$`)

	listRegex,_:=regexp.Compile(`/news/company\.html$`)

	c.OnHTML("div.newIn", func(e *colly.HTMLElement) {
		fmt.Println("文本节点",e.Name)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link:=e.Attr("href")
		//已访问过得列表和详情页面跳过
		if visited[link] && (detailRegex.Match([]byte(link)) || listRegex.Match([]byte(link))){
			return
		}
		//匹配列表和详情才去访问
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			//println("no match!",link)
			return
		}
		time.Sleep(time.Second)
		//println("match",link)
		visited[link] = true
		c.Visit(e.Request.AbsoluteURL(link))
	})

	err:=c.Visit("http://www.jbr.net.cn/")
	if err!=nil {
		fmt.Println(err)
	}
}
