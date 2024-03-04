package grpc

import (
	"io"
	"jwt-service/internal/config"
	"jwt-service/pkg/jwtGRPC"

	"log"
	"net"

	jwtService "jwt-service/internal/service/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type ServerHandlers interface {
	Start()
	io.Closer
}

type gRPCServer struct {
	cfg        config.GRPC
	grpcServer *grpc.Server
}

func New(cfg config.GRPC, service *jwtService.JWTService) ServerHandlers {
	options := buildOptions(cfg)
	server := grpc.NewServer(options...)

	jwtGRPC.RegisterJWTServiceServer(server, service)
	reflection.Register(server)
	return &gRPCServer{
		cfg:        cfg,
		grpcServer: server,
	}
}

func (g gRPCServer) Start() {
	grpcListener, err := net.Listen("tcp", ":"+g.cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] start grpc server success")
	if err := g.grpcServer.Serve(grpcListener); err != nil {
		log.Fatal(err)
	}
}

func buildOptions(cfg config.GRPC) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters(cfg.Keepalive.KeepaliveParams.ServerParameters)),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy(cfg.Keepalive.EnforcementPolicy)),
	}
}

func (g gRPCServer) Close() error {
	g.grpcServer.GracefulStop()
	return nil
}
