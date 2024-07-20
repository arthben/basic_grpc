package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(&lumberjack.Logger{
		Filename:   "log/server.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	})
	log.Logger = log.With().Str("service", "grpc-server").Logger()
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.With().Timestamp().Logger()
}

func InterceptorLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	// write log before request forwarded through handler
	startTime := time.Now()
	// forward request through handler
	resp, err := handler(ctx, req)
	endTime := time.Since(startTime)

	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}

	// if status is error, log will be write in error level
	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}

	logger.Str("protocol", "grpc").
		Str("method", info.FullMethod).
		Int("status", int(statusCode)).
		Str("status-text", statusCode.String()).
		Dur("latency", endTime).
		Interface("request", &req).
		Interface("response", &resp).
		Msg(fmt.Sprintf("%v", resp))

	return resp, err
}
