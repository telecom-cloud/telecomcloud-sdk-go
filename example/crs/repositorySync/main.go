package main

import (
	"context"
	"fmt"
	"os"

	"github.com/telecom-cloud/client-go/pkg/openapi/config"
	crs "github.com/telecom-cloud/telecomcloud-sdk-go/service/crs"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/instance"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/namespace"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/repository"
	"github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/repositorysync"
)

var (
	accessKey     = ""
	secretKey     = ""
	baseDomain    = "https://crs-global.ctapi.ctyun.cn"
	regionId      = ""
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

	// listEnterpriseInstance()
	// listEnterpriseNamespace()
	// listEnterpriseRepository()
	// createSyncTask()
	getSyncTask()
}

func listEnterpriseInstance() {
	req := &instance.ListEnterpriseInstanceRequest{
		RegionId: regionId,
		PageNum:  1,
		PageSize: 100,
	}
	resp, raw, err := client.Instance().ListEnterpriseInstance(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func listEnterpriseNamespace() {
	req := &namespace.ListNamespaceRequest{
		RegionId:   regionId,
		InstanceId: 0,
		PageNum:    1,
		PageSize:   100,
	}
	resp, raw, err := client.Namespace().ListNamespace(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func listEnterpriseRepository() {
	req := &repository.ListRepositoryRequest{
		RegionId:      regionId,
		InstanceId:    0,
		NamespaceName: "test",
		PageNum:       1,
		PageSize:      100,
	}
	resp, raw, err := client.Repository().ListRepository(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func createSyncTask() {
	req := &repositorysync.CreateRepositorySyncTaskRequest{
		RegionId:        regionId,
		SrcInstanceId:   "",
		SrcNamespaceId:  "",
		SrcRepositoryId: "",
		DstInstanceId:   "",
		DstNamespaceId:  "",
	}
	resp, raw, err := client.RepositorySyncTask().CreateRepositorySyncTask(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}

func getSyncTask() {
	req := &repositorysync.GetRepositorySyncTaskRequest{
		RegionId: regionId,
		TaskId:   "",
	}
	resp, raw, err := client.RepositorySyncTask().GetRepositorySyncTask(ctx, req)
	if err != nil {
		fmt.Printf("failed %s", err.Error())
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
