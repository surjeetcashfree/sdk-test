# \TransfersV2API

All URIs are relative to *https://sandbox.cashfree.com/payout*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayoutFetchBatchTransfer**](TransfersV2API.md#PayoutFetchBatchTransfer) | **Get** /transfers/batch | Get Batch Transfer Status V2
[**PayoutFetchTransfer**](TransfersV2API.md#PayoutFetchTransfer) | **Get** /transfers | Get Transfer Status V2
[**PayoutInitiateBatchTransfer**](TransfersV2API.md#PayoutInitiateBatchTransfer) | **Post** /transfers/batch | Batch Transfer V2
[**PayoutInitiateTransfer**](TransfersV2API.md#PayoutInitiateTransfer) | **Post** /transfers | Standard Transfer V2



## PayoutFetchBatchTransfer

> FetchBatchTransferResponse PayoutFetchBatchTransfer(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CfBatchTransferId(cfBatchTransferId).BatchTransferId(batchTransferId).Execute()

Get Batch Transfer Status V2



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
    cfBatchTransferId := "10023" 
    batchTransferId := "Sample_test" 

    resp, r, err := cashfree.PayoutFetchBatchTransfer(&xApiVersion, &xRequestId, &cfBatchTransferId, &batchTransferId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutFetchBatchTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutFetchBatchTransfer`: FetchBatchTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutFetchBatchTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutFetchBatchTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **cfBatchTransferId** | **string** | It is the unique ID created by Cashfree Payments. You receive it in the response of the initiated batch transfer request. (Either cf_batch_transfer_id or batch_transfer_id is mandatory) | 
 **batchTransferId** | **string** | It is the unique ID you created to identify the batch transfer request. | 

### Return type

[**FetchBatchTransferResponse**](FetchBatchTransferResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutFetchTransfer

> CreateTransferResponse PayoutFetchTransfer(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CfTransferId(cfTransferId).TransferId(transferId).Execute()

Get Transfer Status V2



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
    cfTransferId := "10023" 
    transferId := "Sample_test" 

    resp, r, err := cashfree.PayoutFetchTransfer(&xApiVersion, &xRequestId, &cfTransferId, &transferId, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutFetchTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutFetchTransfer`: CreateTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutFetchTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutFetchTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **cfTransferId** | **string** | It is the unique ID created by Cashfree Payments. You receive it in the response of the initiated standard transfer request. (Either cf_transfer_id or transfer_id is mandatory) | 
 **transferId** | **string** | It is the unique ID you created to identify the standard transfer request. | 

### Return type

[**CreateTransferResponse**](CreateTransferResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutInitiateBatchTransfer

> CreateBatchTransferResponse PayoutInitiateBatchTransfer(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CreateBatchTransferRequest(createBatchTransferRequest).Execute()

Batch Transfer V2



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
    createBatchTransferRequest := *cashfree.NewCreateBatchTransferRequest("BatchTransferId_example") 

    resp, r, err := cashfree.PayoutInitiateBatchTransfer(&xApiVersion, &xRequestId, &createBatchTransferRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutInitiateBatchTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutInitiateBatchTransfer`: CreateBatchTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutInitiateBatchTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutInitiateBatchTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **createBatchTransferRequest** | [**CreateBatchTransferRequest**](CreateBatchTransferRequest.md) | Find the request parameters of the API request to transfer money to multiple beneficiaries. | 

### Return type

[**CreateBatchTransferResponse**](CreateBatchTransferResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutInitiateTransfer

> CreateTransferResponse PayoutInitiateTransfer(ctx).XApiVersion(xApiVersion).XRequestId(xRequestId).CreateTransferRequest(createTransferRequest).Execute()

Standard Transfer V2



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
    createTransferRequest := *cashfree.NewCreateTransferRequest("TransferId_example", float64(123)) 

    resp, r, err := cashfree.PayoutInitiateTransfer(&xApiVersion, &xRequestId, &createTransferRequest, nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `cashfree.PayoutInitiateTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayoutInitiateTransfer`: CreateTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `cashfree.PayoutInitiateTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutInitiateTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | It is the API version to be used. The accepted format is YYYY-MM-DD. | [default to &quot;2024-01-01&quot;]
 **xRequestId** | **string** | It is the request ID for the API call. This ID can be used to resolve tech realted issues. Communicate this in your tech related queries to Cashfree Payments. | 
 **createTransferRequest** | [**CreateTransferRequest**](CreateTransferRequest.md) | Find the request parameters of the API request to transfer money to a beneficiary. | 

### Return type

[**CreateTransferResponse**](CreateTransferResponse.md)

### Authorization

[XClientSecret](../README.md#XClientSecret), [XClientID](../README.md#XClientID)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

