package models

import (
	"fmt"

	"github.com/mooncake9527/orange-core/config"
	"google.golang.org/grpc"
)

type ServiceNode struct {
	config.RegisterNode                  //注册节点
	Weight              int              //权重
	failCnt             int              //失败次数
	enable              bool             //是否启用
	grpc                *grpc.ClientConn //grpc连接
}

func (n *ServiceNode) Enable() bool {
	return n.enable
}

func (n *ServiceNode) SetEnable(enable bool) {
	n.enable = enable
}

func (n *ServiceNode) ClearFailCnt() {
	n.failCnt = 0
}

func (n *ServiceNode) IncrFailCnt() {
	n.failCnt++
	if n.failCnt > n.FailLimit {
		n.enable = false
	}
}

func (n *ServiceNode) GetFailCnt() int {
	return n.failCnt
}

func (n *ServiceNode) GetUrl() string {
	return fmt.Sprintf("%s://%s:%d", n.Protocol, n.Addr, n.Port)
}

// func (n *ServiceNode) GetHttpClient() (*http.Client, error) {
// 	//http.NewClient(fmt.Sprintf("%s://%s:%d", n.Protocol, n.Addr, n.Port))
// 	return nil, nil
// }

func (n *ServiceNode) GetGrpcConn() (conn *grpc.ClientConn, err error) {
	if n.grpc != nil {
		conn = n.grpc
	} else {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%d", n.Addr, n.Port), grpc.WithInsecure())
		if err == nil {
			n.grpc = conn
		}
	}
	return
}

// func (n *ServiceNode) GetRpcConn() (*rpc.Client, error) {
// 	if n.Protocol == "tcp" {
// 		return rpc.Dial(n.Protocol, fmt.Sprintf("%s:%d", n.Addr, n.Port))
// 	} else {
// 		return rpc.DialHTTP(n.Protocol, fmt.Sprintf("%s:%d", n.Addr, n.Port))
// 	}
// }

func (n *ServiceNode) Close() {
	n.enable = false
	if n.grpc != nil {
		n.grpc.Close()
	}
}
