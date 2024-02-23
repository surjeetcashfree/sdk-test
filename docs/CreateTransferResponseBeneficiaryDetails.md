# CreateTransferResponseBeneficiaryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryId** | Pointer to **string** | It displays the unique ID to identify the beneficiary to whom you initiated the transfer request. | [optional] 
**BeneficiaryInstrumentDetails** | Pointer to [**CreateTransferResponseBeneficiaryDetailsBeneficiaryInstrumentDetails**](CreateTransferResponseBeneficiaryDetailsBeneficiaryInstrumentDetails.md) |  | [optional] 

## Methods

### NewCreateTransferResponseBeneficiaryDetails

`func NewCreateTransferResponseBeneficiaryDetails() *CreateTransferResponseBeneficiaryDetails`

NewCreateTransferResponseBeneficiaryDetails instantiates a new CreateTransferResponseBeneficiaryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferResponseBeneficiaryDetailsWithDefaults

`func NewCreateTransferResponseBeneficiaryDetailsWithDefaults() *CreateTransferResponseBeneficiaryDetails`

NewCreateTransferResponseBeneficiaryDetailsWithDefaults instantiates a new CreateTransferResponseBeneficiaryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeneficiaryId

`func (o *CreateTransferResponseBeneficiaryDetails) GetBeneficiaryId() string`

GetBeneficiaryId returns the BeneficiaryId field if non-nil, zero value otherwise.

### GetBeneficiaryIdOk

`func (o *CreateTransferResponseBeneficiaryDetails) GetBeneficiaryIdOk() (*string, bool)`

GetBeneficiaryIdOk returns a tuple with the BeneficiaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryId

`func (o *CreateTransferResponseBeneficiaryDetails) SetBeneficiaryId(v string)`

SetBeneficiaryId sets BeneficiaryId field to given value.

### HasBeneficiaryId

`func (o *CreateTransferResponseBeneficiaryDetails) HasBeneficiaryId() bool`

HasBeneficiaryId returns a boolean if a field has been set.

### GetBeneficiaryInstrumentDetails

`func (o *CreateTransferResponseBeneficiaryDetails) GetBeneficiaryInstrumentDetails() CreateTransferResponseBeneficiaryDetailsBeneficiaryInstrumentDetails`

GetBeneficiaryInstrumentDetails returns the BeneficiaryInstrumentDetails field if non-nil, zero value otherwise.

### GetBeneficiaryInstrumentDetailsOk

`func (o *CreateTransferResponseBeneficiaryDetails) GetBeneficiaryInstrumentDetailsOk() (*CreateTransferResponseBeneficiaryDetailsBeneficiaryInstrumentDetails, bool)`

GetBeneficiaryInstrumentDetailsOk returns a tuple with the BeneficiaryInstrumentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryInstrumentDetails

`func (o *CreateTransferResponseBeneficiaryDetails) SetBeneficiaryInstrumentDetails(v CreateTransferResponseBeneficiaryDetailsBeneficiaryInstrumentDetails)`

SetBeneficiaryInstrumentDetails sets BeneficiaryInstrumentDetails field to given value.

### HasBeneficiaryInstrumentDetails

`func (o *CreateTransferResponseBeneficiaryDetails) HasBeneficiaryInstrumentDetails() bool`

HasBeneficiaryInstrumentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


