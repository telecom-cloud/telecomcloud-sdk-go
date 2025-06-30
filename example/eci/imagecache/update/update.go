package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/containergroup"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci/types/imagecache"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://eci-global.ctapi.ctyun.cn"
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

	options := []eci.Option{
		eci.WithClientConfig(config),
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &imagecache.UpdateImageCacheRequest{
		ImageCacheId: "imc-ej5ow9e4sp0igevf",
		Tags: []*containergroup.Tag{
			{
				Key:   "sdk-demo-key-2",
				Value: "sdk-demo-value-3",
			},
		},
		RegionId: "bb9fdb42056f11eda1610242ac110002",
	}

	resp, raw, err := client.ImageCache().UpdateImageCache(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	/**
	raw: {"statusCode":200,"returnObj":{"requestId":"41929e5a-8b05-4783-a434-547e4584c3be"}}
	resp: requestId:"41929e5a-8b05-4783-a434-547e4584c3be"
	*/
}
