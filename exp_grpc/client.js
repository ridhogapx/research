const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = "./proto/greet.proto";

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};

const packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const greeter = grpc.loadPackageDefinition(packageDefinition).proto.Greeter;

const client = new greeter(
    "0.0.0.0:50051",
    grpc.credentials.createInsecure()
);

const stream = client.SayHello({name: "Budi"});
stream.on("data", data => {
    console.log(data);
})