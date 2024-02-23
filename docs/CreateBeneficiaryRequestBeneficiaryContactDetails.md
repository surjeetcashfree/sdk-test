# CreateBeneficiaryRequestBeneficiaryContactDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryEmail** | Pointer to **string** | It is the email address of the beneficiary. The maximum character limit is 200. It should contain dot (.) and at the rate of (@). | [optional] 
**BeneficiaryPhone** | Pointer to **string** | It is the phone number of the beneficiary. Only reigstered Indian phone numbers are allowed. The value should be between 8 and 12 characters after stripping +91. | [optional] 
**BeneficiaryCountryCode** | Pointer to **string** | It is the country code (+91) of the beneficiary&#39;s phone number | [optional] 
**BeneficiaryAddress** | Pointer to **string** | It is the address information of the beneficiary. | [optional] 
**BeneficiaryCity** | Pointer to **string** | It is the name of the city as present in the beneficiary&#39;s address. | [optional] 
**BeneficiaryState** | Pointer to **string** | It is the name of the state as present in the beneficiary&#39;s address. | [optional] 
**BeneficiaryPostalCode** | Pointer to **string** | It is the PIN code information as present in the beneficiary&#39;s address. The maximum character limit is 6. Only numbers are allowed. | [optional] 

## Methods

### NewCreateBeneficiaryRequestBeneficiaryContactDetails

`func NewCreateBeneficiaryRequestBeneficiaryContactDetails() *CreateBeneficiaryRequestBeneficiaryContactDetails`

NewCreateBeneficiaryRequestBeneficiaryContactDetails instantiates a new CreateBeneficiaryRequestBeneficiaryContactDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBeneficiaryRequestBeneficiaryContactDetailsWithDefaults

`func NewCreateBeneficiaryRequestBeneficiaryContactDetailsWithDefaults() *CreateBeneficiaryRequestBeneficiaryContactDetails`

NewCreateBeneficiaryRequestBeneficiaryContactDetailsWithDefaults instantiates a new CreateBeneficiaryRequestBeneficiaryContactDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryEmail

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryEmail() string`

GetBeneficiaryEmail returns the BeneficiaryEmail field if non-nil, zero value otherwise.

### GetBeneficiaryEmailOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryEmailOk() (*string, bool)`

GetBeneficiaryEmailOk returns a tuple with the BeneficiaryEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryEmail

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryEmail(v string)`

SetBeneficiaryEmail sets BeneficiaryEmail field to given value.

### HasBeneficiaryEmail

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryEmail() bool`

HasBeneficiaryEmail returns a boolean if a field has been set.

### GetBeneficiaryPhone

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryPhone() string`

GetBeneficiaryPhone returns the BeneficiaryPhone field if non-nil, zero value otherwise.

### GetBeneficiaryPhoneOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryPhoneOk() (*string, bool)`

GetBeneficiaryPhoneOk returns a tuple with the BeneficiaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPhone

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryPhone(v string)`

SetBeneficiaryPhone sets BeneficiaryPhone field to given value.

### HasBeneficiaryPhone

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryPhone() bool`

HasBeneficiaryPhone returns a boolean if a field has been set.

### GetBeneficiaryCountryCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryCountryCode() string`

GetBeneficiaryCountryCode returns the BeneficiaryCountryCode field if non-nil, zero value otherwise.

### GetBeneficiaryCountryCodeOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryCountryCodeOk() (*string, bool)`

GetBeneficiaryCountryCodeOk returns a tuple with the BeneficiaryCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCountryCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryCountryCode(v string)`

SetBeneficiaryCountryCode sets BeneficiaryCountryCode field to given value.

### HasBeneficiaryCountryCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryCountryCode() bool`

HasBeneficiaryCountryCode returns a boolean if a field has been set.

### GetBeneficiaryAddress

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.

### HasBeneficiaryAddress

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryAddress() bool`

HasBeneficiaryAddress returns a boolean if a field has been set.

### GetBeneficiaryCity

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryCity() string`

GetBeneficiaryCity returns the BeneficiaryCity field if non-nil, zero value otherwise.

### GetBeneficiaryCityOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryCityOk() (*string, bool)`

GetBeneficiaryCityOk returns a tuple with the BeneficiaryCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCity

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryCity(v string)`

SetBeneficiaryCity sets BeneficiaryCity field to given value.

### HasBeneficiaryCity

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryCity() bool`

HasBeneficiaryCity returns a boolean if a field has been set.

### GetBeneficiaryState

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryState() string`

GetBeneficiaryState returns the BeneficiaryState field if non-nil, zero value otherwise.

### GetBeneficiaryStateOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryStateOk() (*string, bool)`

GetBeneficiaryStateOk returns a tuple with the BeneficiaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryState

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryState(v string)`

SetBeneficiaryState sets BeneficiaryState field to given value.

### HasBeneficiaryState

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryState() bool`

HasBeneficiaryState returns a boolean if a field has been set.

### GetBeneficiaryPostalCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryPostalCode() string`

GetBeneficiaryPostalCode returns the BeneficiaryPostalCode field if non-nil, zero value otherwise.

### GetBeneficiaryPostalCodeOk

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) GetBeneficiaryPostalCodeOk() (*string, bool)`

GetBeneficiaryPostalCodeOk returns a tuple with the BeneficiaryPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryPostalCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) SetBeneficiaryPostalCode(v string)`

SetBeneficiaryPostalCode sets BeneficiaryPostalCode field to given value.

### HasBeneficiaryPostalCode

`func (o *CreateBeneficiaryRequestBeneficiaryContactDetails) HasBeneficiaryPostalCode() bool`

HasBeneficiaryPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


