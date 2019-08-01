# cd D:\\gopath\\src\\go-go-go
docker build -t go-go-go -f docker/Dockerfile .
docker tag go-go-go adam077/go-go-go:1
docker push adam077/go-go-go:1
docker rmi adam077/go-go-go:1
docker rmi $(docker images | grep "none" | awk '{print $3}')