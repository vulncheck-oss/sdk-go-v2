# AdvisoryRThreat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**DateSpecified** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to [**AdvisoryIVal**](AdvisoryIVal.md) |  | [optional] 
**ProductID** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **int32** | diff | [optional] 

## Methods

### NewAdvisoryRThreat

`func NewAdvisoryRThreat() *AdvisoryRThreat`

NewAdvisoryRThreat instantiates a new AdvisoryRThreat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRThreatWithDefaults

`func NewAdvisoryRThreatWithDefaults() *AdvisoryRThreat`

NewAdvisoryRThreatWithDefaults instantiates a new AdvisoryRThreat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AdvisoryRThreat) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryRThreat) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryRThreat) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryRThreat) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateSpecified

`func (o *AdvisoryRThreat) GetDateSpecified() bool`

GetDateSpecified returns the DateSpecified field if non-nil, zero value otherwise.

### GetDateSpecifiedOk

`func (o *AdvisoryRThreat) GetDateSpecifiedOk() (*bool, bool)`

GetDateSpecifiedOk returns a tuple with the DateSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSpecified

`func (o *AdvisoryRThreat) SetDateSpecified(v bool)`

SetDateSpecified sets DateSpecified field to given value.

### HasDateSpecified

`func (o *AdvisoryRThreat) HasDateSpecified() bool`

HasDateSpecified returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryRThreat) GetDescription() AdvisoryIVal`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryRThreat) GetDescriptionOk() (*AdvisoryIVal, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryRThreat) SetDescription(v AdvisoryIVal)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryRThreat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductID

`func (o *AdvisoryRThreat) GetProductID() []string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *AdvisoryRThreat) GetProductIDOk() (*[]string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *AdvisoryRThreat) SetProductID(v []string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *AdvisoryRThreat) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryRThreat) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryRThreat) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryRThreat) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryRThreat) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


