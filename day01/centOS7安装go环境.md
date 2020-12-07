### 1.安装：
yum install golang -y 

### 2.环境配置
查看 golang 安装安装 whereis go
1）修改配置文件

vi /etc/profile 

### 3.添加环境变量 在文件后面追加如下文本：
- GOROOT  go env中显示 go 环境地址
  
   export GOROOT=/usr/lib/golang
- GOPATH go 项目地址（可自己设置）
  
   export GOPATH=/opt/go
-  GOPATH bin 
  
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin 
需要立即生效，在终端执行如下命令：

source /etc/profile

在 /opt/go 目录下新建三个目录： src bin pkg

