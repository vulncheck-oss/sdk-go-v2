# AdvisoryLolAdvs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LolJson** | Pointer to **map[string]interface{}** |  | [optional] 
**MitreId** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryLolAdvs

`func NewAdvisoryLolAdvs() *AdvisoryLolAdvs`

NewAdvisoryLolAdvs instantiates a new AdvisoryLolAdvs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryLolAdvsWithDefaults

`func NewAdvisoryLolAdvsWithDefaults() *AdvisoryLolAdvs`

NewAdvisoryLolAdvsWithDefaults instantiates a new AdvisoryLolAdvs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryLolAdvs) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryLolAdvs) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryLolAdvs) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryLolAdvs) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryLolAdvs) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryLolAdvs) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryLolAdvs) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryLolAdvs) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryLolAdvs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryLolAdvs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryLolAdvs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryLolAdvs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryLolAdvs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryLolAdvs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryLolAdvs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryLolAdvs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLolJson

`func (o *AdvisoryLolAdvs) GetLolJson() map[string]interface{}`

GetLolJson returns the LolJson field if non-nil, zero value otherwise.

### GetLolJsonOk

`func (o *AdvisoryLolAdvs) GetLolJsonOk() (*map[string]interface{}, bool)`

GetLolJsonOk returns a tuple with the LolJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLolJson

`func (o *AdvisoryLolAdvs) SetLolJson(v map[string]interface{})`

SetLolJson sets LolJson field to given value.

### HasLolJson

`func (o *AdvisoryLolAdvs) HasLolJson() bool`

HasLolJson returns a boolean if a field has been set.

### GetMitreId

`func (o *AdvisoryLolAdvs) GetMitreId() string`

GetMitreId returns the MitreId field if non-nil, zero value otherwise.

### GetMitreIdOk

`func (o *AdvisoryLolAdvs) GetMitreIdOk() (*string, bool)`

GetMitreIdOk returns a tuple with the MitreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreId

`func (o *AdvisoryLolAdvs) SetMitreId(v string)`

SetMitreId sets MitreId field to given value.

### HasMitreId

`func (o *AdvisoryLolAdvs) HasMitreId() bool`

HasMitreId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryLolAdvs) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryLolAdvs) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryLolAdvs) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryLolAdvs) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryLolAdvs) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryLolAdvs) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryLolAdvs) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryLolAdvs) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryLolAdvs) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryLolAdvs) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryLolAdvs) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryLolAdvs) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


