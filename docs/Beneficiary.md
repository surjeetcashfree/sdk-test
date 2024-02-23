# Beneficiary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryId** | Pointer to **string** | It displays the unique ID you created to identify the beneficiary. | [optional] 
**BeneficaryName** | Pointer to **string** | It displays the name of the beneficiary. | [optional] 
**BeneficiaryInstrumentDetails** | Pointer to [**BeneficiaryBeneficiaryInstrumentDetails**](BeneficiaryBeneficiaryInstrumentDetails.md) |  | [optional] 
**BeneficiaryContactDetails** | Pointer to [**BeneficiaryBeneficiaryContactDetails**](BeneficiaryBeneficiaryContactDetails.md) |  | [optional] 
**BeneficiaryStatus** | Pointer to **string** | It displays the current status of the beneficiary. Possible values are as follows - &#x60;VERIFIED&#x60;: Beneficiary is verified and is available for payouts - &#x60;INVALID&#x60;: Beneficiary is invalid - &#x60;INITIATED&#x60;: Beneficiary verification initiated - &#x60;CANCELLED&#x60;: Beneficiary verification cancelled - &#x60;FAILED&#x60;: Failed to verify beneficiary - &#x60;DELETED&#x60;: Beneficiary is deleted | [optional] 
**AddedOn** | Pointer to **string** | It displays the time of the addition of the beneficiary in UTC. | [optional] 

## Methods

### NewBeneficiary

`func NewBeneficiary() *Beneficiary`

NewBeneficiary instantiates a new Beneficiary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeneficiaryWithDefaults

`func NewBeneficiaryWithDefaults() *Beneficiary`

NewBeneficiaryWithDefaults instantiates a new Beneficiary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryId

`func (o *Beneficiary) GetBeneficiaryId() string`

GetBeneficiaryId returns the BeneficiaryId field if non-nil, zero value otherwise.

### GetBeneficiaryIdOk

`func (o *Beneficiary) GetBeneficiaryIdOk() (*string, bool)`

GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryId

`func (o *Beneficiary) SetBeneficiaryId(v string)`

SetBeneficiaryId sets BeneficiaryId field to given value.

### HasBeneficiaryId

`func (o *Beneficiary) HasBeneficiaryId() bool`

HasBeneficiaryId returns a boolean if a field has been set.

### GetBeneficaryName

`func (o *Beneficiary) GetBeneficaryName() string`

GetBeneficaryName returns the BeneficaryName field if non-nil, zero value otherwise.

### GetBeneficaryNameOk

`func (o *Beneficiary) GetBeneficaryNameOk() (*string, bool)`

GetBeneficaryNameOk returns a tuple with the BeneficaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficaryName

`func (o *Beneficiary) SetBeneficaryName(v string)`

SetBeneficaryName sets BeneficaryName field to given value.

### HasBeneficaryName

`func (o *Beneficiary) HasBeneficaryName() bool`

HasBeneficaryName returns a boolean if a field has been set.

### GetBeneficiaryInstrumentDetails

`func (o *Beneficiary) GetBeneficiaryInstrumentDetails() BeneficiaryBeneficiaryInstrumentDetails`

GetBeneficiaryInstrumentDetails returns the BeneficiaryInstrumentDetails field if non-nil, zero value otherwise.

### GetBeneficiaryInstrumentDetailsOk

`func (o *Beneficiary) GetBeneficiaryInstrumentDetailsOk() (*BeneficiaryBeneficiaryInstrumentDetails, bool)`

GetBeneficiaryInstrumentDetailsOk returns a tuple with the BeneficiaryInstrumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInstrumentDetails

`func (o *Beneficiary) SetBeneficiaryInstrumentDetails(v BeneficiaryBeneficiaryInstrumentDetails)`

SetBeneficiaryInstrumentDetails sets BeneficiaryInstrumentDetails field to given value.

### HasBeneficiaryInstrumentDetails

`func (o *Beneficiary) HasBeneficiaryInstrumentDetails() bool`

HasBeneficiaryInstrumentDetails returns a boolean if a field has been set.

### GetBeneficiaryContactDetails

`func (o *Beneficiary) GetBeneficiaryContactDetails() BeneficiaryBeneficiaryContactDetails`

GetBeneficiaryContactDetails returns the BeneficiaryContactDetails field if non-nil, zero value otherwise.

### GetBeneficiaryContactDetailsOk

`func (o *Beneficiary) GetBeneficiaryContactDetailsOk() (*BeneficiaryBeneficiaryContactDetails, bool)`

GetBeneficiaryContactDetailsOk returns a tuple with the BeneficiaryContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryContactDetails

`func (o *Beneficiary) SetBeneficiaryContactDetails(v BeneficiaryBeneficiaryContactDetails)`

SetBeneficiaryContactDetails sets BeneficiaryContactDetails field to given value.

### HasBeneficiaryContactDetails

`func (o *Beneficiary) HasBeneficiaryContactDetails() bool`

HasBeneficiaryContactDetails returns a boolean if a field has been set.

### GetBeneficiaryStatus

`func (o *Beneficiary) GetBeneficiaryStatus() string`

GetBeneficiaryStatus returns the BeneficiaryStatus field if non-nil, zero value otherwise.

### GetBeneficiaryStatusOk

`func (o *Beneficiary) GetBeneficiaryStatusOk() (*string, bool)`

GetBeneficiaryStatusOk returns a tuple with the BeneficiaryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryStatus

`func (o *Beneficiary) SetBeneficiaryStatus(v string)`

SetBeneficiaryStatus sets BeneficiaryStatus field to given value.

### HasBeneficiaryStatus

`func (o *Beneficiary) HasBeneficiaryStatus() bool`

HasBeneficiaryStatus returns a boolean if a field has been set.

### GetAddedOn

`func (o *Beneficiary) GetAddedOn() string`

GetAddedOn returns the AddedOn field if non-nil, zero value otherwise.

### GetAddedOnOk

`func (o *Beneficiary) GetAddedOnOk() (*string, bool)`

GetAddedOnOk returns a tuple with the AddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedOn

`func (o *Beneficiary) SetAddedOn(v string)`

SetAddedOn sets AddedOn field to given value.

### HasAddedOn

`func (o *Beneficiary) HasAddedOn() bool`

HasAddedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


