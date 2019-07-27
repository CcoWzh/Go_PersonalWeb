# node.js搭建自己的个人网站

个人网站搭建，使用nodejs进行全栈开发

### 版本说明

- 新增博客功能；
- 查看相册大图

---

## 1.我的环境配置

| node.js   | v10.15.1 |
| --------- | -------- |
| npm       | 6.4.1    |
| Mongodb   |          |
| bootstrap | v3       |

使用的是express框架

推荐使用	`JetBrains`的 `WebStorm `

-----


## 2.目录信息

|     目录     |               功能                |   备注   |
| :----------: | :-------------------------------: | :------: |
|    model     | 实现主要的功能，比如调用数据库等  | 累活苦工 |
|    routes    |          负责网站的路由           | 调度资源 |
|    views     | 视图，用于渲染网页，使用模板为ejs | 前端展示 |
|    public    |             静态资源              | 公共资源 |
|     bin      |              bin目录              |          |
| node_modules |            npm 下载包             |          |

-----

## 3.启动

在主目录下，使用命令：

```
node server.js
```

访问：

```
http://127.0.0.1:3000/
```

即可

## 4. 具体说明

#### 4.1 express框架说明

中间件

#### 4.2 路由说明

```
/**
 * 到主目录
 */
app.get("/",user.Homepage);
//=============================================//
/**
 * 登陆&注册界面
 */
app.get("/login",user.Login);
app.get("/logout",user.Logout);
app.post("/login/check",user.check);
app.post("/login/registered",user.registered);
//=============================================//
/**
 * 设置头像页面
 */
app.get("/person",person.getPersonalInfo);
app.get("/upAvatar",person.upAvatar);
app.get("/setAvatar",person.setAvatar);   //剪裁头像
app.post("/dosetAvatar",person.dosetavatar);
app.get("/doCut",person.executeCut);
//=============================================//
/**
 * 写博客页
 */
app.get("/blog",blog.showTalk);
app.get("/blog/doc",blog.writeDoc);
app.post("/blog/saveMd",blog.save);
app.get("/blog/read/:docName",blog.showDoc);
// app.get("/cc",blog.getAllBlog);
//=============================================//
/**
 * 留言页面
 */
app.get("/message",index.liuyan);
app.get("/message/delete",index.delete);
app.post("/message/submitLY",index.doSubmitLY);
app.get("/duliuyan",index.duliuyan);
//=============================================//
//显示相册
app.get("/showIndex",router.showIndex);         //不用加 ()
/**
 * 到 up 路由，上传文件
 * 这个路由是有顺序的，如果这个放在 get("/:albumName") 下面，就会找不到 up 页面
 * 就像流水一样
 */
app.get("/up",router.showUp);  //渲染相册页面
/**
 * 为什么要是 post("/up") 才会可以？明明表单里没有链接
 */
app.post("/up",router.doPost);  //上传文件
/**
 * 到相册目录
 * 在没有写 getAllImaeByAlbumName() 函数之前，随便输入 /** ，都会出现 相册+**
 * 但是写了这个函数后，就会路由寻找到有没有这个 ** 。如果没有，会渲染到 404 页面
 */
app.get("/:albumName",router.showAlbum);  //显示相册
```



#### 4.3 处理请求



#### 4.4 错误处理

