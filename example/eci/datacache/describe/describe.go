package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	dc "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/datacache"
)

func BuildDescribeDataCacheRequest() *dc.DescribeDataCachesRequest {
	return &dc.DescribeDataCachesRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
	}
}

func DescribeDataCaches(ctx context.Context, cli eci.DataCacheClient) (*dc.DescribeDataCachesResponse, *protocol.Response, error) {
	req := BuildDescribeDataCacheRequest()
	return cli.DescribeDataCaches(ctx, req)
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
	resp, raw, err := DescribeDataCaches(ctx, client.DataCache())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
