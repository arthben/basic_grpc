package main

import (
	"fmt"
	"net"
	"os"

	"github.com/arthben/basic_grpc/grpc_server/api/handlers"
	"github.com/arthben/basic_grpc/grpc_server/internal/logging"
	"github.com/arthben/basic_grpc/grpc_server/pkg/contract"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	// get listener port
	port := os.Getenv("BASIC_GRPC_PORT")
	if len(port) == 0 {
		fmt.Println("Env BASIC_GRPC_PORT not found")
		os.Exit(1)
	}

	// init logger
	logging.NewLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.Info().Msg("Start grpc-server on " + listener.Addr().String())
	fmt.Printf("GRPC Server Listen on %s\n", listener.Addr().String())

	if err != nil {
		log.Error().Err(err).Msg("Failed start grpc-server")
		panic(err)
	}

	// middleware to log request & response
	logger := grpc.UnaryInterceptor(logging.InterceptorLogger)

	// service
	serviceHandler := handlers.NewAPIHandlers()

	// spawn server listener
	server := grpc.NewServer(logger)
	contract.RegisterThingsTodoServer(server, serviceHandler)
	if err := server.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Failed start grpc-server")
		panic(err)
	}
}
