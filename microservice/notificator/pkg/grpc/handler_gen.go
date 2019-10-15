// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "github.com/archacon/chacoworks/microservice/notificator/pkg/endpoint"
	pb "github.com/archacon/chacoworks/microservice/notificator/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	sendEmail grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.NotificatorServer {
	return &grpcServer{sendEmail: makeSendEmailHandler(endpoints, options["SendEmail"])}
}
