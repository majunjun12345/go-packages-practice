package resolver

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/coreos/etcd/mvcc/mvccpb"
	etcd3 "go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
)

const (
	watcher_plugin = "watcher"
)

type NodeData struct {
	ServerName string
	Addr       string
	Metadata   map[string]string
}

type Watcher struct {
	key    string
	client *etcd3.Client
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	lock   sync.Mutex
	addrs  []resolver.Address
}

func (w *Watcher) Close() {
	w.cancel()
}

func newWatcher(key string, cli *etcd3.Client) *Watcher {
	ctx, cancel := context.WithCancel(context.Background())
	w := &Watcher{
		key:    key,
		client: cli,
		ctx:    ctx,
		cancel: cancel,
	}
	return w
}

func (w *Watcher) GetAllAddresses() []resolver.Address {
	ret := []resolver.Address{}

	resp, err := w.client.Get(w.ctx, w.key, etcd3.WithPrefix())
	if err == nil {
		addrs := extractAddrs(resp)
		if len(addrs) > 0 {
			for _, addr := range addrs {
				v := addr
				ret = append(ret, resolver.Address{
					Addr:       v.Addr,
					Metadata:   &v.Metadata,
					ServerName: v.ServerName,
				})
			}
		}
	}
	return ret
}

func (w *Watcher) Watch() chan []resolver.Address {
	out := make(chan []resolver.Address, 3)
	w.wg.Add(1)
	go func() {
		defer func() {
			close(out)
			w.wg.Done()
		}()
		w.lock.Lock()
		w.addrs = w.GetAllAddresses()
		if len(w.addrs) != 0 {
			out <- w.cloneAddresses(w.addrs)
		}
		w.lock.Unlock()

		rch := w.client.Watch(w.ctx, w.key, etcd3.WithPrefix())
		for {
			select {
			case wresp, ok := <-rch:
				if !ok {
					fmt.Printf("ch closed,watcher quit")
					return
				}
				for _, ev := range wresp.Events {
					switch ev.Type {
					case mvccpb.PUT:
						nodeData := NodeData{}
						err := json.Unmarshal([]byte(ev.Kv.Value), &nodeData)
						if err != nil {
							fmt.Printf("Parse node data error:%s", err.Error())
							continue
						}
						addr := resolver.Address{ServerName: nodeData.ServerName, Addr: nodeData.Addr, Metadata: &nodeData.Metadata}
						if w.addAddr(addr) {
							if len(w.addrs) != 0 {
								out <- w.cloneAddresses(w.addrs)
							}
						}
					case mvccpb.DELETE:
						nodeData := NodeData{}
						err := json.Unmarshal(ev.Kv.Value, &nodeData)
						if err != nil {
							grpclog.Error("Parse node data error:", err)
							continue
						}
						addr := resolver.Address{ServerName: nodeData.ServerName, Addr: nodeData.Addr, Metadata: &nodeData.Metadata}
						if w.removeAddr(addr) {
							if len(w.addrs) != 0 {
								out <- w.cloneAddresses(w.addrs)
							}
						}
					}
				}
			case <-w.ctx.Done():
				fmt.Printf("watcher quit ")
				return
			}
		}
	}()
	return out
}

func extractAddrs(resp *etcd3.GetResponse) []NodeData {
	addrs := []NodeData{}

	if resp == nil || resp.Kvs == nil {
		return addrs
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			nodeData := NodeData{}
			err := json.Unmarshal(v, &nodeData)
			if err != nil {
				fmt.Printf("Parse node data error:", err.Error())
				continue
			}
			addrs = append(addrs, nodeData)
		}
	}
	return addrs
}

func (w *Watcher) cloneAddresses(in []resolver.Address) []resolver.Address {
	out := make([]resolver.Address, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}
	return out
}

func (w *Watcher) addAddr(addr resolver.Address) bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	for _, v := range w.addrs {
		if addr.Addr == v.Addr {
			return false
		}
	}
	w.addrs = append(w.addrs, addr)
	return true
}

func (w *Watcher) removeAddr(addr resolver.Address) bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	for i, v := range w.addrs {
		if addr.Addr == v.Addr {
			w.addrs = append(w.addrs[:i], w.addrs[i+1:]...)
			return true
		}
	}
	return false
}
