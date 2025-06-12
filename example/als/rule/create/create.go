package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/telecom-cloud/client-go/pkg/client"
	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/als/types/rule"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://ctlts-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []als.Option{
		als.WithClientConfig(config),
		als.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}

	client, err := als.NewClientSet(baseDomain, options...)

	if err != nil {
		fmt.Printf("NewClientSet err: %v\n", err)
		return
	}

	ctx := context.Background()
	req := &rule.CreateRuleRequest{
		RuleName:      "sdkTesting",
		UnitCode:      "6e4512c3d798b0",
		CollectPolicy: "all",
		CuttingMode:   "4",
		AccessType:    1,
		ExtractMode:   3,
		LogPaths:      []string{"stdout"},
		Enable:        true,
		RuleConfig: &rule.CollectRuleConfig{
			MaxPathDepth: 2,
			Containers: &rule.ContainerCollectRule{
				K8SNamespaceRegex: "pod",
				IncludeK8SLabel: map[string]string{
					"app": "test",
				},
			},
		},
		RegionId: "bb9fdb42056f11eda1610242ac110002",
	}

	resp, raw, err := client.Rule().CreateRule(ctx, req)
	if err != nil {
		fmt.Printf("CreateRule err: %v\n", err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
