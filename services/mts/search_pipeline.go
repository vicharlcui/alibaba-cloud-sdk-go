package mts

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

func (client *Client) SearchPipeline(request *SearchPipelineRequest) (response *SearchPipelineResponse, err error) {
	response = CreateSearchPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SearchPipelineWithChan(request *SearchPipelineRequest) (<-chan *SearchPipelineResponse, <-chan error) {
	responseChan := make(chan *SearchPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchPipeline(request)
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

func (client *Client) SearchPipelineWithCallback(request *SearchPipelineRequest, callback func(response *SearchPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchPipelineResponse
		var err error
		defer close(result)
		response, err = client.SearchPipeline(request)
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

type SearchPipelineRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type SearchPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string           `json:"RequestId" xml:"RequestId"`
	TotalCount   requests.Integer `json:"TotalCount" xml:"TotalCount"`
	PageNumber   requests.Integer `json:"PageNumber" xml:"PageNumber"`
	PageSize     requests.Integer `json:"PageSize" xml:"PageSize"`
	PipelineList struct {
		Pipeline []struct {
			Id           string           `json:"Id" xml:"Id"`
			Name         string           `json:"Name" xml:"Name"`
			State        string           `json:"State" xml:"State"`
			Speed        string           `json:"Speed" xml:"Speed"`
			SpeedLevel   requests.Integer `json:"SpeedLevel" xml:"SpeedLevel"`
			Role         string           `json:"Role" xml:"Role"`
			NotifyConfig struct {
				Topic     string `json:"Topic" xml:"Topic"`
				QueueName string `json:"QueueName" xml:"QueueName"`
			} `json:"NotifyConfig" xml:"NotifyConfig"`
		} `json:"Pipeline" xml:"Pipeline"`
	} `json:"PipelineList" xml:"PipelineList"`
}

func CreateSearchPipelineRequest() (request *SearchPipelineRequest) {
	request = &SearchPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchPipeline", "", "")
	return
}

func CreateSearchPipelineResponse() (response *SearchPipelineResponse) {
	response = &SearchPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
