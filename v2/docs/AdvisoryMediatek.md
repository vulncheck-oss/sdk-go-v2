# AdvisoryMediatek

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedChipsets** | Pointer to **[]string** |  | [optional] 
**AffectedSoftware** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMediatek

`func NewAdvisoryMediatek() *AdvisoryMediatek`

NewAdvisoryMediatek instantiates a new AdvisoryMediatek object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMediatekWithDefaults

`func NewAdvisoryMediatekWithDefaults() *AdvisoryMediatek`

NewAdvisoryMediatekWithDefaults instantiates a new AdvisoryMediatek object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedChipsets

`func (o *AdvisoryMediatek) GetAffectedChipsets() []string`

GetAffectedChipsets returns the AffectedChipsets field if non-nil, zero value otherwise.

### GetAffectedChipsetsOk

`func (o *AdvisoryMediatek) GetAffectedChipsetsOk() (*[]string, bool)`

GetAffectedChipsetsOk returns a tuple with the AffectedChipsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedChipsets

`func (o *AdvisoryMediatek) SetAffectedChipsets(v []string)`

SetAffectedChipsets sets AffectedChipsets field to given value.

### HasAffectedChipsets

`func (o *AdvisoryMediatek) HasAffectedChipsets() bool`

HasAffectedChipsets returns a boolean if a field has been set.

### GetAffectedSoftware

`func (o *AdvisoryMediatek) GetAffectedSoftware() []string`

GetAffectedSoftware returns the AffectedSoftware field if non-nil, zero value otherwise.

### GetAffectedSoftwareOk

`func (o *AdvisoryMediatek) GetAffectedSoftwareOk() (*[]string, bool)`

GetAffectedSoftwareOk returns a tuple with the AffectedSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSoftware

`func (o *AdvisoryMediatek) SetAffectedSoftware(v []string)`

SetAffectedSoftware sets AffectedSoftware field to given value.

### HasAffectedSoftware

`func (o *AdvisoryMediatek) HasAffectedSoftware() bool`

HasAffectedSoftware returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryMediatek) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMediatek) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMediatek) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMediatek) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMediatek) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMediatek) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMediatek) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMediatek) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryMediatek) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMediatek) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMediatek) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMediatek) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryMediatek) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryMediatek) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryMediatek) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryMediatek) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMediatek) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMediatek) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMediatek) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMediatek) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryMediatek) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryMediatek) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryMediatek) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryMediatek) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMediatek) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMediatek) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMediatek) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMediatek) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


