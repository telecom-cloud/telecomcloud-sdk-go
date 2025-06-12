package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	crs "github.com/telecom-cloud/telecomcloud-sdk-go/service/crs"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/instance"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = ""
	regionId      = "" // 上海15
	openapiConfig *config.OpenapiConfig
	options       []crs.Option
	client        crs.ClientSet
	err           error
	ctx           context.Context
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

	options = []crs.Option{
		crs.WithClientConfig(openapiConfig),
	}
	client, err = crs.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	subscribeEnterpriseInstance()
}

func subscribeEnterpriseInstance() {
	req := &instance.SubscribeEnterpriseInstanceRequest{
		RegionId:       regionId,
		InstanceName:   "",
		InstanceType:   "enterprise-basic",
		CrUserName:     "",
		CrUserPassword: "",
		OssBucket:      "",
		ReqType:        "2",
		AutoPay:        true,
		BillMode:       "1",
		CycleType:      "3",
		CycleCnt:       3,
	}

	resp, raw, err := client.Instance().SubscribeEnterpriseInstance(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
