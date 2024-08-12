package service

import (
	"context"

	pbv1 "github.com/ThorifArtanel/grpc-sandbox/gen/proto/v1"
	"github.com/ThorifArtanel/grpc-sandbox/internal/app/ddb"
	"github.com/ThorifArtanel/grpc-sandbox/internal/app/models"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	//Embed the unimplemented server and opt-out of forward compatibility
	pbv1.UnimplementedUserServiceServer
}

func UserSrv() *UserService {
	return &UserService{}
}

func (g *UserService) All(ctx context.Context, empty *emptypb.Empty) (*pbv1.UserGetResponse, error) {
	var h pbv1.UserGetResponse

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("UserService - User All Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.Users{
		DB: conn,
	}

	err = m.All(&h)
	if err != nil {
		log.Error().Msgf("UserService - User All Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed retriving data from database")
	}

	log.Info().Msg("UserService - User All: success processing user all")
	return &h, nil
}

func (g *UserService) One(ctx context.Context, req *pbv1.UserOneRequest) (*pbv1.UserOneResponse, error) {
	var h pbv1.UserOneResponse
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Id")
	}

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("UserService - User One Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.Users{
		DB: conn,
	}

	err = m.One(req.Id, &h)
	if err != nil {
		log.Error().Msgf("UserService - User One Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed retriving data from database")
	} else if h.GetUser() == nil {
		log.Info().Msgf("UserService - User One: not found")
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	log.Info().Msg("UserService - User One: success processing user one")
	return &h, nil
}

func (g *UserService) Insert(ctx context.Context, req *pbv1.UserInsertRequest) (*emptypb.Empty, error) {
	if req.GetUser().GetFirstname() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Username")
	}
	if req.GetUser().GetLastname() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Password")
	}

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("UserService - User Insert Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.Users{
		DB: conn,
	}

	err = m.Insert(req)
	if err != nil {
		log.Error().Msgf("UserService - User Insert Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed inserting data to database")
	}

	log.Info().Msg("UserService - User Insert: success user insert")
	return nil, nil
}

func (g *UserService) Update(ctx context.Context, req *pbv1.UserUpdateRequest) (*emptypb.Empty, error) {
	if req.User.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Id")
	}
	if req.GetUser().GetFirstname() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Username")
	}
	if req.GetUser().GetLastname() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Password")
	}

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("UserService - User Update Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.Users{
		DB: conn,
	}

	err = m.Update(req.User.Id, req)
	if err != nil {
		log.Error().Msgf("UserService - User Update Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed updating data to database")
	}

	log.Info().Msg("UserService - User Update: success user updating")
	return nil, nil
}

func (g *UserService) Delete(ctx context.Context, req *pbv1.UserDeleteRequest) (*emptypb.Empty, error) {
	if req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Id")
	}

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("UserService - User Delete Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.Users{
		DB: conn,
	}

	err = m.Delete(req.GetId())
	if err != nil {
		log.Error().Msgf("UserService - User Delete Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed deleting data from database")
	}

	log.Info().Msg("UserService - User Delete: success user delete")
	return nil, nil
}
