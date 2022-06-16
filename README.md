# reapchain-ipfs-grpc


## go-ipfs에 적용하는 방법
### go-ipfs의 go.mod에 아래와 같이 추가
```
require: github.com/mansub1029/reapchain-ipfs/grpc v0.0.0
replace github.com/mansub1029/reapchain-ipfs/grpc v0.0.0 => ../../mansub1029/reapchain-ipfs/grpc (절대경로= $GOPATH/src/github.com/mansub1029/reapchain-ipfs/grpc)
```

### daemon.go에 아래와 같이 추가
```
import reapGrpc "github.com/mansub1029/reapchain-ipfs/grpc"
reapGrpc.ServerInit()
```