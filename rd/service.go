package rd

import (
	"fmt"
	"time"

	"github.com/mooncake9527/npx/config"
	"github.com/mooncake9527/npx/driver/consul"
	"github.com/mooncake9527/npx/driver/etcd"
	"github.com/mooncake9527/npx/models"

	"github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type RDClient interface {
	Register(s *config.RegisterNode) error
	Deregister()
	Watch(s *config.DiscoveryNode) error
	GetService(name string, clientIp string) (*models.ServiceNode, error)
}

func NewRDClient(cfg *config.Config) (client RDClient, err error) {
	if cfg.Driver != "etcd" && cfg.Driver != "consul" {
		err = fmt.Errorf("unsupported driver: %s", cfg.Driver)
		return nil, err
	}
	if cfg.Driver == "etcd" {
		c := clientv3.Config{
			Endpoints:   cfg.Endpoints,
			DialTimeout: cfg.Timeout,
		}
		client, err = etcd.NewClient(&c)
	} else {
		c := api.Config{
			Address:  cfg.Endpoints[0],
			Scheme:   cfg.Scheme,
			WaitTime: cfg.Timeout,
		}
		client, err = consul.NewClient(&c)
	}
	if err != nil {
		return
	}
	for _, rs := range cfg.Registers {
		if rs.Addr == "" || rs.Port <= 0 || rs.Port > 65535 {
			panic("register node addr or port is error")
		}
		if rs.Protocol != "http" && rs.Protocol != "grpc" {
			panic("register node protocol is error")
		}
		if rs.Name == "" {
			panic("register node name is error")
		}
		if rs.Id == "" {
			rs.Id = fmt.Sprintf("%s:%d", rs.Addr, rs.Port)
		}
		if rs.FailLimit <= 0 {
			rs.FailLimit = 3
		}
		if rs.Interval <= 0 {
			rs.Interval = 5 * time.Second
		}
		if rs.Timeout <= 0 {
			rs.Timeout = 10 * time.Second
		}
		err = client.Register(rs)
		if err != nil {
			return
		}
	}
	for _, ds := range cfg.Discoveries {
		if ds.Enable {
			err = client.Watch(ds)
			if err != nil {
				return
			}
		}
	}
	return
}
