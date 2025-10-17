# AdvisoryMicrosoftKb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Kbs** | Pointer to [**[]AdvisoryKb**](AdvisoryKb.md) |  | [optional] 
**Threat** | Pointer to [**AdvisoryKbThreatDescription**](AdvisoryKbThreatDescription.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMicrosoftKb

`func NewAdvisoryMicrosoftKb() *AdvisoryMicrosoftKb`

NewAdvisoryMicrosoftKb instantiates a new AdvisoryMicrosoftKb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMicrosoftKbWithDefaults

`func NewAdvisoryMicrosoftKbWithDefaults() *AdvisoryMicrosoftKb`

NewAdvisoryMicrosoftKbWithDefaults instantiates a new AdvisoryMicrosoftKb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryMicrosoftKb) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMicrosoftKb) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMicrosoftKb) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMicrosoftKb) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMicrosoftKb) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMicrosoftKb) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMicrosoftKb) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMicrosoftKb) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetKbs

`func (o *AdvisoryMicrosoftKb) GetKbs() []AdvisoryKb`

GetKbs returns the Kbs field if non-nil, zero value otherwise.

### GetKbsOk

`func (o *AdvisoryMicrosoftKb) GetKbsOk() (*[]AdvisoryKb, bool)`

GetKbsOk returns a tuple with the Kbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKbs

`func (o *AdvisoryMicrosoftKb) SetKbs(v []AdvisoryKb)`

SetKbs sets Kbs field to given value.

### HasKbs

`func (o *AdvisoryMicrosoftKb) HasKbs() bool`

HasKbs returns a boolean if a field has been set.

### GetThreat

`func (o *AdvisoryMicrosoftKb) GetThreat() AdvisoryKbThreatDescription`

GetThreat returns the Threat field if non-nil, zero value otherwise.

### GetThreatOk

`func (o *AdvisoryMicrosoftKb) GetThreatOk() (*AdvisoryKbThreatDescription, bool)`

GetThreatOk returns a tuple with the Threat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreat

`func (o *AdvisoryMicrosoftKb) SetThreat(v AdvisoryKbThreatDescription)`

SetThreat sets Threat field to given value.

### HasThreat

`func (o *AdvisoryMicrosoftKb) HasThreat() bool`

HasThreat returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMicrosoftKb) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMicrosoftKb) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMicrosoftKb) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMicrosoftKb) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


