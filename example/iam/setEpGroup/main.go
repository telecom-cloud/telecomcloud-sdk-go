package main

import (
	"context"
	"fmt"
	"log"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/client-go/pkg/protocol"
	iam "github.com/telecom-cloud/telecomcloud-sdk-go/service/iam"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/iam/types/delegate"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://ctiam-global.ctapi-internal.ctyun.cn"
	openapiConfig *config.OpenapiConfig
	options       []iam.Option
	client        iam.ClientSet
	err           error
	ctx           context.Context
	accountId     = ""
	policyId      = ""
)

func main() {
	ctx = context.Background()
	openapiConfig = &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options = []iam.Option{
		iam.WithClientConfig(openapiConfig),
		iam.WithClientMiddleware(dumpHttpMiddleware),
	}
	client, err = iam.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	setEpGroup()
}

func setEpGroup() {
	req := &delegate.SetEpGroupRequest{
		Id:        "",
		AccountId: accountId,
		PloyIds:   policyId,
		Value:     "CUS_153_01_0002",
		ProjectId: "0",
	}
	resp, raw, err := client.Delegate().SetEpGroup(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
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
