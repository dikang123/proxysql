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

#### 1.3.创建一个新的用户

```
Action: POST
URL: http://127.0.0.1:3333/api/v1/users
参数： username用户名，password密码
返回结果: 成功返回创建的用户的详细信息
```

#### 示例

```
curl -X POST
     -H 'Content-Type: application/json'
     -d '{"username":"admin","password":"admin"}'
     127.0.0.1:3333/api/v1/users
创建admin用户，密码为admin
{
    "username": "admin",
    "password": "admin",
    "active": 0,
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

#### 1.4.更新用户状态

```
Action: PUT
URL: http://127.0.0.1:333/api/v1/users/status
参数：username用户名，active值为0用户不激活，1用户激活
返回结果：成功返回OK
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"username":"dev","active":1}'
激活dev用户

curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"username":"dev","active":0}'
不激活dev用户
```

#### 1.5.更新用户主机组

```
Action: PUT
URL: http://127.0.0.1:333/api/v1/users/hostgroup
参数：username用户名，default_hostgroup值为主机组id
返回结果：成功返回OK
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"username":"dev","default_hostgroup":1}'

把dev用户归于主机组1中

```

#### 1.6.更新用户schema

```
Action: PUT
URL: http://127.0.0.1:333/api/v1/users/schema
参数：username用户名，schema值为用户默认可以访问的数据库
返回结果：成功返回OK
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"username":"dev","default_schema":"testdb"}'

把dev用户默认访问的数据库设定为testdb

```

#### 1.7.更新用户最大连接数

```
Action: PUT
URL: http://127.0.0.1:333/api/v1/users/maxconnection
参数：username用户名，max_connections用户最大连接数
返回结果：成功返回OK
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"username":"dev","max_connections":10000}'

把dev用户的最大连接数设定成10000

```


### 2.后端数据库服务器相关

#### 2.1.查询所有后端服务节点信息

```
Action: GET
URL: http://127.0.0.1:3333/api/v1/servers
参数：无
返回结果：后端数据库节点信息
```

#### 示例

```
curl -X GET
     -H 'Content-Type: application/json'
     127.0.0.1:3333/api/v1/servers
```

#### 2.2.根据主机组查询后端节点的信息

```
Action: GET
URL: http://127.0.0.1:3333/api/v1/servers/:hostgroup
参数： hostgroup主机组id号
返回结果： 后端数据库节点信息
```

#### 示例

```
curl -X GET
     -H 'Content-Type: application/json'
     127.0.0.1:3333/api/v1/servers/1
查询主机组1中的后端数据库节点信息
```

#### 2.3.查询指定主机的信息

```
Action: PUT
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id号，hostname主机名,port主机端口号
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d '{"hostgroup_id":1,"hostname":"dn03","port":3307}'
     127.0.0.1:3333/api/v1/servers

查询主机组1中主机名为dn03，端口为3307的主机信息

```

#### 2.4.在指定主机组下新建一个主机

```
Action: POST
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id号，hostname主机名，port端口
返回结果： 成功返回OK
```

#### 示例

```
curl -X POST
     -H 'Content-Type: application/json'
     -d '{"hostgroup_id":1,"hostname":"dn03","port":3307}'
     127.0.0.1:3333/api/v1/servers
```

#### 2.5.改变某主机组内主机的状态

```
Action: PUT
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id，hostname主机组名，port主机端口
返回结果：成功返回OK
```

#### 示例

```
curl -X PUT
     -H 'Content-Type: application/json'
     -d
     '{"hostgroup_id":1,"hostname":"dn03","port":3307,"status":"ONLINE|SOFT_OFFLINE|HARD_OFFLINE"}'
     127.0.0.1:3333/api/v1/servers

```

#### 2.6.改变某主机组内主机的权重

```
Action: PUT
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id，hostname主机组名，port主机端口,weight权重
返回结果：成功返回OK
```

#### 示例


```
curl -X PUT
     -H 'Content-Type: application/json'
     -d
     '{"hostgroup_id":1,"hostname":"dn03","port":3307,"weight":100}'
     127.0.0.1:3333/api/v1/servers
```

#### 2.7.改变某主机组内主机的最大连接数

```
Action: PUT
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id，hostname主机组名，port主机端口,max_connections最大
连接数
返回结果：成功返回OK
```

#### 示例


```
curl -X PUT
     -H 'Content-Type: application/json'
     -d
     '{"hostgroup_id":1,"hostname":"dn03","port":3307,"max_connections":100}'
     127.0.0.1:3333/api/v1/servers
```


#### 2.7.删除指定主机

```
Action: DELETE
URL: http://127.0.0.1:3333/api/v1/servers
参数： hostgroup_id主机组id，hostname主机组名，port主机端口
连接数
返回结果：成功返回OK
```

#### 示例


```
curl -X DELETE
     -H 'Content-Type: application/json'
     -d
     '{"hostgroup_id":1,"hostname":"dn03","port":3307}'
     127.0.0.1:3333/api/v1/servers
```




### 3.查询规则相关



### 4.调度器相关
