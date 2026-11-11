
import os
import grpc
import user_pb2
import user_pb2_grpc


def main():
    server_address = os.getenv("GRPC_SERVER", "localhost:50051")

    with grpc.insecure_channel(server_address) as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)

        request = user_pb2.UserRequest(user_id=1)
        response = stub.GetUserProfile(request)

        print("=== Profil Pengguna ===")
        print(f"Nama  : {response.name}")
        print(f"Email : {response.email}")
        print(f"Phone : {response.phone}")


if __name__ == "__main__":
    main()
