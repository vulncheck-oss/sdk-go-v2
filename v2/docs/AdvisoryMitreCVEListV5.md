# AdvisoryMitreCVEListV5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MitreRef** | Pointer to [**AdvisoryMitreCVEListV5Ref**](AdvisoryMitreCVEListV5Ref.md) |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMitreCVEListV5

`func NewAdvisoryMitreCVEListV5() *AdvisoryMitreCVEListV5`

NewAdvisoryMitreCVEListV5 instantiates a new AdvisoryMitreCVEListV5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMitreCVEListV5WithDefaults

`func NewAdvisoryMitreCVEListV5WithDefaults() *AdvisoryMitreCVEListV5`

NewAdvisoryMitreCVEListV5WithDefaults instantiates a new AdvisoryMitreCVEListV5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryMitreCVEListV5) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMitreCVEListV5) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMitreCVEListV5) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMitreCVEListV5) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMitreCVEListV5) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMitreCVEListV5) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMitreCVEListV5) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMitreCVEListV5) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMitreRef

`func (o *AdvisoryMitreCVEListV5) GetMitreRef() AdvisoryMitreCVEListV5Ref`

GetMitreRef returns the MitreRef field if non-nil, zero value otherwise.

### GetMitreRefOk

`func (o *AdvisoryMitreCVEListV5) GetMitreRefOk() (*AdvisoryMitreCVEListV5Ref, bool)`

GetMitreRefOk returns a tuple with the MitreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreRef

`func (o *AdvisoryMitreCVEListV5) SetMitreRef(v AdvisoryMitreCVEListV5Ref)`

SetMitreRef sets MitreRef field to given value.

### HasMitreRef

`func (o *AdvisoryMitreCVEListV5) HasMitreRef() bool`

HasMitreRef returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMitreCVEListV5) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMitreCVEListV5) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMitreCVEListV5) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMitreCVEListV5) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryMitreCVEListV5) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryMitreCVEListV5) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryMitreCVEListV5) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryMitreCVEListV5) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMitreCVEListV5) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMitreCVEListV5) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMitreCVEListV5) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMitreCVEListV5) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryMitreCVEListV5) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryMitreCVEListV5) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryMitreCVEListV5) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryMitreCVEListV5) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMitreCVEListV5) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMitreCVEListV5) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMitreCVEListV5) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMitreCVEListV5) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


