# AdvisoryDebianCVE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **string** |  | [optional] 
**Debianbug** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Releases** | Pointer to [**[]AdvisoryAffectedDebianRelease**](AdvisoryAffectedDebianRelease.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDebianCVE

`func NewAdvisoryDebianCVE() *AdvisoryDebianCVE`

NewAdvisoryDebianCVE instantiates a new AdvisoryDebianCVE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDebianCVEWithDefaults

`func NewAdvisoryDebianCVEWithDefaults() *AdvisoryDebianCVE`

NewAdvisoryDebianCVEWithDefaults instantiates a new AdvisoryDebianCVE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryDebianCVE) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDebianCVE) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDebianCVE) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDebianCVE) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDebianbug

`func (o *AdvisoryDebianCVE) GetDebianbug() int32`

GetDebianbug returns the Debianbug field if non-nil, zero value otherwise.

### GetDebianbugOk

`func (o *AdvisoryDebianCVE) GetDebianbugOk() (*int32, bool)`

GetDebianbugOk returns a tuple with the Debianbug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebianbug

`func (o *AdvisoryDebianCVE) SetDebianbug(v int32)`

SetDebianbug sets Debianbug field to given value.

### HasDebianbug

`func (o *AdvisoryDebianCVE) HasDebianbug() bool`

HasDebianbug returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryDebianCVE) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryDebianCVE) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryDebianCVE) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryDebianCVE) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReleases

`func (o *AdvisoryDebianCVE) GetReleases() []AdvisoryAffectedDebianRelease`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *AdvisoryDebianCVE) GetReleasesOk() (*[]AdvisoryAffectedDebianRelease, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *AdvisoryDebianCVE) SetReleases(v []AdvisoryAffectedDebianRelease)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *AdvisoryDebianCVE) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetScope

`func (o *AdvisoryDebianCVE) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AdvisoryDebianCVE) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AdvisoryDebianCVE) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AdvisoryDebianCVE) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryDebianCVE) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryDebianCVE) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryDebianCVE) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryDebianCVE) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


