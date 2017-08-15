## xsec-dns-server

xsec dns proxy server为一个DNS代理服务器，可以将DNS请求代理到后端的DNS服务器中，在代理的过程中会将dns log写入到数据库中。

### 主要特性如下：

1. 代理DNS请求并记录请求数据
1. 后端支持 sqlite、postgres、mysql和mongodb四种数据库

### 使用说明：

```shell
$ ./xsec-dns-server 
[xorm] [info]  2017/08/15 11:01:24.497380 PING DATABASE mysql
NAME:
   xsec dns proxy server - xsec dns proxy server

USAGE:
   xsec-dns-server [global options] command [command options] [arguments...]
   
VERSION:
   0.1
   
COMMANDS:
     serve    dns proxy Server
     web      web server
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

- serve参数表示启动一个dns代理服务器
- web 参数表示启动一个简单WEB服务器，用来查看dns日志。

### 运行截图

![](https://docs.xsec.io/images/serve.png)

![](https://docs.xsec.io/images/web.png)