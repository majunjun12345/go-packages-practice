package main

import (
	"fmt"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	cli      client.Client
	database = "test_influxdb"
)

func main() {
	defer cli.Close()

	// Insert()

	cmd := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	resp, err := Query(cmd)
	if err != nil {
		panic(err)
	}

	for _, r := range resp.Results {
		fmt.Println(r)
	}
}

func Insert() {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	if err != nil {
		panic(err)
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   201.1,
		"system": 43.3,
		"user":   86.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now()) // cpu_usage 不需要额外创建，这里会自动创建
	if err != nil {
		panic(err)
	}

	bp.AddPoint(pt)

	err = cli.Write(bp)
	if err != nil {
		panic(err)
	}
	fmt.Println("write success")
}

func Query(cmd string) (*client.Response, error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	resp, err := cli.Query(q)
	if err != nil {
		panic(err)
	}
	if resp.Error() != nil {
		panic(resp.Error())
	}
	return resp, nil
}

// Conn influxdb conn
func init() {
	var err error
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
		// Username: username,
		// Password: pwd,
		// Timeout:  time.Second * 5,
	})
	if err != nil {
		panic(err)
	}

	// 建立数据库
	createDbSQL := client.NewQuery(fmt.Sprintf("CREATE DATABASE %s", database), "", "")
	_, err = cli.Query(createDbSQL)
	if err != nil {
		panic(err)
	}

	// 过期策略
	createRPSQL := client.NewQuery(fmt.Sprintf("CREATE RETENTION POLICY default ON %s DURATION 360d REPLICATION 1 DEFAULT", database), database, "")
	_, err = cli.Query(createRPSQL)
	if err != nil {
		panic(err)
	}
}
