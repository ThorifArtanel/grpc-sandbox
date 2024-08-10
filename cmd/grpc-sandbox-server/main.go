package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"grpc-sandbox/internal/hello"

	hellopb_v1 "grpc-sandbox/gen/proto/hello/v1"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	flDev  = flag.Bool("dev", false, "development mode")
	flPort = "8080"
)

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Setup logger
	// if *flDev {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
	// 	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// } else {
	// 	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	// }

	// Env Lookup
	port := flPort
	envPort, present := os.LookupEnv("PORT")
	if present {
		port = envPort
	}

	wg.Add(1)
	// GRPC Server
	go func() {
		lis, err := net.Listen("tcp", ":"+port)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create listener")
		}

		// Create gRPC server with zerolog, and panic recovery middleware
		srv := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				recovery.UnaryServerInterceptor(),
			),
			grpc.ChainStreamInterceptor(
				recovery.StreamServerInterceptor(),
			),
		)

		// // Register your services
		greeter := hello.NewGreeter()
		hellopb_v1.RegisterHelloServiceServer(srv, greeter)

		reflection.Register(srv)

		// Health and reflection service
		healthServer := health.NewServer()
		healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
		healthServer.SetServingStatus(hellopb_v1.HelloService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)
		grpc_health_v1.RegisterHealthServer(srv, healthServer)

		wg.Done()
		log.Info().Msgf("gRPC server listening on :%s", port)
		if err := srv.Serve(lis); err != nil {
			log.Fatal().Err(err).Msg("failed to start gRPC server")
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Handle cleanup here if any

	log.Info().Msg("Server exiting")
}
