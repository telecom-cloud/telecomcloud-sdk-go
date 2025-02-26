package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/eci"
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
	req := &imagecache.DescribeImageCacheRequest{
		ImageCacheId: "imc-ej5ow9e4sp0igevf",
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
	}

	resp, raw, err := client.ImageCache().DescribeImageCache(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
	/**
	raw: {"statusCode":200,"returnObj":{"requestId":"a9bdc9f8-dd22-4cf1-b104-15951b4b1ca9","imageCaches":[{"imageCacheId":"imc-ej5ow9e4sp0igevf","imageCacheName":"sdk-demo-nginx","imageCacheSize":10,"images":["registry-huadong1.crs-internal.ctyun.cn/open-source/nginx:1.25-alpine"],"status":"ready","creationTime":"2025-01-16T09:26:44+08:00","azName":"cn-huadong1-jsnj1A-public-ctcloud"}],"total":1}}
	resp: requestId:"a9bdc9f8-dd22-4cf1-b104-15951b4b1ca9"  imageCaches:{imageCacheId:"imc-ej5ow9e4sp0igevf"  imageCacheName:"sdk-demo-nginx"  imageCacheSize:10  images:"registry-huadong1.crs-internal.ctyun.cn/open-source/nginx:1.25-alpine"  status:"ready"  creationTime:"2025-01-16T09:26:44+08:00"  azName:"cn-huadong1-jsnj1A-public-ctcloud"}  total:1
	*/
}
