# AdvisoryOpenSSLSecAdv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]AdvisoryOpenSSLVulnerability**](AdvisoryOpenSSLVulnerability.md) |  | [optional] 

## Methods

### NewAdvisoryOpenSSLSecAdv

`func NewAdvisoryOpenSSLSecAdv() *AdvisoryOpenSSLSecAdv`

NewAdvisoryOpenSSLSecAdv instantiates a new AdvisoryOpenSSLSecAdv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOpenSSLSecAdvWithDefaults

`func NewAdvisoryOpenSSLSecAdvWithDefaults() *AdvisoryOpenSSLSecAdv`

NewAdvisoryOpenSSLSecAdvWithDefaults instantiates a new AdvisoryOpenSSLSecAdv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryOpenSSLSecAdv) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOpenSSLSecAdv) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOpenSSLSecAdv) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOpenSSLSecAdv) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryOpenSSLSecAdv) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryOpenSSLSecAdv) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryOpenSSLSecAdv) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryOpenSSLSecAdv) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateUpdated

`func (o *AdvisoryOpenSSLSecAdv) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AdvisoryOpenSSLSecAdv) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AdvisoryOpenSSLSecAdv) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AdvisoryOpenSSLSecAdv) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryOpenSSLSecAdv) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryOpenSSLSecAdv) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryOpenSSLSecAdv) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryOpenSSLSecAdv) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilename

`func (o *AdvisoryOpenSSLSecAdv) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdvisoryOpenSSLSecAdv) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdvisoryOpenSSLSecAdv) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AdvisoryOpenSSLSecAdv) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOpenSSLSecAdv) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOpenSSLSecAdv) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOpenSSLSecAdv) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOpenSSLSecAdv) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOpenSSLSecAdv) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOpenSSLSecAdv) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOpenSSLSecAdv) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOpenSSLSecAdv) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryOpenSSLSecAdv) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryOpenSSLSecAdv) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryOpenSSLSecAdv) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryOpenSSLSecAdv) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisoryOpenSSLSecAdv) GetVulnerabilities() []AdvisoryOpenSSLVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisoryOpenSSLSecAdv) GetVulnerabilitiesOk() (*[]AdvisoryOpenSSLVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisoryOpenSSLSecAdv) SetVulnerabilities(v []AdvisoryOpenSSLVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisoryOpenSSLSecAdv) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


