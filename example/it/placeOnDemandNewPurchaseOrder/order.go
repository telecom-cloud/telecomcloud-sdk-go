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

	orderDetail := Order{
		AutoPay: true,
		Source:  "8",
		Orders: []OrderItem{
			{
				CycleType:   101,
				InstanceCnt: 1,
				Items: []Item{
					{
						Master:       true,
						ResourceType: "PAAS_ECI",
						ServiceTag:   "PAAS",
						ItemConfig: map[string]interface{}{
							"version":  "v1",
							"edition":  "basic",
							"billMode": "2",
							"regionId": "b342b77ef26b11ecb0ac0242ac110002",
							"azInfo": []map[string]string{
								{
									"azName": "cn-xinan1-1A",
								},
							},
							"name": "eci-oniy5pupj7urjqnx",
							"extJson": map[string]interface{}{
								"busiChannel": "010",
								"clusterName": "eci-oniy5pupj7urjqnx",
								"envTag":      "198dev",
								"prodId":      "12710101",
								"attrMap": map[string]string{
									"cpu":               "1",
									"memory":            "1",
									"restartPolicy":     "Always",
									"instancePayAmount": "1",
									"vpcUuid":           "vpc-k1w1mwfxli",
									"subnetUuid":        "subnet-n4wenbqywe",
									"securityGroupUuid": "sg-x2cc0bc9ey",
								},
							},
						},
						ItemValue: 1,
					},
					{
						Master:       false,
						ResourceType: "PAAS_ECI_PREMIUM",
						ServiceTag:   "PAAS",
						ItemConfig: map[string]interface{}{
							"version":  "v1",
							"cpuNum":   "1",
							"memSize":  "1",
							"hostType": "s7",
						},
						ItemValue: 1,
					},
				},
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
	req := &order.PlaceOnDemandNewPurchaseOrderRequest{
		// Fill in the request parameters
		OrderDetailJson: orderDetail.Marshal(),
	}
	resp, raw, err := client.Order().PlaceOnDemandNewPurchaseOrder(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
