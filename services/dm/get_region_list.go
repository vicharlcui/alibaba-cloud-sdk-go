package dm

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) GetRegionList(request *GetRegionListRequest) (response *GetRegionListResponse, err error) {
	response = CreateGetRegionListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetRegionListWithChan(request *GetRegionListRequest) (<-chan *GetRegionListResponse, <-chan error) {
	responseChan := make(chan *GetRegionListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRegionList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) GetRegionListWithCallback(request *GetRegionListRequest, callback func(response *GetRegionListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRegionListResponse
		var err error
		defer close(result)
		response, err = client.GetRegionList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type GetRegionListRequest struct {
	*requests.RpcRequest
	Total                string           `position:"Query" name:"Total"`
	PageSize             string           `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
	OffsetCreateTimeDesc string           `position:"Query" name:"OffsetCreateTimeDesc"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Offset               string           `position:"Query" name:"Offset"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OffsetCreateTime     string           `position:"Query" name:"OffsetCreateTime"`
}

type GetRegionListResponse struct {
	*responses.BaseResponse
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Total     requests.Integer `json:"Total" xml:"Total"`
	PageNo    requests.Integer `json:"PageNo" xml:"PageNo"`
	PageSize  requests.Integer `json:"PageSize" xml:"PageSize"`
	Data      struct {
		RegionList []struct {
			Region     string `json:"Region" xml:"Region"`
			RegionDesc string `json:"RegionDesc" xml:"RegionDesc"`
		} `json:"regionList" xml:"regionList"`
	} `json:"data" xml:"data"`
}

func CreateGetRegionListRequest() (request *GetRegionListRequest) {
	request = &GetRegionListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetRegionList", "", "")
	return
}

func CreateGetRegionListResponse() (response *GetRegionListResponse) {
	response = &GetRegionListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
