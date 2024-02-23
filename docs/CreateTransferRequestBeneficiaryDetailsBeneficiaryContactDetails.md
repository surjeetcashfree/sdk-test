# CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryEmail** | Pointer to **string** | It is the email address of the beneficiary. The maximum character limit is 200. It should contain dot (.) and at the rate of (@). | [optional] 
**BeneficiaryPhone** | Pointer to **string** | It is the phone number of the beneficiary. Only registered Indian phone numbers are allowed. The value should be between 8 and 12 characters after stripping +91. | [optional] 
**BeneficiaryCountryCode** | Pointer to **string** | It is the country code (+91) of the beneficiary&#39;s phone number. | [optional] 
**BeneficiaryAddress** | Pointer to **string** | It should contain the address of the beneficiary. The maximum character limit is 150. Alphanumeric characters and whitespaces are allowed. | [optional] 
**BeneficiaryCity** | Pointer to **string** | It is the name of the city as present in the beneficiary address. The maximum character limit is 50. Alphanumeric characters and whitespaces are allowed. | [optional] 
**BeneficiaryState** | Pointer to **string** | It is the name of the state as present in the beneficiary address. The maximum character limit is 50. Alphanumeric characters and whitespaces are allowed. | [optional] 
**BeneficiaryPostalCode** | Pointer to **string** | It is the PIN code as present in the address. It should be a 6 character numeric value. | [optional] 

## Methods

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails() *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetailsWithDefaults

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryEmail

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryEmail() string`

GetBeneficiaryEmail returns the BeneficiaryEmail field if non-nil, zero value otherwise.

### GetBeneficiaryEmailOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryEmailOk() (*string, bool)`

GetBeneficiaryEmailOk returns a tuple with the BeneficiaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryEmail

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryEmail(v string)`

SetBeneficiaryEmail sets BeneficiaryEmail field to given value.

### HasBeneficiaryEmail

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryEmail() bool`

HasBeneficiaryEmail returns a boolean if a field has been set.

### GetBeneficiaryPhone

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryPhone() string`

GetBeneficiaryPhone returns the BeneficiaryPhone field if non-nil, zero value otherwise.

### GetBeneficiaryPhoneOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryPhoneOk() (*string, bool)`

GetBeneficiaryPhoneOk returns a tuple with the BeneficiaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPhone

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryPhone(v string)`

SetBeneficiaryPhone sets BeneficiaryPhone field to given value.

### HasBeneficiaryPhone

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryPhone() bool`

HasBeneficiaryPhone returns a boolean if a field has been set.

### GetBeneficiaryCountryCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryCountryCode() string`

GetBeneficiaryCountryCode returns the BeneficiaryCountryCode field if non-nil, zero value otherwise.

### GetBeneficiaryCountryCodeOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryCountryCodeOk() (*string, bool)`

GetBeneficiaryCountryCodeOk returns a tuple with the BeneficiaryCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCountryCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryCountryCode(v string)`

SetBeneficiaryCountryCode sets BeneficiaryCountryCode field to given value.

### HasBeneficiaryCountryCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryCountryCode() bool`

HasBeneficiaryCountryCode returns a boolean if a field has been set.

### GetBeneficiaryAddress

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.

### HasBeneficiaryAddress

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryAddress() bool`

HasBeneficiaryAddress returns a boolean if a field has been set.

### GetBeneficiaryCity

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryCity() string`

GetBeneficiaryCity returns the BeneficiaryCity field if non-nil, zero value otherwise.

### GetBeneficiaryCityOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryCityOk() (*string, bool)`

GetBeneficiaryCityOk returns a tuple with the BeneficiaryCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCity

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryCity(v string)`

SetBeneficiaryCity sets BeneficiaryCity field to given value.

### HasBeneficiaryCity

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryCity() bool`

HasBeneficiaryCity returns a boolean if a field has been set.

### GetBeneficiaryState

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryState() string`

GetBeneficiaryState returns the BeneficiaryState field if non-nil, zero value otherwise.

### GetBeneficiaryStateOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryStateOk() (*string, bool)`

GetBeneficiaryStateOk returns a tuple with the BeneficiaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryState

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryState(v string)`

SetBeneficiaryState sets BeneficiaryState field to given value.

### HasBeneficiaryState

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryState() bool`

HasBeneficiaryState returns a boolean if a field has been set.

### GetBeneficiaryPostalCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryPostalCode() string`

GetBeneficiaryPostalCode returns the BeneficiaryPostalCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostalCodeOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) GetBeneficiaryPostalCodeOk() (*string, bool)`

GetBeneficiaryPostalCodeOk returns a tuple with the BeneficiaryPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostalCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) SetBeneficiaryPostalCode(v string)`

SetBeneficiaryPostalCode sets BeneficiaryPostalCode field to given value.

### HasBeneficiaryPostalCode

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails) HasBeneficiaryPostalCode() bool`

HasBeneficiaryPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


