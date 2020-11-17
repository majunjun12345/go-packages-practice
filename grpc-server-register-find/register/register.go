package register

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	etcd3 "github.com/coreos/etcd/clientv3"
)

const (
	heart_time   = 10
	grpc_version = "v1.0.0"
)

type EtcdReigistry struct {
	etcd3Client *etcd3.Client
	lease       etcd3.Lease
	kv          etcd3.KV
	key         string
	value       string
	ID          etcd3.LeaseID
	ttl         time.Duration
	ctx         context.Context
	cancel      context.CancelFunc
}

type Option struct {
	EtcdConfig  etcd3.Config
	RegistryDir string
	ServiceName string
	NodeID      string
	NData       NodeData
}

type NodeData struct {
	ServerName string
	Addr       string
	Metadata   map[string]string
}

func InitServiceReg(name, node, regAddr string, addrs []string) error {
	etcdConfg := etcd3.Config{
		Endpoints: addrs,
	}

	registry, err := newRegistry(
		Option{
			EtcdConfig:  etcdConfg,
			RegistryDir: "/grpc-lb",
			ServiceName: name,
			NodeID:      node,
			NData: NodeData{
				ServerName: name,
				Addr:       regAddr,
				Metadata:   make(map[string]string),
			},
		})
	if err != nil {
		return err
	}
	go func() {
		err = registry.Register()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return err
}

func newRegistry(option Option) (*EtcdReigistry, error) {
	client, err := etcd3.New(option.EtcdConfig)
	if err != nil {
		return nil, err
	}
	lease := etcd3.NewLease(client)
	kv := etcd3.NewKV(client)
	val, err := json.Marshal(option.NData)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	registry := &EtcdReigistry{
		lease:       lease,
		kv:          kv,
		etcd3Client: client,
		key:         option.RegistryDir + "/" + option.ServiceName + "/" + grpc_version + "/" + option.NodeID,
		value:       string(val),
		ttl:         heart_time,
		ctx:         ctx,
		cancel:      cancel,
	}
	return registry, nil
}

func (e *EtcdReigistry) Register() error {
	insertFunc := func() error {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				buf := make([]byte, 4096)
				n := runtime.Stack(buf[:], false)
				fmt.Println(string(buf[:n]))
			}
		}()
		if e.ID == 0 {
			if resp, err := e.lease.Grant(context.TODO(), int64(heart_time)); err != nil {
				return err
			} else {
				e.ID = resp.ID
				if _, err := e.kv.Put(context.TODO(), e.key, e.value, etcd3.WithLease(e.ID)); err != nil {
					fmt.Println(err)
				}
			}
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if _, err := e.lease.KeepAliveOnce(ctx, e.ID); err != nil {
			fmt.Println(err)
			e.ID = 0
		} else {
		}
		return nil
	}
	err := insertFunc()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(e.key)
	ticker := time.NewTicker((e.ttl/2 - 1) * time.Second)
	for {
		select {
		case <-ticker.C:
			err = insertFunc()
			if err != nil {
				fmt.Println(err)
			}
		case <-e.ctx.Done():
			ticker.Stop()
			if _, err := e.etcd3Client.Delete(context.Background(), e.key); err != nil {
				fmt.Println(err)
			}
			return nil
		}
	}

	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//	defer cancel()
	//	leaseGrantResp, err := e.lease.Grant(ctx, int64(e.ttl))
	//	if err != nil {
	//		log.ErrorWithFields("Grant error", log.Fields{"error": err})
	//		return err
	//	}
	//
	//	leaseID := leaseGrantResp.ID
	//	kv := etcd3.KV(e.etcd3Client)
	//	_, err = kv.Put(ctx, e.key, e.value, etcd3.WithLease(leaseID))
	//	if err != nil {
	//		log.ErrorWithFields("Put error", log.Fields{"error": err})
	//		return err
	//	}
	//
	//	ctx, cancelFunc := context.WithCancel(context.TODO())
	//	defer cancelFunc()
	//	defer e.lease.Revoke(context.TODO(), leaseID)
	//
	//	keepAliveResp, err := e.lease.KeepAlive(ctx, leaseID)
	//	if err != nil {
	//		log.ErrorWithFields("KeepAlive error", log.Fields{"error": err})
	//		return err
	//	}
	//
	//	for {
	//		select {
	//		case resp := <-keepAliveResp:
	//			if resp == nil {
	//				log.Warn("续约失败")
	//				goto END
	//			}
	//		case <-e.ctx.Done():
	//			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//			defer cancel()
	//			_, err = kv.Delete(ctx, e.key)
	//			if err != nil {
	//				log.ErrorWithFields("Delete error", log.Fields{"error": err})
	//				goto END
	//			}
	//		}
	//	}
	//END:
	//	return nil
}

func (e *EtcdReigistry) DeRegister() error {
	e.cancel()
	return nil
}
