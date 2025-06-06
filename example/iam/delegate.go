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
	baseDomain    = ""
	openapiConfig *config.OpenapiConfig
	options       []iam.Option
	client        iam.ClientSet
	err           error
	ctx           context.Context
	accountId     = ""
	policyId      = ""
	delegateName  = "cceadmintrust"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_AK")
	baseDomain = os.Getenv("CTAPI_DOMAIN")
	accountId = os.Getenv("ACCOUNT_ID")
	policyId = os.Getenv("POLICY_ID")
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
		Name:      "",
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

func switchDelegateRequest() {
	req := &delegate.SwitchDelegateRequest{
		Name:      "cceadmintrust",
		AccountId: accountId,
		ValidTime: 60 * 12,
		SupportS3: true,
	}
	resp, raw, err := client.Delegate().SwitchDelegate(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("resp.AppKey: %s ak: %s \n", resp.AppKey, accessKey)
	sk, err := Decode(resp.AppKey, accessKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ak: %s sk: %s \n", resp.AppId, sk)

	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
