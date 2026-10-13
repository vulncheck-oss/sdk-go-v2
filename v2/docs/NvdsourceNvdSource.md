# NvdsourceNvdSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**CweAcceptanceLevel** | Pointer to [**NvdsourceNvdAcceptanceLevel**](NvdsourceNvdAcceptanceLevel.md) |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourceIdentifiers** | Pointer to **[]string** |  | [optional] 
**V2AcceptanceLevel** | Pointer to [**NvdsourceNvdAcceptanceLevel**](NvdsourceNvdAcceptanceLevel.md) |  | [optional] 
**V3AcceptanceLevel** | Pointer to [**NvdsourceNvdAcceptanceLevel**](NvdsourceNvdAcceptanceLevel.md) |  | [optional] 

## Methods

### NewNvdsourceNvdSource

`func NewNvdsourceNvdSource() *NvdsourceNvdSource`

NewNvdsourceNvdSource instantiates a new NvdsourceNvdSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNvdsourceNvdSourceWithDefaults

`func NewNvdsourceNvdSourceWithDefaults() *NvdsourceNvdSource`

NewNvdsourceNvdSourceWithDefaults instantiates a new NvdsourceNvdSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *NvdsourceNvdSource) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *NvdsourceNvdSource) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *NvdsourceNvdSource) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *NvdsourceNvdSource) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetCreated

`func (o *NvdsourceNvdSource) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NvdsourceNvdSource) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NvdsourceNvdSource) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NvdsourceNvdSource) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCweAcceptanceLevel

`func (o *NvdsourceNvdSource) GetCweAcceptanceLevel() NvdsourceNvdAcceptanceLevel`

GetCweAcceptanceLevel returns the CweAcceptanceLevel field if non-nil, zero value otherwise.

### GetCweAcceptanceLevelOk

`func (o *NvdsourceNvdSource) GetCweAcceptanceLevelOk() (*NvdsourceNvdAcceptanceLevel, bool)`

GetCweAcceptanceLevelOk returns a tuple with the CweAcceptanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCweAcceptanceLevel

`func (o *NvdsourceNvdSource) SetCweAcceptanceLevel(v NvdsourceNvdAcceptanceLevel)`

SetCweAcceptanceLevel sets CweAcceptanceLevel field to given value.

### HasCweAcceptanceLevel

`func (o *NvdsourceNvdSource) HasCweAcceptanceLevel() bool`

HasCweAcceptanceLevel returns a boolean if a field has been set.

### GetLastModified

`func (o *NvdsourceNvdSource) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NvdsourceNvdSource) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NvdsourceNvdSource) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *NvdsourceNvdSource) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *NvdsourceNvdSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NvdsourceNvdSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NvdsourceNvdSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NvdsourceNvdSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceIdentifiers

`func (o *NvdsourceNvdSource) GetSourceIdentifiers() []string`

GetSourceIdentifiers returns the SourceIdentifiers field if non-nil, zero value otherwise.

### GetSourceIdentifiersOk

`func (o *NvdsourceNvdSource) GetSourceIdentifiersOk() (*[]string, bool)`

GetSourceIdentifiersOk returns a tuple with the SourceIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifiers

`func (o *NvdsourceNvdSource) SetSourceIdentifiers(v []string)`

SetSourceIdentifiers sets SourceIdentifiers field to given value.

### HasSourceIdentifiers

`func (o *NvdsourceNvdSource) HasSourceIdentifiers() bool`

HasSourceIdentifiers returns a boolean if a field has been set.

### GetV2AcceptanceLevel

`func (o *NvdsourceNvdSource) GetV2AcceptanceLevel() NvdsourceNvdAcceptanceLevel`

GetV2AcceptanceLevel returns the V2AcceptanceLevel field if non-nil, zero value otherwise.

### GetV2AcceptanceLevelOk

`func (o *NvdsourceNvdSource) GetV2AcceptanceLevelOk() (*NvdsourceNvdAcceptanceLevel, bool)`

GetV2AcceptanceLevelOk returns a tuple with the V2AcceptanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2AcceptanceLevel

`func (o *NvdsourceNvdSource) SetV2AcceptanceLevel(v NvdsourceNvdAcceptanceLevel)`

SetV2AcceptanceLevel sets V2AcceptanceLevel field to given value.

### HasV2AcceptanceLevel

`func (o *NvdsourceNvdSource) HasV2AcceptanceLevel() bool`

HasV2AcceptanceLevel returns a boolean if a field has been set.

### GetV3AcceptanceLevel

`func (o *NvdsourceNvdSource) GetV3AcceptanceLevel() NvdsourceNvdAcceptanceLevel`

GetV3AcceptanceLevel returns the V3AcceptanceLevel field if non-nil, zero value otherwise.

### GetV3AcceptanceLevelOk

`func (o *NvdsourceNvdSource) GetV3AcceptanceLevelOk() (*NvdsourceNvdAcceptanceLevel, bool)`

GetV3AcceptanceLevelOk returns a tuple with the V3AcceptanceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3AcceptanceLevel

`func (o *NvdsourceNvdSource) SetV3AcceptanceLevel(v NvdsourceNvdAcceptanceLevel)`

SetV3AcceptanceLevel sets V3AcceptanceLevel field to given value.

### HasV3AcceptanceLevel

`func (o *NvdsourceNvdSource) HasV3AcceptanceLevel() bool`

HasV3AcceptanceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


