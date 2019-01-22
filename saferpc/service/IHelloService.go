package service

import "net/rpc"

const HelloServiceName = "saferpc/service.HelloService"

type IHelloService interface {
	Hello(request string, reply *string) error
}

var _ IHelloService = (*HelloServiceClient)(nil)
var _ IHelloService = (*HelloService)(nil)

func RegisterHelloService(svc IHelloService) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{c}, nil
}
