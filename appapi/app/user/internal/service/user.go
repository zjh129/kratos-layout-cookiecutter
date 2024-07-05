package service

import (
	"context"
	"kratos-learn/app/user/internal/biz"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "kratos-learn/api/user_service"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.GreeterUsecase
}

func NewUserService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserInfo, error) {
	return &pb.UserInfo{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
