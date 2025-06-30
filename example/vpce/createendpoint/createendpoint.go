package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/common/utils"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/common"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	ve "github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/vpce/types/vpce"
)

var (
	accessKey         = ""
	secretKey         = ""
	baseDomain        = "https://ct-global.ctapi-internal.ctyun.cn"
	accountId         = ""
	userID            = ""
	regionId          = ""
	endpointServiceId = ""
	vpcId             = ""
	subnetId          = ""
	systemAccountId   = ""
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

	createEndpoint(ctx, client)
}

func createEndpoint(ctx context.Context, client ve.ClientSet) {
	endpointName := "ep-test"
	requestId := utils.GetRandomString(32)
	masterOrderId := ""
	masterResourceId := ""
	req := &vpce.CreateEndpointRequest{
		ClientToken:       requestId,
		RegionId:          regionId,
		EndpointServiceID: endpointServiceId,
		CycleType:         "on_demand",
		EndpointName:      endpointName,
		SubnetID:          subnetId,
		VpcID:             vpcId,
		CustomInfo: &common.CustomInfo{
			Type: 2,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
				UserId:    userID,
			},
		},
		ChannelInfo: &common.ChannelInfo{
			PaasAccountId:  systemAccountId,
			MasterOrderId:  masterOrderId,
			PaasResourceId: masterResourceId,
		},
	}
	resp, raw, err := client.VpcEndpoint().CreateEndpoint(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
