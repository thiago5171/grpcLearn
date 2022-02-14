package service

import (
	"context"
	"fmt"
	"github.com/thiago5171/grpc_learn/pb"
	"strings"
)

type UserService struct {
	pb.UnimplementedUserServiceServer // uma importação para que caso add um servico que n existe no pb ele irá parar
}

func NewUserService() *UserService {
	return &UserService{}
}
func (u UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	var name string
	if strings.ToUpper(req.GetName()) == "THIAGO" {
		name = fmt.Sprintf("Aluno: %v", req.GetName())
	} else {
		name = fmt.Sprintf("Professor %v", req.GetName())
	}
	fmt.Println(req.Name)
	return &pb.User{
		Id:    req.GetId(),
		Name:  name,
		Email: req.GetEmail(),
	}, nil
}
