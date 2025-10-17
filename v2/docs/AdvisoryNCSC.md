# AdvisoryNCSC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csaf** | Pointer to [**AdvisoryCSAF**](AdvisoryCSAF.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**SummaryNl** | Pointer to **string** |  | [optional] 
**TitleNl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryNCSC

`func NewAdvisoryNCSC() *AdvisoryNCSC`

NewAdvisoryNCSC instantiates a new AdvisoryNCSC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNCSCWithDefaults

`func NewAdvisoryNCSCWithDefaults() *AdvisoryNCSC`

NewAdvisoryNCSCWithDefaults instantiates a new AdvisoryNCSC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsaf

`func (o *AdvisoryNCSC) GetCsaf() AdvisoryCSAF`

GetCsaf returns the Csaf field if non-nil, zero value otherwise.

### GetCsafOk

`func (o *AdvisoryNCSC) GetCsafOk() (*AdvisoryCSAF, bool)`

GetCsafOk returns a tuple with the Csaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsaf

`func (o *AdvisoryNCSC) SetCsaf(v AdvisoryCSAF)`

SetCsaf sets Csaf field to given value.

### HasCsaf

`func (o *AdvisoryNCSC) HasCsaf() bool`

HasCsaf returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryNCSC) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryNCSC) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryNCSC) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryNCSC) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryNCSC) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryNCSC) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryNCSC) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryNCSC) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryNCSC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryNCSC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryNCSC) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryNCSC) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryNCSC) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryNCSC) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryNCSC) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryNCSC) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummaryNl

`func (o *AdvisoryNCSC) GetSummaryNl() string`

GetSummaryNl returns the SummaryNl field if non-nil, zero value otherwise.

### GetSummaryNlOk

`func (o *AdvisoryNCSC) GetSummaryNlOk() (*string, bool)`

GetSummaryNlOk returns a tuple with the SummaryNl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryNl

`func (o *AdvisoryNCSC) SetSummaryNl(v string)`

SetSummaryNl sets SummaryNl field to given value.

### HasSummaryNl

`func (o *AdvisoryNCSC) HasSummaryNl() bool`

HasSummaryNl returns a boolean if a field has been set.

### GetTitleNl

`func (o *AdvisoryNCSC) GetTitleNl() string`

GetTitleNl returns the TitleNl field if non-nil, zero value otherwise.

### GetTitleNlOk

`func (o *AdvisoryNCSC) GetTitleNlOk() (*string, bool)`

GetTitleNlOk returns a tuple with the TitleNl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleNl

`func (o *AdvisoryNCSC) SetTitleNl(v string)`

SetTitleNl sets TitleNl field to given value.

### HasTitleNl

`func (o *AdvisoryNCSC) HasTitleNl() bool`

HasTitleNl returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryNCSC) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryNCSC) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryNCSC) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryNCSC) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


