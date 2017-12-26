package vpc

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

func (client *Client) ModifyEipAddressAttribute(request *ModifyEipAddressAttributeRequest) (response *ModifyEipAddressAttributeResponse, err error) {
	response = CreateModifyEipAddressAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyEipAddressAttributeWithChan(request *ModifyEipAddressAttributeRequest) (<-chan *ModifyEipAddressAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyEipAddressAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEipAddressAttribute(request)
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

func (client *Client) ModifyEipAddressAttributeWithCallback(request *ModifyEipAddressAttributeRequest, callback func(response *ModifyEipAddressAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEipAddressAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyEipAddressAttribute(request)
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

type ModifyEipAddressAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	Name                 string           `position:"Query" name:"Name"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
}

type ModifyEipAddressAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyEipAddressAttributeRequest() (request *ModifyEipAddressAttributeRequest) {
	request = &ModifyEipAddressAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyEipAddressAttribute", "", "")
	return
}

func CreateModifyEipAddressAttributeResponse() (response *ModifyEipAddressAttributeResponse) {
	response = &ModifyEipAddressAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
