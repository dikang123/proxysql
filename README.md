# ProxySQL RESTful API

### 1.introduce

I think ProxySQL is a best MySQL proxy software.

But I can't control it through a restful api.

So, I write this project.

If you want build this project,You should download some library. 

Such as:

    github.com/go-sql-driver/mysql

    github.com/gin-gonic/gin


### 2.How to build?

```
# go get -u github.com/go-sql-driver/mysql

# go get -u github.com/gin-gonic/gin

# go install proxysql-master  OR go build main.go
```

### 4. How to running?


    # proxysql-master

By default, proxysql-master print all messages to stdout.

By default, proxysql-master use 3333 port to listen connections.

### 5. How to use?

ref : ![api_en-US.md](https://github.com/imSQL/proxysql-master/blob/v1.3/doc/api_en-US.md)
