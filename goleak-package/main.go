package main

import (
	"fmt"
	"time"

	"git.internal.yunify.com/manage/common/etcd"
)

func main() {
	etcdMgr, err := etcd.GetEtcdMgr([]string{"127.0.0.1:2379"}, 3)
	if err != nil {
		panic(err)
	}
	lock := etcdMgr.CreatLock(fmt.Sprintf("/test/redis-ticker/%s", time.Now().Local().Format("2006-01-02:15")))
	defer etcdMgr.Client.Close()
	defer lock.Unlock()
	defer lock.Lease.Close()

	if err := lock.TryLock(int64(60)); err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("===========")
	}
}
