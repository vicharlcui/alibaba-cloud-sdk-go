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

func (client *Client) AddMediaWorkflow(request *AddMediaWorkflowRequest) (response *AddMediaWorkflowResponse, err error) {
	response = CreateAddMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddMediaWorkflowWithChan(request *AddMediaWorkflowRequest) (<-chan *AddMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *AddMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMediaWorkflow(request)
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

func (client *Client) AddMediaWorkflowWithCallback(request *AddMediaWorkflowRequest, callback func(response *AddMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.AddMediaWorkflow(request)
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

type AddMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	Topology             string           `position:"Query" name:"Topology"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type AddMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	MediaWorkflow struct {
		MediaWorkflowId string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
		Name            string `json:"Name" xml:"Name"`
		Topology        string `json:"Topology" xml:"Topology"`
		State           string `json:"State" xml:"State"`
		CreationTime    string `json:"CreationTime" xml:"CreationTime"`
	} `json:"MediaWorkflow" xml:"MediaWorkflow"`
}

func CreateAddMediaWorkflowRequest() (request *AddMediaWorkflowRequest) {
	request = &AddMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddMediaWorkflow", "", "")
	return
}

func CreateAddMediaWorkflowResponse() (response *AddMediaWorkflowResponse) {
	response = &AddMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
