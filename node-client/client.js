
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

// Lokasi file Protocol Buffers
const PROTO_PATH = path.join(__dirname, "..", "user.proto");

// Membaca file user.proto
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

// Mengakses layanan gRPC
const userProto = grpc.loadPackageDefinition(packageDefinition).user;

// Alamat Go Server
const serverAddress = process.env.GRPC_SERVER || "localhost:50051";

// Membuat koneksi ke Go Server
const client = new userProto.UserService(
  serverAddress,
  grpc.credentials.createInsecure()
);

// Meminta data profil pengguna
client.GetUserProfile({ user_id: 1 }, (error, response) => {
  if (error) {
    console.error("Gagal mengambil profil:", error.message);
    client.close();
    return;
  }

  console.log("=== Profil Pengguna ===");
  console.log("Nama  :", response.name);
  console.log("Email :", response.email);
  console.log("Phone :", response.phone);

  client.close();
});
