# 全覆盖测试
```
demo
```

## 简单测试
```
$ cd cmd
$ go build main.go
$ ./main
```

## 数据库环境准备
```
$ git clone https://github.com/pingcap/tidb-docker-compose.git && cd tidb-docker-compose
# 修改compose目录下values.yaml,减少副本（VM），启动
$ sudo docker-compose up -d
```

## TODO
    * 增加期望结果
    * 多数据库支持，用于对比测试
    * 增加代码单元测试
