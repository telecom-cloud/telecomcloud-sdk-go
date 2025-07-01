package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/common"
	// "github.com/telecom-cloud/telecomcloud-sdk-go/test/middleware"

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
		// ve.WithClientMiddleware(middleware.DumpHttpMiddleware),
	}
	client, err := ve.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	listEndpoint(ctx, client)
}

func listEndpoint(ctx context.Context, client ve.ClientSet) {
	//query := "test"
	endpointName := "test-123"
	req := &vpce.ListEndpointRequest{
		RegionID:     regionId,
		EndpointName: endpointName,
		//QueryContent: query,
		CustomInfo: &common.CustomInfo{
			Type: 2,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
				UserId:    userID,
			},
		},
	}
	resp, raw, err := client.VpcEndpoint().ListEndpoint(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
