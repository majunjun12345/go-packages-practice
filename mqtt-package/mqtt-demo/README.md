### 部署 mqtt 服务器
- macOS
    brew install mosquitto
    brew info mosquitto：查看安装信息
    mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf -v：启动 mosquitto 服务，监听端口默认为 1883