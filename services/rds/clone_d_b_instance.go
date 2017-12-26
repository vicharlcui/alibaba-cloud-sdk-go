package rds

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

func (client *Client) CloneDBInstance(request *CloneDBInstanceRequest) (response *CloneDBInstanceResponse, err error) {
	response = CreateCloneDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CloneDBInstanceWithChan(request *CloneDBInstanceRequest) (<-chan *CloneDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CloneDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloneDBInstance(request)
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

func (client *Client) CloneDBInstanceWithCallback(request *CloneDBInstanceRequest, callback func(response *CloneDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloneDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CloneDBInstance(request)
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

type CloneDBInstanceRequest struct {
	*requests.RpcRequest
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	Period                string           `position:"Query" name:"Period"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	BackupId              string           `position:"Query" name:"BackupId"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	RestoreTime           string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	PayType               string           `position:"Query" name:"PayType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

type CloneDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

func CreateCloneDBInstanceRequest() (request *CloneDBInstanceRequest) {
	request = &CloneDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CloneDBInstance", "", "")
	return
}

func CreateCloneDBInstanceResponse() (response *CloneDBInstanceResponse) {
	response = &CloneDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
