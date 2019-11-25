package main

import (
	"github.com/gocolly/colly"
	"github.com/nats-io/nats.go"
	"net/url"
	"os"
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
	c.OnHTML(".post-topheader__info--title", func(e *colly.HTMLElement) {
		println(e.Text)
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

func startConsumer()  {
	nc,err:=nats.Connect(nats.DefaultURL)
	if err!=nil {
		panic(err)
	}
	sub,err:=nc.QueueSubscribeSync("tasks","works")

	if err!=nil {
		//log error
		return
	}

	for {
		msg,err:=sub.NextMsg(time.Hour*100000)
		if err!=nil {
			break
		}
		//开始访问这一条消息
		urlStr:=string(msg.Data)
		println(urlStr)
		ins:=factory(urlStr)
		ins.Visit(urlStr)
	}
}

func main() {
	startConsumer()
}
