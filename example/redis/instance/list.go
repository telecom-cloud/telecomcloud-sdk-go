package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/redis"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/redis/types/instance"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://dcs2-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_redis_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []redis.Option{
		redis.WithClientConfig(config),
	}
	client, err := redis.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &instance.DescribeInstancesRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
	}
	resp, raw, err := client.Instance().DescribeInstances(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
