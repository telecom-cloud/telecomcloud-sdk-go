package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
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
	delegateName  = ""
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

	options = []iam.Option{
		iam.WithClientConfig(openapiConfig),
	}
	client, err = iam.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	createDelegateRole()
}

func createDelegateRole() {
	req := &delegate.CreateAutomateDelegateRoleRequest{
		Name:      delegateName,
		AccountId: accountId,
		PolicyIds: []string{policyId},
		RangeType: "GLOBAL_SERVICE",
		Remark:    "create by openapi",
	}
	resp, raw, err := client.Delegate().CreateAutomateDelegateRole(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
