package grpc_handler

// import (
// 	"context"
// 	"fmt"

// 	grpc_pb "github.com/PaoloEG/terrasense/internal/infrastructure/handlers/grpc/autogen"
// )

// type Server struct {
// 	grpc_pb.UnimplementedTerraSenseServiceServer
// 	productSrv *services.ProductService
// }

// func NewServer(productService *services.ProductService) *Server {
// 	return &Server{productSrv: productService}
// }

// func (s *Server) GetProduct(ctx context.Context, params *pb.ProductSearch) (*pb.Product, error) {
// 	fmt.Println("Entering the GetProduct RPC")
// 	product, err := s.productSrv.GetProductInfo(params.Market, params.Id)
// 	if err != nil {
// 		fmt.Println("Exists the GetProduct RPC with error: ", err.Error())
// 		return &pb.Product{}, err
// 	}
// 	pbProduct := pb_output_mappers.PbProductMapper(product)
// 	return pbProduct, nil
// }