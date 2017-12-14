# ProxySQL RESTful API

### 1.introduce

I think ProxySQL is a best MySQL proxy software.

But I can't control it through a restful api.

So, I write this project.

If you want build this project,You should download some library. 

Such as:

    https://github.com/go-sql-driver/mysql

    https://github.com/labstack/echo


### 2.How to build?

```
# go get -u github.com/go-sql-driver/mysql

# go get -u github.com/labstack/echo/...

# go install proxysql-master
```

### 4. How to running?


    # proxysql-master

By default, proxysql-master print all messages to stdout.


### 5. How to use?

ref : ![api_en-US.md](https://github.com/imSQL/proxysql-master/blob/v1.3/doc/api_en-US.md)
