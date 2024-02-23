# CreateBeneficiaryRequestBeneficiaryInstrumentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccountNumber** | Pointer to **string** | It is the beneficiary bank account number. The value should be between 9 and 18 characters. Alphanumeric characters are allowed. You need to input bank_ifsc if bank_account_number is passed. | [optional] 
**BankIfsc** | Pointer to **string** | It is the IFSC information of the beneficiary&#39;s bank account in the standard IFSC format. The value should be 11 characters. (The first 4 characters should be alphabets, the 5th character should be a 0, and the remaining 6 characters should be numerals.). You need to input bank_account_number if bank_ifsc is passed. | [optional] 
**Vpa** | Pointer to **string** | It is the UPI VPA information of the beneficiary. Only valid UPI VPA information is accepted. It can be an Alphanumeric value with only period (.), hyphen (-), underscore ( _ ), and at the rate of (@). Hyphen (-) is accepted only before at the rate of (@). | [optional] 

## Methods

### NewCreateBeneficiaryRequestBeneficiaryInstrumentDetails

`func NewCreateBeneficiaryRequestBeneficiaryInstrumentDetails() *CreateBeneficiaryRequestBeneficiaryInstrumentDetails`

NewCreateBeneficiaryRequestBeneficiaryInstrumentDetails instantiates a new CreateBeneficiaryRequestBeneficiaryInstrumentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBeneficiaryRequestBeneficiaryInstrumentDetailsWithDefaults

`func NewCreateBeneficiaryRequestBeneficiaryInstrumentDetailsWithDefaults() *CreateBeneficiaryRequestBeneficiaryInstrumentDetails`

NewCreateBeneficiaryRequestBeneficiaryInstrumentDetailsWithDefaults instantiates a new CreateBeneficiaryRequestBeneficiaryInstrumentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccountNumber

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankIfsc

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetBankIfsc() string`

GetBankIfsc returns the BankIfsc field if non-nil, zero value otherwise.

### GetBankIfscOk

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetBankIfscOk() (*string, bool)`

GetBankIfscOk returns a tuple with the BankIfsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIfsc

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) SetBankIfsc(v string)`

SetBankIfsc sets BankIfsc field to given value.

### HasBankIfsc

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) HasBankIfsc() bool`

HasBankIfsc returns a boolean if a field has been set.

### GetVpa

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetVpa() string`

GetVpa returns the Vpa field if non-nil, zero value otherwise.

### GetVpaOk

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) GetVpaOk() (*string, bool)`

GetVpaOk returns a tuple with the Vpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpa

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) SetVpa(v string)`

SetVpa sets Vpa field to given value.

### HasVpa

`func (o *CreateBeneficiaryRequestBeneficiaryInstrumentDetails) HasVpa() bool`

HasVpa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


