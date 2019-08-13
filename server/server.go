package main

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/micro_service/ProductService"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5230"
)

var dataBase = make(map[string]*Product, 10)

//Product is a xxx
type Product struct {
	ProductName    string
	ProductId      string
	ManufacturerId string
	Weight         float64
	ProductionDate int64
	ImportDate     int64
}

type server struct{}

func (s *server) AddProduct(ctx context.Context, request *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	log.Printf("get request from client to add product,request is %s", request)
	productId := strconv.FormatInt(rand.Int63(), 10)
	product := new(Product)
	product.ProductName = request.ProductName
	product.ProductId = productId
	product.ManufacturerId = request.ManufacturerId
	product.Weight = request.Weight
	product.ProductionDate = request.ProductionDate
	product.ImportDate = time.Now().UnixNano()
	dataBase[productId] = product
	return &pb.AddProductResponse{ProductId: productId, Message: "Add product success"}, nil
}

func (s *server) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*pb.EmptyResponse, error) {
	log.Printf("get request from client to add product,request is %s", request)
	productId := request.ProductId
	delete(dataBase, productId)
	return nil, nil
}

func (s *server) QueryProductInfo(ctx context.Context, request *pb.QueryProductRequest) (*pb.ProductInfoResponse, error) {
	log.Printf("get request from client fro query product info,%v", request)
	productId := request.ProductId
	product := dataBase[productId]
	response := new(pb.ProductInfoResponse)
	response.ProductName = product.ProductName
	response.ProductId = product.ProductId
	response.ManufacturerId = product.ManufacturerId
	response.Weight = product.Weight
	response.ProductionDate = product.ProductionDate
	response.ImportDate = product.ImportDate
	return response, nil
}

func (s *server) QueryProductsInfo(ctx context.Context, request *pb.EmptyRequest) (*pb.ProductsInfoResponse, error) {
	// 待实现
	return nil, nil
}

func main() {
	log.Printf("begin to start rpc server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
