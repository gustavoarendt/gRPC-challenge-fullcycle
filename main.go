package main

import (
	"fc01-grpc-product/application/grpc"
	"fc01-grpc-product/infrastructure/db"
	"os"
)

func main() {
	database := db.ConnectDb(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
