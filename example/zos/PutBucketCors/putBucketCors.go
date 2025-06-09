package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/zos"
	zostypes "github.com/telecom-cloud/telecomcloud-sdk-go/service/zos/types/zos"
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

	options := []zos.Option{
		zos.WithClientConfig(config),
		zos.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := zos.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &zostypes.PutBucketCorsRequest{
		RegionID: "b342b77ef26b11ecb0ac0242ac110002",
		Bucket:   "bucket-yyds",
		CORSConfiguration: &zostypes.CORSConfiguration{
			CORSRules: []*zostypes.CORSRule{
				{
					AllowedHeaders: []string{
						"*",
					},
					AllowedMethods: []string{
						"GET",
						"PUT",
						"POST",
						"DELETE",
						"HEAD",
					},
					AllowedOrigins: []string{
						"https://ccse.ctyun.cn",
					},
					ExposeHeaders: []string{
						"*",
					},
					MaxAgeSeconds: 3000,
				},
			},
		},
	}
	resp, raw, err := client.Zos().PutBucketCors(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
