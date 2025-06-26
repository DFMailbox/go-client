# \FederationAPI

All URIs are relative to *http://localhost:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrossSend**](FederationAPI.md#CrossSend) | **Post** /federation/mailbox | Send mail to a plot
[**GetChallenge**](FederationAPI.md#GetChallenge) | **Post** /federation/challenge | Create Challenge
[**MarkKeyAsCompromised**](FederationAPI.md#MarkKeyAsCompromised) | **Delete** /federation/instance | Mark instance key as compromised
[**RefreshToken**](FederationAPI.md#RefreshToken) | **Post** /federation/instance | Refresh identity token
[**VerifyIdentity**](FederationAPI.md#VerifyIdentity) | **Get** /federation/instance | Verify the instance&#39;s ownership of the private key



## CrossSend

> CrossSend200Response CrossSend(ctx).CrossSendRequest(crossSendRequest).Execute()

Send mail to a plot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DFMailbox/go-client"
)

func main() {
	crossSendRequest := *openapiclient.NewCrossSendRequest(int32(41808), int32(41808), interface{}(123)) // CrossSendRequest | Where to send the mail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationAPI.CrossSend(context.Background()).CrossSendRequest(crossSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationAPI.CrossSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrossSend`: CrossSend200Response
	fmt.Fprintf(os.Stdout, "Response from `FederationAPI.CrossSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrossSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crossSendRequest** | [**CrossSendRequest**](CrossSendRequest.md) | Where to send the mail | 

### Return type

[**CrossSend200Response**](CrossSend200Response.md)

### Authorization

[Identity](../README.md#Identity)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChallenge

> GetChallenge201Response GetChallenge(ctx).Execute()

Create Challenge



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DFMailbox/go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationAPI.GetChallenge(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationAPI.GetChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChallenge`: GetChallenge201Response
	fmt.Fprintf(os.Stdout, "Response from `FederationAPI.GetChallenge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetChallengeRequest struct via the builder pattern


### Return type

[**GetChallenge201Response**](GetChallenge201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkKeyAsCompromised

> MarkKeyAsCompromised(ctx).PublicKey(publicKey).Challenge(challenge).Signature(signature).Execute()

Mark instance key as compromised



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DFMailbox/go-client"
)

func main() {
	publicKey := "publicKey_example" // string | Your server public key
	challenge := "challenge_example" // string | A challenge from `POST /federation/challenge`
	signature := "signature_example" // string | The signature of it

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FederationAPI.MarkKeyAsCompromised(context.Background()).PublicKey(publicKey).Challenge(challenge).Signature(signature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationAPI.MarkKeyAsCompromised``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkKeyAsCompromisedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKey** | **string** | Your server public key | 
 **challenge** | **string** | A challenge from &#x60;POST /federation/challenge&#x60; | 
 **signature** | **string** | The signature of it | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshToken

> RefreshToken200Response RefreshToken(ctx).RefreshTokenRequest(refreshTokenRequest).Execute()

Refresh identity token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DFMailbox/go-client"
)

func main() {
	refreshTokenRequest := *openapiclient.NewRefreshTokenRequest("Challenge_example", "7lv2/Z05t53d6LyjA+kXFO5gSIO308sgJ3pX5YAB2Kw9wXP6ZztXrYfkUVSRW0b+cIlPul7F5WQ9dPRPuYH3AA==") // RefreshTokenRequest | A challenge from `POST /federation/challenge` and your public key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationAPI.RefreshToken(context.Background()).RefreshTokenRequest(refreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationAPI.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: RefreshToken200Response
	fmt.Fprintf(os.Stdout, "Response from `FederationAPI.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenRequest** | [**RefreshTokenRequest**](RefreshTokenRequest.md) | A challenge from &#x60;POST /federation/challenge&#x60; and your public key | 

### Return type

[**RefreshToken200Response**](RefreshToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyIdentity

> VerifyIdentity200Response VerifyIdentity(ctx).Challenge(challenge).Execute()

Verify the instance's ownership of the private key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/DFMailbox/go-client"
)

func main() {
	challenge := "challenge_example" // string | A random uuid that will get signed

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationAPI.VerifyIdentity(context.Background()).Challenge(challenge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationAPI.VerifyIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyIdentity`: VerifyIdentity200Response
	fmt.Fprintf(os.Stdout, "Response from `FederationAPI.VerifyIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **challenge** | **string** | A random uuid that will get signed | 

### Return type

[**VerifyIdentity200Response**](VerifyIdentity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

