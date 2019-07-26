
https 通信流程：
好文：https://www.kuacg.com/22672.html

![image.png](https://uploadfiles.nowcoder.com/images/20190413/420034185_1555134777188_4C843633D9B517514370FC7FEF715116)

浏览器里面自带证书机构的公钥，用以对证书解密

![1.png](https://ae01.alicdn.com/kf/HTB1zFVgaUCF3KVjSZJn762nHFXa3.png)

### 制作证书

TLS：安全传输层协议，用于在两个通信应用程序之间提供保密性和数据完整性
- setrver
    生成 key：
        rsa算法：
            openssl genrsa -out server.key 2048
        ECDSA算法：
            openssl ecparam -genkey -name secp384r1 -out server.key
    生成 .crt（自签名公钥 .pem|.crt）：
        openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

    监听 443 端口时，运行程序需要 sudo 权限
    浏览器需要选高级
    客户端需要
    curl 需要带上参数 -k 
- go client 
    client 与 server 进行通信时 client 也要对 server 返回数字证书进行校验
    因为 server 自签证书是无效的 为了 client 与 server 正常通信，通过设置客户端跳过证书校验

### 使用 ca 颁发的证书支持对证书进行校验
    生成 ca 私钥：
        openssl genrsa -out ca.key 2048
    生成CA证书：
        openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt
    生成服务端私钥：
        openssl genrsa -out server.key 2048
    生成证书请求文件：
        openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
    根据CA的私钥和上面的证书请求文件生成服务端证书：
        openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

