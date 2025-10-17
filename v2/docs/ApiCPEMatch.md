# ApiCPEMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe22Uri** | Pointer to **string** |  | [optional] 
**Cpe23Uri** | Pointer to **string** |  | [optional] 
**CpeName** | Pointer to [**[]ApiCPEName**](ApiCPEName.md) |  | [optional] 
**VersionEndExcluding** | Pointer to **string** |  | [optional] 
**VersionEndIncluding** | Pointer to **string** |  | [optional] 
**VersionStartExcluding** | Pointer to **string** |  | [optional] 
**VersionStartIncluding** | Pointer to **string** |  | [optional] 
**Vulnerable** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiCPEMatch

`func NewApiCPEMatch() *ApiCPEMatch`

NewApiCPEMatch instantiates a new ApiCPEMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCPEMatchWithDefaults

`func NewApiCPEMatchWithDefaults() *ApiCPEMatch`

NewApiCPEMatchWithDefaults instantiates a new ApiCPEMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe22Uri

`func (o *ApiCPEMatch) GetCpe22Uri() string`

GetCpe22Uri returns the Cpe22Uri field if non-nil, zero value otherwise.

### GetCpe22UriOk

`func (o *ApiCPEMatch) GetCpe22UriOk() (*string, bool)`

GetCpe22UriOk returns a tuple with the Cpe22Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe22Uri

`func (o *ApiCPEMatch) SetCpe22Uri(v string)`

SetCpe22Uri sets Cpe22Uri field to given value.

### HasCpe22Uri

`func (o *ApiCPEMatch) HasCpe22Uri() bool`

HasCpe22Uri returns a boolean if a field has been set.

### GetCpe23Uri

`func (o *ApiCPEMatch) GetCpe23Uri() string`

GetCpe23Uri returns the Cpe23Uri field if non-nil, zero value otherwise.

### GetCpe23UriOk

`func (o *ApiCPEMatch) GetCpe23UriOk() (*string, bool)`

GetCpe23UriOk returns a tuple with the Cpe23Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe23Uri

`func (o *ApiCPEMatch) SetCpe23Uri(v string)`

SetCpe23Uri sets Cpe23Uri field to given value.

### HasCpe23Uri

`func (o *ApiCPEMatch) HasCpe23Uri() bool`

HasCpe23Uri returns a boolean if a field has been set.

### GetCpeName

`func (o *ApiCPEMatch) GetCpeName() []ApiCPEName`

GetCpeName returns the CpeName field if non-nil, zero value otherwise.

### GetCpeNameOk

`func (o *ApiCPEMatch) GetCpeNameOk() (*[]ApiCPEName, bool)`

GetCpeNameOk returns a tuple with the CpeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeName

`func (o *ApiCPEMatch) SetCpeName(v []ApiCPEName)`

SetCpeName sets CpeName field to given value.

### HasCpeName

`func (o *ApiCPEMatch) HasCpeName() bool`

HasCpeName returns a boolean if a field has been set.

### GetVersionEndExcluding

`func (o *ApiCPEMatch) GetVersionEndExcluding() string`

GetVersionEndExcluding returns the VersionEndExcluding field if non-nil, zero value otherwise.

### GetVersionEndExcludingOk

`func (o *ApiCPEMatch) GetVersionEndExcludingOk() (*string, bool)`

GetVersionEndExcludingOk returns a tuple with the VersionEndExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndExcluding

`func (o *ApiCPEMatch) SetVersionEndExcluding(v string)`

SetVersionEndExcluding sets VersionEndExcluding field to given value.

### HasVersionEndExcluding

`func (o *ApiCPEMatch) HasVersionEndExcluding() bool`

HasVersionEndExcluding returns a boolean if a field has been set.

### GetVersionEndIncluding

`func (o *ApiCPEMatch) GetVersionEndIncluding() string`

GetVersionEndIncluding returns the VersionEndIncluding field if non-nil, zero value otherwise.

### GetVersionEndIncludingOk

`func (o *ApiCPEMatch) GetVersionEndIncludingOk() (*string, bool)`

GetVersionEndIncludingOk returns a tuple with the VersionEndIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndIncluding

`func (o *ApiCPEMatch) SetVersionEndIncluding(v string)`

SetVersionEndIncluding sets VersionEndIncluding field to given value.

### HasVersionEndIncluding

`func (o *ApiCPEMatch) HasVersionEndIncluding() bool`

HasVersionEndIncluding returns a boolean if a field has been set.

### GetVersionStartExcluding

`func (o *ApiCPEMatch) GetVersionStartExcluding() string`

GetVersionStartExcluding returns the VersionStartExcluding field if non-nil, zero value otherwise.

### GetVersionStartExcludingOk

`func (o *ApiCPEMatch) GetVersionStartExcludingOk() (*string, bool)`

GetVersionStartExcludingOk returns a tuple with the VersionStartExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartExcluding

`func (o *ApiCPEMatch) SetVersionStartExcluding(v string)`

SetVersionStartExcluding sets VersionStartExcluding field to given value.

### HasVersionStartExcluding

`func (o *ApiCPEMatch) HasVersionStartExcluding() bool`

HasVersionStartExcluding returns a boolean if a field has been set.

### GetVersionStartIncluding

`func (o *ApiCPEMatch) GetVersionStartIncluding() string`

GetVersionStartIncluding returns the VersionStartIncluding field if non-nil, zero value otherwise.

### GetVersionStartIncludingOk

`func (o *ApiCPEMatch) GetVersionStartIncludingOk() (*string, bool)`

GetVersionStartIncludingOk returns a tuple with the VersionStartIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartIncluding

`func (o *ApiCPEMatch) SetVersionStartIncluding(v string)`

SetVersionStartIncluding sets VersionStartIncluding field to given value.

### HasVersionStartIncluding

`func (o *ApiCPEMatch) HasVersionStartIncluding() bool`

HasVersionStartIncluding returns a boolean if a field has been set.

### GetVulnerable

`func (o *ApiCPEMatch) GetVulnerable() bool`

GetVulnerable returns the Vulnerable field if non-nil, zero value otherwise.

### GetVulnerableOk

`func (o *ApiCPEMatch) GetVulnerableOk() (*bool, bool)`

GetVulnerableOk returns a tuple with the Vulnerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerable

`func (o *ApiCPEMatch) SetVulnerable(v bool)`

SetVulnerable sets Vulnerable field to given value.

### HasVulnerable

`func (o *ApiCPEMatch) HasVulnerable() bool`

HasVulnerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


