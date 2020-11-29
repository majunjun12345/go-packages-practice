module testGoScripts

go 1.12

require (
	git.internal.yunify.com/majun/pipe2ignite v0.0.0-20191212091909-f23661d9be46
	git.internal.yunify.com/manage/common v0.0.0-20201117011148-37d215f37072
	github.com/Shopify/sarama v1.19.0
	github.com/Workiva/go-datastructures v1.0.50
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/alexedwards/scs/v2 v2.4.0 // indirect
	github.com/arl/statsviz v0.2.1
	github.com/astaxie/beego v1.12.0
	github.com/aws/aws-sdk-go v1.23.0
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/coreos/etcd v3.3.17+incompatible
	github.com/dgraph-io/badger v1.6.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/didip/tollbooth_gin v0.0.0-20170928041415-5752492be505
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-echarts/statsview v0.2.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.3
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-multierror v0.0.0-20161216184304-ed905158d874
	github.com/influxdata/influxdb1-client v0.0.0-20200827194710-b269163b24ab
	github.com/jinzhu/gorm v1.9.10
	github.com/judwhite/go-svc v1.1.2
	github.com/julienschmidt/httprouter v1.2.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/muesli/cache2go v0.0.0-20191019095710-4098a3aa8c94
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/olivere/elastic/v7 v7.0.6
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/panjf2000/ants/v2 v2.4.1
	github.com/panjf2000/gnet v1.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paulbellamy/ratecounter v0.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/qingcloud-iot/edge-app-go v0.0.0-20201125020549-7967cdc21332 // indirect
	github.com/qingcloud-iot/edge-driver-go v0.0.0-20201030073905-730148ce9dd2
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/robfig/cron v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smallnest/goframe v1.0.0
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/stretchr/testify v1.6.1
	github.com/tealeg/xlsx v1.0.5
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	github.com/ulule/limiter v2.2.2+incompatible
	github.com/ulule/limiter/v3 v3.7.0
	github.com/unknwon/com v1.0.1
	github.com/urfave/negroni v1.0.0
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/vmihailenco/treemux v0.3.0
	github.com/xiaomeng79/go-log v2.0.4+incompatible
	github.com/yunify/qingcloud-sdk-go v2.0.0-alpha.35+incompatible
	go.etcd.io/etcd v3.3.13+incompatible
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
	golang.org/x/text v0.3.3
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc v1.25.1
	gopkg.in/fsnotify.v1 v1.4.7
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
)
