# GO环境

## GO目录结构划分
**适用于个人**
![](http://tuku.dcgamer.top/20190521171259.png?lamber_tuku)
**适用于公司**
![](http://tuku.dcgamer.top/20190521171459.png?lamber_tuku)

## 包的概念

- 和Python一样，把相同功能的代码放到一个目录，称之为包
- 包可以被其他包引用
- main包是用来生成可执行文件，每个程序只有一个main包，每个程序只能有一个main包和一个main函数。定义多了就直接报错了。
- 包的主要用途是提高代码的可复用性。

## 简单Go命令介绍

```shell
go run
```
作用是可以像脚本一样快速执行，它的内部原理其实也是先编译在执行，相当于综合了两个步骤，编译的临时结果会放在一个临时目录，比如tmp目录下。

```shell
go build
```
编译后执行，其实就是编译后生成了机器码然后再进行执行。

```shell
go install
```
安装可执行文件到bin目录，相当于先编译，然后把编译完的文件拷贝到bin目录下。
```shell
go test
```
执行单元测试或者压力测试
```shell
go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/lamber/Library/Caches/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/lamber/go_workspace/Go_study_notebook"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/Cellar/go/1.11.1/libexec"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.11.1/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/8l/g95nllln61j4ly_zm_tqj2m40000gn/T/go-build992691579=/tmp/go-build -gno-record-gcc-switches -fno-common"
```
显示go相关的环境变量
```shell
# go fmt格式化的单位也是包，所以格式和go build是一致的
go fmt
```
格式化源代码，比如go fmt github.com/luffy/lession1

## Go程序结构

- go源码按package进行组织，并且package要放到非注释的第一行
- 一个程序只有一个main包和一个main函数
- main函数是程序的执行入口

## 注释

- 单行注释： //
- 多行注释： /* */
