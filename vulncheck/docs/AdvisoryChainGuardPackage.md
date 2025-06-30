# AdvisoryChainGuardPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Secfixes** | Pointer to [**[]AdvisoryChainGuardSecFix**](AdvisoryChainGuardSecFix.md) |  | [optional] 

## Methods

### NewAdvisoryChainGuardPackage

`func NewAdvisoryChainGuardPackage() *AdvisoryChainGuardPackage`

NewAdvisoryChainGuardPackage instantiates a new AdvisoryChainGuardPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryChainGuardPackageWithDefaults

`func NewAdvisoryChainGuardPackageWithDefaults() *AdvisoryChainGuardPackage`

NewAdvisoryChainGuardPackageWithDefaults instantiates a new AdvisoryChainGuardPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvisoryChainGuardPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryChainGuardPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryChainGuardPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryChainGuardPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecfixes

`func (o *AdvisoryChainGuardPackage) GetSecfixes() []AdvisoryChainGuardSecFix`

GetSecfixes returns the Secfixes field if non-nil, zero value otherwise.

### GetSecfixesOk

`func (o *AdvisoryChainGuardPackage) GetSecfixesOk() (*[]AdvisoryChainGuardSecFix, bool)`

GetSecfixesOk returns a tuple with the Secfixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecfixes

`func (o *AdvisoryChainGuardPackage) SetSecfixes(v []AdvisoryChainGuardSecFix)`

SetSecfixes sets Secfixes field to given value.

### HasSecfixes

`func (o *AdvisoryChainGuardPackage) HasSecfixes() bool`

HasSecfixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


