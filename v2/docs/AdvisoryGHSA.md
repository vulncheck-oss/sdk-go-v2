# AdvisoryGHSA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Ghsa** | Pointer to [**AdvisoryOriginalGHSA**](AdvisoryOriginalGHSA.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGHSA

`func NewAdvisoryGHSA() *AdvisoryGHSA`

NewAdvisoryGHSA instantiates a new AdvisoryGHSA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGHSAWithDefaults

`func NewAdvisoryGHSAWithDefaults() *AdvisoryGHSA`

NewAdvisoryGHSAWithDefaults instantiates a new AdvisoryGHSA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryGHSA) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGHSA) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGHSA) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGHSA) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGHSA) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGHSA) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGHSA) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGHSA) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetGhsa

`func (o *AdvisoryGHSA) GetGhsa() AdvisoryOriginalGHSA`

GetGhsa returns the Ghsa field if non-nil, zero value otherwise.

### GetGhsaOk

`func (o *AdvisoryGHSA) GetGhsaOk() (*AdvisoryOriginalGHSA, bool)`

GetGhsaOk returns a tuple with the Ghsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhsa

`func (o *AdvisoryGHSA) SetGhsa(v AdvisoryOriginalGHSA)`

SetGhsa sets Ghsa field to given value.

### HasGhsa

`func (o *AdvisoryGHSA) HasGhsa() bool`

HasGhsa returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryGHSA) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryGHSA) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryGHSA) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryGHSA) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


