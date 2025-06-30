package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it/types/order"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = ""
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []it.Option{
		it.WithClientConfig(config),
		it.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := it.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &order.QueryResourceInfoByMasterOrderIdRequest{
		MasterOrderId: "c4dc07308b0f45f49bd405fa0c2db694",
	}
	resp, raw, err := client.Order().QueryResourceInfoByMasterOrderId(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
