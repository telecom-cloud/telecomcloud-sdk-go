package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cceone"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cceone/types/cluster"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://cceone-global.ctapi.ctyun.cn"
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

	options := []cceone.Option{
		cceone.WithClientConfig(config),
		cceone.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := cceone.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()

	req := &cluster.ListClustersPageRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		PageNow:  1,
		PageSize: 100,
	}

	resp, _, err := client.Cluster().ListClustersPage(ctx, req)
	if err != nil {
		fmt.Printf("List cluster page err: %v\n", err)
		return
	}

	//fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	fmt.Println(resp.Records)
}
