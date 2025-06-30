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

func BuildCreateDataCacheRequest() *dc.CreateDataCacheRequest {
	return &dc.CreateDataCacheRequest{
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		SecurityGroupId: "sg-nz9vppn0fo",
		SubnetId:        "subnet-uisqhhfqep",
		Bucket:          "testbucket",
		Path:            "testpath",
		DataCacheName:   "testedc",
		Size:            20,
		Tags: []*cg.Tag{
			{
				Key:   "testkey",
				Value: "testvalue",
			},
		},
		Eip: &cg.Eip{
			AutoCreateEip: true,
			EipBandwidth:  1,
		},
		AzName: "cn-huadong1-jsnj1A-public-ctcloud",
		VpcId:  "vpc-717d87ek7m",
		DataSource: &dc.DataSource{
			DataSourceType: "GIT",
			DatasourceOption: &dc.DataSourceOption{
				Git: &dc.Git{
					Url:         "https://hf-mirror.com/datasets/fka/awesome-chatgpt-prompts",
					Branch:      "",
					AccessToken: "",
				},
			},
		},
	}
}

func CreateDataCache(ctx context.Context, cli eci.DataCacheClient) (*dc.CreateDataCacheResponse, *protocol.Response, error) {
	req := BuildCreateDataCacheRequest()
	return cli.CreateDataCache(ctx, req)
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
	resp, raw, err := CreateDataCache(ctx, client.DataCache())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
