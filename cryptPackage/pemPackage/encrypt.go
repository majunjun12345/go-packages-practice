package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
	pem包实现了PEM数据编码（源自保密增强邮件协议），目前PEM编码主要用于TLS密钥和证书。
	也就是这个包实现了其他数据格式与 pem 数据格式之间的编解码。

	编码操作将原始数据生成 公私钥 .pem 文件
	加密时需要解码 public.pem 文件内容生成 key 对 相应内容 进行加密；
	解密时，需要解码 private.pem 文件内容获取 私钥对象，对 相应内容 进行解密；
*/

// 全局变量
var privateKey, publicKey []byte
var privite_key_path = "private.pem"
var public_key_path = "public.pem"

/*
	解码操作需要
*/
// func init() {
// 	var err error
// 	publicKey, err = ioutil.ReadFile("public.pem")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	privateKey, err = ioutil.ReadFile("private.pem")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func main() {
	/*
		编码操作
	*/
	var bits int
	flag.IntVar(&bits, "b", 1024, "密钥长度，默认是1024")
	flag.Parse()

	if GenRsaKey(bits) != nil {
		log.Fatalln("密钥文件生成失败！")
	}
	log.Println("密钥文件生成成功！")

	/*
		解码操作
	*/
	// 获取rsa 公钥长度
	pubKeyLen, _ := GetPubKeyLen(publicKey)
	fmt.Println(pubKeyLen)

	// 获取rsa 私钥长度
	privateLen, _ := GetPriKeyLen(privateKey)
	fmt.Println(privateLen)

	/*
		利用生成的公私钥进行加解密操作
	*/
	// 公钥加密
	src := []byte(`{"name":"酷走天涯"}`)
	cryptoData, error := rsaEncrypt(src)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("加密:", cryptoData)

	// 私钥解密
	dst, error := rsaDecrypt(cryptoData)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("解密:", string(dst))
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	/* 核心代码开始 */
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	fi, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(fi, block)
	if err != nil {
		return err
	}
	/* 核心代码结束 */
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	/* 核心代码开始 */
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	fi, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(fi, block)
	/* 核心代码结束 */
	return err
}

/**
 * 功能：获取RSA公钥长度
 * 参数：public
 * 返回：成功则返回 RSA 公钥长度，失败返回 error 错误信息
 */
func GetPubKeyLen(pubKey []byte) (int, error) {
	if pubKey == nil {
		return 0, errors.New("input arguments error")
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return 0, errors.New("public rsaKey error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return 0, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	// fmt.Println("pub:", pub)
	return pub.N.BitLen(), nil
}

/*
   获取RSA私钥长度
   PriKey
   成功返回 RSA 私钥长度，失败返回error
*/
func GetPriKeyLen(priKey []byte) (int, error) {
	if priKey == nil {
		return 0, errors.New("input arguments error")
	}
	block, _ := pem.Decode(priKey)
	if block == nil {
		return 0, errors.New("private rsaKey error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return 0, err
	}
	return priv.N.BitLen(), nil
}

// 公钥加密
func rsaEncrypt(origData []byte) ([]byte, error) {

	// 从文件中读取公钥编码字节流
	file, error := os.Open(public_key_path)
	if error != nil {
		log.Fatal(error)
	}

	publicKey, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatal(error)
	}
	// 解码对应的block块数据
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 获取公钥key值
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	// 加密数据
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func rsaDecrypt(ciphertext []byte) ([]byte, error) {

	// 从文件中读取私钥pem字节流
	file, error := os.Open(privite_key_path)
	if error != nil {
		log.Fatal(error)
	}
	privateKey, error := ioutil.ReadAll(file)
	if error != nil {
		log.Fatal(error)
	}
	// 解码出对应的block值
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	// 获取私钥对象
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 解密文件
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
