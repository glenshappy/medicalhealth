package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func incr()  {
	client:=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		DB:0,
		Password:"",
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	var lockkey,counterkey string
	lockkey="lock_key"
	counterkey="counter_key"
	lockrs:=client.SetNX(lockkey,1,5*time.Second)
	lockrs1,err:=lockrs.Result()
	if err!=nil || !lockrs1 {
		fmt.Println("lock failed!")
		return
	}
	//加锁成功
	cntResp:=client.Get(counterkey)
	cntvalue,err:=cntResp.Int64()
	if err==nil {
		cntvalue++
		resp:=client.Set(counterkey,cntvalue,0)
		_,err=resp.Result()
		if err!=nil {
			println("set value error")
		}
		fmt.Printf("current counter is:%d\n",cntvalue)
	}

	delResp:=client.Del(lockkey)
	unlocksuccess,err:=delResp.Result()
	if unlocksuccess>0 || err==nil {
		println("unlock success")
	}else{
		println("unlock failed")
	}
}

func main() {
	var wg sync.WaitGroup
	for i:=0;i<100;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
