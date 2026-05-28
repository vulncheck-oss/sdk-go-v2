# AdvisoryCirclAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CirclRef** | Pointer to [**AdvisoryCirclRef**](AdvisoryCirclRef.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Gcve** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCirclAdvisory

`func NewAdvisoryCirclAdvisory() *AdvisoryCirclAdvisory`

NewAdvisoryCirclAdvisory instantiates a new AdvisoryCirclAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCirclAdvisoryWithDefaults

`func NewAdvisoryCirclAdvisoryWithDefaults() *AdvisoryCirclAdvisory`

NewAdvisoryCirclAdvisoryWithDefaults instantiates a new AdvisoryCirclAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCirclRef

`func (o *AdvisoryCirclAdvisory) GetCirclRef() AdvisoryCirclRef`

GetCirclRef returns the CirclRef field if non-nil, zero value otherwise.

### GetCirclRefOk

`func (o *AdvisoryCirclAdvisory) GetCirclRefOk() (*AdvisoryCirclRef, bool)`

GetCirclRefOk returns a tuple with the CirclRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirclRef

`func (o *AdvisoryCirclAdvisory) SetCirclRef(v AdvisoryCirclRef)`

SetCirclRef sets CirclRef field to given value.

### HasCirclRef

`func (o *AdvisoryCirclAdvisory) HasCirclRef() bool`

HasCirclRef returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCirclAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCirclAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCirclAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCirclAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCirclAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCirclAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCirclAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCirclAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetGcve

`func (o *AdvisoryCirclAdvisory) GetGcve() string`

GetGcve returns the Gcve field if non-nil, zero value otherwise.

### GetGcveOk

`func (o *AdvisoryCirclAdvisory) GetGcveOk() (*string, bool)`

GetGcveOk returns a tuple with the Gcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcve

`func (o *AdvisoryCirclAdvisory) SetGcve(v string)`

SetGcve sets Gcve field to given value.

### HasGcve

`func (o *AdvisoryCirclAdvisory) HasGcve() bool`

HasGcve returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCirclAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCirclAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCirclAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCirclAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryCirclAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryCirclAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryCirclAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryCirclAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCirclAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCirclAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCirclAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCirclAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryCirclAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryCirclAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryCirclAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryCirclAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCirclAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCirclAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCirclAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCirclAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


