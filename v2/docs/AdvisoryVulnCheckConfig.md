# AdvisoryVulnCheckConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryVulnCheckConfig

`func NewAdvisoryVulnCheckConfig() *AdvisoryVulnCheckConfig`

NewAdvisoryVulnCheckConfig instantiates a new AdvisoryVulnCheckConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnCheckConfigWithDefaults

`func NewAdvisoryVulnCheckConfigWithDefaults() *AdvisoryVulnCheckConfig`

NewAdvisoryVulnCheckConfigWithDefaults instantiates a new AdvisoryVulnCheckConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AdvisoryVulnCheckConfig) GetConfig() []AdvisoryNVD20Configuration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AdvisoryVulnCheckConfig) GetConfigOk() (*[]AdvisoryNVD20Configuration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AdvisoryVulnCheckConfig) SetConfig(v []AdvisoryNVD20Configuration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AdvisoryVulnCheckConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryVulnCheckConfig) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryVulnCheckConfig) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryVulnCheckConfig) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryVulnCheckConfig) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryVulnCheckConfig) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryVulnCheckConfig) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryVulnCheckConfig) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryVulnCheckConfig) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


