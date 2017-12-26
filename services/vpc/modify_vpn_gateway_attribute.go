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

func (client *Client) ModifyVpnGatewayAttribute(request *ModifyVpnGatewayAttributeRequest) (response *ModifyVpnGatewayAttributeResponse, err error) {
	response = CreateModifyVpnGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyVpnGatewayAttributeWithChan(request *ModifyVpnGatewayAttributeRequest) (<-chan *ModifyVpnGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVpnGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpnGatewayAttribute(request)
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

func (client *Client) ModifyVpnGatewayAttributeWithCallback(request *ModifyVpnGatewayAttributeRequest, callback func(response *ModifyVpnGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpnGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpnGatewayAttribute(request)
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

type ModifyVpnGatewayAttributeRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	Name                 string           `position:"Query" name:"Name"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyVpnGatewayAttributeResponse struct {
	*responses.BaseResponse
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	VpnGatewayId   string           `json:"VpnGatewayId" xml:"VpnGatewayId"`
	VpcId          string           `json:"VpcId" xml:"VpcId"`
	VSwitchId      string           `json:"VSwitchId" xml:"VSwitchId"`
	InternetIp     string           `json:"InternetIp" xml:"InternetIp"`
	IntranetIp     string           `json:"IntranetIp" xml:"IntranetIp"`
	CreateTime     requests.Integer `json:"CreateTime" xml:"CreateTime"`
	EndTime        requests.Integer `json:"EndTime" xml:"EndTime"`
	Spec           string           `json:"Spec" xml:"Spec"`
	Name           string           `json:"Name" xml:"Name"`
	Description    string           `json:"Description" xml:"Description"`
	Status         string           `json:"Status" xml:"Status"`
	BusinessStatus string           `json:"BusinessStatus" xml:"BusinessStatus"`
}

func CreateModifyVpnGatewayAttributeRequest() (request *ModifyVpnGatewayAttributeRequest) {
	request = &ModifyVpnGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnGatewayAttribute", "", "")
	return
}

func CreateModifyVpnGatewayAttributeResponse() (response *ModifyVpnGatewayAttributeResponse) {
	response = &ModifyVpnGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
