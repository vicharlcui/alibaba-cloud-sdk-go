package dm

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

func (client *Client) DescAccountSummary(request *DescAccountSummaryRequest) (response *DescAccountSummaryResponse, err error) {
	response = CreateDescAccountSummaryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescAccountSummaryWithChan(request *DescAccountSummaryRequest) (<-chan *DescAccountSummaryResponse, <-chan error) {
	responseChan := make(chan *DescAccountSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescAccountSummary(request)
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

func (client *Client) DescAccountSummaryWithCallback(request *DescAccountSummaryRequest, callback func(response *DescAccountSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescAccountSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescAccountSummary(request)
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

type DescAccountSummaryRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescAccountSummaryResponse struct {
	*responses.BaseResponse
	RequestId     string           `json:"RequestId" xml:"RequestId"`
	UserStatus    requests.Integer `json:"UserStatus" xml:"UserStatus"`
	QuotaLevel    requests.Integer `json:"QuotaLevel" xml:"QuotaLevel"`
	MaxQuotaLevel requests.Integer `json:"MaxQuotaLevel" xml:"MaxQuotaLevel"`
	DailyQuota    requests.Integer `json:"DailyQuota" xml:"DailyQuota"`
	MonthQuota    requests.Integer `json:"MonthQuota" xml:"MonthQuota"`
	Tags          requests.Integer `json:"Tags" xml:"Tags"`
	Domains       requests.Integer `json:"Domains" xml:"Domains"`
	MailAddresses requests.Integer `json:"MailAddresses" xml:"MailAddresses"`
	Receivers     requests.Integer `json:"Receivers" xml:"Receivers"`
	Templates     requests.Integer `json:"Templates" xml:"Templates"`
	DayuStatus    requests.Integer `json:"DayuStatus" xml:"DayuStatus"`
	SmsTemplates  requests.Integer `json:"SmsTemplates" xml:"SmsTemplates"`
	SmsRecord     requests.Integer `json:"SmsRecord" xml:"SmsRecord"`
	SmsSign       requests.Integer `json:"SmsSign" xml:"SmsSign"`
	EnableTimes   requests.Integer `json:"EnableTimes" xml:"EnableTimes"`
}

func CreateDescAccountSummaryRequest() (request *DescAccountSummaryRequest) {
	request = &DescAccountSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DescAccountSummary", "", "")
	return
}

func CreateDescAccountSummaryResponse() (response *DescAccountSummaryResponse) {
	response = &DescAccountSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
