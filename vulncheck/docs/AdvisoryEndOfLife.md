# AdvisoryEndOfLife

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**Cycles** | Pointer to [**[]AdvisoryCycle**](AdvisoryCycle.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryEndOfLife

`func NewAdvisoryEndOfLife() *AdvisoryEndOfLife`

NewAdvisoryEndOfLife instantiates a new AdvisoryEndOfLife object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryEndOfLifeWithDefaults

`func NewAdvisoryEndOfLifeWithDefaults() *AdvisoryEndOfLife`

NewAdvisoryEndOfLifeWithDefaults instantiates a new AdvisoryEndOfLife object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryEndOfLife) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryEndOfLife) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryEndOfLife) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryEndOfLife) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCycles

`func (o *AdvisoryEndOfLife) GetCycles() []AdvisoryCycle`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *AdvisoryEndOfLife) GetCyclesOk() (*[]AdvisoryCycle, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *AdvisoryEndOfLife) SetCycles(v []AdvisoryCycle)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *AdvisoryEndOfLife) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryEndOfLife) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryEndOfLife) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryEndOfLife) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryEndOfLife) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryEndOfLife) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryEndOfLife) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryEndOfLife) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryEndOfLife) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryEndOfLife) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryEndOfLife) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryEndOfLife) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryEndOfLife) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


