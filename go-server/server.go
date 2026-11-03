
package main

import (
	"context"
	"log"
	"net"

	pb "grpc-polyglot-lab/go-server/pb"

	"google.golang.org/grpc"
)

// UserServer menangani permintaan profil pengguna.
type userServer struct {
	pb.UnimplementedUserServiceServer
}

// GetUserProfile mengirimkan data profil berdasarkan ID pengguna.
func (s *userServer) GetUserProfile(
	ctx context.Context,
	req *pb.UserRequest,
) (*pb.UserResponse, error) {

	if req.GetUserId() == 1 {
		return &pb.UserResponse{
			Name:  "Elisabet Kunmop",
			Email: "elisabet@example.com",
			Phone: "000000000000",
		}, nil
	}

	return &pb.UserResponse{
		Name:  "User Tidak Ditemukan",
		Email: "-",
		Phone: "-",
	}, nil
}

func main() {
	// Membuka port untuk komunikasi gRPC.
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	// Membuat server gRPC.
	server := grpc.NewServer()

	// Mendaftarkan layanan profil pengguna.
	pb.RegisterUserServiceServer(server, &userServer{})

	log.Println("gRPC Server berjalan pada port 50051")
	log.Println("Nama pengguna: Elisabet Kunmop")

	// Menjalankan server hingga dihentikan.
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server gagal berjalan: %v", err)
	}
}
