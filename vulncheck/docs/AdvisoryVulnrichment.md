# AdvisoryVulnrichment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MitreRef** | Pointer to [**AdvisoryVulnrichmentCVERef**](AdvisoryVulnrichmentCVERef.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryVulnrichment

`func NewAdvisoryVulnrichment() *AdvisoryVulnrichment`

NewAdvisoryVulnrichment instantiates a new AdvisoryVulnrichment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnrichmentWithDefaults

`func NewAdvisoryVulnrichmentWithDefaults() *AdvisoryVulnrichment`

NewAdvisoryVulnrichmentWithDefaults instantiates a new AdvisoryVulnrichment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryVulnrichment) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryVulnrichment) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryVulnrichment) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryVulnrichment) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryVulnrichment) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryVulnrichment) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryVulnrichment) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryVulnrichment) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMitreRef

`func (o *AdvisoryVulnrichment) GetMitreRef() AdvisoryVulnrichmentCVERef`

GetMitreRef returns the MitreRef field if non-nil, zero value otherwise.

### GetMitreRefOk

`func (o *AdvisoryVulnrichment) GetMitreRefOk() (*AdvisoryVulnrichmentCVERef, bool)`

GetMitreRefOk returns a tuple with the MitreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreRef

`func (o *AdvisoryVulnrichment) SetMitreRef(v AdvisoryVulnrichmentCVERef)`

SetMitreRef sets MitreRef field to given value.

### HasMitreRef

`func (o *AdvisoryVulnrichment) HasMitreRef() bool`

HasMitreRef returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryVulnrichment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryVulnrichment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryVulnrichment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryVulnrichment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


