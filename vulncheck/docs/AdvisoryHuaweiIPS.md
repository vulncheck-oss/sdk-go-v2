# AdvisoryHuaweiIPS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cnnvd** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**ThreatId** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHuaweiIPS

`func NewAdvisoryHuaweiIPS() *AdvisoryHuaweiIPS`

NewAdvisoryHuaweiIPS instantiates a new AdvisoryHuaweiIPS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHuaweiIPSWithDefaults

`func NewAdvisoryHuaweiIPSWithDefaults() *AdvisoryHuaweiIPS`

NewAdvisoryHuaweiIPSWithDefaults instantiates a new AdvisoryHuaweiIPS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnnvd

`func (o *AdvisoryHuaweiIPS) GetCnnvd() string`

GetCnnvd returns the Cnnvd field if non-nil, zero value otherwise.

### GetCnnvdOk

`func (o *AdvisoryHuaweiIPS) GetCnnvdOk() (*string, bool)`

GetCnnvdOk returns a tuple with the Cnnvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnnvd

`func (o *AdvisoryHuaweiIPS) SetCnnvd(v string)`

SetCnnvd sets Cnnvd field to given value.

### HasCnnvd

`func (o *AdvisoryHuaweiIPS) HasCnnvd() bool`

HasCnnvd returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHuaweiIPS) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHuaweiIPS) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHuaweiIPS) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHuaweiIPS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHuaweiIPS) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHuaweiIPS) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHuaweiIPS) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHuaweiIPS) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateUpdated

`func (o *AdvisoryHuaweiIPS) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AdvisoryHuaweiIPS) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AdvisoryHuaweiIPS) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AdvisoryHuaweiIPS) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryHuaweiIPS) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryHuaweiIPS) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryHuaweiIPS) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryHuaweiIPS) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryHuaweiIPS) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryHuaweiIPS) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryHuaweiIPS) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryHuaweiIPS) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetThreatId

`func (o *AdvisoryHuaweiIPS) GetThreatId() int32`

GetThreatId returns the ThreatId field if non-nil, zero value otherwise.

### GetThreatIdOk

`func (o *AdvisoryHuaweiIPS) GetThreatIdOk() (*int32, bool)`

GetThreatIdOk returns a tuple with the ThreatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatId

`func (o *AdvisoryHuaweiIPS) SetThreatId(v int32)`

SetThreatId sets ThreatId field to given value.

### HasThreatId

`func (o *AdvisoryHuaweiIPS) HasThreatId() bool`

HasThreatId returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryHuaweiIPS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryHuaweiIPS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryHuaweiIPS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryHuaweiIPS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryHuaweiIPS) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryHuaweiIPS) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryHuaweiIPS) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryHuaweiIPS) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


