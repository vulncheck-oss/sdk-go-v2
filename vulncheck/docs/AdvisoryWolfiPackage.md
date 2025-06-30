# AdvisoryWolfiPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Secfixes** | Pointer to [**[]AdvisoryWolfiSecFix**](AdvisoryWolfiSecFix.md) |  | [optional] 

## Methods

### NewAdvisoryWolfiPackage

`func NewAdvisoryWolfiPackage() *AdvisoryWolfiPackage`

NewAdvisoryWolfiPackage instantiates a new AdvisoryWolfiPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryWolfiPackageWithDefaults

`func NewAdvisoryWolfiPackageWithDefaults() *AdvisoryWolfiPackage`

NewAdvisoryWolfiPackageWithDefaults instantiates a new AdvisoryWolfiPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdvisoryWolfiPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryWolfiPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryWolfiPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryWolfiPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecfixes

`func (o *AdvisoryWolfiPackage) GetSecfixes() []AdvisoryWolfiSecFix`

GetSecfixes returns the Secfixes field if non-nil, zero value otherwise.

### GetSecfixesOk

`func (o *AdvisoryWolfiPackage) GetSecfixesOk() (*[]AdvisoryWolfiSecFix, bool)`

GetSecfixesOk returns a tuple with the Secfixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecfixes

`func (o *AdvisoryWolfiPackage) SetSecfixes(v []AdvisoryWolfiSecFix)`

SetSecfixes sets Secfixes field to given value.

### HasSecfixes

`func (o *AdvisoryWolfiPackage) HasSecfixes() bool`

HasSecfixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


