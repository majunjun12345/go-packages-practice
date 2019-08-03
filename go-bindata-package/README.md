https://blog.csdn.net/chen195822080/article/details/78111672

go get -u github.com/jteeuwen/go-bindata/...

- 生成静态文件的go文件
    go-bindata -o=./asset/asset.go -pkg=asset static/...
    -o:需要生成的路径和文件
    -pkg:包名
    - static/...
        需要执行的目录,这里是 static 路径下所有的文件