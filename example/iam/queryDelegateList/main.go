package main

import (
	"context"
	"fmt"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	iam "github.com/telecom-cloud/telecomcloud-sdk-go/service/iam"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/iam/types/delegate"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://ctiam-global.ctapi.ctyun.cn"
	openapiConfig *config.OpenapiConfig
	options       []iam.Option
	client        iam.ClientSet
	err           error
	ctx           context.Context
	accountId     = ""
	delegateName  = ""
)

func main() {
	ctx = context.Background()
	openapiConfig = &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options = []iam.Option{
		iam.WithClientConfig(openapiConfig),
	}
	client, err = iam.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryDelegateList()
}

func queryDelegateList() {
	req := &delegate.QueryDelegateListRequest{
		Name:      delegateName,
		AccountId: accountId,
		Type:      "3",
	}
	resp, raw, err := client.Delegate().QueryDelegateList(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
