package service

import (
	"context"
	"fmt"
	"net"

	api "alior-digital/api/proto/grpc"
	"alior-digital/src/config"
	"alior-digital/src/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	TerminateCode = 1
)

type Database interface {
	Insert(ctx context.Context, user *types.User) (userID int32, err error)
	CheckPassword(ctx context.Context, email, password string) (int32, error)
	GetUserByID(ctx context.Context, id int32) (*types.User, error)
}

func NewService(connectionPool Database) *Service { return &Service{DB: connectionPool} }

type Service struct {
	DB Database
	api.UnimplementedAuthServer
}

func (s *Service) Start(grpcEndpoint *config.Endpoint) (chan int, error) {
	var exitChan chan int
	server := grpc.NewServer()
	api.RegisterAuthServer(server, s)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", grpcEndpoint.ListenAddr, grpcEndpoint.Port))
	if err != nil {
		return nil, err
	}

	go func() {
		for code := range exitChan {
			if code == TerminateCode {
				server.Stop()
				return
			}
		}
	}()
	go func() {
		err = server.Serve(lis)
		if err != nil {
			return
		}
	}()
	return exitChan, nil

}

func (s *Service) CheckUserPassword(ctx context.Context, in *api.CheckRequest) (*api.CheckResponse, error) {
	uid, err := s.DB.CheckPassword(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("rpc error: %s", err))
	}
	return &api.CheckResponse{
		ID: uid,
	}, status.Error(codes.OK, "OK")
}

func (s *Service) castProtoUserToTypesUser(proto *api.User) *types.User {
	return &types.User{
		ID:          proto.UID,
		FullName:    proto.FullName,
		PhoneNumber: proto.PhoneNumber,
		Email:       proto.Email,
	}
}
func (s *Service) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	uid, err := s.DB.Insert(ctx, s.castProtoUserToTypesUser(in.User))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("rpc error: %s", err))
	}
	return &api.CreateUserResponse{
		UID: uid,
	}, status.Error(codes.OK, "OK")
}

func (s *Service) GetUserByID(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	user, err := s.DB.GetUserByID(ctx, in.UID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("rpc error: %s", err))
	}
	return &api.GetUserResponse{
		User: &api.User{
			Email:    user.Email,
			Password: user.Password,
			UID:      in.UID,
			FullName: user.FullName,
		},
	}, status.Error(codes.OK, "OK")
}
