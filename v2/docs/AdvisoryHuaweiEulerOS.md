# AdvisoryHuaweiEulerOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Packages** | Pointer to **string** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Synopsis** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHuaweiEulerOS

`func NewAdvisoryHuaweiEulerOS() *AdvisoryHuaweiEulerOS`

NewAdvisoryHuaweiEulerOS instantiates a new AdvisoryHuaweiEulerOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHuaweiEulerOSWithDefaults

`func NewAdvisoryHuaweiEulerOSWithDefaults() *AdvisoryHuaweiEulerOS`

NewAdvisoryHuaweiEulerOSWithDefaults instantiates a new AdvisoryHuaweiEulerOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryHuaweiEulerOS) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHuaweiEulerOS) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHuaweiEulerOS) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHuaweiEulerOS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHuaweiEulerOS) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHuaweiEulerOS) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHuaweiEulerOS) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHuaweiEulerOS) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryHuaweiEulerOS) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryHuaweiEulerOS) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryHuaweiEulerOS) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryHuaweiEulerOS) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryHuaweiEulerOS) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryHuaweiEulerOS) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryHuaweiEulerOS) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryHuaweiEulerOS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryHuaweiEulerOS) GetPackages() string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryHuaweiEulerOS) GetPackagesOk() (*string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryHuaweiEulerOS) SetPackages(v string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryHuaweiEulerOS) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisoryHuaweiEulerOS) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisoryHuaweiEulerOS) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisoryHuaweiEulerOS) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisoryHuaweiEulerOS) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryHuaweiEulerOS) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryHuaweiEulerOS) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryHuaweiEulerOS) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryHuaweiEulerOS) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSynopsis

`func (o *AdvisoryHuaweiEulerOS) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *AdvisoryHuaweiEulerOS) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *AdvisoryHuaweiEulerOS) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *AdvisoryHuaweiEulerOS) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryHuaweiEulerOS) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryHuaweiEulerOS) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryHuaweiEulerOS) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryHuaweiEulerOS) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryHuaweiEulerOS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryHuaweiEulerOS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryHuaweiEulerOS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryHuaweiEulerOS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


