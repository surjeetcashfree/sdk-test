# ErrorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the error received | [optional] 
**Code** | Pointer to **string** | Code of the error received | [optional] 
**Message** | Pointer to **string** | Detailed message explaining the response | [optional] 

## Methods

### NewErrorV2

`func NewErrorV2() *ErrorV2`

NewErrorV2 instantiates a new ErrorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorV2WithDefaults

`func NewErrorV2WithDefaults() *ErrorV2`

NewErrorV2WithDefaults instantiates a new ErrorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ErrorV2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorV2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorV2) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorV2) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorV2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorV2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorV2) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorV2) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


