package emr

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

// DeleteClusterScript invokes the emr.DeleteClusterScript API synchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterscript.html
func (client *Client) DeleteClusterScript(request *DeleteClusterScriptRequest) (response *DeleteClusterScriptResponse, err error) {
	response = CreateDeleteClusterScriptResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteClusterScriptWithChan invokes the emr.DeleteClusterScript API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterScriptWithChan(request *DeleteClusterScriptRequest) (<-chan *DeleteClusterScriptResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteClusterScript(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteClusterScriptWithCallback invokes the emr.DeleteClusterScript API asynchronously
// api document: https://help.aliyun.com/api/emr/deleteclusterscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterScriptWithCallback(request *DeleteClusterScriptRequest, callback func(response *DeleteClusterScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterScriptResponse
		var err error
		defer close(result)
		response, err = client.DeleteClusterScript(request)
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

// DeleteClusterScriptRequest is the request struct for api DeleteClusterScript
type DeleteClusterScriptRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Id              string           `position:"Query" name:"Id"`
}

// DeleteClusterScriptResponse is the response struct for api DeleteClusterScript
type DeleteClusterScriptResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteClusterScriptRequest creates a request to invoke DeleteClusterScript API
func CreateDeleteClusterScriptRequest() (request *DeleteClusterScriptRequest) {
	request = &DeleteClusterScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteClusterScript", "emr", "openAPI")
	return
}

// CreateDeleteClusterScriptResponse creates a response to parse from DeleteClusterScript response
func CreateDeleteClusterScriptResponse() (response *DeleteClusterScriptResponse) {
	response = &DeleteClusterScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
