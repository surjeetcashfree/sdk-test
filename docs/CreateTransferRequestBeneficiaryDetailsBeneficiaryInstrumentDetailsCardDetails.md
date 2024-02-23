# CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | Pointer to **string** | It is the tokenised card number or card token for this transfer. | [optional] 
**CardNetworkType** | Pointer to **string** | It is the network type of the card - VISA/MASTERCARD. | [optional] 
**CardCryptogram** | Pointer to **string** | It is the formatted chip/cryptogram data relating to the token cryptogram. The maximum character limit is 600. It is optional for MASTERCARD and not required for VISA. | [optional] 
**CardTokenExpiry** | Pointer to **string** | It is applicable only for MASTERCARD. The format for the valid token expiry date should be YYYY-MM. It cannot be null. Provide a valid tokenExpiry if collected from the customers. If unavailable, populate a static value with a forward year and month in the correct format (YYYY-MM). The maximum character limit is 10. | [optional] 
**CardType** | Pointer to **string** | It is the type of the card. DEBIT and CREDIT are the only values allowed. The default value is CREDIT if the parameter does not exist or not specified. | [optional] 
**CardTokenPANSequenceNumber** | Pointer to **string** | A maximum of 3 alphanumeric characters are allowed. It is an optional parameter for MASTERCARD. | [optional] 

## Methods

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetailsWithDefaults

`func NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetailsWithDefaults() *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails`

NewCreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetailsWithDefaults instantiates a new CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetCardNetworkType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardNetworkType() string`

GetCardNetworkType returns the CardNetworkType field if non-nil, zero value otherwise.

### GetCardNetworkTypeOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardNetworkTypeOk() (*string, bool)`

GetCardNetworkTypeOk returns a tuple with the CardNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetworkType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardNetworkType(v string)`

SetCardNetworkType sets CardNetworkType field to given value.

### HasCardNetworkType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardNetworkType() bool`

HasCardNetworkType returns a boolean if a field has been set.

### GetCardCryptogram

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardCryptogram() string`

GetCardCryptogram returns the CardCryptogram field if non-nil, zero value otherwise.

### GetCardCryptogramOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardCryptogramOk() (*string, bool)`

GetCardCryptogramOk returns a tuple with the CardCryptogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCryptogram

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardCryptogram(v string)`

SetCardCryptogram sets CardCryptogram field to given value.

### HasCardCryptogram

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardCryptogram() bool`

HasCardCryptogram returns a boolean if a field has been set.

### GetCardTokenExpiry

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenExpiry() string`

GetCardTokenExpiry returns the CardTokenExpiry field if non-nil, zero value otherwise.

### GetCardTokenExpiryOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenExpiryOk() (*string, bool)`

GetCardTokenExpiryOk returns a tuple with the CardTokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTokenExpiry

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardTokenExpiry(v string)`

SetCardTokenExpiry sets CardTokenExpiry field to given value.

### HasCardTokenExpiry

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardTokenExpiry() bool`

HasCardTokenExpiry returns a boolean if a field has been set.

### GetCardType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardTokenPANSequenceNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenPANSequenceNumber() string`

GetCardTokenPANSequenceNumber returns the CardTokenPANSequenceNumber field if non-nil, zero value otherwise.

### GetCardTokenPANSequenceNumberOk

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) GetCardTokenPANSequenceNumberOk() (*string, bool)`

GetCardTokenPANSequenceNumberOk returns a tuple with the CardTokenPANSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTokenPANSequenceNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) SetCardTokenPANSequenceNumber(v string)`

SetCardTokenPANSequenceNumber sets CardTokenPANSequenceNumber field to given value.

### HasCardTokenPANSequenceNumber

`func (o *CreateTransferRequestBeneficiaryDetailsBeneficiaryInstrumentDetailsCardDetails) HasCardTokenPANSequenceNumber() bool`

HasCardTokenPANSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


