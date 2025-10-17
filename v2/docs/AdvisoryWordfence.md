# AdvisoryWordfence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssScore** | Pointer to **string** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Fixed** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryWordfence

`func NewAdvisoryWordfence() *AdvisoryWordfence`

NewAdvisoryWordfence instantiates a new AdvisoryWordfence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryWordfenceWithDefaults

`func NewAdvisoryWordfenceWithDefaults() *AdvisoryWordfence`

NewAdvisoryWordfenceWithDefaults instantiates a new AdvisoryWordfence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryWordfence) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryWordfence) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryWordfence) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryWordfence) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryWordfence) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryWordfence) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryWordfence) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryWordfence) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryWordfence) GetCvssScore() string`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryWordfence) GetCvssScoreOk() (*string, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryWordfence) SetCvssScore(v string)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryWordfence) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryWordfence) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryWordfence) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryWordfence) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryWordfence) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryWordfence) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryWordfence) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryWordfence) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryWordfence) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryWordfence) GetFixed() []string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryWordfence) GetFixedOk() (*[]string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryWordfence) SetFixed(v []string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryWordfence) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryWordfence) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryWordfence) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryWordfence) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryWordfence) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryWordfence) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryWordfence) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryWordfence) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryWordfence) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryWordfence) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryWordfence) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryWordfence) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryWordfence) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryWordfence) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryWordfence) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryWordfence) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryWordfence) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryWordfence) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryWordfence) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryWordfence) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryWordfence) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


