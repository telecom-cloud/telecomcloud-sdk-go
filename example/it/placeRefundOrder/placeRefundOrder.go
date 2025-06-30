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

	resourceDetail := Order{
		AutoApproval: true,
		Source:       "8",
		RefundReason: "手动退订",
		Resources: []Resource{
			{
				ResourceIds: []string{"1c68415975914dd2b2166d8e57f0f88d"},
			},
		},
		CustomInfo: CustomInfo{
			Type: 2,
			Identity: Identity{
				AccountId: "a3f2b12f917d4ea3afa6e3c7e9553694",
			},
		},
	}

	ctx := context.Background()
	req := &order.PlaceRefundOrderRequest{
		ResourceDetailJson: resourceDetail.Marshal(),
		Type:               2,
	}
	resp, raw, err := client.Order().PlaceRefundOrder(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
