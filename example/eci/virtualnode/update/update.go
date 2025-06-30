package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	cg "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
	vn "github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/virtualnode"
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

func BuildUpdateVirtualNodeRequest() *vn.UpdateVirtualNodeRequest {
	return &vn.UpdateVirtualNodeRequest{
		RegionId:      "b342b77ef26b11ecb0ac0242ac110002",
		VirtualNodeId: "vnd-8y5m277r8b1csi7u",
		Tags: []*cg.Tag{
			{
				Key:   "key-update",
				Value: "value-update",
			},
		},
	}
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
	req := BuildUpdateVirtualNodeRequest()
	resp, raw, err := client.VirtualNode().UpdateVirtualNode(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\n\nresp: %v\n", string(raw.Body()), resp)
}
