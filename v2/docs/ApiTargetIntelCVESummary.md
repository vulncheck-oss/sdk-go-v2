# ApiTargetIntelCVESummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Asns** | Pointer to [**[]ApiTargetIntelCVESummaryAsnsInner**](ApiTargetIntelCVESummaryAsnsInner.md) |  | [optional] 
**CountryCodes** | Pointer to [**[]ApiTargetIntelCVESummaryAsnsInner**](ApiTargetIntelCVESummaryAsnsInner.md) |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]ApiTargetIntelCVESummaryPortsInner**](ApiTargetIntelCVESummaryPortsInner.md) |  | [optional] 
**Vendors** | Pointer to [**[]ApiTargetIntelCVESummaryAsnsInner**](ApiTargetIntelCVESummaryAsnsInner.md) |  | [optional] 

## Methods

### NewApiTargetIntelCVESummary

`func NewApiTargetIntelCVESummary() *ApiTargetIntelCVESummary`

NewApiTargetIntelCVESummary instantiates a new ApiTargetIntelCVESummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTargetIntelCVESummaryWithDefaults

`func NewApiTargetIntelCVESummaryWithDefaults() *ApiTargetIntelCVESummary`

NewApiTargetIntelCVESummaryWithDefaults instantiates a new ApiTargetIntelCVESummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApiTargetIntelCVESummary) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiTargetIntelCVESummary) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiTargetIntelCVESummary) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiTargetIntelCVESummary) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAsns

`func (o *ApiTargetIntelCVESummary) GetAsns() []ApiTargetIntelCVESummaryAsnsInner`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *ApiTargetIntelCVESummary) GetAsnsOk() (*[]ApiTargetIntelCVESummaryAsnsInner, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *ApiTargetIntelCVESummary) SetAsns(v []ApiTargetIntelCVESummaryAsnsInner)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *ApiTargetIntelCVESummary) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetCountryCodes

`func (o *ApiTargetIntelCVESummary) GetCountryCodes() []ApiTargetIntelCVESummaryAsnsInner`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ApiTargetIntelCVESummary) GetCountryCodesOk() (*[]ApiTargetIntelCVESummaryAsnsInner, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ApiTargetIntelCVESummary) SetCountryCodes(v []ApiTargetIntelCVESummaryAsnsInner)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ApiTargetIntelCVESummary) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetCve

`func (o *ApiTargetIntelCVESummary) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiTargetIntelCVESummary) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiTargetIntelCVESummary) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiTargetIntelCVESummary) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetPorts

`func (o *ApiTargetIntelCVESummary) GetPorts() []ApiTargetIntelCVESummaryPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApiTargetIntelCVESummary) GetPortsOk() (*[]ApiTargetIntelCVESummaryPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApiTargetIntelCVESummary) SetPorts(v []ApiTargetIntelCVESummaryPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApiTargetIntelCVESummary) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetVendors

`func (o *ApiTargetIntelCVESummary) GetVendors() []ApiTargetIntelCVESummaryAsnsInner`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *ApiTargetIntelCVESummary) GetVendorsOk() (*[]ApiTargetIntelCVESummaryAsnsInner, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *ApiTargetIntelCVESummary) SetVendors(v []ApiTargetIntelCVESummaryAsnsInner)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *ApiTargetIntelCVESummary) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


