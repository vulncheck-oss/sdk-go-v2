# AdvisoryRemediationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to **[]string** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**ProductIds** | Pointer to **[]string** |  | [optional] 
**RestartRequired** | Pointer to [**AdvisoryRestartData**](AdvisoryRestartData.md) |  | [optional] 

## Methods

### NewAdvisoryRemediationData

`func NewAdvisoryRemediationData() *AdvisoryRemediationData`

NewAdvisoryRemediationData instantiates a new AdvisoryRemediationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRemediationDataWithDefaults

`func NewAdvisoryRemediationDataWithDefaults() *AdvisoryRemediationData`

NewAdvisoryRemediationDataWithDefaults instantiates a new AdvisoryRemediationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AdvisoryRemediationData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryRemediationData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryRemediationData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryRemediationData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDate

`func (o *AdvisoryRemediationData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryRemediationData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryRemediationData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryRemediationData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryRemediationData) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryRemediationData) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryRemediationData) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryRemediationData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEntitlements

`func (o *AdvisoryRemediationData) GetEntitlements() []string`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *AdvisoryRemediationData) GetEntitlementsOk() (*[]string, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *AdvisoryRemediationData) SetEntitlements(v []string)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *AdvisoryRemediationData) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetGroupIds

`func (o *AdvisoryRemediationData) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *AdvisoryRemediationData) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *AdvisoryRemediationData) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *AdvisoryRemediationData) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetProductIds

`func (o *AdvisoryRemediationData) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *AdvisoryRemediationData) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *AdvisoryRemediationData) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *AdvisoryRemediationData) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetRestartRequired

`func (o *AdvisoryRemediationData) GetRestartRequired() AdvisoryRestartData`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *AdvisoryRemediationData) GetRestartRequiredOk() (*AdvisoryRestartData, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *AdvisoryRemediationData) SetRestartRequired(v AdvisoryRestartData)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *AdvisoryRemediationData) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


