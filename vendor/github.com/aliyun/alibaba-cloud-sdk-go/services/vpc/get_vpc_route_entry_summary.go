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

// GetVpcRouteEntrySummary invokes the vpc.GetVpcRouteEntrySummary API synchronously
func (client *Client) GetVpcRouteEntrySummary(request *GetVpcRouteEntrySummaryRequest) (response *GetVpcRouteEntrySummaryResponse, err error) {
	response = CreateGetVpcRouteEntrySummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetVpcRouteEntrySummaryWithChan invokes the vpc.GetVpcRouteEntrySummary API asynchronously
func (client *Client) GetVpcRouteEntrySummaryWithChan(request *GetVpcRouteEntrySummaryRequest) (<-chan *GetVpcRouteEntrySummaryResponse, <-chan error) {
	responseChan := make(chan *GetVpcRouteEntrySummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVpcRouteEntrySummary(request)
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

// GetVpcRouteEntrySummaryWithCallback invokes the vpc.GetVpcRouteEntrySummary API asynchronously
func (client *Client) GetVpcRouteEntrySummaryWithCallback(request *GetVpcRouteEntrySummaryRequest, callback func(response *GetVpcRouteEntrySummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVpcRouteEntrySummaryResponse
		var err error
		defer close(result)
		response, err = client.GetVpcRouteEntrySummary(request)
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

// GetVpcRouteEntrySummaryRequest is the request struct for api GetVpcRouteEntrySummary
type GetVpcRouteEntrySummaryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteEntryType       string           `position:"Query" name:"RouteEntryType"`
	RouteTableId         string           `position:"Query" name:"RouteTableId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
}

// GetVpcRouteEntrySummaryResponse is the response struct for api GetVpcRouteEntrySummary
type GetVpcRouteEntrySummaryResponse struct {
	*responses.BaseResponse
	RequestId          string                   `json:"RequestId" xml:"RequestId"`
	RouteEntrySummarys []RouteEntrySummarysItem `json:"RouteEntrySummarys" xml:"RouteEntrySummarys"`
}

// CreateGetVpcRouteEntrySummaryRequest creates a request to invoke GetVpcRouteEntrySummary API
func CreateGetVpcRouteEntrySummaryRequest() (request *GetVpcRouteEntrySummaryRequest) {
	request = &GetVpcRouteEntrySummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "GetVpcRouteEntrySummary", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVpcRouteEntrySummaryResponse creates a response to parse from GetVpcRouteEntrySummary response
func CreateGetVpcRouteEntrySummaryResponse() (response *GetVpcRouteEntrySummaryResponse) {
	response = &GetVpcRouteEntrySummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
