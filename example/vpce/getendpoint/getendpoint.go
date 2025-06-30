package main

import (
	"context"
	"fmt"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/common"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	ve "github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/vpce"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ct-global.ctapi-internal.ctyun.cn"
	accountId  = ""
	userID     = ""
	regionId   = ""
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	ctx := context.Background()
	openapiConfig := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []ve.Option{
		ve.WithClientConfig(openapiConfig),
	}
	client, err := ve.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	getEndpoint(ctx, client)
}

func getEndpoint(ctx context.Context, client ve.ClientSet) {
	endpointId := ""
	req := &vpce.GetEndpointRequest{
		RegionID:   regionId,
		EndpointID: endpointId,
		CustomInfo: &common.CustomInfo{
			Type: 2,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
				UserId:    userID,
			},
		},
	}
	resp, raw, err := client.VpcEndpoint().GetEndpoint(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
