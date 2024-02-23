# CreateTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | **string** | It is the unique ID you create to identify the transfer. You can use a maximum of 40 characters to create a transfer_id.  Alphanumeric and underscore ( _ ) are allowed. | 
**TransferAmount** | **float64** | It is the transfer amount. Decimal values are allowed. The minimum value should be equal to or greater than 1.00. (&gt;&#x3D; 1.00) | 
**TransferCurrency** | Pointer to **string** | It is the currency of the transfer amount. The default value is INR. | [optional] 
**TransferMode** | Pointer to **string** | It is the mode of transfer. Allowed values are banktransfer, imps, neft, rtgs, upi, paytm, amazonpay, card. The default transfer_mode is banktransfer. | [optional] 
**BeneficiaryDetails** | Pointer to [**CreateTransferRequestBeneficiaryDetails**](CreateTransferRequestBeneficiaryDetails.md) |  | [optional] 
**TransferRemarks** | Pointer to **string** | It can contain any additional remarks for the transfer. Alphanumeric and whitespaces are allowed. The maximum character limit is 70. | [optional] 
**FundsourceId** | Pointer to **string** | It is the ID of the fund source from which the transfer amount will be debited. | [optional] 

## Methods

### NewCreateTransferRequest

`func NewCreateTransferRequest(transferId string, transferAmount float64, ) *CreateTransferRequest`

NewCreateTransferRequest instantiates a new CreateTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestWithDefaults

`func NewCreateTransferRequestWithDefaults() *CreateTransferRequest`

NewCreateTransferRequestWithDefaults instantiates a new CreateTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *CreateTransferRequest) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *CreateTransferRequest) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *CreateTransferRequest) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetTransferAmount

`func (o *CreateTransferRequest) GetTransferAmount() float64`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *CreateTransferRequest) GetTransferAmountOk() (*float64, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *CreateTransferRequest) SetTransferAmount(v float64)`

SetTransferAmount sets TransferAmount field to given value.


### GetTransferCurrency

`func (o *CreateTransferRequest) GetTransferCurrency() string`

GetTransferCurrency returns the TransferCurrency field if non-nil, zero value otherwise.

### GetTransferCurrencyOk

`func (o *CreateTransferRequest) GetTransferCurrencyOk() (*string, bool)`

GetTransferCurrencyOk returns a tuple with the TransferCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferCurrency

`func (o *CreateTransferRequest) SetTransferCurrency(v string)`

SetTransferCurrency sets TransferCurrency field to given value.

### HasTransferCurrency

`func (o *CreateTransferRequest) HasTransferCurrency() bool`

HasTransferCurrency returns a boolean if a field has been set.

### GetTransferMode

`func (o *CreateTransferRequest) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *CreateTransferRequest) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *CreateTransferRequest) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *CreateTransferRequest) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### GetBeneficiaryDetails

`func (o *CreateTransferRequest) GetBeneficiaryDetails() CreateTransferRequestBeneficiaryDetails`

GetBeneficiaryDetails returns the BeneficiaryDetails field if non-nil, zero value otherwise.

### GetBeneficiaryDetailsOk

`func (o *CreateTransferRequest) GetBeneficiaryDetailsOk() (*CreateTransferRequestBeneficiaryDetails, bool)`

GetBeneficiaryDetailsOk returns a tuple with the BeneficiaryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryDetails

`func (o *CreateTransferRequest) SetBeneficiaryDetails(v CreateTransferRequestBeneficiaryDetails)`

SetBeneficiaryDetails sets BeneficiaryDetails field to given value.

### HasBeneficiaryDetails

`func (o *CreateTransferRequest) HasBeneficiaryDetails() bool`

HasBeneficiaryDetails returns a boolean if a field has been set.

### GetTransferRemarks

`func (o *CreateTransferRequest) GetTransferRemarks() string`

GetTransferRemarks returns the TransferRemarks field if non-nil, zero value otherwise.

### GetTransferRemarksOk

`func (o *CreateTransferRequest) GetTransferRemarksOk() (*string, bool)`

GetTransferRemarksOk returns a tuple with the TransferRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferRemarks

`func (o *CreateTransferRequest) SetTransferRemarks(v string)`

SetTransferRemarks sets TransferRemarks field to given value.

### HasTransferRemarks

`func (o *CreateTransferRequest) HasTransferRemarks() bool`

HasTransferRemarks returns a boolean if a field has been set.

### GetFundsourceId

`func (o *CreateTransferRequest) GetFundsourceId() string`

GetFundsourceId returns the FundsourceId field if non-nil, zero value otherwise.

### GetFundsourceIdOk

`func (o *CreateTransferRequest) GetFundsourceIdOk() (*string, bool)`

GetFundsourceIdOk returns a tuple with the FundsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsourceId

`func (o *CreateTransferRequest) SetFundsourceId(v string)`

SetFundsourceId sets FundsourceId field to given value.

### HasFundsourceId

`func (o *CreateTransferRequest) HasFundsourceId() bool`

HasFundsourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


