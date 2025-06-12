package main

import (
	"context"
	"fmt"
	"log"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
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
		zosService.WithClientMiddleware(dumpHttpMiddleware),
	}
	client, err = zosService.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	getBucketAcl()

}

func getBucketAcl() {
	req := &zos.GetBucketAclRequest{
		RegionID: regionId,
		Bucket:   "openapi-0610",
	}
	resp, raw, err := client.Zos().GetBucketAcl(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func dumpHttpMiddleware(next cli.Endpoint) cli.Endpoint {
	return func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error) {
		log.Printf("proxy-request  method %s path %s header %s body %s", req.Method(), string(req.Host())+string(req.Path()), req.Header.String(), string(req.Body()))
		err = next(ctx, req, resp)
		if err != nil {
			return err
		}
		log.Printf("proxy-response status %d  header %s body %s", resp.StatusCode(), string(resp.Header.Header()), string(resp.Body()))
		return nil
	}
}
