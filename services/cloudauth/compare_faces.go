package cloudauth

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

// CompareFaces invokes the cloudauth.CompareFaces API synchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaces.html
func (client *Client) CompareFaces(request *CompareFacesRequest) (response *CompareFacesResponse, err error) {
	response = CreateCompareFacesResponse()
	err = client.DoAction(request, response)
	return
}

// CompareFacesWithChan invokes the cloudauth.CompareFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompareFacesWithChan(request *CompareFacesRequest) (<-chan *CompareFacesResponse, <-chan error) {
	responseChan := make(chan *CompareFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompareFaces(request)
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

// CompareFacesWithCallback invokes the cloudauth.CompareFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompareFacesWithCallback(request *CompareFacesRequest, callback func(response *CompareFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompareFacesResponse
		var err error
		defer close(result)
		response, err = client.CompareFaces(request)
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

// CompareFacesRequest is the request struct for api CompareFaces
type CompareFacesRequest struct {
	*requests.RpcRequest
	SourceImageType  string           `position:"Body" name:"SourceImageType"`
	ResourceOwnerId  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	TargetImageType  string           `position:"Body" name:"TargetImageType"`
	SourceImageValue string           `position:"Body" name:"SourceImageValue"`
	TargetImageValue string           `position:"Body" name:"TargetImageValue"`
}

// CompareFacesResponse is the response struct for api CompareFaces
type CompareFacesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCompareFacesRequest creates a request to invoke CompareFaces API
func CreateCompareFacesRequest() (request *CompareFacesRequest) {
	request = &CompareFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "CompareFaces", "CloudAuth", "openAPI")
	return
}

// CreateCompareFacesResponse creates a response to parse from CompareFaces response
func CreateCompareFacesResponse() (response *CompareFacesResponse) {
	response = &CompareFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
