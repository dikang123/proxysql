# ProxySQL RESTful API

### 1.介绍

ProxySQL是我觉得特别好用的MySQL数据库代理软件，其强大的功能深深的吸引了我。

ProxySQL项目地址为： https://github.com/sysown/proxysql

ProxySQL还是一个年轻的项目，其必然有一些不尽如人意的地方。比如对它的管理还得使用命令行来完成。

处于ProxySQL易用性考虑，我决定为它编写一套RESTful风格的api。

本api使用golang开发，调用的第三方库有：

    https://github.com/go-sql-driver/mysql

    https://github.com/labstack/echo

由于我的golang基本上从0起步，所以代码写的很烂，希望能与高手一起学习。

### 2.编译及使用：

```
# go get -u github.com/go-sql-driver/mysql

# go get -u github.com/labstack/echo/...

# go install proxysql-master
```

### 3.程序参数解释

    proxysql-master 
    -s MySQL连接风格的URI。例如：'username/password@hostname:port/dbname' 
    -p 程序运行时监听的端口
    -l 程序运行时输出的日志文件路径。

### 4.运行程序

    proxysql-master -s 'admin/admin@localhost:6032/main' -p 3322 -l /tmp/pm.log

### 5.访问api：

参考doc/api.md
