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

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
	response = CreateCreateDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateDBInstanceWithChan(request *CreateDBInstanceRequest) (<-chan *CreateDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstance(request)
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

func (client *Client) CreateDBInstanceWithCallback(request *CreateDBInstanceRequest, callback func(response *CreateDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstance(request)
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

type CreateDBInstanceRequest struct {
	*requests.RpcRequest
	UsedTime              string           `position:"Query" name:"UsedTime"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	PayType               string           `position:"Query" name:"PayType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	Engine                string           `position:"Query" name:"Engine"`
	DBInstanceNetType     string           `position:"Query" name:"DBInstanceNetType"`
	Period                string           `position:"Query" name:"Period"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	SecurityIPList        string           `position:"Query" name:"SecurityIPList"`
	SystemDBCharset       string           `position:"Query" name:"SystemDBCharset"`
	ConnectionMode        string           `position:"Query" name:"ConnectionMode"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	TunnelId              string           `position:"Query" name:"TunnelId"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

type CreateDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

func CreateCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstance", "", "")
	return
}

func CreateCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
