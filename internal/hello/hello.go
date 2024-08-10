package hello

import (
	"context"

	hellopb_v1 "github.com/ThorifArtanel/grpc-sandbox/gen/proto/hello/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

//Implement the generated HelloServiceServer gRPC interface

type Greeter struct {
	//Embed the unimplemented server and opt-out of forward compatibility
	hellopb_v1.UnimplementedHelloServiceServer
}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Hello(ctx context.Context, empty *emptypb.Empty) (*hellopb_v1.HelloResponse, error) {
	h := hellopb_v1.HelloResponse{Hello: "Hello world!"}
	log.Info().Msg("replying to the greeting")
	return &h, nil
}
