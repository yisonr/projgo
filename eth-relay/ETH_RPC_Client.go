package main

import (
	"fmt"
	"net/rpc"
)

type ETHRPClient struct {
	NodeUrl string
	client  *rpc.Client // rcp 客户端句柄实例
}

func NewETHRPCClient(nodeUrl string) *ETHRPClient {
	client := &ETHRPClient{
		NodeUrl: nodeUrl,
	}
	client.initRpc()
	return client
}

func (erc *ETHRPClient) initRpc() {
	rpcClient, err := rpc.DialHTTP("tcp", erc.NodeUrl)
	if err != nil {
		errInfo := fmt.Errorf("init rpc client err: %v", err).Error()
		panic(errInfo)
	}

	erc.client = rpcClient
}

func (erc *ETHRPClient) GetRpc() *rpc.Client {
	if erc.client == nil {
		erc.initRpc()
	}
	return erc.client
}
