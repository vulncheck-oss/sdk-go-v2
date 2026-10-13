# NvdcpesNvdCpe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeName** | Pointer to **string** |  | [optional] 
**CpeNameId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Deprecated** | Pointer to **bool** |  | [optional] 
**DeprecatedBy** | Pointer to [**[]NvdcpesNvdCpeDeprecation**](NvdcpesNvdCpeDeprecation.md) |  | [optional] 
**Deprecates** | Pointer to [**[]NvdcpesNvdCpeDeprecation**](NvdcpesNvdCpeDeprecation.md) |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Refs** | Pointer to [**[]NvdcpesNvdCpeReference**](NvdcpesNvdCpeReference.md) |  | [optional] 
**Titles** | Pointer to [**[]NvdcpesNvdCpeTitle**](NvdcpesNvdCpeTitle.md) |  | [optional] 

## Methods

### NewNvdcpesNvdCpe

`func NewNvdcpesNvdCpe() *NvdcpesNvdCpe`

NewNvdcpesNvdCpe instantiates a new NvdcpesNvdCpe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNvdcpesNvdCpeWithDefaults

`func NewNvdcpesNvdCpeWithDefaults() *NvdcpesNvdCpe`

NewNvdcpesNvdCpeWithDefaults instantiates a new NvdcpesNvdCpe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeName

`func (o *NvdcpesNvdCpe) GetCpeName() string`

GetCpeName returns the CpeName field if non-nil, zero value otherwise.

### GetCpeNameOk

`func (o *NvdcpesNvdCpe) GetCpeNameOk() (*string, bool)`

GetCpeNameOk returns a tuple with the CpeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeName

`func (o *NvdcpesNvdCpe) SetCpeName(v string)`

SetCpeName sets CpeName field to given value.

### HasCpeName

`func (o *NvdcpesNvdCpe) HasCpeName() bool`

HasCpeName returns a boolean if a field has been set.

### GetCpeNameId

`func (o *NvdcpesNvdCpe) GetCpeNameId() string`

GetCpeNameId returns the CpeNameId field if non-nil, zero value otherwise.

### GetCpeNameIdOk

`func (o *NvdcpesNvdCpe) GetCpeNameIdOk() (*string, bool)`

GetCpeNameIdOk returns a tuple with the CpeNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeNameId

`func (o *NvdcpesNvdCpe) SetCpeNameId(v string)`

SetCpeNameId sets CpeNameId field to given value.

### HasCpeNameId

`func (o *NvdcpesNvdCpe) HasCpeNameId() bool`

HasCpeNameId returns a boolean if a field has been set.

### GetCreated

`func (o *NvdcpesNvdCpe) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NvdcpesNvdCpe) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NvdcpesNvdCpe) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NvdcpesNvdCpe) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeprecated

`func (o *NvdcpesNvdCpe) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *NvdcpesNvdCpe) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *NvdcpesNvdCpe) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *NvdcpesNvdCpe) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecatedBy

`func (o *NvdcpesNvdCpe) GetDeprecatedBy() []NvdcpesNvdCpeDeprecation`

GetDeprecatedBy returns the DeprecatedBy field if non-nil, zero value otherwise.

### GetDeprecatedByOk

`func (o *NvdcpesNvdCpe) GetDeprecatedByOk() (*[]NvdcpesNvdCpeDeprecation, bool)`

GetDeprecatedByOk returns a tuple with the DeprecatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedBy

`func (o *NvdcpesNvdCpe) SetDeprecatedBy(v []NvdcpesNvdCpeDeprecation)`

SetDeprecatedBy sets DeprecatedBy field to given value.

### HasDeprecatedBy

`func (o *NvdcpesNvdCpe) HasDeprecatedBy() bool`

HasDeprecatedBy returns a boolean if a field has been set.

### GetDeprecates

`func (o *NvdcpesNvdCpe) GetDeprecates() []NvdcpesNvdCpeDeprecation`

GetDeprecates returns the Deprecates field if non-nil, zero value otherwise.

### GetDeprecatesOk

`func (o *NvdcpesNvdCpe) GetDeprecatesOk() (*[]NvdcpesNvdCpeDeprecation, bool)`

GetDeprecatesOk returns a tuple with the Deprecates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecates

`func (o *NvdcpesNvdCpe) SetDeprecates(v []NvdcpesNvdCpeDeprecation)`

SetDeprecates sets Deprecates field to given value.

### HasDeprecates

`func (o *NvdcpesNvdCpe) HasDeprecates() bool`

HasDeprecates returns a boolean if a field has been set.

### GetLastModified

`func (o *NvdcpesNvdCpe) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NvdcpesNvdCpe) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NvdcpesNvdCpe) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *NvdcpesNvdCpe) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetRefs

`func (o *NvdcpesNvdCpe) GetRefs() []NvdcpesNvdCpeReference`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *NvdcpesNvdCpe) GetRefsOk() (*[]NvdcpesNvdCpeReference, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *NvdcpesNvdCpe) SetRefs(v []NvdcpesNvdCpeReference)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *NvdcpesNvdCpe) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### GetTitles

`func (o *NvdcpesNvdCpe) GetTitles() []NvdcpesNvdCpeTitle`

GetTitles returns the Titles field if non-nil, zero value otherwise.

### GetTitlesOk

`func (o *NvdcpesNvdCpe) GetTitlesOk() (*[]NvdcpesNvdCpeTitle, bool)`

GetTitlesOk returns a tuple with the Titles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitles

`func (o *NvdcpesNvdCpe) SetTitles(v []NvdcpesNvdCpeTitle)`

SetTitles sets Titles field to given value.

### HasTitles

`func (o *NvdcpesNvdCpe) HasTitles() bool`

HasTitles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


