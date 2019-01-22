package service

import "example-rpc/protorpc"

type HelloService struct {
}

func (p *HelloService) Hello(request pb.String, reply *pb.String) error {
	reply.Value = "Hello: " + request.GetValue()
	return nil
}
