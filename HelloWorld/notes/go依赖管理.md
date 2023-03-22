# go依赖管理

在go项目的根目录下用`go mod init 包名`初始化一个包，生成一个`go.mod`文件

```
└─learn_go
    │  .gitignore
    │  go.mod
    │  README.md
    │      
    ├─HelloWorld
    │  ├─1.1
    │  │      helloworld.go
    │  │      
    │  ├─1.2
    │  │  │  echo1.go
    │  │  │  echo2.go
    │  │  │  echo3.go
    │  │  │  echo4.go
    │  │  │  echo5.go
    │  │  │  
    │  │  └─stringtest
    │  │          addstring.go
    │  │          addstring_test.go
    │  │          
    │  ├─1.3
    │  │      1.txt
    │  │      2.txt
    │  │      dup1.go
    │  │      dup2.go
    │  │      dup3.go
    │  │      dup4.go
    │  │      
    │  └─notes
    │          go依赖管理.md
    │          
    └─pkg
        └─util
                dup_util.go
```

`go.mod`文件内容

```go
module "github/tianluoding/learn_go"

go 1.18
```

