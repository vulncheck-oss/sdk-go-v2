# AdvisoryK8S

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**IssueId** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryK8S

`func NewAdvisoryK8S() *AdvisoryK8S`

NewAdvisoryK8S instantiates a new AdvisoryK8S object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryK8SWithDefaults

`func NewAdvisoryK8SWithDefaults() *AdvisoryK8S`

NewAdvisoryK8SWithDefaults instantiates a new AdvisoryK8S object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AdvisoryK8S) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AdvisoryK8S) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AdvisoryK8S) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AdvisoryK8S) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryK8S) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryK8S) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryK8S) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryK8S) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryK8S) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryK8S) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryK8S) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryK8S) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetIssueId

`func (o *AdvisoryK8S) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *AdvisoryK8S) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *AdvisoryK8S) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *AdvisoryK8S) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryK8S) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryK8S) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryK8S) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryK8S) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryK8S) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryK8S) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryK8S) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryK8S) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


