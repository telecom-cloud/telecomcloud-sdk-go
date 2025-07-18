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

package als

import (
	"context"
	"fmt"
	"net/http"

	"github.com/telecom-cloud/client-go/pkg/common/config"
	"github.com/telecom-cloud/client-go/pkg/openapi"
	"github.com/telecom-cloud/client-go/pkg/protocol"

	project "github.com/telecom-cloud/telecomcloud-sdk-go/service/als/types/project"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type ProjectClient interface {
	CreateProject(context context.Context, req *project.CreateProjectRequest, reqOpt ...config.RequestOption) (resp *project.CreateProjectResponse, rawResponse *protocol.Response, err error)

	DeleteProject(context context.Context, req *project.DeleteProjectRequest, reqOpt ...config.RequestOption) (resp *project.DeleteProjectResponse, rawResponse *protocol.Response, err error)

	GetProject(context context.Context, req *project.GetProjectRequest, reqOpt ...config.RequestOption) (resp *project.GetProjectResponse, rawResponse *protocol.Response, err error)

	ListProject(context context.Context, req *project.ListProjectRequest, reqOpt ...config.RequestOption) (resp *project.ListProjectResponse, rawResponse *protocol.Response, err error)
}

type projectClient struct {
	client *HttpClient
}

func NewProjectClient(hostUrl string, ops ...Option) (ProjectClient, error) {
	opts := GetOptions(append(ops, WithHostUrl(hostUrl))...)
	cli, err := NewHttpClient(opts)
	if err != nil {
		return nil, err
	}
	return &projectClient{
		client: cli,
	}, nil
}

func (s *projectClient) CreateProject(ctx context.Context, req *project.CreateProjectRequest, reqOpt ...config.RequestOption) (resp *project.CreateProjectResponse, rawResponse *protocol.Response, err error) {
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
		Execute(http.MethodPost, "/v1/project/create")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *projectClient) DeleteProject(ctx context.Context, req *project.DeleteProjectRequest, reqOpt ...config.RequestOption) (resp *project.DeleteProjectResponse, rawResponse *protocol.Response, err error) {
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
		Execute(http.MethodPost, "/v1/project/delete")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *projectClient) GetProject(ctx context.Context, req *project.GetProjectRequest, reqOpt ...config.RequestOption) (resp *project.GetProjectResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	queryParams := map[string]interface{}{
		"projectName": req.GetProjectName(),
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
		Execute(http.MethodGet, "/v1/project/getCodeByName")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *projectClient) ListProject(ctx context.Context, req *project.ListProjectRequest, reqOpt ...config.RequestOption) (resp *project.ListProjectResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.Response{}
	openapiResp.ReturnObj = &resp

	queryParams := map[string]interface{}{
		"projectName": req.GetProjectName(),
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
		Execute(http.MethodGet, "/v1/project/list")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

var defaultProjectClient, _ = NewProjectClient(baseDomain)

func ConfigDefaultProjectClient(ops ...Option) (err error) {
	defaultProjectClient, err = NewProjectClient(baseDomain, ops...)
	return
}

func CreateProject(context context.Context, req *project.CreateProjectRequest, reqOpt ...config.RequestOption) (resp *project.CreateProjectResponse, rawResponse *protocol.Response, err error) {
	return defaultProjectClient.CreateProject(context, req, reqOpt...)
}

func DeleteProject(context context.Context, req *project.DeleteProjectRequest, reqOpt ...config.RequestOption) (resp *project.DeleteProjectResponse, rawResponse *protocol.Response, err error) {
	return defaultProjectClient.DeleteProject(context, req, reqOpt...)
}

func GetProject(context context.Context, req *project.GetProjectRequest, reqOpt ...config.RequestOption) (resp *project.GetProjectResponse, rawResponse *protocol.Response, err error) {
	return defaultProjectClient.GetProject(context, req, reqOpt...)
}

func ListProject(context context.Context, req *project.ListProjectRequest, reqOpt ...config.RequestOption) (resp *project.ListProjectResponse, rawResponse *protocol.Response, err error) {
	return defaultProjectClient.ListProject(context, req, reqOpt...)
}
