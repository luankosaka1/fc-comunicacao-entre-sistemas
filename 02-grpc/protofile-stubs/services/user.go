package services

import (
	"context"
	"fmt"

	"github.com/luankosaka1/golang-protofile-stubs/pb"
)

// Garantir que o servico exista e nao comprometa o protobuffer
type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
