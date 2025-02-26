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

func BuildDescribePriceRequest() *price.DescribeContainerGroupPriceRequest {
	return &price.DescribeContainerGroupPriceRequest{
		RegionId: "b342b77ef26b11ecb0ac0242ac110002",
		//Cpu:      1,
		//Memory:   2,
		FlavorName: "s7.large.2",
		AzName:     "cn-xinan1-2A",
	}
}

func DescribeContainerGroupPrice(ctx context.Context, cli eci.PriceClient) (*price.DescribeContainerGroupPriceResponse, *protocol.Response, error) {
	req := BuildDescribePriceRequest()
	return cli.DescribePrice(ctx, req)
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
	resp, raw, err := DescribeContainerGroupPrice(ctx, client.Price())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
