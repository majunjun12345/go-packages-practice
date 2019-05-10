package main

import (
	"beeworker-utils/encrypt"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/kingpin"
)

var (
	observer  = kingpin.Command("server", "dump image from idverify_db ")
	enc       = observer.Flag("enc", "encrypt the images").Bool()          // 添加
	secretDir = observer.Flag("secretDir", "public keys dirname").String() // 添加
	dnc       = observer.Flag("dnc", "dncrypt the images").Bool()          // 添加
)

func main() {
	switch kingpin.Parse() {
	case "server":
		if *enc {
			err := EncryptFileV3(filepath.Join("images", "id.jpg"), filepath.Join("images", "id.jpg.meg"))
			checkError1(err)
			// err = EncryptFileV3(filepath.Join("images", "site.jpg"), filepath.Join("images", "site.jpg.meg"))
			// checkError1(err)
			fmt.Println("加密成功")
		}
		if *dnc {
			err := decryptV3(filepath.Join("images", "id.jpg.meg"), filepath.Join("images", "id.jpg"))
			if err != nil {
				checkError1(err)
				return
			}
			fmt.Println("解密成功！")
		}
	default:
	}
}

func choosePublicPemFile() (int64, []byte) {
	fileInfos, err := ioutil.ReadDir(*secretDir)
	checkError1(err)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIndex := r.Intn(len(fileInfos))
	fileName := fileInfos[randIndex].Name() // 获取随机 publickey 文件
	fmt.Println(fileName)
	publickeyid, err := strconv.ParseInt(strings.Split(fileName, ".")[0], 10, 64)
	checkError1(err)

	pubk, err := ioutil.ReadFile(filepath.Join(*secretDir, fileName))
	checkError1(err)

	return publickeyid, pubk
}

func EncryptFileV3(originFile, encFile string) error {
	pubKeyID, pubk := choosePublicPemFile()
	// 加密
	enc, err := encrypt.NewEncryptorV3(pubKeyID, pubk) // pubkeyID是密钥对ID
	if err != nil {
		return err
	}
	fmt.Println(originFile)
	fi, _ := os.Stat(originFile)
	f, _ := os.Open(originFile)

	encReader, err := enc.GetEncryptStreamReader(f, uint32(fi.Size()))
	if err != nil {
		return err
	}

	encFp, err := os.Create(encFile)
	if err != nil {
		return err
	}
	defer encFp.Close()
	io.Copy(encFp, encReader)

	return err
}

func decryptV3(encFile, originFile string) error {
	privPem, _ := ioutil.ReadFile("private.pem")
	dec, err := encrypt.NewDecryptor(privPem)
	if err != nil {
		return err
	}

	decf, err := os.Open(encFile)
	if err != nil {
		return err
	}

	decReader, err := dec.GetDecryptStreamReader(decf)
	if err != nil {
		return err
	}
	ff, _ := os.Create(encFile + ".dec")
	io.Copy(ff, decReader)
	ff.Close()
	if !fileCompare(originFile, encFile+".dec") {
		return err
	}

	return nil
}

func fileCompare(file1, file2 string) bool {
	// Check file size ...
	const chunkSize = 64000

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}

func checkError1(err error) {
	if err != nil {
		panic(err)
	}
}
