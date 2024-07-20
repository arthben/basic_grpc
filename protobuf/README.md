# Generate Protobuf
Open terminal window, and run 
```console
make gen
```

## integrate protobuf to project
After success with generate protobuf, copy contract folder to grpc_server folder
```console
cd protobuf
cp -pR contract ../../grpc_server/pkg
```