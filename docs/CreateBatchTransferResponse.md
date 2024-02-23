# CreateBatchTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchTransferId** | Pointer to **string** | It displays the unique ID you created to identify the batch transfer request. | [optional] 
**CfBatchTransferId** | Pointer to **string** | It displays the unique ID created by Cashfree Payments. You receive this ID in the response after initiating the batch transfer request. | [optional] 
**Status** | Pointer to **string** | It displays the status of the API request. | [optional] 

## Methods

### NewCreateBatchTransferResponse

`func NewCreateBatchTransferResponse() *CreateBatchTransferResponse`

NewCreateBatchTransferResponse instantiates a new CreateBatchTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchTransferResponseWithDefaults

`func NewCreateBatchTransferResponseWithDefaults() *CreateBatchTransferResponse`

NewCreateBatchTransferResponseWithDefaults instantiates a new CreateBatchTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchTransferId

`func (o *CreateBatchTransferResponse) GetBatchTransferId() string`

GetBatchTransferId returns the BatchTransferId field if non-nil, zero value otherwise.

### GetBatchTransferIdOk

`func (o *CreateBatchTransferResponse) GetBatchTransferIdOk() (*string, bool)`

GetBatchTransferIdOk returns a tuple with the BatchTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchTransferId

`func (o *CreateBatchTransferResponse) SetBatchTransferId(v string)`

SetBatchTransferId sets BatchTransferId field to given value.

### HasBatchTransferId

`func (o *CreateBatchTransferResponse) HasBatchTransferId() bool`

HasBatchTransferId returns a boolean if a field has been set.

### GetCfBatchTransferId

`func (o *CreateBatchTransferResponse) GetCfBatchTransferId() string`

GetCfBatchTransferId returns the CfBatchTransferId field if non-nil, zero value otherwise.

### GetCfBatchTransferIdOk

`func (o *CreateBatchTransferResponse) GetCfBatchTransferIdOk() (*string, bool)`

GetCfBatchTransferIdOk returns a tuple with the CfBatchTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfBatchTransferId

`func (o *CreateBatchTransferResponse) SetCfBatchTransferId(v string)`

SetCfBatchTransferId sets CfBatchTransferId field to given value.

### HasCfBatchTransferId

`func (o *CreateBatchTransferResponse) HasCfBatchTransferId() bool`

HasCfBatchTransferId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateBatchTransferResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBatchTransferResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBatchTransferResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateBatchTransferResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


