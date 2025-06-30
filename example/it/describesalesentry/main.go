package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it/types/sales"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://bss-global.ctapi-internal.ctyun.cn"
)

func init() {
	//accessKey = os.Getenv("CTAPI_AK")
	//secretKey = os.Getenv("CTAPI_SK")
	//domain := os.Getenv("CTAPI_DOMAIN")
	//if domain != "" {
	//	baseDomain = domain
	//}
}

func main() {
	ctx := context.Background()
	openapiConfig := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []it.Option{
		it.WithClientConfig(openapiConfig),
	}
	client, err := it.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strings.Repeat("=", 50))
	describeSalesEntry(ctx, client, "PAAS_ECI_BASIC")
	fmt.Println(strings.Repeat("=", 50))
	describeSalesEntry(ctx, client, "PAAS_ECI_PREMIUM")
	fmt.Println(strings.Repeat("=", 50))
	describeSalesEntry(ctx, client, "PAAS_ECI_EBS")
}

func describeSalesEntry(ctx context.Context, client it.ClientSet, resourceType string) {
	req := &sales.DescribeSalesEntryRequest{
		PageSize:     200,
		PageNum:      1,
		ServiceTag:   "PAAS",
		ResourceType: resourceType,
	}
	resp, raw, err := client.SalesEntry().DescribeSalesEntry(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
