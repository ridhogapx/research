"use strict";
const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = path.resolve(__dirname, "../auth.proto");
const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};
const packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const proto = grpc.loadPackageDefinition(packageDefinition).proto;
const client = new proto.AuthService("localhost:50051", grpc.credentials.createInsecure());
const stream = client.SignUp({
    email: "devnodejs@mail.com",
    password: "qweasdxzc"
}, (err, res) => {
    if (err) {
        console.log(err);
    }
    else {
        console.log(res);
    }
});
