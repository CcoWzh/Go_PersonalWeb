# git基本使用方法（一）

### 1. 安装

参考这个[链接方法](http://www.cnblogs.com/ximiaomiao/p/7140456.html)

#### 2. 使用

git 初始化：

```
设置用户名：git  config -- global  user.name  '你再github上注册的用户名';
设置用户邮箱：git  config -- global  user.email  '注册时候的邮箱';

git config --list       //查看是否配置成功
```



使用当前目录作为Git仓库，我们只需使它初始化`init`：

```
git init            //初始化，会生成 ".git"的文件夹
git init 指定目录    //初始化后，会在指定目录下会出现一个名为 .git 的目录
```



提交暂存区仓库的一系列操作`add`和`commit`：

```
$ git status          //命令用于查看项目的当前状态
$ git add 文件名       //执行 git add 命令来添加文件，提交到暂存区
$ git add .
//$ git rm 'test.txt'
$ git commit -m "delete test.txt"  //提交到工作区（git仓库）
$ git push           //提交到git远程仓库，如果没有设置git remote命令，需要先设置
```



git的`push`命令：

```
$ git remote add origin https://github.com/HanlaoTwo/SparkStudy.git  #添加远程仓库
“https://github.com/HanlaoTwo/SparkStudy.git”为你的远程仓库地址
#报错
fatal: remote origin already exists.
表明这个仓库已经存在了，不用再添加到本地git里了。如果非要添加，可以先删除在添加
#删除
git remote rm origin
#再添加
git remote add origin https://github.com/HanlaoTwo/SparkStudy.git
```

```
$ git push origin master
```

上面命令表示，将本地的`master`分支推送到`origin`主机的`master`分支。如果`master`不存在，则会被新建。

git push origin master的意思就是上传本地当前分支代码到master分支。git push是上传本地所有分支代码到远程对应的分支上

master是主分支，还可以建一些其他的分支用于开发。



删除`rm`：

```
使用git删除：$ git rm '文件名'
然后提交操作:$ git commit -m "delete test.txt"
更新到远程仓库：$ git push
强行推送到远程仓库：$ git push -f origin master
```



查看远程仓库：

```
Administrator@HP-PC MINGW64 ~/Desktop/tianlang (master)
$ git remote
origin

Administrator@HP-PC MINGW64 ~/Desktop/tianlang (master)
$ git remote show origin
Enter passphrase for key '/c/Users/Administrator/.ssh/id_rsa':
* remote origin
  Fetch URL: git@192.168.10.113:wuzhihui/tianlang.git
  Push  URL: git@192.168.10.113:wuzhihui/tianlang.git
  HEAD branch: master
  Remote branch:
    master tracked
  Local branch configured for 'git pull':
    master merges with remote master
  Local ref configured for 'git push':
    master pushes to master (up to date)
```



改仓库名：

使用`mv`命令

```
Administrator@HP-PC MINGW64 ~/Desktop/tianlang (master)
$ git mv tianlang_java_sol Tl_src

Administrator@HP-PC MINGW64 ~/Desktop/tianlang (master)
$ git commit -m "change name 1"
[master 3a163f8] change name 1
 2 files changed, 0 insertions(+), 0 deletions(-)
 rename {tianlang_java_sol => Tl_src}/Integral.sol (100%)
 rename {tianlang_java_sol => Tl_src}/testIntegral02.java (100%)

Administrator@HP-PC MINGW64 ~/Desktop/tianlang (master)
$ git push
Enter passphrase for key '/c/Users/Administrator/.ssh/id_rsa':
Enumerating objects: 3, done.
Counting objects: 100% (3/3), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (2/2), 257 bytes | 257.00 KiB/s, done.
Total 2 (delta 0), reused 0 (delta 0)
To 192.168.10.113:wuzhihui/tianlang.git
   97bc76b..3a163f8  master -> master
```

每一次更改都得提交到暂存区，并更新到远程仓库才有效。

可以屯一波，到时候一起push，但一定要提交到暂存区。



`.git`目录的文件分析：



### 3. 分支管理

```
查看状态
Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract
$ git status
fatal: not a git repository (or any of the parent directories): .git

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract
$ ls
tianlang-smart-contract/

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract
$ cd tianlang-smart-contract/

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (master)
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean

查看主支和分支
Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (master)
$ git branch
* master

拉取远程代码
Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (master)
$ git pull
Enter passphrase for key '/c/Users/Administrator/.ssh/id_rsa':
Already up to date.


Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (master)
$ git branch -a
* master
  remotes/origin/HEAD -> origin/master
  remotes/origin/feature
  remotes/origin/master

切换到分支上
Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (master)
$ git checkout feature
Switched to a new branch 'feature'
Branch 'feature' set up to track remote branch 'feature' from 'origin'.

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git branch
* feature
  master
```

```
Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git add Tl_src/ README.md
warning: LF will be replaced by CRLF in README.md.
The file will have its original line endings in your working directory
warning: LF will be replaced by CRLF in Tl_src/testIntegral02.java.
The file will have its original line endings in your working directory

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git status
On branch feature
Your branch is up to date with 'origin/feature'.

Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        modified:   README.md
        new file:   Tl_src/Integral.sol
        new file:   Tl_src/testIntegral02.java


Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git commit -m "First update."
[feature 7295968] First update.
 3 files changed, 339 insertions(+), 2 deletions(-)
 create mode 100644 Tl_src/Integral.sol
 create mode 100644 Tl_src/testIntegral02.java

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git branch -a
* feature
  master
  remotes/origin/HEAD -> origin/master
  remotes/origin/feature
  remotes/origin/master

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
$ git push origin feature
Enter passphrase for key '/c/Users/Administrator/.ssh/id_rsa':
Enumerating objects: 8, done.
Counting objects: 100% (8/8), done.
Delta compression using up to 8 threads
Compressing objects: 100% (6/6), done.
Writing objects: 100% (6/6), 4.16 KiB | 1.04 MiB/s, done.
Total 6 (delta 0), reused 0 (delta 0)
remote:
remote: To create a merge request for feature, visit:
remote:   http://192.168.10.113/tianlang/tianlang-smart-contract/merge_requests/new?merge_request%5Bsource_branch%5D=feature
remote:
To 192.168.10.113:tianlang/tianlang-smart-contract.git
   648949b..7295968  feature -> feature

Administrator@HP-PC MINGW64 ~/Desktop/tianlang-contract/tianlang-smart-contract (feature)
```

冲突：

```
git checkout feature 命令的时候，将提示出错：
error: Your local changes to the following files would be overwritten by checkout:
        readme.txt
Please commit your changes or stash them before you switch branches.
（请在切换分支之前提交您的更改或隐藏它们）

出现这个问题的原因是其他人修改了xxx.php并提交到版本库中去了，而你本地也修改了xxx.php，这时候你进行git pull操作就好出现冲突了，解决方法，在上面的提示中也说的很明确了。

1）直接commit本地的修改
2）通过git stash
```

