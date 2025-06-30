# AdvisoryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductID** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryStatus

`func NewAdvisoryStatus() *AdvisoryStatus`

NewAdvisoryStatus instantiates a new AdvisoryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryStatusWithDefaults

`func NewAdvisoryStatusWithDefaults() *AdvisoryStatus`

NewAdvisoryStatusWithDefaults instantiates a new AdvisoryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductID

`func (o *AdvisoryStatus) GetProductID() []string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *AdvisoryStatus) GetProductIDOk() (*[]string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *AdvisoryStatus) SetProductID(v []string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *AdvisoryStatus) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


