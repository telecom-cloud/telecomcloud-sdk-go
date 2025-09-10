package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it/types/common"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/it/types/price"
	"github.com/telecom-cloud/telecomcloud-sdk-go/test/middleware"
)

var (
	accountId  = ""
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://apiproxy-global.ctapi-internal.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

// FIXME 运行测试用例会报proto注册冲突，添加环境变量 GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
func main() {
	ctx := context.Background()
	openapiConfig := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []it.Option{
		it.WithClientConfig(openapiConfig),
		it.WithClientMiddleware(middleware.DumpHttpMiddleware),
	}
	client, err := it.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	//queryEciPremiumPrice(ctx, client)
	fmt.Println(strings.Repeat("-", 50))
	queryEciEBSPrice(ctx, client)
	fmt.Println(strings.Repeat("-", 50))
	queryEciBasicPrice(ctx, client)
}

const (
	IdentifiedByAccountId = 2
	DefaultCount          = 1
	SSDVolumeType         = "SSD"
	CycleCntByHour        = 1
)

const (
	ServiceTag     = "PAAS"
	PaasVersion    = "v1"
	PaasECIEdition = "basic"
	BillMode       = "2" // 按需
)

const (
	CPUType    = "cpu"
	MemoryType = "mem"
)

const (
	PAAS_ECI_BASIC   = "PAAS_ECI_BASIC"
	PAAS_ECI_PREMIUM = "PAAS_ECI_PREMIUM"
	PAAS_ECI_EBS     = "PAAS_ECI_EBS"
	PAAS_ECI         = "PAAS_ECI"
)

func queryEciPremiumPrice(ctx context.Context, client it.ClientSet) {
	detail := &price.QueryPriceRequestDetail{
		Orders: []*price.OrderDetail{
			{
				InstanceCnt: DefaultCount,
				CycleCnt:    CycleCntByHour,
				Items: []*price.Item{
					{
						Master:       true,
						ResourceType: PAAS_ECI,
						ServiceTag:   ServiceTag,
						ItemValue:    DefaultCount,
						ItemConfig: &price.ItemConfig{
							Version:  PaasVersion,
							Edition:  PaasECIEdition,
							BillMode: BillMode,
						},
					},
					{
						Master:       false,
						ResourceType: PAAS_ECI_PREMIUM,
						ServiceTag:   ServiceTag,
						ItemConfig: &price.ItemConfig{
							// 指定规格
							CpuNum:  fmt.Sprintf("%g", 1),
							MemSize: fmt.Sprintf("%g", 2),
							// HostType: 取 flavorName 前缀：pi7.4xlarge.4 中的 pi7
							HostType: "pi7",
							Version:  PaasVersion,
						},
						ItemValue: DefaultCount,
					},
				},
			},
		},
		CustomInfo: &common.CustomInfo{
			Type: IdentifiedByAccountId,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
			},
		},
	}

	data, _ := json.Marshal(detail)
	req := &price.QueryPriceRequest{
		OrderDetailJson: string(data),
	}

	resp, raw, err := client.Price().QueryPrice(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}

func queryEciEBSPrice(ctx context.Context, client it.ClientSet) {
	detail := &price.QueryPriceRequestDetail{
		Orders: []*price.OrderDetail{
			{
				InstanceCnt: DefaultCount,
				CycleCnt:    CycleCntByHour,
				Items: []*price.Item{
					{
						Master:       true,
						ResourceType: PAAS_ECI,
						ServiceTag:   ServiceTag,
						ItemValue:    DefaultCount,
						ItemConfig: &price.ItemConfig{
							Version:  PaasVersion,
							Edition:  PaasECIEdition,
							BillMode: BillMode,
						},
					},
					{
						Master:       false,
						ResourceType: PAAS_ECI_EBS,
						ServiceTag:   ServiceTag,
						ItemValue:    DefaultCount,
						ItemConfig: &price.ItemConfig{
							Version:    PaasVersion,
							VolumeType: SSDVolumeType,
						},
					},
				},
			},
		},
		CustomInfo: &common.CustomInfo{
			Type: IdentifiedByAccountId,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
			},
		},
	}

	data, _ := json.Marshal(detail)
	req := &price.QueryPriceRequest{
		OrderDetailJson: string(data),
	}

	resp, raw, err := client.Price().QueryPrice(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}

func queryEciBasicPrice(ctx context.Context, client it.ClientSet) {
	detail := &price.QueryPriceRequestDetail{
		Orders: []*price.OrderDetail{
			{
				InstanceCnt: DefaultCount,
				CycleCnt:    CycleCntByHour,
				Items: []*price.Item{
					{
						Master:       true,
						ResourceType: PAAS_ECI,
						ServiceTag:   ServiceTag,
						ItemValue:    DefaultCount,
						ItemConfig: &price.ItemConfig{
							Version:  PaasVersion,
							Edition:  PaasECIEdition,
							BillMode: BillMode,
						},
					},
					{
						Master:       false,
						ResourceType: PAAS_ECI_BASIC,
						ServiceTag:   ServiceTag,
						ItemConfig: &price.ItemConfig{
							Type:    CPUType,
							Version: PaasVersion,
						},
						ItemValue: DefaultCount,
					},
					{
						Master:       false,
						ResourceType: PAAS_ECI_BASIC,
						ServiceTag:   ServiceTag,
						ItemConfig: &price.ItemConfig{
							Type:    MemoryType,
							Version: PaasVersion,
						},
						ItemValue: DefaultCount,
					},
				},
			},
		},
		CustomInfo: &common.CustomInfo{
			Type: IdentifiedByAccountId,
			Identity: &common.TenantIdentity{
				AccountId: accountId,
			},
		},
	}

	data, _ := json.Marshal(detail)
	req := &price.QueryPriceRequest{
		OrderDetailJson: string(data),
	}

	resp, raw, err := client.Price().QueryPrice(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp:  %v\n", string(raw.Body()), resp)
}
