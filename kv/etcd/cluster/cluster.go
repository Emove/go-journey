package cluster

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitClusterClient(addr []string) (*clientv3.Client, error) {
	//return clientv3.NewFromURLs(addr)
	return clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 3 * time.Second,
	})
}
