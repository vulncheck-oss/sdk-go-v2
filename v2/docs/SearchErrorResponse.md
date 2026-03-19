# SearchErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchErrorResponse

`func NewSearchErrorResponse() *SearchErrorResponse`

NewSearchErrorResponse instantiates a new SearchErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchErrorResponseWithDefaults

`func NewSearchErrorResponseWithDefaults() *SearchErrorResponse`

NewSearchErrorResponseWithDefaults instantiates a new SearchErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SearchErrorResponse) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchErrorResponse) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchErrorResponse) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *SearchErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrors

`func (o *SearchErrorResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SearchErrorResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SearchErrorResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SearchErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


