# AdvisoryMendix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Fixed** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MendixId** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMendix

`func NewAdvisoryMendix() *AdvisoryMendix`

NewAdvisoryMendix instantiates a new AdvisoryMendix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMendixWithDefaults

`func NewAdvisoryMendixWithDefaults() *AdvisoryMendix`

NewAdvisoryMendixWithDefaults instantiates a new AdvisoryMendix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryMendix) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryMendix) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryMendix) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryMendix) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryMendix) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMendix) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMendix) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMendix) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMendix) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMendix) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMendix) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMendix) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryMendix) GetFixed() []string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryMendix) GetFixedOk() (*[]string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryMendix) SetFixed(v []string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryMendix) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryMendix) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryMendix) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryMendix) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryMendix) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMendixId

`func (o *AdvisoryMendix) GetMendixId() string`

GetMendixId returns the MendixId field if non-nil, zero value otherwise.

### GetMendixIdOk

`func (o *AdvisoryMendix) GetMendixIdOk() (*string, bool)`

GetMendixIdOk returns a tuple with the MendixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMendixId

`func (o *AdvisoryMendix) SetMendixId(v string)`

SetMendixId sets MendixId field to given value.

### HasMendixId

`func (o *AdvisoryMendix) HasMendixId() bool`

HasMendixId returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryMendix) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryMendix) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryMendix) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryMendix) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMendix) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMendix) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMendix) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMendix) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMendix) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMendix) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMendix) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMendix) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


