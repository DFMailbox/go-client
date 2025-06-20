# \PlotAPI

All URIs are relative to *http://localhost:8080/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCleanMailbox**](PlotAPI.md#CheckCleanMailbox) | **Delete** /plot/mailbox | Check and cleanup mailbox
[**CheckMailbox**](PlotAPI.md#CheckMailbox) | **Get** /plot/mailbox | Check mailbox for items
[**CreateApiKey**](PlotAPI.md#CreateApiKey) | **Post** /plot/api-key | Create an API Key
[**GetApiKeyHashes**](PlotAPI.md#GetApiKeyHashes) | **Get** /plot/api-key | List all API key hashes
[**GetPlotInfo**](PlotAPI.md#GetPlotInfo) | **Get** /plot | Get plot info
[**GetTrusted**](PlotAPI.md#GetTrusted) | **Get** /plot/trust | Get trusted plots
[**LookupPlot**](PlotAPI.md#LookupPlot) | **Get** /plots/{plot_id} | Get another plot&#39;s info
[**Query**](PlotAPI.md#Query) | **Post** /plot/query | Run mailbox query
[**RegisterPlot**](PlotAPI.md#RegisterPlot) | **Post** /plot | Register plot
[**RevokeAllApiKeys**](PlotAPI.md#RevokeAllApiKeys) | **Delete** /plot/api-key | Revoke all API keys
[**SendMail**](PlotAPI.md#SendMail) | **Post** /plots/{plot_id}/mailbox | Send an item into another mailbox
[**SendMailToSelf**](PlotAPI.md#SendMailToSelf) | **Post** /plot/mailbox | Send items to mailbox
[**TrustPlots**](PlotAPI.md#TrustPlots) | **Post** /plot/trust | Trust plots
[**UnregisterPlot**](PlotAPI.md#UnregisterPlot) | **Delete** /plot | Unregister plot
[**UntrustPlot**](PlotAPI.md#UntrustPlot) | **Delete** /plot/trust | Un-trust plots
[**UpdateInstance**](PlotAPI.md#UpdateInstance) | **Put** /plot | Change plot instance



## CheckCleanMailbox

> FetchMailboxResult CheckCleanMailbox(ctx).MsgId(msgId).Limit(limit).Execute()

Check and cleanup mailbox



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
	msgId := int64(0) // int64 | The id to get messages after
	limit := int64(50) // int64 | The max amount of messsages that can be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.CheckCleanMailbox(context.Background()).MsgId(msgId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.CheckCleanMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCleanMailbox`: FetchMailboxResult
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.CheckCleanMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCleanMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgId** | **int64** | The id to get messages after | 
 **limit** | **int64** | The max amount of messsages that can be returned | 

### Return type

[**FetchMailboxResult**](FetchMailboxResult.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckMailbox

> FetchMailboxResult CheckMailbox(ctx).MsgId(msgId).Limit(limit).Execute()

Check mailbox for items



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
	msgId := int64(0) // int64 | The id to get messages after
	limit := int64(50) // int64 | The max amount of messsages that can be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.CheckMailbox(context.Background()).MsgId(msgId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.CheckMailbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMailbox`: FetchMailboxResult
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.CheckMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msgId** | **int64** | The id to get messages after | 
 **limit** | **int64** | The max amount of messsages that can be returned | 

### Return type

[**FetchMailboxResult**](FetchMailboxResult.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiKey

> CreateApiKey200Response CreateApiKey(ctx).Execute()

Create an API Key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.CreateApiKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.CreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiKey`: CreateApiKey200Response
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


### Return type

[**CreateApiKey200Response**](CreateApiKey200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyHashes

> interface{} GetApiKeyHashes(ctx).Execute()

List all API key hashes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.GetApiKeyHashes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.GetApiKeyHashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeyHashes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.GetApiKeyHashes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyHashesRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlotInfo

> Plot GetPlotInfo(ctx).Execute()

Get plot info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.GetPlotInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.GetPlotInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlotInfo`: Plot
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.GetPlotInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlotInfoRequest struct via the builder pattern


### Return type

[**Plot**](Plot.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrusted

> interface{} GetTrusted(ctx).Execute()

Get trusted plots



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.GetTrusted(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.GetTrusted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrusted`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.GetTrusted`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustedRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupPlot

> Plot LookupPlot(ctx, plotId).Execute()

Get another plot's info



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
	plotId := int32(56) // int32 | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot <plot_id>

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.LookupPlot(context.Background(), plotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.LookupPlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupPlot`: Plot
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.LookupPlot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plotId** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupPlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plot**](Plot.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Query

> interface{} Query(ctx).Body(body).Execute()

Run mailbox query



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
	body := interface{}([{"type":"send","to":42998,"value":[{"type":"chat","name":"DynamicCake","msg":"The `id: chat` field isn't defined by diamondfire, rather by the plots themselves"},{"type":"chat","name":"DynamicCake","msg":"This field is merely a convention, not a standard"}]},{"type":"peek_clean","after":55}]) // interface{} | A list of mailbox query operations

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.Query(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.Query``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Query`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.Query`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | A list of mailbox query operations | 

### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPlot

> RegisterPlot(ctx).UpdateInstanceRequest(updateInstanceRequest).Execute()

Register plot



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
	updateInstanceRequest := *openapiclient.NewUpdateInstanceRequest("Instance_example") // UpdateInstanceRequest | The public key of the owned instance. If registering to this instance, use a null

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlotAPI.RegisterPlot(context.Background()).UpdateInstanceRequest(updateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.RegisterPlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInstanceRequest** | [**UpdateInstanceRequest**](UpdateInstanceRequest.md) | The public key of the owned instance. If registering to this instance, use a null | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeAllApiKeys

> RevokeAllApiKeys(ctx).Execute()

Revoke all API keys



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlotAPI.RevokeAllApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.RevokeAllApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeAllApiKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMail

> InlineObject SendMail(ctx, plotId).Body(body).Execute()

Send an item into another mailbox



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
	plotId := int32(56) // int32 | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot <plot_id>
	body := interface{}([{id=str, val=Hello DFMailbox}, {id=vec, x=1, y=2, z=3}]) // interface{} | Items to send. The first item goes in first, last item goes in last.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.SendMail(context.Background(), plotId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.SendMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMail`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.SendMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plotId** | **int32** | An id assigned by DiamondFire to identify a plot, this ID can be used in /plot &lt;plot_id&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | Items to send. The first item goes in first, last item goes in last. | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMailToSelf

> InlineObject SendMailToSelf(ctx).Body(body).Execute()

Send items to mailbox



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
	body := interface{}([{id=str, val=Hello DFMailbox}, {id=vec, x=1, y=2, z=3}]) // interface{} | Items to send. The first item goes in first, last item goes in last.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.SendMailToSelf(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.SendMailToSelf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMailToSelf`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.SendMailToSelf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMailToSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | Items to send. The first item goes in first, last item goes in last. | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustPlots

> TrustPlots(ctx).Body(body).Execute()

Trust plots



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
	body := interface{}(987) // interface{} | Plots to trust

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlotAPI.TrustPlots(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.TrustPlots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrustPlotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **interface{}** | Plots to trust | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterPlot

> UnregisterPlot(ctx).Execute()

Unregister plot



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlotAPI.UnregisterPlot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.UnregisterPlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterPlotRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UntrustPlot

> interface{} UntrustPlot(ctx).Plots(plots).Execute()

Un-trust plots



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
	plots := "plots_example" // string | A list of comma separated plots to un-trust, empty means un-trust all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlotAPI.UntrustPlot(context.Background()).Plots(plots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.UntrustPlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UntrustPlot`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlotAPI.UntrustPlot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUntrustPlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plots** | **string** | A list of comma separated plots to un-trust, empty means un-trust all | 

### Return type

**interface{}**

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> UpdateInstance(ctx).UpdateInstanceRequest(updateInstanceRequest).Execute()

Change plot instance



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
	updateInstanceRequest := *openapiclient.NewUpdateInstanceRequest("Instance_example") // UpdateInstanceRequest | The public key of your new instance, null for this instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlotAPI.UpdateInstance(context.Background()).UpdateInstanceRequest(updateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlotAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInstanceRequest** | [**UpdateInstanceRequest**](UpdateInstanceRequest.md) | The public key of your new instance, null for this instance | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [Plot](../README.md#Plot)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

