package ecs

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

// SecurityGroup is a nested struct in ecs response
type SecurityGroup struct {
	CreationTime            string                       `json:"CreationTime" xml:"CreationTime"`
	VpcId                   string                       `json:"VpcId" xml:"VpcId"`
	ServiceManaged          bool                         `json:"ServiceManaged" xml:"ServiceManaged"`
	Description             string                       `json:"Description" xml:"Description"`
	SecurityGroupId         string                       `json:"SecurityGroupId" xml:"SecurityGroupId"`
	ResourceGroupId         string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SecurityGroupName       string                       `json:"SecurityGroupName" xml:"SecurityGroupName"`
	EcsCount                int                          `json:"EcsCount" xml:"EcsCount"`
	ServiceID               int64                        `json:"ServiceID" xml:"ServiceID"`
	SecurityGroupType       string                       `json:"SecurityGroupType" xml:"SecurityGroupType"`
	AvailableInstanceAmount int                          `json:"AvailableInstanceAmount" xml:"AvailableInstanceAmount"`
	Tags                    TagsInDescribeSecurityGroups `json:"Tags" xml:"Tags"`
}
