# QR Transfer

## 简介

[LiveDemo]([qrtransfer (whl.red)](https://qr.whl.red/ui/))

实时转发动态二维码，一次只能有一个（群）人用。

## 编译指南

```bash
cd ./qrtransfer
npm run build
cd ../back
go build
```

编译后得到的单个二进制文件包含了资源文件，可以直接上传服务器运行，默认使用8888端口，修改源码以更改默认端口。

## 部署提示

部署时请使用nginx等软件套一层https上去，不然分享功能无法使用。