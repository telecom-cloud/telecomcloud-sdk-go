package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/flavor"
)

func BuildDescribeAvailableResourceRequest() *flavor.DescribeAvailableResourceRequest {
	return &flavor.DescribeAvailableResourceRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		AzName:     "cn-huadong1-jsnj1A-public-ctcloud",
		Cpu:        2,
		Memory:     4,
		FlavorName: "",
	}
}

func DescribeFlavorResource(ctx context.Context, cli eci.FlavorClient) (*flavor.DescribeAvailableResourceResponse, *protocol.Response, error) {
	req := BuildDescribeAvailableResourceRequest()
	return cli.DescribeAvailableResource(ctx, req)
}

func main() {
	baseDomain := "https://eci-global.ctapi.ctyun.cn"
	config := &config.OpenapiConfig{
		AccessKey: os.Getenv("CTAPI_AK"),
		SecretKey: os.Getenv("CTAPI_SK"),
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	resp, raw, err := DescribeFlavorResource(ctx, client.Flavor())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
