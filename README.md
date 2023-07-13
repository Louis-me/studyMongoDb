## 说明

本次代码为go中对mongodb进行增删改查

## 环境

- win10, go 1.20.4
- 本地启动mongodb服务

## 运行代码

```
go run main.go
```

## 客户端测试

```python
import requests

data ={"name": "test111", "password": "1123456", "id": 2, "key": "t_key2"}
resp = requests.post("http://127.0.0.1:8000/UserAdd", json=data)
print(resp.text)


data ={"name": "test1131", "password": "123456811", "id": 1112, "key": "t_key1"}
resp = requests.post("http://127.0.0.1:8000/UserEditOne", json=data)
print(resp.text)
#
data ={"key": "t_key1"}
resp = requests.post("http://127.0.0.1:8000/UserDelete", json=data)
print(resp.text)


resp = requests.get("http://127.0.0.1:8000/UserQuery/t_key2/")
print(resp.text)

resp = requests.get("http://127.0.0.1:8000/UserList")
print(resp.text)

```

