# \InstanceAPI

All URIs are relative to *http://localhost:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntroduceInstance**](InstanceAPI.md#IntroduceInstance) | **Post** /instance | Introduce an instance
[**LookupInstanceAddress**](InstanceAPI.md#LookupInstanceAddress) | **Get** /instance | Get instance address by key



## IntroduceInstance

> IntroduceInstance(ctx).IntroduceInstanceRequest(introduceInstanceRequest).Execute()

Introduce an instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	introduceInstanceRequest := *openapiclient.NewIntroduceInstanceRequest("0nqH2kJLWxfqdH0QIsKJp84Ck9OhPWCHZ3uMmcoknSY", "api.dfmailbox.dev") // IntroduceInstanceRequest | The instance information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.IntroduceInstance(context.Background()).IntroduceInstanceRequest(introduceInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.IntroduceInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntroduceInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **introduceInstanceRequest** | [**IntroduceInstanceRequest**](IntroduceInstanceRequest.md) | The instance information | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupInstanceAddress

> LookupInstanceAddress200Response LookupInstanceAddress(ctx).PublicKey(publicKey).Execute()

Get instance address by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	publicKey := "publicKey_example" // string | The server key to check the address of

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.LookupInstanceAddress(context.Background()).PublicKey(publicKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.LookupInstanceAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupInstanceAddress`: LookupInstanceAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.LookupInstanceAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupInstanceAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string** | The server key to check the address of | 

### Return type

[**LookupInstanceAddress200Response**](LookupInstanceAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

