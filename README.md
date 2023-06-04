# API文档



## POST /user

### 接口描述

创建用户根目录

### 请求体

​	格式：JSON

​	参数：		

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| accountAddr | string | 用户的钱包账号地址 |

### 返回

| 参数名称 | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| code     | int    | HTTP状态码             |
| message  | string | 返回消息               |
| data     | object | 返回数据，该接口为null |

### 示例

  请求：

```ssh
curl --location 'http://127.0.0.1:9000/user' \
--header 'Content-Type: application/json' \
--data '{"accountAddress": "0x10163d42008C943FA23a611D487f71a1d8f82a79"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": null
}
```





## GET /data/dir

### 接口描述

获取指定目录下的文件列表（包含目录）

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求参数（Params）	

| 参数名称 | 类型   | 描述               |
| -------- | ------ | ------------------ |
| path     | string | 指定的查询目录路径 |

### 返回

| 参数名称 | 类型   | 描述       |
| -------- | ------ | ---------- |
| code     | int    | HTTP状态码 |
| message  | string | 返回消息   |
| data     | object | 返回数据   |

#### 返回数据（data）

  类型：[]object

  object参数：

| 参数名称 | 类型   | 描述                             |
| -------- | ------ | -------------------------------- |
| Name     | string | 文件/目录名称                    |
| Type     | int    | 对象类型<br />0=文件<br />1=目录 |
| Size     | int    | 对象大小（字节）                 |
| Hash     | string | CID                              |

### 示例

  请求：

```ssh
curl --location 'http://127.0.0.1:9000/data/dir?path=%2F' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": [
        {
            "Name": "test-dir",
            "Type": 1,
            "Size": 0,
            "Hash": "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn"
        },
        {
            "Name": "截屏2023-03-23 14.21.47.png",
            "Type": 0,
            "Size": 1132590,
            "Hash": "QmPiRyBa1uQQPu6RzwjWpFnxXc2NaFFoKcrEGRjTALHU35"
        }
    ]
}
```





## POST /data/dir

### 接口描述

在指定路径下创建空目录

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：JSON

​	参数：		

| 参数名称 | 类型   | 描述           |
| -------- | ------ | -------------- |
| path     | string | 指定的绝对路径 |

### 返回

| 参数名称 | 类型   | 描述       |
| -------- | ------ | ---------- |
| code     | int    | HTTP状态码 |
| message  | string | 返回消息   |
| data     | object | 返回数据   |

#### 返回数据（data）

  类型：object

  object参数：

| 参数名称 | 类型   | 描述             |
| -------- | ------ | ---------------- |
| path     | string | 指定的绝对路径   |
| cid      | string | CID              |
| Size     | int    | 对象大小（字节） |

### 示例

  请求：

```ssh
curl --location 'http://127.0.0.1:9000/data/dir' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--header 'Content-Type: application/json' \
--data '{"path":"/test-dir"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": {
        "path": "/test-dir",
        "cid": "QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn"
    }
}
```





## DELETE /data/dir

### 接口描述

删除指定的目录

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：JSON

​	参数：		

| 参数名称 | 类型   | 描述           |
| -------- | ------ | -------------- |
| path     | string | 指定的绝对路径 |

### 返回

| 参数名称 | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| code     | int    | HTTP状态码             |
| message  | string | 返回消息               |
| data     | object | 返回数据，该接口为null |

### 示例

  请求：

```ssh
curl --location --request DELETE 'http://127.0.0.1:9000/data/dir' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--header 'Content-Type: application/json' \
--data '{"path":"/test-dir"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": null
}
```



## GET /data/file

### 接口描述

下载指定目录文件

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：JSON

​	参数：		

| 参数名称 | 类型   | 描述           |
| -------- | ------ | -------------- |
| path     | string | 指定的绝对路径 |

### 返回

  文件字节流

### 示例

  请求：

```ssh
curl --output test.png --location 'http://127.0.0.1:9000/data/file' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--form 'files=@"/Users/xuqiang/Documents/截屏2023-03-23 14.21.47.png"' \
--form 'path="/"'
```

  返回：

  test.png





## GET /data/file/stat

### 接口描述

获取指定路径文件信息

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

格式：JSON

| 参数名称 | 类型   | 描述               |
| -------- | ------ | ------------------ |
| path     | string | 指定的查询目录路径 |

### 返回

| 参数名称 | 类型   | 描述       |
| -------- | ------ | ---------- |
| code     | int    | HTTP状态码 |
| message  | string | 返回消息   |
| data     | object | 返回数据   |

#### 返回数据（data）

  类型：object

  object参数：

| 参数名称 | 类型   | 描述             |
| -------- | ------ | ---------------- |
| path     | string | 文件路径         |
| cid      | string | cid              |
| size     | int    | 对象大小（字节） |

### 示例

  请求：

```ssh
curl --location --request GET 'http://127.0.0.1:9000/data/file/stat' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--header 'Content-Type: application/json' \
--data '{"path":"/截屏2023-03-23 14.21.47.png"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": {
        "path": "/截屏2023-03-23 14.21.47.png",
        "cid": "QmPiRyBa1uQQPu6RzwjWpFnxXc2NaFFoKcrEGRjTALHU35",
        "size": 1132590
    }
}
```







## POST /data/file

### 接口描述

上传文件到指定目录

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：form-data

​	参数：		

| 参数名称 | 类型   | 描述           |
| -------- | ------ | -------------- |
| path     | string | 指定的绝对路径 |
| files    | []file | 文件           |

### 返回

| 参数名称 | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| code     | int    | HTTP状态码             |
| message  | string | 返回消息               |
| data     | object | 返回数据，该接口为null |

### 示例

  请求：

```ssh
curl --location 'http://127.0.0.1:9000/data/file' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--form 'files=@"/Users/xuqiang/Documents/截屏2023-03-23 14.21.47.png"' \
--form 'path="/"'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": null
}
```





## DELETE /data/file

### 接口描述

删除指定的目录

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：JSON

​	参数：		

| 参数名称 | 类型   | 描述           |
| -------- | ------ | -------------- |
| path     | string | 指定的绝对路径 |

### 返回

| 参数名称 | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| code     | int    | HTTP状态码             |
| message  | string | 返回消息               |
| data     | object | 返回数据，该接口为null |

### 示例

  请求：

```ssh
curl --location --request DELETE 'http://127.0.0.1:9000/data/dir' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--header 'Content-Type: application/json' \
--data '{"path":"/test-dir"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": null
}
```





## POST /data/file

### 接口描述

删除指定文件

### 请求头（Headers）

| 参数名称    | 类型   | 描述               |
| ----------- | ------ | ------------------ |
| AccountAddr | string | 用户的钱包账号地址 |

### 请求体

​	格式：JSON

​	参数：		

| 参数名称 | 类型   | 描述               |
| -------- | ------ | ------------------ |
| path     | string | 指定的文件绝对路径 |

### 返回

| 参数名称 | 类型   | 描述                   |
| -------- | ------ | ---------------------- |
| code     | int    | HTTP状态码             |
| message  | string | 返回消息               |
| data     | object | 返回数据，该接口为null |

### 示例

  请求：

```ssh
curl --location --request DELETE 'http://127.0.0.1:9000/data/file' \
--header 'AccountAddr: 0x10163d42008C943FA23a611D487f71a1d8f82a79' \
--header 'Content-Type: application/json' \
--data '{"path":"/截屏2023-03-23 14.21.47.png"}'
```

  返回：

```
{
    "code": 200,
    "message": "OK",
    "data": null
}
```



