package ecs

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

func (client *Client) RebootInstance(request *RebootInstanceRequest) (response *RebootInstanceResponse, err error) {
	response = CreateRebootInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RebootInstanceWithChan(request *RebootInstanceRequest) (<-chan *RebootInstanceResponse, <-chan error) {
	responseChan := make(chan *RebootInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebootInstance(request)
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

func (client *Client) RebootInstanceWithCallback(request *RebootInstanceRequest, callback func(response *RebootInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebootInstanceResponse
		var err error
		defer close(result)
		response, err = client.RebootInstance(request)
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

type RebootInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ForceStop            requests.Boolean `position:"Query" name:"ForceStop"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

type RebootInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateRebootInstanceRequest() (request *RebootInstanceRequest) {
	request = &RebootInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RebootInstance", "", "")
	return
}

func CreateRebootInstanceResponse() (response *RebootInstanceResponse) {
	response = &RebootInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
