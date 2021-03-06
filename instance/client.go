package instance

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	cdshttp "github.com/capitalonline/cds-gic-sdk-go/common/http"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
)

const ApiVersion = "2019-08-08"

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}
func NewAddInstanceRequest() (request *AddInstanceRequest) {
	request = &AddInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", ApiVersion, "CreateInstance")
	return
}

func NewAddInstanceResponse() (response *AddInstanceResponse) {
	response = &AddInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DescribeInstances")
	return
}
func NewDescribeInstanceResponse() (response *DescribeInstanceReponse) {
	response = &DescribeInstanceReponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DeleteInstance")
	return
}
func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyInstanceName")
	return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyInstanceSpecRequest() (request *ModifyInstanceSpecRequest) {
	request = &ModifyInstanceSpecRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyInstanceSpec")
	return
}

func NewModifyInstanceSpecResponse() (response *ModifyInstanceSpecResponse) {
	response = &ModifyInstanceSpecResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewCreateDiskRequest() (request *CreateDiskRequest) {
	request = &CreateDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "CreateDisk")
	return
}
func NewCreateDiskResponse() (response *CreateDiskResponse) {
	response = &CreateDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewResizeDiskRequest() (request *ResizeDiskRequest) {
	request = &ResizeDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ResizeDisk")
	return
}
func NewResizeDiskResponse() (response *ResizeDiskResponse) {
	response = &ResizeDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewDeleteDiskRequest() (request *DeleteDiskRequest) {
	request = &DeleteDiskRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "DeleteDisk")
	return
}
func NewDeleteDiskResponse() (response *DeleteDiskResponse) {
	response = &DeleteDiskResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

func NewModifyIpRequest() (request *ModifyIpRequest) {
	request = &ModifyIpRequest{
		BaseRequest: &cdshttp.BaseRequest{},
	}
	request.Init().WithApiInfo("CCS", ApiVersion, "ModifyIpAddress")
	return
}
func NewModifyIpResponse() (response *ModifyIpResponse) {
	response = &ModifyIpResponse{
		BaseResponse: &cdshttp.BaseResponse{},
	}
	return
}

// Create Instance
func (c *Client) CreateInstance(request *AddInstanceRequest) (response *AddInstanceResponse, err error) {
	if request == nil {
		request = NewAddInstanceRequest()
	}
	response = NewAddInstanceResponse()
	err = c.Send(request, response)
	return
}

// Describe Instance
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceReponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	response = NewDescribeInstanceResponse()
	err = c.Send(request, response)
	return
}

// Delete Instance
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

// Modify Instance Name
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}

// ModifyInstanceSpec cpu, ram
func (c *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (response *ModifyInstanceSpecResponse, err error) {
	if request == nil {
		request = NewModifyInstanceSpecRequest()
	}
	response = NewModifyInstanceSpecResponse()
	err = c.Send(request, response)
	return
}

// Create Disk
func (c *Client) CreateDisk(request *CreateDiskRequest) (response *CreateDiskResponse, err error) {
	if request == nil {
		request = NewCreateDiskRequest()
	}
	response = NewCreateDiskResponse()
	err = c.Send(request, response)
	return
}

// Resize Disk
func (c *Client) ResizeDisk(request *ResizeDiskRequest) (response *ResizeDiskResponse, err error) {
	if request == nil {
		request = NewResizeDiskRequest()
	}
	response = NewResizeDiskResponse()
	err = c.Send(request, response)
	return
}

// Delete Disk
func (c *Client) DeleteDisk(request *DeleteDiskRequest) (response *DeleteDiskResponse, err error) {
	if request == nil {
		request = NewDeleteDiskRequest()
	}
	response = NewDeleteDiskResponse()
	err = c.Send(request, response)
	return
}

// Delete Disk
func (c *Client) ModifyIpAddress(request *ModifyIpRequest) (response *ModifyIpResponse, err error) {
	if request == nil {
		request = NewModifyIpRequest()
	}
	response = NewModifyIpResponse()
	err = c.Send(request, response)
	return
}
