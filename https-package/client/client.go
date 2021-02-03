package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

}

func httpsClient() {
	/*
		InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true,
		则不会校验证书以及证书中的主机名和服务器主机名是否一致。
	*/
	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }

	// 校验服务端证书
	pool := x509.NewCertPool()
	caCertPath := "../keys/ca/ca.crt"
	//调用ca.crt文件
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	//解析证书
	pool.AppendCertsFromPEM(caCrt)

	client := &http.Client{
		Transport: &http.Transport{
			////把从服务器传过来的非叶子证书，添加到中间证书的池中，使用设置的根证书和中间证书对叶子证书进行验证。
			TLSClientConfig: &tls.Config{RootCAs: pool},
		}}
	resp, err := client.Get("https://localhost")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
