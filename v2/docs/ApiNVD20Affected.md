# ApiNVD20Affected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedData** | Pointer to [**[]ApiNVD20AffectedProduct**](ApiNVD20AffectedProduct.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20Affected

`func NewApiNVD20Affected() *ApiNVD20Affected`

NewApiNVD20Affected instantiates a new ApiNVD20Affected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20AffectedWithDefaults

`func NewApiNVD20AffectedWithDefaults() *ApiNVD20Affected`

NewApiNVD20AffectedWithDefaults instantiates a new ApiNVD20Affected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedData

`func (o *ApiNVD20Affected) GetAffectedData() []ApiNVD20AffectedProduct`

GetAffectedData returns the AffectedData field if non-nil, zero value otherwise.

### GetAffectedDataOk

`func (o *ApiNVD20Affected) GetAffectedDataOk() (*[]ApiNVD20AffectedProduct, bool)`

GetAffectedDataOk returns a tuple with the AffectedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedData

`func (o *ApiNVD20Affected) SetAffectedData(v []ApiNVD20AffectedProduct)`

SetAffectedData sets AffectedData field to given value.

### HasAffectedData

`func (o *ApiNVD20Affected) HasAffectedData() bool`

HasAffectedData returns a boolean if a field has been set.

### GetSource

`func (o *ApiNVD20Affected) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiNVD20Affected) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiNVD20Affected) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiNVD20Affected) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


