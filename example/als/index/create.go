package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als/types/index"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ctlts-global.ctapi.ctyun.cn"
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
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []als.Option{
		als.WithClientConfig(config),
		als.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := als.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()
	req := &index.CreateIndexRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		ProjectCode: "c47dc9e57bc1fd",
		UnitCode:    "df5b2b7dfcfb14",
		//FullTextIndex: false,
		Fields: []*index.AlsIndexField{
			{
				FieldName: "__tag__podName",
				DataType:  "text",
			},
		},
	}

	resp, raw, err := client.Index().CreateIndex(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
