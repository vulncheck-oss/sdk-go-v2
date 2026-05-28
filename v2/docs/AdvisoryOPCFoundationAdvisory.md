# AdvisoryOPCFoundationAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csaf** | Pointer to [**AdvisoryOPCFoundationRef**](AdvisoryOPCFoundationRef.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Gcve** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOPCFoundationAdvisory

`func NewAdvisoryOPCFoundationAdvisory() *AdvisoryOPCFoundationAdvisory`

NewAdvisoryOPCFoundationAdvisory instantiates a new AdvisoryOPCFoundationAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOPCFoundationAdvisoryWithDefaults

`func NewAdvisoryOPCFoundationAdvisoryWithDefaults() *AdvisoryOPCFoundationAdvisory`

NewAdvisoryOPCFoundationAdvisoryWithDefaults instantiates a new AdvisoryOPCFoundationAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsaf

`func (o *AdvisoryOPCFoundationAdvisory) GetCsaf() AdvisoryOPCFoundationRef`

GetCsaf returns the Csaf field if non-nil, zero value otherwise.

### GetCsafOk

`func (o *AdvisoryOPCFoundationAdvisory) GetCsafOk() (*AdvisoryOPCFoundationRef, bool)`

GetCsafOk returns a tuple with the Csaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsaf

`func (o *AdvisoryOPCFoundationAdvisory) SetCsaf(v AdvisoryOPCFoundationRef)`

SetCsaf sets Csaf field to given value.

### HasCsaf

`func (o *AdvisoryOPCFoundationAdvisory) HasCsaf() bool`

HasCsaf returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryOPCFoundationAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOPCFoundationAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOPCFoundationAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOPCFoundationAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryOPCFoundationAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryOPCFoundationAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryOPCFoundationAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryOPCFoundationAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetGcve

`func (o *AdvisoryOPCFoundationAdvisory) GetGcve() string`

GetGcve returns the Gcve field if non-nil, zero value otherwise.

### GetGcveOk

`func (o *AdvisoryOPCFoundationAdvisory) GetGcveOk() (*string, bool)`

GetGcveOk returns a tuple with the Gcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcve

`func (o *AdvisoryOPCFoundationAdvisory) SetGcve(v string)`

SetGcve sets Gcve field to given value.

### HasGcve

`func (o *AdvisoryOPCFoundationAdvisory) HasGcve() bool`

HasGcve returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOPCFoundationAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOPCFoundationAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOPCFoundationAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOPCFoundationAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOPCFoundationAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOPCFoundationAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOPCFoundationAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOPCFoundationAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOPCFoundationAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOPCFoundationAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOPCFoundationAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOPCFoundationAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryOPCFoundationAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryOPCFoundationAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryOPCFoundationAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryOPCFoundationAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryOPCFoundationAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryOPCFoundationAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryOPCFoundationAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryOPCFoundationAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


