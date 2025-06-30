package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/price"
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

func BuildBatchDescribePriceRequest() *price.BatchDescribeContainerGroupPriceRequest {
	return &price.BatchDescribeContainerGroupPriceRequest{
		RegionId: "",
		Flavors: []*price.FlavorParam{
			{
				FlavorName: "s7.large.2",
				AzName:     "cn-xinan1-2A",
			},
			{
				FlavorName: "s7.large.4",
				AzName:     "cn-xinan1-2A",
			},
		},
	}
}

func BatchDescribeContainerGroupPrice(ctx context.Context, cli eci.PriceClient) (*price.BatchDescribeContainerGroupPriceResponse, *protocol.Response, error) {
	req := BuildBatchDescribePriceRequest()
	return cli.BatchDescribePrice(ctx, req)
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
	resp, raw, err := BatchDescribeContainerGroupPrice(ctx, client.Price())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
