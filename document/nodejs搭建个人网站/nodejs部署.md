# node.js在云服务器上部署遇到的问题

#### 1. 如何不间断的运行一个服务

使用命令：

```
nohup node server.js  >/dev/null 2>&1 &
```

使用`&`命令后，作业被提交到后台运行，当前控制台没有被占用，但是一但把当前控制台关掉(退出帐户时)，作业就会停止运行。

`nohup`命令可以在你退出帐户之后继续运行相应的进程。`nohup`就是不挂起的意思( no hang up)。该命令的一般形式为：` nohup command &`

如果使用`nohup`命令提交作业，那么在 缺省情况 下该作业的所有输出都被重定向到一个名为`nohup.out`的文件中，除非另外指定了输出文件：

```
nohup command > myout.file 2>&1
```

> #### >/dev/null

这条命令的作用是将标准输出1重定向到`/dev/null`中。`/dev/null`代表 linux 的空设备文件，所有往这个文件里面写入的内容都会丢失，俗称“黑洞”。那么执行了`>/dev/null`之后，标准输出就会不再存在，没有任何地方能够找到输出的内容。

>####2>&1

这条命令用到了重定向绑定，采用&可以将两个输出绑定在一起。这条命令的作用是错误输出将和标准输出同用一个文件描述符，说人话就是错误输出将会和标准输出输出到同一个地方。

> ####> /dev/null 2>&1 

的作用就是让标准输出重定向到`/dev/null`中（丢弃标准输出），然后错误输出由于重用了标准输出的描述符，所以错误输出也被定向到了00000`/dev/null`中，错误输出同样也被丢弃了。执行了这条命令之后，**该条shell命令将不会输出任何信息到控制台，也不会有任何信息输出到文件中**。



#### 2.如何查看进程

查看进程

```
ps a 显示现行终端机下的所有程序，包括其他用户的程序
ps aux|grep server.js  //查看servse.js的程序运行进程
```

杀死进程;

```
使用kill命令结束进程：kill xxx
常用：kill －9 25743
```

> app.listen(8080,"172.27.0.4");  

查看使用某端口的进程

```
`lsof -i:8090`
```





#### 3.MongoVUE里的collections为空

win 下的启动方法，解决方法：

```
mongod -storageEngine mmapv1 -dbpath F:\mongotestdb
```



#### 4.URL里莫名出现%20问题

本来应该是：

```
http://127.0.0.1:3000/message/delete?id=5c88abf840258d10c4c67c01
```

但是，却出现了：

```
http://127.0.0.1:3000/message/delete?id=%205c88abf840258d10c4c67c01
```

原因是：URL里出现%20的原因的地址中存在的空格被转码成了%20，所以在**写模板的时候要注意了**，有空格的话把它删掉就可以了。



#### 5 rar 解压

用于在Linux上解压win上的文件

```
rar x etc.rar 
```



#### 6 安装gm

上官网下载gm linux版本

安装插件
```
yum install -y libjpeg-devel libjpeg
yum install -y libpng-devel libpng
yum install -y giflib-devel giflib
```
安装GraphicsMagick
```
tar -zxvf GraphicsMagick-1.3.25.tar.gz 
mv GraphicsMagick-1.3.31 /usr/local/
cd GraphicsMagick-1.3.31
./configure --prefix=/usr/local/GraphicsMagick-1.3.31 --with-quantum-depth=8   --enable-shared --enable-static

make
make install
```
设置环境变量

>vim /etc/profile  在最后添加如下配置
```
export GMAGICK_HOME="/usr/local/GraphicsMagick-1.3.31"
export PATH="$GMAGICK_HOME/bin:$PATH"
```

使环境生效

> source /etc/profile 



#### 7  mongodb的安装和登陆问题

如果mongodb不加密登陆的话，云服务器会有破解的危险



#### 8 从本地上传文件到服务器 从服务器下载文件到本地

在终端输入

```
scp -r 本地文件路径 服务器帐号名@服务器的adress:想要保存的路径 #从本地到服务器 
scp -r 服务器帐号名@服务器的adress:文件路径 本地保存路径 #从服务器到本地
```



#### 9 如何递归地压缩一个目录及目录下的文件？

Linux下压缩一个文件：

```
zip -r var-log-dir.zip /var/log/ 
```



#### 10 npm 强制清除缓存

```
npm cache clean --force
```

