# ApiOSSPackageHashInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | See OSSPackageHashInfoAlgo* consts | [optional] 
**Type** | Pointer to **string** | See OSSPackageHashInfoType* consts | [optional] 
**Value** | Pointer to **string** | hex string digest or link to hash | [optional] 

## Methods

### NewApiOSSPackageHashInfo

`func NewApiOSSPackageHashInfo() *ApiOSSPackageHashInfo`

NewApiOSSPackageHashInfo instantiates a new ApiOSSPackageHashInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOSSPackageHashInfoWithDefaults

`func NewApiOSSPackageHashInfoWithDefaults() *ApiOSSPackageHashInfo`

NewApiOSSPackageHashInfoWithDefaults instantiates a new ApiOSSPackageHashInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *ApiOSSPackageHashInfo) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ApiOSSPackageHashInfo) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ApiOSSPackageHashInfo) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ApiOSSPackageHashInfo) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetType

`func (o *ApiOSSPackageHashInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiOSSPackageHashInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiOSSPackageHashInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiOSSPackageHashInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ApiOSSPackageHashInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiOSSPackageHashInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiOSSPackageHashInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiOSSPackageHashInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


