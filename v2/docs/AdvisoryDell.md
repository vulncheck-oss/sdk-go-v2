# AdvisoryDell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleNumber** | Pointer to **string** |  | [optional] 
**CombinedProductList** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DellCves** | Pointer to [**[]AdvisoryDellCVE**](AdvisoryDellCVE.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDell

`func NewAdvisoryDell() *AdvisoryDell`

NewAdvisoryDell instantiates a new AdvisoryDell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDellWithDefaults

`func NewAdvisoryDellWithDefaults() *AdvisoryDell`

NewAdvisoryDellWithDefaults instantiates a new AdvisoryDell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleNumber

`func (o *AdvisoryDell) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *AdvisoryDell) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *AdvisoryDell) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *AdvisoryDell) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetCombinedProductList

`func (o *AdvisoryDell) GetCombinedProductList() string`

GetCombinedProductList returns the CombinedProductList field if non-nil, zero value otherwise.

### GetCombinedProductListOk

`func (o *AdvisoryDell) GetCombinedProductListOk() (*string, bool)`

GetCombinedProductListOk returns a tuple with the CombinedProductList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedProductList

`func (o *AdvisoryDell) SetCombinedProductList(v string)`

SetCombinedProductList sets CombinedProductList field to given value.

### HasCombinedProductList

`func (o *AdvisoryDell) HasCombinedProductList() bool`

HasCombinedProductList returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDell) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDell) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDell) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDell) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryDell) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryDell) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryDell) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryDell) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDellCves

`func (o *AdvisoryDell) GetDellCves() []AdvisoryDellCVE`

GetDellCves returns the DellCves field if non-nil, zero value otherwise.

### GetDellCvesOk

`func (o *AdvisoryDell) GetDellCvesOk() (*[]AdvisoryDellCVE, bool)`

GetDellCvesOk returns a tuple with the DellCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDellCves

`func (o *AdvisoryDell) SetDellCves(v []AdvisoryDellCVE)`

SetDellCves sets DellCves field to given value.

### HasDellCves

`func (o *AdvisoryDell) HasDellCves() bool`

HasDellCves returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryDell) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryDell) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryDell) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryDell) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDell) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDell) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDell) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDell) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryDell) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryDell) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryDell) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryDell) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryDell) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryDell) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryDell) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryDell) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


