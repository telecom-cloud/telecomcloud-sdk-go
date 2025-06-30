package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	cg "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
	dc "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/datacache"
)

func BuildUpdateDataCacheRequest() *dc.UpdateDataCacheRequest {
	return &dc.UpdateDataCacheRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		DataCacheId: "edc-3oytf3ix1hbph9zx",
		Tags: []*cg.Tag{
			{
				Key:   "testkey",
				Value: "testvalue",
			},
		},
	}
}

func UpdateDataCache(ctx context.Context, cli eci.DataCacheClient) (*dc.UpdateDataCacheResponse, *protocol.Response, error) {
	req := BuildUpdateDataCacheRequest()
	return cli.UpdateDataCache(ctx, req)
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
	resp, raw, err := UpdateDataCache(ctx, client.DataCache())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
