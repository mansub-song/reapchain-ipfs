# reapchain-ipfs-grpc


## go-ipfs에 적용하는 방법
### go-ipfs의 go.mod에 아래와 같이 추가
```
require: github.com/reappay/reapchain-ipfs/grpc v0.0.0
replace github.com/reappay/reapchain-ipfs/grpc v0.0.0 => ../../reappay/reapchain-ipfs/grpc (절대경로= $GOPATH/src/github.com/reappay/reapchain-ipfs/grpc)
```

### daemon.go에 아래와 같이 추가
```
import reapGrpc "github.com/reappay/reapchain-ipfs/grpc"
reapGrpc.ServerInit()
```