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

### 关于gqlgen 与 proto模型的问题
- 问题1: gqlgen的FLOAT转换为go是float  而proto的float 转换为go是float32 
- 问题2: proto的 int 拥有默认值0 gqlgen前端就无法判断 当前这个是默认值还是就是0

##### 问题1处理 就是在yml指定当前标量类型 (标量定义参考: https://gqlgen.com/reference/scalars/)
##### 问题2处理:
参考文献: https://tomasbasham.dev/development/2017/09/13/protocol-buffers-and-optional-values.html
```
message Int {
    int val = 1;
}
message Resp {
    Int id = 1;  // grpc message 默认是传地址  (这样就可以通过包装类型解决默认值的问题)
}
```


### S1 评估 Evaluation
为检测对 Twirp 框架的理解，请回答以下问题：

如何理解 RPC？这和 GraphQL、REST 有何不同？优势在哪？
Twirp 为我们生成了两个文件(.pb.go 和 .twirp.go)，请问这两个文件分别提供了什么？我们在实现服务端和客户端时，都用到了它们的哪些部分?
在教程提供的 proto 文件里，有以下这行代码，请问这行代码的作用是什么：

option go_package = "haberdasher";

如果 proto 文件需要引用其他 proto 文件里定义的类型，应该怎么做？(`import "base.proto";`)用 Protoc 生成代码时需要注意增加哪些参数?

如果我们需要让 Twirp 使用 HTTP2，而不是当前的 HTTP1.1，应该怎么做？
自行调查有哪些其他的 Go 语言 RPC 框架？他们之间各有什么区别？各自有什么优势？
什么是 protobuf？请用三句话介绍它。
在用 protobuf 定义服务时，有哪些命名规范？

以上问题并不一定能从提供的学习资料中找到答案。

#### proto命名规范
- message使用驼峰式命名, 首字母大写,成员数据使用下划线分隔命名
``` 
message SongServerRequest {
  required string song_name = 1;
}
```
- 枚举类型使用驼峰式命名, 首字母大写,每一项使用下划线大写分隔命名
``` 
enum Foo {
  FIRST_VALUE = 0;
  SECOND_VALUE = 1;
}
```
- grpc的函数接口使用驼峰式命名,首字母大写, 成员数据使用驼峰式命名
``` 
service FooService {
  rpc GetSomething(FooRequest) returns (FooResponse);
}
```

### S2 评估 Evaluation