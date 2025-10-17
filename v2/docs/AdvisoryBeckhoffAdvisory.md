# AdvisoryBeckhoffAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeckhoffId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cwe** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateLastRevised** | Pointer to **string** | if in the future we can delete this great - it&#39;s just a dupe to normalize the field names | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vde** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryBeckhoffAdvisory

`func NewAdvisoryBeckhoffAdvisory() *AdvisoryBeckhoffAdvisory`

NewAdvisoryBeckhoffAdvisory instantiates a new AdvisoryBeckhoffAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryBeckhoffAdvisoryWithDefaults

`func NewAdvisoryBeckhoffAdvisoryWithDefaults() *AdvisoryBeckhoffAdvisory`

NewAdvisoryBeckhoffAdvisoryWithDefaults instantiates a new AdvisoryBeckhoffAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeckhoffId

`func (o *AdvisoryBeckhoffAdvisory) GetBeckhoffId() string`

GetBeckhoffId returns the BeckhoffId field if non-nil, zero value otherwise.

### GetBeckhoffIdOk

`func (o *AdvisoryBeckhoffAdvisory) GetBeckhoffIdOk() (*string, bool)`

GetBeckhoffIdOk returns a tuple with the BeckhoffId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeckhoffId

`func (o *AdvisoryBeckhoffAdvisory) SetBeckhoffId(v string)`

SetBeckhoffId sets BeckhoffId field to given value.

### HasBeckhoffId

`func (o *AdvisoryBeckhoffAdvisory) HasBeckhoffId() bool`

HasBeckhoffId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryBeckhoffAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryBeckhoffAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryBeckhoffAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryBeckhoffAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryBeckhoffAdvisory) GetCwe() []string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryBeckhoffAdvisory) GetCweOk() (*[]string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryBeckhoffAdvisory) SetCwe(v []string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryBeckhoffAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryBeckhoffAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryBeckhoffAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryBeckhoffAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryBeckhoffAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateLastRevised

`func (o *AdvisoryBeckhoffAdvisory) GetDateLastRevised() string`

GetDateLastRevised returns the DateLastRevised field if non-nil, zero value otherwise.

### GetDateLastRevisedOk

`func (o *AdvisoryBeckhoffAdvisory) GetDateLastRevisedOk() (*string, bool)`

GetDateLastRevisedOk returns a tuple with the DateLastRevised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastRevised

`func (o *AdvisoryBeckhoffAdvisory) SetDateLastRevised(v string)`

SetDateLastRevised sets DateLastRevised field to given value.

### HasDateLastRevised

`func (o *AdvisoryBeckhoffAdvisory) HasDateLastRevised() bool`

HasDateLastRevised returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryBeckhoffAdvisory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryBeckhoffAdvisory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryBeckhoffAdvisory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryBeckhoffAdvisory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryBeckhoffAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryBeckhoffAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryBeckhoffAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryBeckhoffAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryBeckhoffAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryBeckhoffAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryBeckhoffAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryBeckhoffAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVde

`func (o *AdvisoryBeckhoffAdvisory) GetVde() []string`

GetVde returns the Vde field if non-nil, zero value otherwise.

### GetVdeOk

`func (o *AdvisoryBeckhoffAdvisory) GetVdeOk() (*[]string, bool)`

GetVdeOk returns a tuple with the Vde field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVde

`func (o *AdvisoryBeckhoffAdvisory) SetVde(v []string)`

SetVde sets Vde field to given value.

### HasVde

`func (o *AdvisoryBeckhoffAdvisory) HasVde() bool`

HasVde returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


