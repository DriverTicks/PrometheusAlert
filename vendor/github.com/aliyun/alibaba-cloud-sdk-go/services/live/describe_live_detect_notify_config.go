package live

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

// DescribeLiveDetectNotifyConfig invokes the live.DescribeLiveDetectNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/describelivedetectnotifyconfig.html
func (client *Client) DescribeLiveDetectNotifyConfig(request *DescribeLiveDetectNotifyConfigRequest) (response *DescribeLiveDetectNotifyConfigResponse, err error) {
	response = CreateDescribeLiveDetectNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDetectNotifyConfigWithChan invokes the live.DescribeLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDetectNotifyConfigWithChan(request *DescribeLiveDetectNotifyConfigRequest) (<-chan *DescribeLiveDetectNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDetectNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDetectNotifyConfig(request)
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

// DescribeLiveDetectNotifyConfigWithCallback invokes the live.DescribeLiveDetectNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describelivedetectnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveDetectNotifyConfigWithCallback(request *DescribeLiveDetectNotifyConfigRequest, callback func(response *DescribeLiveDetectNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDetectNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDetectNotifyConfig(request)
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

// DescribeLiveDetectNotifyConfigRequest is the request struct for api DescribeLiveDetectNotifyConfig
type DescribeLiveDetectNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeLiveDetectNotifyConfigResponse is the response struct for api DescribeLiveDetectNotifyConfig
type DescribeLiveDetectNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	LiveDetectNotifyConfig LiveDetectNotifyConfig `json:"LiveDetectNotifyConfig" xml:"LiveDetectNotifyConfig"`
}

// CreateDescribeLiveDetectNotifyConfigRequest creates a request to invoke DescribeLiveDetectNotifyConfig API
func CreateDescribeLiveDetectNotifyConfigRequest() (request *DescribeLiveDetectNotifyConfigRequest) {
	request = &DescribeLiveDetectNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDetectNotifyConfig", "live", "openAPI")
	return
}

// CreateDescribeLiveDetectNotifyConfigResponse creates a response to parse from DescribeLiveDetectNotifyConfig response
func CreateDescribeLiveDetectNotifyConfigResponse() (response *DescribeLiveDetectNotifyConfigResponse) {
	response = &DescribeLiveDetectNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}