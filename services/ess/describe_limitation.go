package ess

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

func (client *Client) DescribeLimitation(request *DescribeLimitationRequest) (response *DescribeLimitationResponse, err error) {
	response = CreateDescribeLimitationResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLimitationWithChan(request *DescribeLimitationRequest) (<-chan *DescribeLimitationResponse, <-chan error) {
	responseChan := make(chan *DescribeLimitationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLimitation(request)
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

func (client *Client) DescribeLimitationWithCallback(request *DescribeLimitationRequest, callback func(response *DescribeLimitationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLimitationResponse
		var err error
		defer close(result)
		response, err = client.DescribeLimitation(request)
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

type DescribeLimitationRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeLimitationResponse struct {
	*responses.BaseResponse
	MaxNumberOfScalingGroups         requests.Integer `json:"MaxNumberOfScalingGroups" xml:"MaxNumberOfScalingGroups"`
	MaxNumberOfScalingConfigurations requests.Integer `json:"MaxNumberOfScalingConfigurations" xml:"MaxNumberOfScalingConfigurations"`
	MaxNumberOfScalingRules          requests.Integer `json:"MaxNumberOfScalingRules" xml:"MaxNumberOfScalingRules"`
	MaxNumberOfScheduledTasks        requests.Integer `json:"MaxNumberOfScheduledTasks" xml:"MaxNumberOfScheduledTasks"`
	MaxNumberOfScalingInstances      requests.Integer `json:"MaxNumberOfScalingInstances" xml:"MaxNumberOfScalingInstances"`
	MaxNumberOfDBInstances           requests.Integer `json:"MaxNumberOfDBInstances" xml:"MaxNumberOfDBInstances"`
	MaxNumberOfLoadBalancers         requests.Integer `json:"MaxNumberOfLoadBalancers" xml:"MaxNumberOfLoadBalancers"`
	MaxNumberOfMinSize               requests.Integer `json:"MaxNumberOfMinSize" xml:"MaxNumberOfMinSize"`
	MaxNumberOfMaxSize               requests.Integer `json:"MaxNumberOfMaxSize" xml:"MaxNumberOfMaxSize"`
}

func CreateDescribeLimitationRequest() (request *DescribeLimitationRequest) {
	request = &DescribeLimitationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeLimitation", "", "")
	return
}

func CreateDescribeLimitationResponse() (response *DescribeLimitationResponse) {
	response = &DescribeLimitationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
