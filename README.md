# 图片压缩程序

## 使用Kratos开发的一个图片压缩程序

## 启动方式
```·
go build 了之后直接把二进制文件丢到服务器运行即可。
```
## 获取图片链接
```
<scema>://url/getcropimg?url=<图片地址>&width=<显示图片宽度>
```
注意：图片地址需要先encodeURIComponent一下，否则路径可能有问题
## 直接把图片路径放在img的src上即刻显示图片
```
 <img src="<scema>://url/getcropimg?url=<图片地址>&width=<显示图片宽度>"/>
```
## 大概流程

<img src="./flow.jpg">