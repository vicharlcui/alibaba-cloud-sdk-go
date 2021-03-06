package bssopenapi

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

// QueryInstanceGaapCost invokes the bssopenapi.QueryInstanceGaapCost API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancegaapcost.html
func (client *Client) QueryInstanceGaapCost(request *QueryInstanceGaapCostRequest) (response *QueryInstanceGaapCostResponse, err error) {
	response = CreateQueryInstanceGaapCostResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceGaapCostWithChan invokes the bssopenapi.QueryInstanceGaapCost API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancegaapcost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInstanceGaapCostWithChan(request *QueryInstanceGaapCostRequest) (<-chan *QueryInstanceGaapCostResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceGaapCostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstanceGaapCost(request)
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

// QueryInstanceGaapCostWithCallback invokes the bssopenapi.QueryInstanceGaapCost API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancegaapcost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInstanceGaapCostWithCallback(request *QueryInstanceGaapCostRequest, callback func(response *QueryInstanceGaapCostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceGaapCostResponse
		var err error
		defer close(result)
		response, err = client.QueryInstanceGaapCost(request)
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

// QueryInstanceGaapCostRequest is the request struct for api QueryInstanceGaapCost
type QueryInstanceGaapCostRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	BillingCycle     string           `position:"Query" name:"BillingCycle"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	ProductType      string           `position:"Query" name:"ProductType"`
}

// QueryInstanceGaapCostResponse is the response struct for api QueryInstanceGaapCost
type QueryInstanceGaapCostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryInstanceGaapCostRequest creates a request to invoke QueryInstanceGaapCost API
func CreateQueryInstanceGaapCostRequest() (request *QueryInstanceGaapCostRequest) {
	request = &QueryInstanceGaapCostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryInstanceGaapCost", "", "")
	return
}

// CreateQueryInstanceGaapCostResponse creates a response to parse from QueryInstanceGaapCost response
func CreateQueryInstanceGaapCostResponse() (response *QueryInstanceGaapCostResponse) {
	response = &QueryInstanceGaapCostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
