package cloudphoto

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

func (client *Client) GetPhotosByMd5s(request *GetPhotosByMd5sRequest) (response *GetPhotosByMd5sResponse, err error) {
	response = CreateGetPhotosByMd5sResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetPhotosByMd5sWithChan(request *GetPhotosByMd5sRequest) (<-chan *GetPhotosByMd5sResponse, <-chan error) {
	responseChan := make(chan *GetPhotosByMd5sResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPhotosByMd5s(request)
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

func (client *Client) GetPhotosByMd5sWithCallback(request *GetPhotosByMd5sRequest, callback func(response *GetPhotosByMd5sResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPhotosByMd5sResponse
		var err error
		defer close(result)
		response, err = client.GetPhotosByMd5s(request)
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

type GetPhotosByMd5sRequest struct {
	*requests.RpcRequest
	State     string    `position:"Query" name:"State"`
	Md5       *[]string `position:"Query" name:"Md5"  type:"Repeated"`
	LibraryId string    `position:"Query" name:"LibraryId"`
	StoreName string    `position:"Query" name:"StoreName"`
}

type GetPhotosByMd5sResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
	Photos    []struct {
		Id              requests.Integer `json:"Id" xml:"Id"`
		Title           string           `json:"Title" xml:"Title"`
		FileId          string           `json:"FileId" xml:"FileId"`
		State           string           `json:"State" xml:"State"`
		Md5             string           `json:"Md5" xml:"Md5"`
		IsVideo         requests.Boolean `json:"IsVideo" xml:"IsVideo"`
		Remark          string           `json:"Remark" xml:"Remark"`
		Width           requests.Integer `json:"Width" xml:"Width"`
		Height          requests.Integer `json:"Height" xml:"Height"`
		Ctime           requests.Integer `json:"Ctime" xml:"Ctime"`
		Mtime           requests.Integer `json:"Mtime" xml:"Mtime"`
		TakenAt         requests.Integer `json:"TakenAt" xml:"TakenAt"`
		ShareExpireTime requests.Integer `json:"ShareExpireTime" xml:"ShareExpireTime"`
	} `json:"Photos" xml:"Photos"`
}

func CreateGetPhotosByMd5sRequest() (request *GetPhotosByMd5sRequest) {
	request = &GetPhotosByMd5sRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPhotosByMd5s", "", "")
	return
}

func CreateGetPhotosByMd5sResponse() (response *GetPhotosByMd5sResponse) {
	response = &GetPhotosByMd5sResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
