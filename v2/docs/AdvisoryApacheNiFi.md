# AdvisoryApacheNiFi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedVersion** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**FixedVersions** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryApacheNiFi

`func NewAdvisoryApacheNiFi() *AdvisoryApacheNiFi`

NewAdvisoryApacheNiFi instantiates a new AdvisoryApacheNiFi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryApacheNiFiWithDefaults

`func NewAdvisoryApacheNiFiWithDefaults() *AdvisoryApacheNiFi`

NewAdvisoryApacheNiFiWithDefaults instantiates a new AdvisoryApacheNiFi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedVersion

`func (o *AdvisoryApacheNiFi) GetAffectedVersion() string`

GetAffectedVersion returns the AffectedVersion field if non-nil, zero value otherwise.

### GetAffectedVersionOk

`func (o *AdvisoryApacheNiFi) GetAffectedVersionOk() (*string, bool)`

GetAffectedVersionOk returns a tuple with the AffectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersion

`func (o *AdvisoryApacheNiFi) SetAffectedVersion(v string)`

SetAffectedVersion sets AffectedVersion field to given value.

### HasAffectedVersion

`func (o *AdvisoryApacheNiFi) HasAffectedVersion() bool`

HasAffectedVersion returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryApacheNiFi) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryApacheNiFi) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryApacheNiFi) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryApacheNiFi) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryApacheNiFi) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryApacheNiFi) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryApacheNiFi) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryApacheNiFi) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixedVersions

`func (o *AdvisoryApacheNiFi) GetFixedVersions() string`

GetFixedVersions returns the FixedVersions field if non-nil, zero value otherwise.

### GetFixedVersionsOk

`func (o *AdvisoryApacheNiFi) GetFixedVersionsOk() (*string, bool)`

GetFixedVersionsOk returns a tuple with the FixedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersions

`func (o *AdvisoryApacheNiFi) SetFixedVersions(v string)`

SetFixedVersions sets FixedVersions field to given value.

### HasFixedVersions

`func (o *AdvisoryApacheNiFi) HasFixedVersions() bool`

HasFixedVersions returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryApacheNiFi) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryApacheNiFi) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryApacheNiFi) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryApacheNiFi) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryApacheNiFi) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryApacheNiFi) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryApacheNiFi) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryApacheNiFi) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryApacheNiFi) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryApacheNiFi) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryApacheNiFi) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryApacheNiFi) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryApacheNiFi) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryApacheNiFi) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryApacheNiFi) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryApacheNiFi) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


