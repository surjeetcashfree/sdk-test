# CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | It is the beneficiary bank account number. The value should be between 9 and 18 characters. Alphanumeric characters are allowed. This value is required if beneficiary_id is not present and if the transfer_mode is banktransfer, imps, neft, rtgs mode. | [optional] 
**BankIfsc** | Pointer to **string** | It is the IFSC information of the beneficiary&#39;s bank account in the standard IFSC format. The value should be 11 characters. (The first 4 characters should be alphabets, the 5th character should be a 0, and the remaining 6 characters should be numerals.) | [optional] 
**Vpa** | Pointer to **string** | It is the UPI VPA information of the beneficiary. Only valid UPI VPA information is accepted. It can be an Alphanumeric value with only period (.), hyphen (-), underscore ( _ ), and at the rate of (@). Hyphen (-) is accepted only before at the rate of (@). This value is required if the transfer_mode is upi. | [optional] 

## Methods

### NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails

`func NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails() *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails`

NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails instantiates a new CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults

`func NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults() *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails`

NewCreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults instantiates a new CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankIfsc

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfsc() string`

GetBankIfsc returns the BankIfsc field if non-nil, zero value otherwise.

### GetBankIfscOk

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfscOk() (*string, bool)`

GetBankIfscOk returns a tuple with the BankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIfsc

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankIfsc(v string)`

SetBankIfsc sets BankIfsc field to given value.

### HasBankIfsc

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankIfsc() bool`

HasBankIfsc returns a boolean if a field has been set.

### GetVpa

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) SetVpa(v string)`

SetVpa sets Vpa field to given value.

### HasVpa

`func (o *CreateBatchTransferRequestTransfersInnerBeneficiaryDetailsBeneficiaryInstrumentDetails) HasVpa() bool`

HasVpa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


