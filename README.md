# Twirp
### proto file
``` 
syntax = "proto3";
package hello;
option go_package = ".;helloworld";  
// Go 中：默认使用package名作为包名，除非指定了option go_package选项 
// 注意: protoc-gen-go的未来版本将要求指定此内容
// go_package 需要明确指定生成在那里  如果是当前目录 option go_package = ",;package_name"


message HelloResp {
  string Body = 2;
}

message HelloReq {
  string Name = 1;
}

service HelloWorld {
  rpc Hello(HelloReq) returns (HelloResp);
}

// proto --go_out=plugins=grpc:. *.proto
// protoc --twirp_out=. --go_out=. *.proto
```

### 关于HTTP2 TLS
``` 
h2，基于TLS之上构建的HTTP/2，作为ALPN的标识符，两个字节表示，0x68, 0x32，即https
h2c，直接在TCP之上构建的HTTP/2，缺乏安全保证，即http
在HTTP/2 RFC文档出现之前，以上版本字段需要添加上草案版本号，类似于h2-11,h2c-17
```