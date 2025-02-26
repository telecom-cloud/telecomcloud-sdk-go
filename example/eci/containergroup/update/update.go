package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	cg "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
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

func BuildUpdateContainerGroupRequest() *cg.UpdateContainerGroupRequest {
	return &cg.UpdateContainerGroupRequest{
		ContainerGroupId: "eci-xvcv485qgvxnlhzz",
		RegionId:         "b342b77ef26b11ecb0ac0242ac110002",
		Tags: []*cg.Tag{
			{
				Key:   "hello",
				Value: "world",
			},
		},
	}
}

func UpdateContainerGroup(ctx context.Context, cli eci.ContainerGroupClient) (*cg.UpdateContainerGroupResponse, *protocol.Response, error) {
	req := BuildUpdateContainerGroupRequest()
	return cli.UpdateContainerGroup(ctx, req)
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	resp, raw, err := UpdateContainerGroup(ctx, client.ContainerGroup())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
