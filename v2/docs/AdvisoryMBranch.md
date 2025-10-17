# AdvisoryMBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**[]AdvisoryMBranch**](AdvisoryMBranch.md) |  | [optional] 
**FullProductName** | Pointer to [**[]AdvisoryMFullProductName**](AdvisoryMFullProductName.md) |  | [optional] 
**Items** | Pointer to [**[]AdvisoryMItem**](AdvisoryMItem.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | diff | [optional] 

## Methods

### NewAdvisoryMBranch

`func NewAdvisoryMBranch() *AdvisoryMBranch`

NewAdvisoryMBranch instantiates a new AdvisoryMBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMBranchWithDefaults

`func NewAdvisoryMBranchWithDefaults() *AdvisoryMBranch`

NewAdvisoryMBranchWithDefaults instantiates a new AdvisoryMBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *AdvisoryMBranch) GetBranch() []AdvisoryMBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AdvisoryMBranch) GetBranchOk() (*[]AdvisoryMBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AdvisoryMBranch) SetBranch(v []AdvisoryMBranch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AdvisoryMBranch) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetFullProductName

`func (o *AdvisoryMBranch) GetFullProductName() []AdvisoryMFullProductName`

GetFullProductName returns the FullProductName field if non-nil, zero value otherwise.

### GetFullProductNameOk

`func (o *AdvisoryMBranch) GetFullProductNameOk() (*[]AdvisoryMFullProductName, bool)`

GetFullProductNameOk returns a tuple with the FullProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProductName

`func (o *AdvisoryMBranch) SetFullProductName(v []AdvisoryMFullProductName)`

SetFullProductName sets FullProductName field to given value.

### HasFullProductName

`func (o *AdvisoryMBranch) HasFullProductName() bool`

HasFullProductName returns a boolean if a field has been set.

### GetItems

`func (o *AdvisoryMBranch) GetItems() []AdvisoryMItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdvisoryMBranch) GetItemsOk() (*[]AdvisoryMItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdvisoryMBranch) SetItems(v []AdvisoryMItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *AdvisoryMBranch) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryMBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryMBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryMBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryMBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryMBranch) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryMBranch) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryMBranch) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryMBranch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


