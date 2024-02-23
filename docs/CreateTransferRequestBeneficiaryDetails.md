# CreateTransferRequestBeneficiaryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryId** | Pointer to **string** | It is the unique ID you created to identify the beneficiary. Alphanumeric characters are allowed. | [optional] 
**BeneficiaryName** | Pointer to **string** | It is the name of the beneficiary. The maximum character limit is 100.  Only alphabets and whitespaces are allowed. It is required if beneficiary_id is not present. | [optional] 
**BeneficiaryInstrumentDetails** | Pointer to [**CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails**](CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails.md) |  | [optional] 
**BeneficiaryContactDetails** | Pointer to [**CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails**](CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails.md) |  | [optional] 

## Methods

### NewCreateTransferRequestBeneficiaryDetails

`func NewCreateTransferRequestBeneficiaryDetails() *CreateTransferRequestBeneficiaryDetails`

NewCreateTransferRequestBeneficiaryDetails instantiates a new CreateTransferRequestBeneficiaryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestBeneficiaryDetailsWithDefaults

`func NewCreateTransferRequestBeneficiaryDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetails`

NewCreateTransferRequestBeneficiaryDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryId

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryId() string`

GetBeneficiaryId returns the BeneficiaryId field if non-nil, zero value otherwise.

### GetBeneficiaryIdOk

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryIdOk() (*string, bool)`

GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryId

`func (o *CreateTransferRequestBeneficiaryDetails) SetBeneficiaryId(v string)`

SetBeneficiaryId sets BeneficiaryId field to given value.

### HasBeneficiaryId

`func (o *CreateTransferRequestBeneficiaryDetails) HasBeneficiaryId() bool`

HasBeneficiaryId returns a boolean if a field has been set.

### GetBeneficiaryName

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *CreateTransferRequestBeneficiaryDetails) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *CreateTransferRequestBeneficiaryDetails) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### GetBeneficiaryInstrumentDetails

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryInstrumentDetails() CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails`

GetBeneficiaryInstrumentDetails returns the BeneficiaryInstrumentDetails field if non-nil, zero value otherwise.

### GetBeneficiaryInstrumentDetailsOk

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryInstrumentDetailsOk() (*CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails, bool)`

GetBeneficiaryInstrumentDetailsOk returns a tuple with the BeneficiaryInstrumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInstrumentDetails

`func (o *CreateTransferRequestBeneficiaryDetails) SetBeneficiaryInstrumentDetails(v CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetails)`

SetBeneficiaryInstrumentDetails sets BeneficiaryInstrumentDetails field to given value.

### HasBeneficiaryInstrumentDetails

`func (o *CreateTransferRequestBeneficiaryDetails) HasBeneficiaryInstrumentDetails() bool`

HasBeneficiaryInstrumentDetails returns a boolean if a field has been set.

### GetBeneficiaryContactDetails

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryContactDetails() CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails`

GetBeneficiaryContactDetails returns the BeneficiaryContactDetails field if non-nil, zero value otherwise.

### GetBeneficiaryContactDetailsOk

`func (o *CreateTransferRequestBeneficiaryDetails) GetBeneficiaryContactDetailsOk() (*CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails, bool)`

GetBeneficiaryContactDetailsOk returns a tuple with the BeneficiaryContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryContactDetails

`func (o *CreateTransferRequestBeneficiaryDetails) SetBeneficiaryContactDetails(v CreateTransferRequestBeneficiaryDetailsBeneficiaryContactDetails)`

SetBeneficiaryContactDetails sets BeneficiaryContactDetails field to given value.

### HasBeneficiaryContactDetails

`func (o *CreateTransferRequestBeneficiaryDetails) HasBeneficiaryContactDetails() bool`

HasBeneficiaryContactDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


