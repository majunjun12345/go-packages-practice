package pool

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testGoScripts/grpc-server-register-find/resolver"
	"testGoScripts/grpc-server-register-find/tool/tracer"
	"time"

	"git.internal.yunify.com/manage/common/etcd/balancer"
	"github.com/opentracing/opentracing-go"
	"github.com/xiaomeng79/go-log"
	etcd "go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	GRPC_MAX_CONNECT  = 1
	GRPC_IDLE_TIMEOUT = 30 * time.Second
	GRPC_VERSION      = "v1.0.0"
)

var pool *ClientPool
var once sync.Once

//buid client
func Init(addr ...string) {
	once.Do(func() {
		pool = NewClientPool(addr)
	})
}
func GetConPool() *ClientPool {
	return pool
}

type ClientPool struct {
	Addr    []string
	Lock    sync.RWMutex
	Clients map[string]*grpc.ClientConn
}

//create client pool
func NewClientPool(address []string) *ClientPool {
	c := &ClientPool{}
	c.Addr = address
	c.Clients = make(map[string]*grpc.ClientConn)
	return c
}

func withOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(balancer.RoundRobin),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 10, // 如果这段时间内 client 没有发送消息，那么将发送 ping 包，最小值为 10s
			Timeout:             time.Second * 20, // 如果ping ack 1s之内未返回则认为连接已断开
			PermitWithoutStream: true,             // 果没有active的stream,是否允许发送ping
		}),
		grpc.WithUnaryInterceptor(
			tracer.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
		),
	}
}

func (c *ClientPool) FindServer(name string) error {
	etcdConfg := etcd.Config{
		Endpoints: c.Addr,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resolver.RegisterResolver("etcd", etcdConfg, name, GRPC_VERSION)
	conn, err := grpc.DialContext(ctx, "etcd:///", grpc.WithInsecure(),
		grpc.WithInsecure(),
		grpc.WithBalancerName(balancer.RoundRobin),
		grpc.WithUnaryInterceptor(
			tracer.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
		))

	if err != nil {
		log.Warnf("grpc dial service(%s) error:%s", name, err.Error())
		return err
	}
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.Clients[name] = conn
	return nil
}

func (c *ClientPool) OmitServer(name string) {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	if v, ok := c.Clients[name]; ok {
		v.Close()
	}
}
func (c *ClientPool) GetClient(name string) (interface{}, error) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	if v, ok := c.Clients[name]; ok {
		//con, err := v.Get()
		//if err != nil {
		//	return nil, err
		//}
		////log.Infof("GetClient con:%+v,num:%d", con, v.len())
		//defer v.Put(con)
		return v, nil
	} else {
		return nil, errors.New(fmt.Sprintf("service %s is not existed", name))
	}
}
