package cdn

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

func (client *Client) DeleteLiveDomainMapping(request *DeleteLiveDomainMappingRequest) (response *DeleteLiveDomainMappingResponse, err error) {
	response = CreateDeleteLiveDomainMappingResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteLiveDomainMappingWithChan(request *DeleteLiveDomainMappingRequest) (<-chan *DeleteLiveDomainMappingResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveDomainMappingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveDomainMapping(request)
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

func (client *Client) DeleteLiveDomainMappingWithCallback(request *DeleteLiveDomainMappingRequest, callback func(response *DeleteLiveDomainMappingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveDomainMappingResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveDomainMapping(request)
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

type DeleteLiveDomainMappingRequest struct {
	*requests.RpcRequest
	PullDomain    string           `position:"Query" name:"PullDomain"`
	PushDomain    string           `position:"Query" name:"PushDomain"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DeleteLiveDomainMappingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteLiveDomainMappingRequest() (request *DeleteLiveDomainMappingRequest) {
	request = &DeleteLiveDomainMappingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveDomainMapping", "", "")
	return
}

func CreateDeleteLiveDomainMappingResponse() (response *DeleteLiveDomainMappingResponse) {
	response = &DeleteLiveDomainMappingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
