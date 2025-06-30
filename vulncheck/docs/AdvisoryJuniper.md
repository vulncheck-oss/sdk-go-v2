# AdvisoryJuniper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss3Score** | Pointer to **string** |  | [optional] 
**Cvss3Vector** | Pointer to **string** |  | [optional] 
**Cvss4Score** | Pointer to **string** |  | [optional] 
**Cvss4Vector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Fixed** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryJuniper

`func NewAdvisoryJuniper() *AdvisoryJuniper`

NewAdvisoryJuniper instantiates a new AdvisoryJuniper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryJuniperWithDefaults

`func NewAdvisoryJuniperWithDefaults() *AdvisoryJuniper`

NewAdvisoryJuniperWithDefaults instantiates a new AdvisoryJuniper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryJuniper) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryJuniper) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryJuniper) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryJuniper) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryJuniper) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryJuniper) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryJuniper) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryJuniper) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss3Score

`func (o *AdvisoryJuniper) GetCvss3Score() string`

GetCvss3Score returns the Cvss3Score field if non-nil, zero value otherwise.

### GetCvss3ScoreOk

`func (o *AdvisoryJuniper) GetCvss3ScoreOk() (*string, bool)`

GetCvss3ScoreOk returns a tuple with the Cvss3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Score

`func (o *AdvisoryJuniper) SetCvss3Score(v string)`

SetCvss3Score sets Cvss3Score field to given value.

### HasCvss3Score

`func (o *AdvisoryJuniper) HasCvss3Score() bool`

HasCvss3Score returns a boolean if a field has been set.

### GetCvss3Vector

`func (o *AdvisoryJuniper) GetCvss3Vector() string`

GetCvss3Vector returns the Cvss3Vector field if non-nil, zero value otherwise.

### GetCvss3VectorOk

`func (o *AdvisoryJuniper) GetCvss3VectorOk() (*string, bool)`

GetCvss3VectorOk returns a tuple with the Cvss3Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Vector

`func (o *AdvisoryJuniper) SetCvss3Vector(v string)`

SetCvss3Vector sets Cvss3Vector field to given value.

### HasCvss3Vector

`func (o *AdvisoryJuniper) HasCvss3Vector() bool`

HasCvss3Vector returns a boolean if a field has been set.

### GetCvss4Score

`func (o *AdvisoryJuniper) GetCvss4Score() string`

GetCvss4Score returns the Cvss4Score field if non-nil, zero value otherwise.

### GetCvss4ScoreOk

`func (o *AdvisoryJuniper) GetCvss4ScoreOk() (*string, bool)`

GetCvss4ScoreOk returns a tuple with the Cvss4Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss4Score

`func (o *AdvisoryJuniper) SetCvss4Score(v string)`

SetCvss4Score sets Cvss4Score field to given value.

### HasCvss4Score

`func (o *AdvisoryJuniper) HasCvss4Score() bool`

HasCvss4Score returns a boolean if a field has been set.

### GetCvss4Vector

`func (o *AdvisoryJuniper) GetCvss4Vector() string`

GetCvss4Vector returns the Cvss4Vector field if non-nil, zero value otherwise.

### GetCvss4VectorOk

`func (o *AdvisoryJuniper) GetCvss4VectorOk() (*string, bool)`

GetCvss4VectorOk returns a tuple with the Cvss4Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss4Vector

`func (o *AdvisoryJuniper) SetCvss4Vector(v string)`

SetCvss4Vector sets Cvss4Vector field to given value.

### HasCvss4Vector

`func (o *AdvisoryJuniper) HasCvss4Vector() bool`

HasCvss4Vector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryJuniper) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryJuniper) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryJuniper) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryJuniper) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryJuniper) GetFixed() string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryJuniper) GetFixedOk() (*string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryJuniper) SetFixed(v string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryJuniper) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryJuniper) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryJuniper) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryJuniper) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryJuniper) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryJuniper) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryJuniper) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryJuniper) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryJuniper) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryJuniper) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryJuniper) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryJuniper) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryJuniper) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryJuniper) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryJuniper) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryJuniper) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryJuniper) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryJuniper) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryJuniper) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryJuniper) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryJuniper) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryJuniper) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryJuniper) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryJuniper) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryJuniper) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


