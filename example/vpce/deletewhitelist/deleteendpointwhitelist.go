package main

import (
	"context"
	"fmt"
	"github.com/telecom-cloud/client-go/pkg/common/utils"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/common"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	ve "github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/vpce"
)

var (
	regionId          = ""
	endpointServiceId = ""
	accessKey         = ""
	secretKey         = ""
	baseDomain        = "https://ct-global.ctapi-internal.ctyun.cn"
	bssAccountId      = ""
	systemAccountId   = ""
	systemUserId      = ""
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

	deleteEndpointWhitelist(ctx, client)
}

func deleteEndpointWhitelist(ctx context.Context, client ve.ClientSet) {
	requestId := utils.GetRandomString(32)
	req := &vpce.DeleteEndpointWhitelistRequest{
		ClientToken:       requestId,
		RegionID:          regionId,
		EndpointServiceID: endpointServiceId,
		BssAccountID:      bssAccountId,
		CustomInfo: &common.CustomInfo{
			Type: 2,
			Identity: &common.TenantIdentity{
				AccountId: systemAccountId,
				UserId:    systemUserId,
			},
		},
	}
	resp, raw, err := client.VpcEndpoint().DeleteEndpointWhitelist(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp:  %v\nstautsCode:%v\n", string(raw.Body()), resp, raw.StatusCode())
}
