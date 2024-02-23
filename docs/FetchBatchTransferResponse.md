# FetchBatchTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchTransferId** | Pointer to **string** | It displays the unique ID you created to identify the batch transfer. | [optional] 
**CfBatchTransferId** | Pointer to **string** | It displays the unique ID created by Cashfree Payments. You receive this ID in the response after initiating the batch transfer request. | [optional] 
**Status** | Pointer to **string** | It displays the status of the batch transfer. | [optional] 
**Transfers** | Pointer to [**[]CreateTransferResponse**](CreateTransferResponse.md) |  | [optional] 

## Methods

### NewFetchBatchTransferResponse

`func NewFetchBatchTransferResponse() *FetchBatchTransferResponse`

NewFetchBatchTransferResponse instantiates a new FetchBatchTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchBatchTransferResponseWithDefaults

`func NewFetchBatchTransferResponseWithDefaults() *FetchBatchTransferResponse`

NewFetchBatchTransferResponseWithDefaults instantiates a new FetchBatchTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchTransferId

`func (o *FetchBatchTransferResponse) GetBatchTransferId() string`

GetBatchTransferId returns the BatchTransferId field if non-nil, zero value otherwise.

### GetBatchTransferIdOk

`func (o *FetchBatchTransferResponse) GetBatchTransferIdOk() (*string, bool)`

GetBatchTransferIdOk returns a tuple with the BatchTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchTransferId

`func (o *FetchBatchTransferResponse) SetBatchTransferId(v string)`

SetBatchTransferId sets BatchTransferId field to given value.

### HasBatchTransferId

`func (o *FetchBatchTransferResponse) HasBatchTransferId() bool`

HasBatchTransferId returns a boolean if a field has been set.

### GetCfBatchTransferId

`func (o *FetchBatchTransferResponse) GetCfBatchTransferId() string`

GetCfBatchTransferId returns the CfBatchTransferId field if non-nil, zero value otherwise.

### GetCfBatchTransferIdOk

`func (o *FetchBatchTransferResponse) GetCfBatchTransferIdOk() (*string, bool)`

GetCfBatchTransferIdOk returns a tuple with the CfBatchTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfBatchTransferId

`func (o *FetchBatchTransferResponse) SetCfBatchTransferId(v string)`

SetCfBatchTransferId sets CfBatchTransferId field to given value.

### HasCfBatchTransferId

`func (o *FetchBatchTransferResponse) HasCfBatchTransferId() bool`

HasCfBatchTransferId returns a boolean if a field has been set.

### GetStatus

`func (o *FetchBatchTransferResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchBatchTransferResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchBatchTransferResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FetchBatchTransferResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransfers

`func (o *FetchBatchTransferResponse) GetTransfers() []CreateTransferResponse`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *FetchBatchTransferResponse) GetTransfersOk() (*[]CreateTransferResponse, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *FetchBatchTransferResponse) SetTransfers(v []CreateTransferResponse)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *FetchBatchTransferResponse) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


