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

func (client *Client) DescribeLiveStreamRoomBitRate(request *DescribeLiveStreamRoomBitRateRequest) (response *DescribeLiveStreamRoomBitRateResponse, err error) {
	response = CreateDescribeLiveStreamRoomBitRateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamRoomBitRateWithChan(request *DescribeLiveStreamRoomBitRateRequest) (<-chan *DescribeLiveStreamRoomBitRateResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRoomBitRateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRoomBitRate(request)
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

func (client *Client) DescribeLiveStreamRoomBitRateWithCallback(request *DescribeLiveStreamRoomBitRateRequest, callback func(response *DescribeLiveStreamRoomBitRateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRoomBitRateResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRoomBitRate(request)
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

type DescribeLiveStreamRoomBitRateRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamRoomBitRateResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	FrameRateAndBitRateInfos struct {
		FrameRateAndBitRateInfo []struct {
			StreamUrl      string         `json:"StreamUrl" xml:"StreamUrl"`
			VideoFrameRate requests.Float `json:"VideoFrameRate" xml:"VideoFrameRate"`
			AudioFrameRate requests.Float `json:"AudioFrameRate" xml:"AudioFrameRate"`
			BitRate        requests.Float `json:"BitRate" xml:"BitRate"`
			Time           string         `json:"Time" xml:"Time"`
		} `json:"FrameRateAndBitRateInfo" xml:"FrameRateAndBitRateInfo"`
	} `json:"FrameRateAndBitRateInfos" xml:"FrameRateAndBitRateInfos"`
}

func CreateDescribeLiveStreamRoomBitRateRequest() (request *DescribeLiveStreamRoomBitRateRequest) {
	request = &DescribeLiveStreamRoomBitRateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRoomBitRate", "", "")
	return
}

func CreateDescribeLiveStreamRoomBitRateResponse() (response *DescribeLiveStreamRoomBitRateResponse) {
	response = &DescribeLiveStreamRoomBitRateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
