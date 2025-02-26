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

func BuildCreateVirtualNodeRequest() *vn.CreateVirtualNodeRequest {
	return &vn.CreateVirtualNodeRequest{
		RegionId: "b342b77ef26b11ecb0ac0242ac110002",
		AzInfo: []*cg.AzInfo{
			{
				AzName: "cn-xinan1-2A",
			},
		},
		VpcId:           "vpc-yehu9qjjol",
		VSwitchId:       "subnet-z73ymwzk87",
		SecurityGroupId: "sg-0mpilsidy4",
		Tags: []*cg.Tag{
			{
				Key:   "key-test",
				Value: "value-test",
			},
		},
		Taints: []*vn.Taint{
			{
				Key:    "key-test",
				Value:  "value-test",
				Effect: "NoSchedule",
			},
		},
		KubeConfig: "xxxx",
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
	req := BuildCreateVirtualNodeRequest()
	resp, raw, err := client.VirtualNode().CreateVirtualNode(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)

}
