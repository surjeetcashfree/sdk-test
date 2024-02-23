# CreateBeneficiaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryId** | **string** | It is the unique ID you create to identify the beneficiary. Alphanumeric, underscore ( _ ), pipe ( | ), and dot ( . ) are allowed. | 
**BeneficaryName** | **string** | It is the name of the beneficiary. The maximum character limit is 100. Only alphabets and whitespaces are allowed. | 
**BeneficiaryInstrumentDetails** | Pointer to [**CreateBeneficiaryRequestBeneficiaryInstrumentDetails**](CreateBeneficiaryRequestBeneficiaryInstrumentDetails.md) |  | [optional] 
**BeneficiaryContactDetails** | Pointer to [**CreateBeneficiaryRequestBeneficiaryContactDetails**](CreateBeneficiaryRequestBeneficiaryContactDetails.md) |  | [optional] 

## Methods

### NewCreateBeneficiaryRequest

`func NewCreateBeneficiaryRequest(beneficiaryId string, beneficaryName string, ) *CreateBeneficiaryRequest`

NewCreateBeneficiaryRequest instantiates a new CreateBeneficiaryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBeneficiaryRequestWithDefaults

`func NewCreateBeneficiaryRequestWithDefaults() *CreateBeneficiaryRequest`

NewCreateBeneficiaryRequestWithDefaults instantiates a new CreateBeneficiaryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryId

`func (o *CreateBeneficiaryRequest) GetBeneficiaryId() string`

GetBeneficiaryId returns the BeneficiaryId field if non-nil, zero value otherwise.

### GetBeneficiaryIdOk

`func (o *CreateBeneficiaryRequest) GetBeneficiaryIdOk() (*string, bool)`

GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryId

`func (o *CreateBeneficiaryRequest) SetBeneficiaryId(v string)`

SetBeneficiaryId sets BeneficiaryId field to given value.


### GetBeneficaryName

`func (o *CreateBeneficiaryRequest) GetBeneficaryName() string`

GetBeneficaryName returns the BeneficaryName field if non-nil, zero value otherwise.

### GetBeneficaryNameOk

`func (o *CreateBeneficiaryRequest) GetBeneficaryNameOk() (*string, bool)`

GetBeneficaryNameOk returns a tuple with the BeneficaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficaryName

`func (o *CreateBeneficiaryRequest) SetBeneficaryName(v string)`

SetBeneficaryName sets BeneficaryName field to given value.


### GetBeneficiaryInstrumentDetails

`func (o *CreateBeneficiaryRequest) GetBeneficiaryInstrumentDetails() CreateBeneficiaryRequestBeneficiaryInstrumentDetails`

GetBeneficiaryInstrumentDetails returns the BeneficiaryInstrumentDetails field if non-nil, zero value otherwise.

### GetBeneficiaryInstrumentDetailsOk

`func (o *CreateBeneficiaryRequest) GetBeneficiaryInstrumentDetailsOk() (*CreateBeneficiaryRequestBeneficiaryInstrumentDetails, bool)`

GetBeneficiaryInstrumentDetailsOk returns a tuple with the BeneficiaryInstrumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInstrumentDetails

`func (o *CreateBeneficiaryRequest) SetBeneficiaryInstrumentDetails(v CreateBeneficiaryRequestBeneficiaryInstrumentDetails)`

SetBeneficiaryInstrumentDetails sets BeneficiaryInstrumentDetails field to given value.

### HasBeneficiaryInstrumentDetails

`func (o *CreateBeneficiaryRequest) HasBeneficiaryInstrumentDetails() bool`

HasBeneficiaryInstrumentDetails returns a boolean if a field has been set.

### GetBeneficiaryContactDetails

`func (o *CreateBeneficiaryRequest) GetBeneficiaryContactDetails() CreateBeneficiaryRequestBeneficiaryContactDetails`

GetBeneficiaryContactDetails returns the BeneficiaryContactDetails field if non-nil, zero value otherwise.

### GetBeneficiaryContactDetailsOk

`func (o *CreateBeneficiaryRequest) GetBeneficiaryContactDetailsOk() (*CreateBeneficiaryRequestBeneficiaryContactDetails, bool)`

GetBeneficiaryContactDetailsOk returns a tuple with the BeneficiaryContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryContactDetails

`func (o *CreateBeneficiaryRequest) SetBeneficiaryContactDetails(v CreateBeneficiaryRequestBeneficiaryContactDetails)`

SetBeneficiaryContactDetails sets BeneficiaryContactDetails field to given value.

### HasBeneficiaryContactDetails

`func (o *CreateBeneficiaryRequest) HasBeneficiaryContactDetails() bool`

HasBeneficiaryContactDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


