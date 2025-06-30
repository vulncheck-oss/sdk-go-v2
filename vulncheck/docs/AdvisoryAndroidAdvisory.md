# AdvisoryAndroidAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryAndroidAffected**](AdvisoryAndroidAffected.md) |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryAndroidReference**](AdvisoryAndroidReference.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAndroidAdvisory

`func NewAdvisoryAndroidAdvisory() *AdvisoryAndroidAdvisory`

NewAdvisoryAndroidAdvisory instantiates a new AdvisoryAndroidAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAndroidAdvisoryWithDefaults

`func NewAdvisoryAndroidAdvisoryWithDefaults() *AdvisoryAndroidAdvisory`

NewAdvisoryAndroidAdvisoryWithDefaults instantiates a new AdvisoryAndroidAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryAndroidAdvisory) GetAffected() []AdvisoryAndroidAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryAndroidAdvisory) GetAffectedOk() (*[]AdvisoryAndroidAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryAndroidAdvisory) SetAffected(v []AdvisoryAndroidAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryAndroidAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryAndroidAdvisory) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryAndroidAdvisory) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryAndroidAdvisory) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryAndroidAdvisory) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAndroidAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAndroidAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAndroidAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAndroidAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAndroidAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAndroidAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAndroidAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAndroidAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryAndroidAdvisory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryAndroidAdvisory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryAndroidAdvisory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryAndroidAdvisory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryAndroidAdvisory) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryAndroidAdvisory) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryAndroidAdvisory) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryAndroidAdvisory) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryAndroidAdvisory) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryAndroidAdvisory) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryAndroidAdvisory) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryAndroidAdvisory) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryAndroidAdvisory) GetReferences() []AdvisoryAndroidReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryAndroidAdvisory) GetReferencesOk() (*[]AdvisoryAndroidReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryAndroidAdvisory) SetReferences(v []AdvisoryAndroidReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryAndroidAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryAndroidAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryAndroidAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryAndroidAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryAndroidAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


