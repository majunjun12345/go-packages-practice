[InFluxDB数据库使用手册](cnblogs.com/huang-yc/p/10487455.html)

开源的时序数据库，用 go 编写，被广泛应用于存储系统的监控数据，IOT 行业的实时数据等场景;

influxDB名词  传统数据库概念
database     数据库
measurement  数据表
point        数据行
timestamp    主键
tag          有索引的列
field        没有索引的列

- point
  相当于数据库里面的一条数据，由时间戳(timestamp)、数据(field)、标签(tag)组成;
  time: 每个数据记录时间，是数据库中的主索引
  field: 各种记录值(没有索引的属性)，例如温度、湿度
  tags: 各种有索引的属性 

- series
  Series 相当于是 InfluxDB 中一些数据的集合，在同一个 database 中，retention policy、measurement、tag sets 完全相同的数据同属于一个 series，同一个 series 的数据在物理上会按照时间顺序排列存储在一起。
  show series from cpu

- 使用的端口
  8083: Web admin管理服务的端口, http://localhost:8083
  8086: HTTP API的端口
  8089: 集群端口(目前还不是很清楚, 配置在全局的bind-address，默认不配置就是开启的)


- 创建数据库
  docker exec -it 77 /bin/sh
  bash
  influx
  show databases; / CREATE DATABASE test_influxdb; 

- 数据保存策略
  查看数据库的保存策略: show retention policies on "db_name"
  创建数据库的保存策略: create retention policy "rp_name" on "db_name" duration 3w replication 1 default  // 副本一般为 1 即可;
  rp_name: 3w：保存3周;