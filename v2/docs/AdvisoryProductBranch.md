# AdvisoryProductBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]AdvisoryProductBranch**](AdvisoryProductBranch.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Product** | Pointer to [**AdvisoryProduct**](AdvisoryProduct.md) |  | [optional] 
**Relationships** | Pointer to [**[]AdvisoryCSAFRelationship**](AdvisoryCSAFRelationship.md) |  | [optional] 

## Methods

### NewAdvisoryProductBranch

`func NewAdvisoryProductBranch() *AdvisoryProductBranch`

NewAdvisoryProductBranch instantiates a new AdvisoryProductBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryProductBranchWithDefaults

`func NewAdvisoryProductBranchWithDefaults() *AdvisoryProductBranch`

NewAdvisoryProductBranchWithDefaults instantiates a new AdvisoryProductBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *AdvisoryProductBranch) GetBranches() []AdvisoryProductBranch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *AdvisoryProductBranch) GetBranchesOk() (*[]AdvisoryProductBranch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *AdvisoryProductBranch) SetBranches(v []AdvisoryProductBranch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *AdvisoryProductBranch) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetCategory

`func (o *AdvisoryProductBranch) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryProductBranch) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryProductBranch) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryProductBranch) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryProductBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryProductBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryProductBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryProductBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryProductBranch) GetProduct() AdvisoryProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryProductBranch) GetProductOk() (*AdvisoryProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryProductBranch) SetProduct(v AdvisoryProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryProductBranch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRelationships

`func (o *AdvisoryProductBranch) GetRelationships() []AdvisoryCSAFRelationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdvisoryProductBranch) GetRelationshipsOk() (*[]AdvisoryCSAFRelationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdvisoryProductBranch) SetRelationships(v []AdvisoryCSAFRelationship)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdvisoryProductBranch) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


