package main

import (
	"context"
	"crypto/tls"
	"fmt"
	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/ccse"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/ccse/types/plugin"
	"os"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ccse-global.ctapi.ctyun.cn"
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

	options := []ccse.Option{
		ccse.WithClientConfig(config),
		ccse.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := ccse.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()

	req := &plugin.CheckPluginExistRequest{
		ClusterId:    "09d36e22398a4bad80c9f190275d1182",
		InstanceName: "cube-prometheus",
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
	}

	resp, _, err := client.Plugin().CheckPluginExist(ctx, req)
	if err != nil {
		fmt.Printf("CheckPluginExist err: %v\n", err)
		return
	}

	//fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	fmt.Println(resp.Data)
}
