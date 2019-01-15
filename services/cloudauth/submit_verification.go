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

// SubmitVerification invokes the cloudauth.SubmitVerification API synchronously
// api document: https://help.aliyun.com/api/cloudauth/submitverification.html
func (client *Client) SubmitVerification(request *SubmitVerificationRequest) (response *SubmitVerificationResponse, err error) {
	response = CreateSubmitVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitVerificationWithChan invokes the cloudauth.SubmitVerification API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/submitverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitVerificationWithChan(request *SubmitVerificationRequest) (<-chan *SubmitVerificationResponse, <-chan error) {
	responseChan := make(chan *SubmitVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitVerification(request)
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

// SubmitVerificationWithCallback invokes the cloudauth.SubmitVerification API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/submitverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitVerificationWithCallback(request *SubmitVerificationRequest, callback func(response *SubmitVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitVerificationResponse
		var err error
		defer close(result)
		response, err = client.SubmitVerification(request)
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

// SubmitVerificationRequest is the request struct for api SubmitVerification
type SubmitVerificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	Biz             string                        `position:"Body" name:"Biz"`
	SourceIp        string                        `position:"Query" name:"SourceIp"`
	Material        *[]SubmitVerificationMaterial `position:"Body" name:"Material"  type:"Repeated"`
	TicketId        string                        `position:"Body" name:"TicketId"`
}

// SubmitVerificationMaterial is a repeated param struct in SubmitVerificationRequest
type SubmitVerificationMaterial struct {
	MaterialType string `name:"MaterialType"`
	Value        string `name:"Value"`
}

// SubmitVerificationResponse is the response struct for api SubmitVerification
type SubmitVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitVerificationRequest creates a request to invoke SubmitVerification API
func CreateSubmitVerificationRequest() (request *SubmitVerificationRequest) {
	request = &SubmitVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "SubmitVerification", "CloudAuth", "openAPI")
	return
}

// CreateSubmitVerificationResponse creates a response to parse from SubmitVerification response
func CreateSubmitVerificationResponse() (response *SubmitVerificationResponse) {
	response = &SubmitVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
