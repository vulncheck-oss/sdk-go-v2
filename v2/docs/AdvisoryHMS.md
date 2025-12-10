# AdvisoryHMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Mitigation** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHMS

`func NewAdvisoryHMS() *AdvisoryHMS`

NewAdvisoryHMS instantiates a new AdvisoryHMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHMSWithDefaults

`func NewAdvisoryHMSWithDefaults() *AdvisoryHMS`

NewAdvisoryHMSWithDefaults instantiates a new AdvisoryHMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryHMS) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryHMS) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryHMS) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryHMS) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHMS) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHMS) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHMS) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHMS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHMS) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHMS) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHMS) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHMS) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryHMS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryHMS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryHMS) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryHMS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisoryHMS) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisoryHMS) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisoryHMS) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisoryHMS) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetMitigation

`func (o *AdvisoryHMS) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *AdvisoryHMS) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *AdvisoryHMS) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *AdvisoryHMS) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryHMS) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryHMS) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryHMS) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryHMS) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryHMS) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryHMS) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryHMS) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryHMS) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryHMS) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryHMS) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryHMS) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryHMS) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryHMS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryHMS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryHMS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryHMS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


