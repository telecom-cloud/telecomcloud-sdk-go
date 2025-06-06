// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Telecom Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package crs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/telecom-cloud/client-go/pkg/common/config"
	"github.com/telecom-cloud/client-go/pkg/openapi"
	"github.com/telecom-cloud/client-go/pkg/protocol"

	namespace "github.com/telecom-cloud/telecomcloud-sdk-go/service/crs/types/namespace"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type NamespaceClient interface {
	CreateNamespace(context context.Context, req *namespace.CreateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.CreateNamespaceResponse, rawResponse *protocol.Response, err error)

	ListNamespace(context context.Context, req *namespace.ListNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.ListNamespaceResponse, rawResponse *protocol.Response, err error)

	GetNamespace(context context.Context, req *namespace.GetNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.GetNamespaceResponse, rawResponse *protocol.Response, err error)

	UpdateNamespace(context context.Context, req *namespace.UpdateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.UpdateNamespaceResponse, rawResponse *protocol.Response, err error)
}

type namespaceClient struct {
	client *HttpClient
}

func NewNamespaceClient(hostUrl string, ops ...Option) (NamespaceClient, error) {
	opts := GetOptions(append(ops, WithHostUrl(hostUrl))...)
	cli, err := NewHttpClient(opts)
	if err != nil {
		return nil, err
	}
	return &namespaceClient{
		client: cli,
	}, nil
}

func (s *namespaceClient) CreateNamespace(ctx context.Context, req *namespace.CreateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.CreateNamespaceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	ret, err := s.client.R().
		SetContext(ctx).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/v1/createNamespace")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *namespaceClient) ListNamespace(ctx context.Context, req *namespace.ListNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.ListNamespaceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	queryParams := map[string]interface{}{
		"instanceId":    req.GetInstanceId(),
		"namespaceName": req.GetNamespaceName(),
		"pageNum":       req.GetPageNum(),
		"pageSize":      req.GetPageSize(),
		"orderBy":       req.GetOrderBy(),
		"order":         req.GetOrder(),
	}
	OptimizeQueryParams(queryParams)
	ret, err := s.client.R().
		SetContext(ctx).
		SetQueryParams(queryParams).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodGet, "/v1/listNamespace")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *namespaceClient) GetNamespace(ctx context.Context, req *namespace.GetNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.GetNamespaceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	queryParams := map[string]interface{}{
		"instanceId":    req.GetInstanceId(),
		"namespaceName": req.GetNamespaceName(),
	}
	OptimizeQueryParams(queryParams)
	ret, err := s.client.R().
		SetContext(ctx).
		SetQueryParams(queryParams).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodGet, "/v1/getNamespace")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *namespaceClient) UpdateNamespace(ctx context.Context, req *namespace.UpdateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.UpdateNamespaceResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	ret, err := s.client.R().
		SetContext(ctx).
		AddHeaders(map[string]string{
			"regionId": req.GetRegionId(),
		}).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/v1/updateNamespace")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

var defaultNamespaceClient, _ = NewNamespaceClient(baseDomain)

func ConfigDefaultNamespaceClient(ops ...Option) (err error) {
	defaultNamespaceClient, err = NewNamespaceClient(baseDomain, ops...)
	return
}

func CreateNamespace(context context.Context, req *namespace.CreateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.CreateNamespaceResponse, rawResponse *protocol.Response, err error) {
	return defaultNamespaceClient.CreateNamespace(context, req, reqOpt...)
}

func ListNamespace(context context.Context, req *namespace.ListNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.ListNamespaceResponse, rawResponse *protocol.Response, err error) {
	return defaultNamespaceClient.ListNamespace(context, req, reqOpt...)
}

func GetNamespace(context context.Context, req *namespace.GetNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.GetNamespaceResponse, rawResponse *protocol.Response, err error) {
	return defaultNamespaceClient.GetNamespace(context, req, reqOpt...)
}

func UpdateNamespace(context context.Context, req *namespace.UpdateNamespaceRequest, reqOpt ...config.RequestOption) (resp *namespace.UpdateNamespaceResponse, rawResponse *protocol.Response, err error) {
	return defaultNamespaceClient.UpdateNamespace(context, req, reqOpt...)
}
