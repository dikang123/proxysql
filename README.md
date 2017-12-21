# ProxySQL RESTful API

### 1.introduce

I think ProxySQL is a best MySQL proxy software.

But I can't control it through a restful api.

So, I write this project.

If you want build this project,You should download some libraries. 

Such as:

    github.com/go-sql-driver/mysql

    github.com/gin-gonic/gin


### 2.How to build?

#### 2.1 build

You can quickly build proxysql-master on your OS.

I recommend go version >= 1.6.

```
# go get -u github.com/go-sql-driver/mysql

# go get -u github.com/gin-gonic/gin

# go install proxysql-master 
```

#### 2.2. docker build

You can quickly build proxysql-master with Docker.

The Dockerfile in docker directory.

```
# docker build -t proxysql-master .

```

### 4. How to running?


    # proxysql-master

By default, proxysql-master print all messages to stdout.

By default, proxysql-master use 3333 port to listen connections.

If you running proxysql-master by docker , execute this command:

    # docker run -it proxysql-master

### 5. How to use?

ref : ![api_en-US.md](https://github.com/imSQL/proxysql-master/blob/v1.3/doc/api_en-US.md)
