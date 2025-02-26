package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://eci-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &containergroup.DescribeContainerGroupMetricRequest{
		ContainerGroupId: "eci-aslfae6ub29wf2hm",
		RegionId:         "bb9fdb42056f11eda1610242ac110002",
	}

	resp, raw, err := client.ContainerGroup().Monitor(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	/**
	raw: {"statusCode":200,"returnObj":{"name":"eci-aslfae6ub29wf2hm","records":[{"timestamp":"1737012703","network":{},"cpu":{},"memory":{"memoryUsage":5.64117431640625},"gpu":{},"fileSystem":[{"name":"/dev/root","fileSystemUsage":8},{"name":"/dev/shm"},{"name":"/run","fileSystemUsage":1},{"name":"/sys/fs/cgroup"},{"name":"/tmp"}]},{"timestamp":"1737012718","network":{"interfaces":[{"name":"eth0","networkReceiveRate":10913100,"networkTransmitRate":30763},{"name":"sit0"},{"name":"tunl0"}]},"cpu":{"cpuUsage":69.77746282},"memory":{"memoryUsage":9.301948547363281},"gpu":{},"fileSystem":[{"name":"/dev/root","fileSystemUsage":9},{"name":"/dev/shm"},{"name":"/run","fileSystemUsage":1},{"name":"/sys/fs/cgroup"},{"name":"/tmp"}]},{"timestamp":"1737012733","network":{"interfaces":[{"name":"eth0","networkReceiveRate":10913100,"networkTransmitRate":30763},{"name":"sit0"},{"name":"tunl0"}]},"cpu":{"cpuUsage":69.77746282},"memory":{"memoryUsage":9.301948547363281},"gpu":{},"fileSystem":[{"name":"/dev/root","fileSystemUsage":9},{"name":"/dev/shm"},{"name":"/run","fileSystemUsage":1},{"name":"/sys/fs/cgroup"},{"name":"/tmp"}]},{"timestamp":"1737012748","network":{"interfaces":[{"name":"eth0","networkReceiveRate":10913100,"networkTransmitRate":30763},{"name":"sit0"},{"name":"tunl0"}]},"cpu":{"cpuUsage":69.77746282},"memory":{"memoryUsage":9.301948547363281},"gpu":{},"fileSystem":[{"name":"/dev/root","fileSystemUsage":9},{"name":"/dev/shm"},{"name":"/run","fileSystemUsage":1},{"name":"/sys/fs/cgroup"},{"name":"/tmp"}]}],"requestId":"a8422a06-02e5-4e9c-a14a-4ac2aedf4998"}}
	resp: name:"eci-aslfae6ub29wf2hm"  records:{timestamp:"1737012703"  network:{}  cpu:{}  memory:{memoryUsage:5.64117431640625}  gpu:{}  fileSystem:{name:"/dev/root"  fileSystemUsage:8}  fileSystem:{name:"/dev/shm"}  fileSystem:{name:"/run"  fileSystemUsage:1}  fileSystem:{name:"/sys/fs/cgroup"}  fileSystem:{name:"/tmp"}}  records:{timestamp:"1737012718"  network:{interfaces:{name:"eth0"  networkReceiveRate:10913100  networkTransmitRate:30763}  interfaces:{name:"sit0"}  interfaces:{name:"tunl0"}}  cpu:{cpuUsage:69.77746282}  memory:{memoryUsage:9.301948547363281}  gpu:{}  fileSystem:{name:"/dev/root"  fileSystemUsage:9}  fileSystem:{name:"/dev/shm"}  fileSystem:{name:"/run"  fileSystemUsage:1}  fileSystem:{name:"/sys/fs/cgroup"}  fileSystem:{name:"/tmp"}}  records:{timestamp:"1737012733"  network:{interfaces:{name:"eth0"  networkReceiveRate:10913100  networkTransmitRate:30763}  interfaces:{name:"sit0"}  interfaces:{name:"tunl0"}}  cpu:{cpuUsage:69.77746282}  memory:{memoryUsage:9.301948547363281}  gpu:{}  fileSystem:{name:"/dev/root"  fileSystemUsage:9}  fileSystem:{name:"/dev/shm"}  fileSystem:{name:"/run"  fileSystemUsage:1}  fileSystem:{name:"/sys/fs/cgroup"}  fileSystem:{name:"/tmp"}}  records:{timestamp:"1737012748"  network:{interfaces:{name:"eth0"  networkReceiveRate:10913100  networkTransmitRate:30763}  interfaces:{name:"sit0"}  interfaces:{name:"tunl0"}}  cpu:{cpuUsage:69.77746282}  memory:{memoryUsage:9.301948547363281}  gpu:{}  fileSystem:{name:"/dev/root"  fileSystemUsage:9}  fileSystem:{name:"/dev/shm"}  fileSystem:{name:"/run"  fileSystemUsage:1}  fileSystem:{name:"/sys/fs/cgroup"}  fileSystem:{name:"/tmp"}}  requestId:"a8422a06-02e5-4e9c-a14a-4ac2aedf4998"
	*/
}
