# CreateBatchTransferRequestTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | **string** | It is the unique ID you create to identify the transfer. You can use a maximum of 40 characters to create a transfer_id.  Alphanumeric and underscore ( _ ) are allowed. | 
**TransferAmount** | **float64** | It is the transfer amount. Decimal values are allowed. The minimum value should be equal to or greater than 1.00. (&gt;&#x3D; 1.00) | 
**TransferCurrency** | Pointer to **string** | It is the currency of the transfer amount. The default value is INR. | [optional] 
**TransferMode** | Pointer to **string** | It is the mode of transfer. Allowed values are banktransfer, imps, neft, rtgs, upi, paytm, amazonpay, card. The default transfer_mode is banktransfer. | [optional] 
**BeneficiaryDetails** | Pointer to [**CreateBatchTransferRequestTransfersInnerBeneficiaryDetails**](CreateBatchTransferRequestTransfersInnerBeneficiaryDetails.md) |  | [optional] 
**TransferRemarks** | Pointer to **string** | It can contain any additional remarks for the transfer. Alphanumeric and whitespaces are allowed. The maximum character limit is 70. | [optional] 
**FundsourceId** | Pointer to **string** | It is the ID of the fund source from where you want to debit the transfer amount. | [optional] 

## Methods

### NewCreateBatchTransferRequestTransfersInner

`func NewCreateBatchTransferRequestTransfersInner(transferId string, transferAmount float64, ) *CreateBatchTransferRequestTransfersInner`

NewCreateBatchTransferRequestTransfersInner instantiates a new CreateBatchTransferRequestTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchTransferRequestTransfersInnerWithDefaults

`func NewCreateBatchTransferRequestTransfersInnerWithDefaults() *CreateBatchTransferRequestTransfersInner`

NewCreateBatchTransferRequestTransfersInnerWithDefaults instantiates a new CreateBatchTransferRequestTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *CreateBatchTransferRequestTransfersInner) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetTransferAmount

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferAmount() float64`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferAmountOk() (*float64, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *CreateBatchTransferRequestTransfersInner) SetTransferAmount(v float64)`

SetTransferAmount sets TransferAmount field to given value.


### GetTransferCurrency

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferCurrency() string`

GetTransferCurrency returns the TransferCurrency field if non-nil, zero value otherwise.

### GetTransferCurrencyOk

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferCurrencyOk() (*string, bool)`

GetTransferCurrencyOk returns a tuple with the TransferCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferCurrency

`func (o *CreateBatchTransferRequestTransfersInner) SetTransferCurrency(v string)`

SetTransferCurrency sets TransferCurrency field to given value.

### HasTransferCurrency

`func (o *CreateBatchTransferRequestTransfersInner) HasTransferCurrency() bool`

HasTransferCurrency returns a boolean if a field has been set.

### GetTransferMode

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *CreateBatchTransferRequestTransfersInner) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *CreateBatchTransferRequestTransfersInner) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### GetBeneficiaryDetails

`func (o *CreateBatchTransferRequestTransfersInner) GetBeneficiaryDetails() CreateBatchTransferRequestTransfersInnerBeneficiaryDetails`

GetBeneficiaryDetails returns the BeneficiaryDetails field if non-nil, zero value otherwise.

### GetBeneficiaryDetailsOk

`func (o *CreateBatchTransferRequestTransfersInner) GetBeneficiaryDetailsOk() (*CreateBatchTransferRequestTransfersInnerBeneficiaryDetails, bool)`

GetBeneficiaryDetailsOk returns a tuple with the BeneficiaryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryDetails

`func (o *CreateBatchTransferRequestTransfersInner) SetBeneficiaryDetails(v CreateBatchTransferRequestTransfersInnerBeneficiaryDetails)`

SetBeneficiaryDetails sets BeneficiaryDetails field to given value.

### HasBeneficiaryDetails

`func (o *CreateBatchTransferRequestTransfersInner) HasBeneficiaryDetails() bool`

HasBeneficiaryDetails returns a boolean if a field has been set.

### GetTransferRemarks

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferRemarks() string`

GetTransferRemarks returns the TransferRemarks field if non-nil, zero value otherwise.

### GetTransferRemarksOk

`func (o *CreateBatchTransferRequestTransfersInner) GetTransferRemarksOk() (*string, bool)`

GetTransferRemarksOk returns a tuple with the TransferRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferRemarks

`func (o *CreateBatchTransferRequestTransfersInner) SetTransferRemarks(v string)`

SetTransferRemarks sets TransferRemarks field to given value.

### HasTransferRemarks

`func (o *CreateBatchTransferRequestTransfersInner) HasTransferRemarks() bool`

HasTransferRemarks returns a boolean if a field has been set.

### GetFundsourceId

`func (o *CreateBatchTransferRequestTransfersInner) GetFundsourceId() string`

GetFundsourceId returns the FundsourceId field if non-nil, zero value otherwise.

### GetFundsourceIdOk

`func (o *CreateBatchTransferRequestTransfersInner) GetFundsourceIdOk() (*string, bool)`

GetFundsourceIdOk returns a tuple with the FundsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsourceId

`func (o *CreateBatchTransferRequestTransfersInner) SetFundsourceId(v string)`

SetFundsourceId sets FundsourceId field to given value.

### HasFundsourceId

`func (o *CreateBatchTransferRequestTransfersInner) HasFundsourceId() bool`

HasFundsourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


