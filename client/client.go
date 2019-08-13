package main

import (
	"log"
	"time"

	pb "github.com/micro_service/ProductService"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5230"
)

func main() {
	// 建立一个与服务端的连接.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	response, err := client.AddProduct(ctx, &pb.AddProductRequest{ProductName: "phone"})
	if nil != err {
		log.Fatalf("add product failed, %v", err)
	}
	log.Printf("add product success,%s", response)

	productId := response.ProductId
	queryResp, err := client.QueryProductInfo(ctx, &pb.QueryProductRequest{ProductId: productId})
	if nil != err {
		log.Fatalf("query product info failed,%v", err)
	}
	log.Printf("Product info is %v", queryResp)

	defer cancel()
}
