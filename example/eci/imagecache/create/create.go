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
	req := &imagecache.CreateImageCacheRequest{
		ImageCacheName:  "sdk-demo-nginx",
		ImageCacheSize:  10,
		Images:          []string{"registry-huadong1.crs-internal.ctyun.cn/open-source/nginx:1.25-alpine"},
		VpcId:           "vpc-mv2e552mn4",
		VSwitchId:       "subnet-vlvv43t8i3",
		SecurityGroupId: "sg-qm3fz92p5p",
		AzInfo: &containergroup.AzInfo{
			AzId:   1,
			AzName: "cn-huadong1-jsnj1A-public-ctcloud",
		},
		RetentionDays: 0,
		RegionId:      "bb9fdb42056f11eda1610242ac110002",
	}
	resp, raw, err := client.ImageCache().CreateImageCache(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	/**
	raw: {"statusCode":200,"returnObj":{"requestId":"13e0ec99-5a1c-4cc7-a5e2-34cdcdc623fa","imageCacheId":"imc-ej5ow9e4sp0igevf","imageCacheName":"sdk-demo-nginx","imageCacheSize":10,"images":["registry-huadong1.crs-internal.ctyun.cn/open-source/nginx:1.25-alpine"]}}
	resp: requestId:"13e0ec99-5a1c-4cc7-a5e2-34cdcdc623fa" imageCacheId:"imc-ej5ow9e4sp0igevf" imageCacheName:"sdk-demo-nginx" imageCacheSize:10 images:"registry-huadong1.crs-internal.ctyun.cn/open-source/nginx:1.25-alpine"
	*/
}
