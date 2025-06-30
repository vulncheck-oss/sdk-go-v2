# AdvisorySiemensAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsafUrl** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvrfUrl** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdate** | Pointer to **string** | could potentially kill this in the future as it&#39;s a dupe | [optional] 
**PdfUrl** | Pointer to **string** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**Ssa** | Pointer to [**AdvisorySSASource**](AdvisorySSASource.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TxtUrl** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySiemensAdvisory

`func NewAdvisorySiemensAdvisory() *AdvisorySiemensAdvisory`

NewAdvisorySiemensAdvisory instantiates a new AdvisorySiemensAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySiemensAdvisoryWithDefaults

`func NewAdvisorySiemensAdvisoryWithDefaults() *AdvisorySiemensAdvisory`

NewAdvisorySiemensAdvisoryWithDefaults instantiates a new AdvisorySiemensAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsafUrl

`func (o *AdvisorySiemensAdvisory) GetCsafUrl() string`

GetCsafUrl returns the CsafUrl field if non-nil, zero value otherwise.

### GetCsafUrlOk

`func (o *AdvisorySiemensAdvisory) GetCsafUrlOk() (*string, bool)`

GetCsafUrlOk returns a tuple with the CsafUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsafUrl

`func (o *AdvisorySiemensAdvisory) SetCsafUrl(v string)`

SetCsafUrl sets CsafUrl field to given value.

### HasCsafUrl

`func (o *AdvisorySiemensAdvisory) HasCsafUrl() bool`

HasCsafUrl returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySiemensAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySiemensAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySiemensAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySiemensAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvrfUrl

`func (o *AdvisorySiemensAdvisory) GetCvrfUrl() string`

GetCvrfUrl returns the CvrfUrl field if non-nil, zero value otherwise.

### GetCvrfUrlOk

`func (o *AdvisorySiemensAdvisory) GetCvrfUrlOk() (*string, bool)`

GetCvrfUrlOk returns a tuple with the CvrfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvrfUrl

`func (o *AdvisorySiemensAdvisory) SetCvrfUrl(v string)`

SetCvrfUrl sets CvrfUrl field to given value.

### HasCvrfUrl

`func (o *AdvisorySiemensAdvisory) HasCvrfUrl() bool`

HasCvrfUrl returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySiemensAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySiemensAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySiemensAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySiemensAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *AdvisorySiemensAdvisory) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *AdvisorySiemensAdvisory) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *AdvisorySiemensAdvisory) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *AdvisorySiemensAdvisory) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetId

`func (o *AdvisorySiemensAdvisory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisorySiemensAdvisory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisorySiemensAdvisory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisorySiemensAdvisory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *AdvisorySiemensAdvisory) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *AdvisorySiemensAdvisory) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *AdvisorySiemensAdvisory) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *AdvisorySiemensAdvisory) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetPdfUrl

`func (o *AdvisorySiemensAdvisory) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *AdvisorySiemensAdvisory) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *AdvisorySiemensAdvisory) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.

### HasPdfUrl

`func (o *AdvisorySiemensAdvisory) HasPdfUrl() bool`

HasPdfUrl returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisorySiemensAdvisory) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisorySiemensAdvisory) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisorySiemensAdvisory) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisorySiemensAdvisory) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSsa

`func (o *AdvisorySiemensAdvisory) GetSsa() AdvisorySSASource`

GetSsa returns the Ssa field if non-nil, zero value otherwise.

### GetSsaOk

`func (o *AdvisorySiemensAdvisory) GetSsaOk() (*AdvisorySSASource, bool)`

GetSsaOk returns a tuple with the Ssa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsa

`func (o *AdvisorySiemensAdvisory) SetSsa(v AdvisorySSASource)`

SetSsa sets Ssa field to given value.

### HasSsa

`func (o *AdvisorySiemensAdvisory) HasSsa() bool`

HasSsa returns a boolean if a field has been set.

### GetTags

`func (o *AdvisorySiemensAdvisory) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvisorySiemensAdvisory) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvisorySiemensAdvisory) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvisorySiemensAdvisory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySiemensAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySiemensAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySiemensAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySiemensAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTxtUrl

`func (o *AdvisorySiemensAdvisory) GetTxtUrl() string`

GetTxtUrl returns the TxtUrl field if non-nil, zero value otherwise.

### GetTxtUrlOk

`func (o *AdvisorySiemensAdvisory) GetTxtUrlOk() (*string, bool)`

GetTxtUrlOk returns a tuple with the TxtUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtUrl

`func (o *AdvisorySiemensAdvisory) SetTxtUrl(v string)`

SetTxtUrl sets TxtUrl field to given value.

### HasTxtUrl

`func (o *AdvisorySiemensAdvisory) HasTxtUrl() bool`

HasTxtUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisorySiemensAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisorySiemensAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisorySiemensAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisorySiemensAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


