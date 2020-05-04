package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	pb "github.com/verbruggenjesse/grpc-consumer/gen"
	"github.com/verbruggenjesse/grpc-consumer/infrastructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	var logger *infrastructure.Logger
	var err error

	configLoader := &infrastructure.AppConfigLoader{
		Logger: infrastructure.DefaultLogger(),
	}

	config := configLoader.LoadFromEnv()

	if logLevel, err := strconv.Atoi(config["LOG_LEVEL"]); err != nil {
		logger = infrastructure.DefaultLogger()
	} else {
		logger = infrastructure.NewLogger(logLevel)
	}

	var subscriber *infrastructure.RedisSubscriber
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	redisOptions := &redis.Options{
		Addr:     config["REDIS_ADDR"],
		Password: config["REDIS_PASSWORD"],
		DB:       0,
	}

	if config["REDIS_TLS"] == "enabled" {
		redisOptions.TLSConfig = &tls.Config{}
	}

	if subscriber, err = infrastructure.NewRedisSubscriber(redisOptions); err != nil {
		logger.Fatal(fmt.Sprintf("could not initialize publisher: %s", err.Error()), 1)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))

	if err != nil {
		logger.Fatal(err.Error(), 1)
	}

	var opts []grpc.ServerOption

	opts = []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
	}

	server := grpc.NewServer(opts...)

	consumer := infrastructure.NewEventConsumerServer(subscriber, logger)

	pb.RegisterConsumerServer(server, consumer)

	logger.Info(fmt.Sprintf("grpc-consumer started, listening for incoming subscriptions at %s", lis.Addr().String()))
	if err := server.Serve(lis); err != nil {
		logger.Fatal(err.Error(), 1)
	}
}
