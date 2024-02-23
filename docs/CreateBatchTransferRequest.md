# CreateBatchTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchTransferId** | **string** | It is the unique ID you create to identify the batch transfer request. The maximum character limit is 60. Alphanumeric and underscore ( _ ) characters are allowed. | 
**Transfers** | Pointer to [**[]CreateBatchTransferRequestTransfersInner**](CreateBatchTransferRequestTransfersInner.md) |  | [optional] 

## Methods

### NewCreateBatchTransferRequest

`func NewCreateBatchTransferRequest(batchTransferId string, ) *CreateBatchTransferRequest`

NewCreateBatchTransferRequest instantiates a new CreateBatchTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchTransferRequestWithDefaults

`func NewCreateBatchTransferRequestWithDefaults() *CreateBatchTransferRequest`

NewCreateBatchTransferRequestWithDefaults instantiates a new CreateBatchTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchTransferId

`func (o *CreateBatchTransferRequest) GetBatchTransferId() string`

GetBatchTransferId returns the BatchTransferId field if non-nil, zero value otherwise.

### GetBatchTransferIdOk

`func (o *CreateBatchTransferRequest) GetBatchTransferIdOk() (*string, bool)`

GetBatchTransferIdOk returns a tuple with the BatchTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchTransferId

`func (o *CreateBatchTransferRequest) SetBatchTransferId(v string)`

SetBatchTransferId sets BatchTransferId field to given value.


### GetTransfers

`func (o *CreateBatchTransferRequest) GetTransfers() []CreateBatchTransferRequestTransfersInner`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *CreateBatchTransferRequest) GetTransfersOk() (*[]CreateBatchTransferRequestTransfersInner, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *CreateBatchTransferRequest) SetTransfers(v []CreateBatchTransferRequestTransfersInner)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *CreateBatchTransferRequest) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


