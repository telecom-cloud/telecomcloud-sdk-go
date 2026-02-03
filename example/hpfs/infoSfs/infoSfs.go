package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	hpfsService "github.com/telecom-cloud/telecomcloud-sdk-go/service/hpfs"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/hpfs/types/hpfs"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://hpfs-global.ctapi.ctyun.cn"
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

	options := []hpfsService.Option{
		hpfsService.WithClientConfig(config),
	}
	client, err := hpfsService.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	req := &hpfs.InfoSfsRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
		SfsUID:  "577993b4-e91f-5b36-9d90-32e7e6af5860",
	}
	resp, raw, err := client.Hpfs().InfoSfs(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
