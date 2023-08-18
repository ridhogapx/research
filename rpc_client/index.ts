const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.resolve(__dirname, "auth.proto");

const options = {
    keepCase: true,
    longs: String, 
    enums: String, 
    defaults: true,
    oneofs: true,
};

const packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const proto = grpc.loadPackageDefinition(packageDefinition).proto;

const client = new proto.AuthService(
    "localhost:50052",
    grpc.credentials.createInsecure()
)

client.SignUp({
    email: "fromjs@dev.com",
    password: "12345"
}, (err: any, result: any) => {
    if(err) console.log(err);
    console.log(result);
})