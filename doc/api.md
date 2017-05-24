# 本文档主要说明restful-api的使用方法

### 1.用户相关

#### 1.1.查看当前代理保存的用户列表

```
Action: GET
URL: http://127.0.0.1:3333/api/v1/users
参数：无
返回结果： 用户详细信息列表

```
#### 示例

```
curl -X GET 
     -H 'Content-Type: application/json'  \
     127.0.0.1:3333/api/v1/users
返回结果：
[
    {
        "username": "dev",
        "password": "dev",
        "active": 1,
        "use_ssl": 0,
        "default_hostgroup": 0,
        "default_schema": "tpcc",
        "schema_locked": 0,
        "transaction_persistent": 0,
        "fast_forward": 0,
        "backend": 1,
        "frontend": 1,
        "max_connections":10000
    },
    {
        "username":"tianlei2",
        "password":"111111",
        "active":1,
        "use_ssl":0,
        "default_hostgroup":0,
        "default_schema":"tpcc",
        "schema_locked":0,
        "transaction_persistent":0,
        "fast_forward":0,
        "backend":1,
        "frontend":1,
        "max_connections":10000
    }
]
```

#### 1.2.只查询一个用户的详细信息

```
Action: GET
URL: http://127.0.0.1:3333/api/v1/users/:username
参数：无
返回结果： 指定用户相信信息
```

#### 示例

```
curl -X GET
     -H 'Content-Type: application/json'
     127.0.0.1:3333/api/v1/users/dev
查看dev用户的详细信息

返回结果
{
    "username": "dev",
    "password": "dev",
    "active": 1,
    "use_ssl": 0,
    "default_hostgroup": 0,
    "default_schema": "",
    "schema_locked": 0,
    "transaction_persistent": 0,
    "fast_forward": 0,
    "backend": 0,
    "frontend": 0,
    "max_connections": 0
}
```



### 2.后端数据库服务器相关


### 3.查询规则相关


### 4.调度器相关
