# AdvisoryCompassSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsncId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Effect** | Pointer to **string** |  | [optional] 
**Introduction** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCompassSecurity

`func NewAdvisoryCompassSecurity() *AdvisoryCompassSecurity`

NewAdvisoryCompassSecurity instantiates a new AdvisoryCompassSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCompassSecurityWithDefaults

`func NewAdvisoryCompassSecurityWithDefaults() *AdvisoryCompassSecurity`

NewAdvisoryCompassSecurityWithDefaults instantiates a new AdvisoryCompassSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsncId

`func (o *AdvisoryCompassSecurity) GetCsncId() string`

GetCsncId returns the CsncId field if non-nil, zero value otherwise.

### GetCsncIdOk

`func (o *AdvisoryCompassSecurity) GetCsncIdOk() (*string, bool)`

GetCsncIdOk returns a tuple with the CsncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsncId

`func (o *AdvisoryCompassSecurity) SetCsncId(v string)`

SetCsncId sets CsncId field to given value.

### HasCsncId

`func (o *AdvisoryCompassSecurity) HasCsncId() bool`

HasCsncId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCompassSecurity) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCompassSecurity) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCompassSecurity) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCompassSecurity) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCompassSecurity) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCompassSecurity) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCompassSecurity) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCompassSecurity) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetEffect

`func (o *AdvisoryCompassSecurity) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AdvisoryCompassSecurity) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AdvisoryCompassSecurity) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *AdvisoryCompassSecurity) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetIntroduction

`func (o *AdvisoryCompassSecurity) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *AdvisoryCompassSecurity) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *AdvisoryCompassSecurity) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *AdvisoryCompassSecurity) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryCompassSecurity) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryCompassSecurity) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryCompassSecurity) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryCompassSecurity) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCompassSecurity) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCompassSecurity) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCompassSecurity) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCompassSecurity) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRisk

`func (o *AdvisoryCompassSecurity) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AdvisoryCompassSecurity) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AdvisoryCompassSecurity) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AdvisoryCompassSecurity) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryCompassSecurity) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryCompassSecurity) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryCompassSecurity) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryCompassSecurity) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCompassSecurity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCompassSecurity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCompassSecurity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCompassSecurity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCompassSecurity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCompassSecurity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCompassSecurity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCompassSecurity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryCompassSecurity) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryCompassSecurity) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryCompassSecurity) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryCompassSecurity) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


