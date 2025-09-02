package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cloudaudit"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/cloudaudit/types/manager"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://cloudaudit-global.ctapi.ctyun.cn"
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
		AccessKey: "",
		SecretKey: "",
	}

	options := []cloudaudit.Option{
		cloudaudit.WithClientConfig(config),
		cloudaudit.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := cloudaudit.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()
	req := &manager.ManagerServiceRequest{
		RegionId:  "",
		AccountId: "",
		UserId:    "",
	}

	resp, raw, err := client.Manager().CheckServiceStatus(ctx, req)
	if err != nil {
		fmt.Printf("CheckService err: %v\n", err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
	fmt.Println(resp.Data)
}
