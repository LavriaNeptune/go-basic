module example
// 通过 go mod init example 生成

go 1.18

require (
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/rs/zerolog v1.28.0 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
  "package/custom v1.2.3
)
// 通过 go install github.com/rs/zerolog 命令安装第三官方模块,自动生成上述依赖