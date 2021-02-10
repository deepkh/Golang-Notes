# {{classname}}

All URIs are relative to *http://api.example.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileuploadPost**](DefaultApi.md#FileuploadPost) | **Post** /fileupload | Upload a file
[**FormdatauploadPost**](DefaultApi.md#FormdatauploadPost) | **Post** /formdataupload | Upload by using \&quot;multipart/form-data\&quot;
[**UserPost**](DefaultApi.md#UserPost) | **Post** /user | Creates a new user
[**UserUserIdsortsortGet**](DefaultApi.md#UserUserIdsortsortGet) | **Get** /user/{userId}?sort&#x3D;{sort} | Returns a user information by using userId

# **FileuploadPost**
> Status FileuploadPost(ctx, contentDisposition, optional)
Upload a file

Upload a file description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentDisposition** | **string**| attachment; filename&#x3D;\&quot;file_name\&quot; | 
 **optional** | ***DefaultApiFileuploadPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFileuploadPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Object**](Object.md)|  | 

### Return type

[**Status**](status.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FormdatauploadPost**
> string FormdatauploadPost(ctx, optional)
Upload by using \"multipart/form-data\"

Upload by using \"multipart/form-data\"  description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiFormdatauploadPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFormdatauploadPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.**|  | 
 **userId** | **optional.**|  | 
 **fileName** | **optional.Interface of *os.File****optional.**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPost**
> Status UserPost(ctx, body)
Creates a new user

Creates a new user description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 

### Return type

[**Status**](status.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdsortsortGet**
> User UserUserIdsortsortGet(ctx, sort, userId, xAccessToken)
Returns a user information by using userId

Returns a user information by using userId (description)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sort** | **string**| Sort order:  * &#x60;asc&#x60; - Ascending, from A to Z  * &#x60;desc&#x60; - Descending, from Z to A  | 
  **userId** | **int32**| Query the user information from User Id | 
  **xAccessToken** | [**string**](.md)| Access Token which encoded by base64 | 

### Return type

[**User**](user.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

