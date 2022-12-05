# go-scaffold

## Clean Architecture

![Clean Architecture](docs/img/clean-architecture.jpg)

![Layers](docs/img/layers.png)

## Project Layout

```
go-easy
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── app/
│       └── app.go
│   └── pkg/
│       └── cfg/
│       └── db/
│       └── log/
├── pkg/
│   └── logger/
│       └── logger.go
└── vendor/
│   ├── github.com/
│   │   ├── golang/
│   │   ├── prometheus/
│   └── golang.org/
├── go.mod
├── go.sum
```

### `/cmd`

项目中的所有你将要编译成可执行程序的入口代码都放在 `cmd/` 文件夹里，这些代码和业务没有关系。每个程序对应一个文件夹，文件夹的名称应该以程序的名称命名。

每个文件夹必须有一个main包的源文件，该源文件的名称也最好命名成可执行程序的名称，当然也可以保留main文件名。在此会导入和调用 `internal/` 和 `pkg/` 等其他文件夹中相关的代码。

### `/internal`

在项目中不被复用，也不能被其他项目导入，仅被本项目内部使用的代码包即私有的代码包都应该放在 `internal` 文件夹下。该文件夹下的所有包及相应文件都有一个项目保护级别，即其他项目是不能导入这些包的，仅仅是该项目内部使用。

如果你在其他项目中导入另一个项目的 `internal` 的代码包，保存或 `go build` 编译时会报 `use of internal package ... not allowed` 错误，该特性是在 [go 1.4](https://golang.org/doc/go1.4#internalpackages) 版本开始支持的，编译时强行校验。

你可以在 `internal` 文件夹下添加额外的目录来区分可共享和不可共享的内部代码，比如你的实际应用程序代码可以放在 `/internal/app` 目录下(e.g. `/internal/app/myapp`)，这些应用程序共享的代码可以放在 `/internal/pkg` 目录下(e.g. `/internal/pkg/myprivlib`)。

### `/pkg`

如果你把代码包放在 `/pkg` 目录下，其他项目是可以直接导入的，即这里的代码包是开放的(当然你的项目本身也可以直接访问)，所以在 `/pkg` 目录下放代码之前要三思，有没必要这样做。

注意，`/internal` 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的；而 `/pkg` 目录仍然是明确传达该目录中的代码可供其他人安全使用的好方法。如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，这样做是合适的，如果是你自己或公司的某一个项目，个人经验，基本上用不上 `/pkg` 目录。

### `/vendor`

`/vendor` 文件夹包含了所有依赖的三方的源代码，它是go项目最早的依赖包的管理方式。目前大都用的 go mod 的依赖包管理，相对vendor，能指定版本，并且你不用特意手动下载更新依赖包，通过正常的 go build 命令会自动处理，这样会减少项目本身的容量大小。

你可以用命令 `go mod vendor` 来创建你项目的 vendor 目录。如果你项目中既要用到之前的vendor又要用到go mod，你可以使用 `-mod=vendor` 参数进行编译，但是在 go1.14 就不用了，当你用 go build 时，会自动检查项目根目录下有无 vendor 并进行编译。

注意，自从 1.13 以后，Go 还启用了模块代理功能(默认使用 https://proxy.golang.org 作为他们的模块代理服务器)，在国内，模块代理功能默认是被墙的，可以通过以下命令进行设置：

``` sh
$ go env -w GOPROXY=https://goproxy.io,direct
```

## 参考链接

- [Project Layout](https://github.com/golang-standards/project-layout)
- [Go Clean template](https://github.com/evrone/go-clean-template)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)