# ApiCveItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configurations** | Pointer to [**ApiConfigurations**](ApiConfigurations.md) |  | [optional] 
**Cve** | Pointer to [**ApiCVE**](ApiCVE.md) |  | [optional] 
**Impact** | Pointer to [**ApiImpact**](ApiImpact.md) |  | [optional] 
**LastModifiedDate** | Pointer to **string** |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**VcConfigurations** | Pointer to [**ApiConfigurations**](ApiConfigurations.md) |  | [optional] 
**VcVulnerableCPEs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiCveItems

`func NewApiCveItems() *ApiCveItems`

NewApiCveItems instantiates a new ApiCveItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCveItemsWithDefaults

`func NewApiCveItemsWithDefaults() *ApiCveItems`

NewApiCveItemsWithDefaults instantiates a new ApiCveItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurations

`func (o *ApiCveItems) GetConfigurations() ApiConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ApiCveItems) GetConfigurationsOk() (*ApiConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ApiCveItems) SetConfigurations(v ApiConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ApiCveItems) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCve

`func (o *ApiCveItems) GetCve() ApiCVE`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiCveItems) GetCveOk() (*ApiCVE, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiCveItems) SetCve(v ApiCVE)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiCveItems) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetImpact

`func (o *ApiCveItems) GetImpact() ApiImpact`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *ApiCveItems) GetImpactOk() (*ApiImpact, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *ApiCveItems) SetImpact(v ApiImpact)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *ApiCveItems) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ApiCveItems) GetLastModifiedDate() string`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ApiCveItems) GetLastModifiedDateOk() (*string, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ApiCveItems) SetLastModifiedDate(v string)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ApiCveItems) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetPublishedDate

`func (o *ApiCveItems) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ApiCveItems) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ApiCveItems) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *ApiCveItems) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetVcConfigurations

`func (o *ApiCveItems) GetVcConfigurations() ApiConfigurations`

GetVcConfigurations returns the VcConfigurations field if non-nil, zero value otherwise.

### GetVcConfigurationsOk

`func (o *ApiCveItems) GetVcConfigurationsOk() (*ApiConfigurations, bool)`

GetVcConfigurationsOk returns a tuple with the VcConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurations

`func (o *ApiCveItems) SetVcConfigurations(v ApiConfigurations)`

SetVcConfigurations sets VcConfigurations field to given value.

### HasVcConfigurations

`func (o *ApiCveItems) HasVcConfigurations() bool`

HasVcConfigurations returns a boolean if a field has been set.

### GetVcVulnerableCPEs

`func (o *ApiCveItems) GetVcVulnerableCPEs() []string`

GetVcVulnerableCPEs returns the VcVulnerableCPEs field if non-nil, zero value otherwise.

### GetVcVulnerableCPEsOk

`func (o *ApiCveItems) GetVcVulnerableCPEsOk() (*[]string, bool)`

GetVcVulnerableCPEsOk returns a tuple with the VcVulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcVulnerableCPEs

`func (o *ApiCveItems) SetVcVulnerableCPEs(v []string)`

SetVcVulnerableCPEs sets VcVulnerableCPEs field to given value.

### HasVcVulnerableCPEs

`func (o *ApiCveItems) HasVcVulnerableCPEs() bool`

HasVcVulnerableCPEs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


