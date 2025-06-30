package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	zosService "github.com/telecom-cloud/telecomcloud-sdk-go/service/zos"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/zos/types/zos"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://zos-global.ctapi.ctyun.cn"
	regionId   = ""

	err           error
	ctx           context.Context
	openapiConfig *config.OpenapiConfig
	options       []zosService.Option
	client        zosService.ClientSet
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_CRS_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	ctx = context.Background()
	openapiConfig = &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options = []zosService.Option{
		zosService.WithClientConfig(openapiConfig),
	}
	client, err = zosService.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	getOssServiceStatus()
}

func getOssServiceStatus() {
	req := &zos.GetOssServiceStatusRequest{
		RegionId: regionId,
	}
	resp, raw, err := client.Zos().GetOssServiceStatus(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
