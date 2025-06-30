# AdvisoryRustsecAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | Pointer to [**AdvisoryRustsecFrontMatterAdvisory**](AdvisoryRustsecFrontMatterAdvisory.md) |  | [optional] 
**Affected** | Pointer to [**AdvisoryRustsecAffected**](AdvisoryRustsecAffected.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to [**AdvisoryRustsecFrontMatterVersions**](AdvisoryRustsecFrontMatterVersions.md) |  | [optional] 

## Methods

### NewAdvisoryRustsecAdvisory

`func NewAdvisoryRustsecAdvisory() *AdvisoryRustsecAdvisory`

NewAdvisoryRustsecAdvisory instantiates a new AdvisoryRustsecAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRustsecAdvisoryWithDefaults

`func NewAdvisoryRustsecAdvisoryWithDefaults() *AdvisoryRustsecAdvisory`

NewAdvisoryRustsecAdvisoryWithDefaults instantiates a new AdvisoryRustsecAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisory

`func (o *AdvisoryRustsecAdvisory) GetAdvisory() AdvisoryRustsecFrontMatterAdvisory`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *AdvisoryRustsecAdvisory) GetAdvisoryOk() (*AdvisoryRustsecFrontMatterAdvisory, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *AdvisoryRustsecAdvisory) SetAdvisory(v AdvisoryRustsecFrontMatterAdvisory)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *AdvisoryRustsecAdvisory) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryRustsecAdvisory) GetAffected() AdvisoryRustsecAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryRustsecAdvisory) GetAffectedOk() (*AdvisoryRustsecAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryRustsecAdvisory) SetAffected(v AdvisoryRustsecAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryRustsecAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryRustsecAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryRustsecAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryRustsecAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryRustsecAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryRustsecAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryRustsecAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryRustsecAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryRustsecAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryRustsecAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryRustsecAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryRustsecAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryRustsecAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryRustsecAdvisory) GetVersions() AdvisoryRustsecFrontMatterVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryRustsecAdvisory) GetVersionsOk() (*AdvisoryRustsecFrontMatterVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryRustsecAdvisory) SetVersions(v AdvisoryRustsecFrontMatterVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryRustsecAdvisory) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


