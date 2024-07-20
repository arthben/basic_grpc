# gRPC Server
Start gRPC server using docker
```console
cd grpc_server
docker build -t basic_grpc_server:v1.0 .

# on the Dockerfile, from image expose same port with the env BASIC_GRPC_PORT which is port 8099
# specify port to publish with -p parameter ( -p HOST_PORT:CONTAINER_PORT)
docker run -d -p 8111:8099 basic_grpc_server:v1.0

# check the container
docker ps | grep basic_grpc_server
```

Stop gRPC server
```console
# find the PID
docker ps | grep basic_grpc_server | awk '{print $1}'

# stop (PID may different on each machine)
docker stop f511ab800f47
```

## Note
Any update on contract (ex: regenerate after add other service), contract folder should be copied to pkg directory