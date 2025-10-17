# AdvisoryAIX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssScore** | Pointer to **string** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAIX

`func NewAdvisoryAIX() *AdvisoryAIX`

NewAdvisoryAIX instantiates a new AdvisoryAIX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAIXWithDefaults

`func NewAdvisoryAIXWithDefaults() *AdvisoryAIX`

NewAdvisoryAIXWithDefaults instantiates a new AdvisoryAIX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryAIX) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAIX) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAIX) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAIX) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryAIX) GetCvssScore() string`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryAIX) GetCvssScoreOk() (*string, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryAIX) SetCvssScore(v string)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryAIX) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryAIX) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryAIX) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryAIX) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryAIX) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAIX) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAIX) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAIX) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAIX) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryAIX) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryAIX) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryAIX) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryAIX) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryAIX) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryAIX) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryAIX) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryAIX) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryAIX) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryAIX) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryAIX) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryAIX) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryAIX) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryAIX) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryAIX) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryAIX) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryAIX) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryAIX) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryAIX) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryAIX) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


