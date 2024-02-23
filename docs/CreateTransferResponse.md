# CreateTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | Pointer to **string** | It displays the unique ID you created to identify the transfer. | [optional] 
**CfTransferId** | Pointer to **string** | It displays the unique ID created by Cashfree Payments. You receive this ID in the response after initiating the standard transfer request. | [optional] 
**Status** | Pointer to **string** | It displays the status of the transfer. | [optional] 
**BeneficiaryDetails** | Pointer to [**CreateTransferResponseBeneficiaryDetails**](CreateTransferResponseBeneficiaryDetails.md) |  | [optional] 
**TransferAmount** | Pointer to **float64** | It displays the transfer amount initiated in the request. | [optional] 
**TransferServiceCharge** | Pointer to **float64** | It displays the service charge applicable for the successful transfer request. | [optional] 
**TransferServiceTax** | Pointer to **float64** | It displays the service tax applicable for the successful transfer request. | [optional] 
**TransferMode** | Pointer to **string** | It displays the mode of the transfer. | [optional] 
**TransferUtr** | Pointer to **string** | It displays the unique number that is generated to recognise any fund transfer that is created by the bank that facilitates the transfer. | [optional] 
**FundsourceId** | Pointer to **string** | It displays the ID of the fund source from where the money was debited for this transfer request. | [optional] 
**AddedOn** | Pointer to **string** | It displays the time of when the transfer request was added to the system. | [optional] 
**UpdatedOn** | Pointer to **string** | It displays the updated time for the transfer. | [optional] 

## Methods

### NewCreateTransferResponse

`func NewCreateTransferResponse() *CreateTransferResponse`

NewCreateTransferResponse instantiates a new CreateTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferResponseWithDefaults

`func NewCreateTransferResponseWithDefaults() *CreateTransferResponse`

NewCreateTransferResponseWithDefaults instantiates a new CreateTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *CreateTransferResponse) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *CreateTransferResponse) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *CreateTransferResponse) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *CreateTransferResponse) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### GetCfTransferId

`func (o *CreateTransferResponse) GetCfTransferId() string`

GetCfTransferId returns the CfTransferId field if non-nil, zero value otherwise.

### GetCfTransferIdOk

`func (o *CreateTransferResponse) GetCfTransferIdOk() (*string, bool)`

GetCfTransferIdOk returns a tuple with the CfTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfTransferId

`func (o *CreateTransferResponse) SetCfTransferId(v string)`

SetCfTransferId sets CfTransferId field to given value.

### HasCfTransferId

`func (o *CreateTransferResponse) HasCfTransferId() bool`

HasCfTransferId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateTransferResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTransferResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTransferResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTransferResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBeneficiaryDetails

`func (o *CreateTransferResponse) GetBeneficiaryDetails() CreateTransferResponseBeneficiaryDetails`

GetBeneficiaryDetails returns the BeneficiaryDetails field if non-nil, zero value otherwise.

### GetBeneficiaryDetailsOk

`func (o *CreateTransferResponse) GetBeneficiaryDetailsOk() (*CreateTransferResponseBeneficiaryDetails, bool)`

GetBeneficiaryDetailsOk returns a tuple with the BeneficiaryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryDetails

`func (o *CreateTransferResponse) SetBeneficiaryDetails(v CreateTransferResponseBeneficiaryDetails)`

SetBeneficiaryDetails sets BeneficiaryDetails field to given value.

### HasBeneficiaryDetails

`func (o *CreateTransferResponse) HasBeneficiaryDetails() bool`

HasBeneficiaryDetails returns a boolean if a field has been set.

### GetTransferAmount

`func (o *CreateTransferResponse) GetTransferAmount() float64`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *CreateTransferResponse) GetTransferAmountOk() (*float64, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *CreateTransferResponse) SetTransferAmount(v float64)`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *CreateTransferResponse) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.

### GetTransferServiceCharge

`func (o *CreateTransferResponse) GetTransferServiceCharge() float64`

GetTransferServiceCharge returns the TransferServiceCharge field if non-nil, zero value otherwise.

### GetTransferServiceChargeOk

`func (o *CreateTransferResponse) GetTransferServiceChargeOk() (*float64, bool)`

GetTransferServiceChargeOk returns a tuple with the TransferServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferServiceCharge

`func (o *CreateTransferResponse) SetTransferServiceCharge(v float64)`

SetTransferServiceCharge sets TransferServiceCharge field to given value.

### HasTransferServiceCharge

`func (o *CreateTransferResponse) HasTransferServiceCharge() bool`

HasTransferServiceCharge returns a boolean if a field has been set.

### GetTransferServiceTax

`func (o *CreateTransferResponse) GetTransferServiceTax() float64`

GetTransferServiceTax returns the TransferServiceTax field if non-nil, zero value otherwise.

### GetTransferServiceTaxOk

`func (o *CreateTransferResponse) GetTransferServiceTaxOk() (*float64, bool)`

GetTransferServiceTaxOk returns a tuple with the TransferServiceTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferServiceTax

`func (o *CreateTransferResponse) SetTransferServiceTax(v float64)`

SetTransferServiceTax sets TransferServiceTax field to given value.

### HasTransferServiceTax

`func (o *CreateTransferResponse) HasTransferServiceTax() bool`

HasTransferServiceTax returns a boolean if a field has been set.

### GetTransferMode

`func (o *CreateTransferResponse) GetTransferMode() string`

GetTransferMode returns the TransferMode field if non-nil, zero value otherwise.

### GetTransferModeOk

`func (o *CreateTransferResponse) GetTransferModeOk() (*string, bool)`

GetTransferModeOk returns a tuple with the TransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMode

`func (o *CreateTransferResponse) SetTransferMode(v string)`

SetTransferMode sets TransferMode field to given value.

### HasTransferMode

`func (o *CreateTransferResponse) HasTransferMode() bool`

HasTransferMode returns a boolean if a field has been set.

### GetTransferUtr

`func (o *CreateTransferResponse) GetTransferUtr() string`

GetTransferUtr returns the TransferUtr field if non-nil, zero value otherwise.

### GetTransferUtrOk

`func (o *CreateTransferResponse) GetTransferUtrOk() (*string, bool)`

GetTransferUtrOk returns a tuple with the TransferUtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferUtr

`func (o *CreateTransferResponse) SetTransferUtr(v string)`

SetTransferUtr sets TransferUtr field to given value.

### HasTransferUtr

`func (o *CreateTransferResponse) HasTransferUtr() bool`

HasTransferUtr returns a boolean if a field has been set.

### GetFundsourceId

`func (o *CreateTransferResponse) GetFundsourceId() string`

GetFundsourceId returns the FundsourceId field if non-nil, zero value otherwise.

### GetFundsourceIdOk

`func (o *CreateTransferResponse) GetFundsourceIdOk() (*string, bool)`

GetFundsourceIdOk returns a tuple with the FundsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsourceId

`func (o *CreateTransferResponse) SetFundsourceId(v string)`

SetFundsourceId sets FundsourceId field to given value.

### HasFundsourceId

`func (o *CreateTransferResponse) HasFundsourceId() bool`

HasFundsourceId returns a boolean if a field has been set.

### GetAddedOn

`func (o *CreateTransferResponse) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *CreateTransferResponse) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *CreateTransferResponse) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *CreateTransferResponse) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *CreateTransferResponse) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *CreateTransferResponse) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *CreateTransferResponse) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *CreateTransferResponse) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


