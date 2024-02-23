# \BeneficiaryV2API

All URIs are relative to *https://sandbox.cashfree.com/payout*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayoutCreateBeneficiary**](BeneficiaryV2API.md#PayoutCreateBeneficiary) | **Post** /beneficiary | Create Beneficiary V2
[**PayoutDeleteBeneficiary**](BeneficiaryV2API.md#PayoutDeleteBeneficiary) | **Delete** /beneficiary | Remove Beneficiary V2
[**PayoutFetchBeneficiary**](BeneficiaryV2API.md#PayoutFetchBeneficiary) | **Get** /beneficiary | Get Beneficiary V2



## PayoutCreateBeneficiary

> Beneficiary PayoutCreateBeneficiary(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CreateBeneficiaryRequest(createBeneficiaryRequest).Execute()

Create Beneficiary V2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/surjeetcashfree/cashfree-pg/v4"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2024-01-01" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    createBeneficiaryRequest := *cashfree.NewCreateBeneficiaryRequest("JOHN18011343", "John Doe") 

    resp, r, err := cashfree.PayoutCreateBeneficiary(&xApiVersion, &xRequestId, &createBeneficiaryRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutCreateBeneficiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutCreateBeneficiary`: Beneficiary
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutCreateBeneficiary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutCreateBeneficiaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **createBeneficiaryRequest** | [**CreateBeneficiaryRequest**](CreateBeneficiaryRequest.md) | Find the request parameters to create a beneficiary | 

### Return type

[**Beneficiary**](Beneficiary.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: applicate/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutDeleteBeneficiary

> Beneficiary PayoutDeleteBeneficiary(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).BeneficiaryId(beneficiaryId).Execute()

Remove Beneficiary V2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/surjeetcashfree/cashfree-pg/v4"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2024-01-01" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    beneficiaryId := "sample_test" 

    resp, r, err := cashfree.PayoutDeleteBeneficiary(&xApiVersion, &xRequestId, &beneficiaryId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutDeleteBeneficiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutDeleteBeneficiary`: Beneficiary
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutDeleteBeneficiary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutDeleteBeneficiaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **beneficiaryId** | **string** | It is the unique ID you create to identify the beneficiary. The maximum character limit is 50. Only alphabets and whitespaces are allowed. | 

### Return type

[**Beneficiary**](Beneficiary.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutFetchBeneficiary

> Beneficiary PayoutFetchBeneficiary(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).BeneficiaryId(beneficiaryId).BankAccountNumber(bankAccountNumber).BankIfsc(bankIfsc).Execute()

Get Beneficiary V2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    cashfree "github.com/surjeetcashfree/cashfree-pg/v4"
)

func main() {

    clientId := "<x-client-id>"
	clientSecret := "<x-client-secret>"
	cashfree.XClientId = &clientId
	cashfree.XClientSecret = &clientSecret
	cashfree.XEnvironment = cashfree.SANDBOX

    xApiVersion := "2024-01-01" 
    xRequestId := "4dfb9780-46fe-11ee-be56-0242ac120002" 
    beneficiaryId := "JOHN18011343" 
    bankAccountNumber := "00001111222233" 
    bankIfsc := "HDFC0000001" 

    resp, r, err := cashfree.PayoutFetchBeneficiary(&xApiVersion, &xRequestId, &beneficiaryId, &bankAccountNumber, &bankIfsc, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutFetchBeneficiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutFetchBeneficiary`: Beneficiary
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutFetchBeneficiary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutFetchBeneficiaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **beneficiaryId** | **string** | It is the unique ID you created to identify the beneficary while creating the beneficiary. | 
 **bankAccountNumber** | **string** | It is the bank account number of the beneficiary. If bank_account_number is passed in query, bank_ifsc should also be passed. | 
 **bankIfsc** | **string** | It is the IFSC information as present in the beneficiary&#39;s bank account information. If bank_ifsc is passed in query, bank_account_number should also be passed. | 

### Return type

[**Beneficiary**](Beneficiary.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

