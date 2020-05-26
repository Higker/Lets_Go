首先我们需要安装kafak
这个可以去官方网站下载
然后解压即可安装

1.先修改配置文件 zeekeeper的
zookeeper.properties
dataDir=/root/kafka/zookeeper 数据存放目录

2.修改kafka配置
server.properties
log.dirs=/root/kafka/kafka-logs 日志存放目录
advertised.listeners=PLAINTEXT://104.215.49.91:9092 远程本机ip端口绑定


3.启动2个中间件命令
./kafka-server-start.sh -daemon ../config/server.properties
./zookeeper-server-start.sh -daemon ../config/zookeeper.properties
