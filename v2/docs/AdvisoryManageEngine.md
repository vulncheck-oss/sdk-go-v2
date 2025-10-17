# AdvisoryManageEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADVISORY** | Pointer to **string** |  | [optional] 
**AddedTime** | Pointer to **string** |  | [optional] 
**CVEDetailsLink** | Pointer to [**AdvisoryCVEDetailsLink**](AdvisoryCVEDetailsLink.md) |  | [optional] 
**CVE_ID** | Pointer to **string** |  | [optional] 
**CVSSSeverityRating** | Pointer to **string** |  | [optional] 
**Fixed** | Pointer to **string** |  | [optional] 
**ForProductSearch** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Product** | Pointer to [**AdvisoryMEProduct**](AdvisoryMEProduct.md) |  | [optional] 
**ProductList** | Pointer to [**[]AdvisoryMEProduct**](AdvisoryMEProduct.md) |  | [optional] 
**ProductSpecificDetails** | Pointer to [**[]AdvisoryProductSpecificDetail**](AdvisoryProductSpecificDetail.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**IndexField** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryManageEngine

`func NewAdvisoryManageEngine() *AdvisoryManageEngine`

NewAdvisoryManageEngine instantiates a new AdvisoryManageEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryManageEngineWithDefaults

`func NewAdvisoryManageEngineWithDefaults() *AdvisoryManageEngine`

NewAdvisoryManageEngineWithDefaults instantiates a new AdvisoryManageEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetADVISORY

`func (o *AdvisoryManageEngine) GetADVISORY() string`

GetADVISORY returns the ADVISORY field if non-nil, zero value otherwise.

### GetADVISORYOk

`func (o *AdvisoryManageEngine) GetADVISORYOk() (*string, bool)`

GetADVISORYOk returns a tuple with the ADVISORY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVISORY

`func (o *AdvisoryManageEngine) SetADVISORY(v string)`

SetADVISORY sets ADVISORY field to given value.

### HasADVISORY

`func (o *AdvisoryManageEngine) HasADVISORY() bool`

HasADVISORY returns a boolean if a field has been set.

### GetAddedTime

`func (o *AdvisoryManageEngine) GetAddedTime() string`

GetAddedTime returns the AddedTime field if non-nil, zero value otherwise.

### GetAddedTimeOk

`func (o *AdvisoryManageEngine) GetAddedTimeOk() (*string, bool)`

GetAddedTimeOk returns a tuple with the AddedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTime

`func (o *AdvisoryManageEngine) SetAddedTime(v string)`

SetAddedTime sets AddedTime field to given value.

### HasAddedTime

`func (o *AdvisoryManageEngine) HasAddedTime() bool`

HasAddedTime returns a boolean if a field has been set.

### GetCVEDetailsLink

`func (o *AdvisoryManageEngine) GetCVEDetailsLink() AdvisoryCVEDetailsLink`

GetCVEDetailsLink returns the CVEDetailsLink field if non-nil, zero value otherwise.

### GetCVEDetailsLinkOk

`func (o *AdvisoryManageEngine) GetCVEDetailsLinkOk() (*AdvisoryCVEDetailsLink, bool)`

GetCVEDetailsLinkOk returns a tuple with the CVEDetailsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVEDetailsLink

`func (o *AdvisoryManageEngine) SetCVEDetailsLink(v AdvisoryCVEDetailsLink)`

SetCVEDetailsLink sets CVEDetailsLink field to given value.

### HasCVEDetailsLink

`func (o *AdvisoryManageEngine) HasCVEDetailsLink() bool`

HasCVEDetailsLink returns a boolean if a field has been set.

### GetCVE_ID

`func (o *AdvisoryManageEngine) GetCVE_ID() string`

GetCVE_ID returns the CVE_ID field if non-nil, zero value otherwise.

### GetCVE_IDOk

`func (o *AdvisoryManageEngine) GetCVE_IDOk() (*string, bool)`

GetCVE_IDOk returns a tuple with the CVE_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVE_ID

`func (o *AdvisoryManageEngine) SetCVE_ID(v string)`

SetCVE_ID sets CVE_ID field to given value.

### HasCVE_ID

`func (o *AdvisoryManageEngine) HasCVE_ID() bool`

HasCVE_ID returns a boolean if a field has been set.

### GetCVSSSeverityRating

`func (o *AdvisoryManageEngine) GetCVSSSeverityRating() string`

GetCVSSSeverityRating returns the CVSSSeverityRating field if non-nil, zero value otherwise.

### GetCVSSSeverityRatingOk

`func (o *AdvisoryManageEngine) GetCVSSSeverityRatingOk() (*string, bool)`

GetCVSSSeverityRatingOk returns a tuple with the CVSSSeverityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVSSSeverityRating

`func (o *AdvisoryManageEngine) SetCVSSSeverityRating(v string)`

SetCVSSSeverityRating sets CVSSSeverityRating field to given value.

### HasCVSSSeverityRating

`func (o *AdvisoryManageEngine) HasCVSSSeverityRating() bool`

HasCVSSSeverityRating returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryManageEngine) GetFixed() string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryManageEngine) GetFixedOk() (*string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryManageEngine) SetFixed(v string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryManageEngine) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetForProductSearch

`func (o *AdvisoryManageEngine) GetForProductSearch() string`

GetForProductSearch returns the ForProductSearch field if non-nil, zero value otherwise.

### GetForProductSearchOk

`func (o *AdvisoryManageEngine) GetForProductSearchOk() (*string, bool)`

GetForProductSearchOk returns a tuple with the ForProductSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForProductSearch

`func (o *AdvisoryManageEngine) SetForProductSearch(v string)`

SetForProductSearch sets ForProductSearch field to given value.

### HasForProductSearch

`func (o *AdvisoryManageEngine) HasForProductSearch() bool`

HasForProductSearch returns a boolean if a field has been set.

### GetID

`func (o *AdvisoryManageEngine) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *AdvisoryManageEngine) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *AdvisoryManageEngine) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *AdvisoryManageEngine) HasID() bool`

HasID returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryManageEngine) GetProduct() AdvisoryMEProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryManageEngine) GetProductOk() (*AdvisoryMEProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryManageEngine) SetProduct(v AdvisoryMEProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryManageEngine) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductList

`func (o *AdvisoryManageEngine) GetProductList() []AdvisoryMEProduct`

GetProductList returns the ProductList field if non-nil, zero value otherwise.

### GetProductListOk

`func (o *AdvisoryManageEngine) GetProductListOk() (*[]AdvisoryMEProduct, bool)`

GetProductListOk returns a tuple with the ProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductList

`func (o *AdvisoryManageEngine) SetProductList(v []AdvisoryMEProduct)`

SetProductList sets ProductList field to given value.

### HasProductList

`func (o *AdvisoryManageEngine) HasProductList() bool`

HasProductList returns a boolean if a field has been set.

### GetProductSpecificDetails

`func (o *AdvisoryManageEngine) GetProductSpecificDetails() []AdvisoryProductSpecificDetail`

GetProductSpecificDetails returns the ProductSpecificDetails field if non-nil, zero value otherwise.

### GetProductSpecificDetailsOk

`func (o *AdvisoryManageEngine) GetProductSpecificDetailsOk() (*[]AdvisoryProductSpecificDetail, bool)`

GetProductSpecificDetailsOk returns a tuple with the ProductSpecificDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSpecificDetails

`func (o *AdvisoryManageEngine) SetProductSpecificDetails(v []AdvisoryProductSpecificDetail)`

SetProductSpecificDetails sets ProductSpecificDetails field to given value.

### HasProductSpecificDetails

`func (o *AdvisoryManageEngine) HasProductSpecificDetails() bool`

HasProductSpecificDetails returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryManageEngine) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryManageEngine) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryManageEngine) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryManageEngine) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryManageEngine) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryManageEngine) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryManageEngine) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryManageEngine) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIndexField

`func (o *AdvisoryManageEngine) GetIndexField() string`

GetIndexField returns the IndexField field if non-nil, zero value otherwise.

### GetIndexFieldOk

`func (o *AdvisoryManageEngine) GetIndexFieldOk() (*string, bool)`

GetIndexFieldOk returns a tuple with the IndexField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexField

`func (o *AdvisoryManageEngine) SetIndexField(v string)`

SetIndexField sets IndexField field to given value.

### HasIndexField

`func (o *AdvisoryManageEngine) HasIndexField() bool`

HasIndexField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


