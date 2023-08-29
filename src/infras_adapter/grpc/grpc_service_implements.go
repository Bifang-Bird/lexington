/*
*

	@author: junwang
	@since: 2023/8/25
	@desc: //TODO

*
*/

package grpc

import (
	"codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-proto.git/gen"
	"lexington/src/usecases"

	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 这是一个dome
var _ gen.SmsServiceServer = (*GRPCServerImpl)(nil)

var SmsGRPCServerSet = wire.NewSet(NewGRPCServerImpl)

type GRPCServerImpl struct {
	gen.UnimplementedSmsServiceServer
	uc usecases.UseCase
}

func NewGRPCServerImpl(
	grpcServer *grpc.Server,
	uc usecases.UseCase,
) gen.SmsServiceServer {
	svc := GRPCServerImpl{
		uc: uc,
	}
	gen.RegisterSmsServiceServer(grpcServer, &svc)
	reflection.Register(grpcServer)
	return &svc
}
