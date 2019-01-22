package service

import "net/rpc"

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
