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

type DDBService struct {
	//Embed the unimplemented server and opt-out of forward compatibility
	pbv1.UnimplementedDuckdbServiceServer
}

func DDBSrv() *DDBService {
	return &DDBService{}
}

func (g *DDBService) ReGenDB(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	var h emptypb.Empty

	conn, err := ddb.Conn()
	if err != nil {
		log.Error().Msgf("DuckDBService - ReGenDB Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed opening database")
	}
	defer conn.Close()

	m := models.DDB{
		DB: conn,
	}

	err = m.ReGenDB()
	if err != nil {
		log.Error().Msgf("DuckDBService - ReGenDB Failed: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "failed retriving data from database")
	}

	log.Info().Msg("DuckDBService - ReGenDB: success generating DuckDB table")
	return &h, nil
}
