# AdvisoryCSAFRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**FullProductName** | Pointer to [**AdvisoryProduct**](AdvisoryProduct.md) |  | [optional] 
**ProductReference** | Pointer to **string** |  | [optional] 
**RelatesToProductReference** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCSAFRelationship

`func NewAdvisoryCSAFRelationship() *AdvisoryCSAFRelationship`

NewAdvisoryCSAFRelationship instantiates a new AdvisoryCSAFRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCSAFRelationshipWithDefaults

`func NewAdvisoryCSAFRelationshipWithDefaults() *AdvisoryCSAFRelationship`

NewAdvisoryCSAFRelationshipWithDefaults instantiates a new AdvisoryCSAFRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AdvisoryCSAFRelationship) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryCSAFRelationship) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryCSAFRelationship) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryCSAFRelationship) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFullProductName

`func (o *AdvisoryCSAFRelationship) GetFullProductName() AdvisoryProduct`

GetFullProductName returns the FullProductName field if non-nil, zero value otherwise.

### GetFullProductNameOk

`func (o *AdvisoryCSAFRelationship) GetFullProductNameOk() (*AdvisoryProduct, bool)`

GetFullProductNameOk returns a tuple with the FullProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullProductName

`func (o *AdvisoryCSAFRelationship) SetFullProductName(v AdvisoryProduct)`

SetFullProductName sets FullProductName field to given value.

### HasFullProductName

`func (o *AdvisoryCSAFRelationship) HasFullProductName() bool`

HasFullProductName returns a boolean if a field has been set.

### GetProductReference

`func (o *AdvisoryCSAFRelationship) GetProductReference() string`

GetProductReference returns the ProductReference field if non-nil, zero value otherwise.

### GetProductReferenceOk

`func (o *AdvisoryCSAFRelationship) GetProductReferenceOk() (*string, bool)`

GetProductReferenceOk returns a tuple with the ProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductReference

`func (o *AdvisoryCSAFRelationship) SetProductReference(v string)`

SetProductReference sets ProductReference field to given value.

### HasProductReference

`func (o *AdvisoryCSAFRelationship) HasProductReference() bool`

HasProductReference returns a boolean if a field has been set.

### GetRelatesToProductReference

`func (o *AdvisoryCSAFRelationship) GetRelatesToProductReference() string`

GetRelatesToProductReference returns the RelatesToProductReference field if non-nil, zero value otherwise.

### GetRelatesToProductReferenceOk

`func (o *AdvisoryCSAFRelationship) GetRelatesToProductReferenceOk() (*string, bool)`

GetRelatesToProductReferenceOk returns a tuple with the RelatesToProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatesToProductReference

`func (o *AdvisoryCSAFRelationship) SetRelatesToProductReference(v string)`

SetRelatesToProductReference sets RelatesToProductReference field to given value.

### HasRelatesToProductReference

`func (o *AdvisoryCSAFRelationship) HasRelatesToProductReference() bool`

HasRelatesToProductReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


