package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nats-io/nats.go"
	"net/url"
	"os"
	"regexp"
	"time"
)

var nc *nats.Conn

var domain2Collector = map[string]*colly.Collector{}

func initCollector1()  *colly.Collector{
	c:=colly.NewCollector(
		colly.AllowedDomains("www.jbr.net.cn"),
		colly.MaxDepth(1),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {


	})
	return c
}

func initCollector2()  *colly.Collector{
	c:=colly.NewCollector(
		colly.AllowedDomains("segmentfault.com"),
		colly.MaxDepth(1),
	)
	c.OnResponse(func(resp *colly.Response) {
		//做一些爬完之后的善后工作
		//比如页面已爬完的确认存进MYSQL
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//基本的反爬虫策略
		time.Sleep(time.Second*2)

		//TODO ，正则match列表页的话，就visit
		//TODO，正则match落地页的话，就发消息队列
		regList,err:=regexp.Compile(`channel/\w+$`)
		if err!=nil {
			panic(err)
		}
		regDetail,err:=regexp.Compile(`a/\d+$`)
		if err!=nil {
			panic(err)
		}
		link:=e.Attr("href")
		if regList.Match([]byte(link)){
			fmt.Println("list:",link)
			c.Visit(e.Request.AbsoluteURL(link))
		}else if(regDetail.Match([]byte(link))){
			fmt.Println("detail:",link)
			nc.Publish("tasks",[]byte(e.Request.AbsoluteURL(link)))
		}else{
			fmt.Println("other:",link)
		}
	})
	return c
}

func init()  {
	domain2Collector["www.jbr.net.cn"] = initCollector1()
	domain2Collector["segmentfault.com"] = initCollector2()

	var err error
	nc,err=nats.Connect(nats.DefaultURL)
	if err!=nil {
		os.Exit(1)
	}
}

func factory(urlStr string)  *colly.Collector{
	u,_:=url.Parse(urlStr)
	return domain2Collector[u.Host]
}

func main() {
	urls:=[]string{"http://www.jbr.net.cn","https://segmentfault.com"}
	for _,url:=range urls {
		instance:=factory(url)
		instance.Visit(url)
	}
}
