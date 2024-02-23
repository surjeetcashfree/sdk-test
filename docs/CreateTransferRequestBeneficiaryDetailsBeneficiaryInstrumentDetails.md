# CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | It is the beneficiary bank account number. The value should be between 9 and 18 characters. Alphanumeric characters are allowed. This value is required if beneficiary_id is not present and if the transfer_mode is banktransfer, imps, neft,rtgs mode. | [optional] 
**BankIfsc** | Pointer to **string** | It is the IFSC information of the beneficiary&#39;s bank account in the standard IFSC format. The value should be 11 characters. (The first 4 characters should be alphabets, the 5th character should be a 0, and the remaining 6 characters should be numerals.) | [optional] 
**Vpa** | Pointer to **string** | It is the UPI VPA information of the beneficiary. Only valid UPI VPA information is accepted. It can be an Alphanumeric value with only period (.), hyphen (-), underscore ( _ ), and at the rate of (@). Hyphen (-) is accepted only before at the rate of (@). This value is required if the transfer_mode is upi. | [optional] 
**CardDetails** | Pointer to [**CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails**](CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails.md) |  | [optional] 

## Methods

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankIfsc

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfsc() string`

GetBankIfsc returns the BankIfsc field if non-nil, zero value otherwise.

### GetBankIfscOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetBankIfscOk() (*string, bool)`

GetBankIfscOk returns a tuple with the BankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIfsc

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetBankIfsc(v string)`

SetBankIfsc sets BankIfsc field to given value.

### HasBankIfsc

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasBankIfsc() bool`

HasBankIfsc returns a boolean if a field has been set.

### GetVpa

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetVpa(v string)`

SetVpa sets Vpa field to given value.

### HasVpa

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasVpa() bool`

HasVpa returns a boolean if a field has been set.

### GetCardDetails

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetCardDetails() CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) GetCardDetailsOk() (*CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) SetCardDetails(v CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails)`

SetCardDetails sets CardDetails field to given value.

### HasCardDetails

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails) HasCardDetails() bool`

HasCardDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


