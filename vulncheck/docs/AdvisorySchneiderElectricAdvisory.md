# AdvisorySchneiderElectricAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsafUrl** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cwe** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**PdfUrl** | Pointer to **string** |  | [optional] 
**SchneiderCves** | Pointer to [**[]AdvisorySchneiderCVE**](AdvisorySchneiderCVE.md) |  | [optional] 
**SchneiderElectricId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySchneiderElectricAdvisory

`func NewAdvisorySchneiderElectricAdvisory() *AdvisorySchneiderElectricAdvisory`

NewAdvisorySchneiderElectricAdvisory instantiates a new AdvisorySchneiderElectricAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySchneiderElectricAdvisoryWithDefaults

`func NewAdvisorySchneiderElectricAdvisoryWithDefaults() *AdvisorySchneiderElectricAdvisory`

NewAdvisorySchneiderElectricAdvisoryWithDefaults instantiates a new AdvisorySchneiderElectricAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsafUrl

`func (o *AdvisorySchneiderElectricAdvisory) GetCsafUrl() string`

GetCsafUrl returns the CsafUrl field if non-nil, zero value otherwise.

### GetCsafUrlOk

`func (o *AdvisorySchneiderElectricAdvisory) GetCsafUrlOk() (*string, bool)`

GetCsafUrlOk returns a tuple with the CsafUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsafUrl

`func (o *AdvisorySchneiderElectricAdvisory) SetCsafUrl(v string)`

SetCsafUrl sets CsafUrl field to given value.

### HasCsafUrl

`func (o *AdvisorySchneiderElectricAdvisory) HasCsafUrl() bool`

HasCsafUrl returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySchneiderElectricAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySchneiderElectricAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySchneiderElectricAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySchneiderElectricAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisorySchneiderElectricAdvisory) GetCwe() []string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisorySchneiderElectricAdvisory) GetCweOk() (*[]string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisorySchneiderElectricAdvisory) SetCwe(v []string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisorySchneiderElectricAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySchneiderElectricAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySchneiderElectricAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySchneiderElectricAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySchneiderElectricAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetPdfUrl

`func (o *AdvisorySchneiderElectricAdvisory) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *AdvisorySchneiderElectricAdvisory) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *AdvisorySchneiderElectricAdvisory) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.

### HasPdfUrl

`func (o *AdvisorySchneiderElectricAdvisory) HasPdfUrl() bool`

HasPdfUrl returns a boolean if a field has been set.

### GetSchneiderCves

`func (o *AdvisorySchneiderElectricAdvisory) GetSchneiderCves() []AdvisorySchneiderCVE`

GetSchneiderCves returns the SchneiderCves field if non-nil, zero value otherwise.

### GetSchneiderCvesOk

`func (o *AdvisorySchneiderElectricAdvisory) GetSchneiderCvesOk() (*[]AdvisorySchneiderCVE, bool)`

GetSchneiderCvesOk returns a tuple with the SchneiderCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchneiderCves

`func (o *AdvisorySchneiderElectricAdvisory) SetSchneiderCves(v []AdvisorySchneiderCVE)`

SetSchneiderCves sets SchneiderCves field to given value.

### HasSchneiderCves

`func (o *AdvisorySchneiderElectricAdvisory) HasSchneiderCves() bool`

HasSchneiderCves returns a boolean if a field has been set.

### GetSchneiderElectricId

`func (o *AdvisorySchneiderElectricAdvisory) GetSchneiderElectricId() string`

GetSchneiderElectricId returns the SchneiderElectricId field if non-nil, zero value otherwise.

### GetSchneiderElectricIdOk

`func (o *AdvisorySchneiderElectricAdvisory) GetSchneiderElectricIdOk() (*string, bool)`

GetSchneiderElectricIdOk returns a tuple with the SchneiderElectricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchneiderElectricId

`func (o *AdvisorySchneiderElectricAdvisory) SetSchneiderElectricId(v string)`

SetSchneiderElectricId sets SchneiderElectricId field to given value.

### HasSchneiderElectricId

`func (o *AdvisorySchneiderElectricAdvisory) HasSchneiderElectricId() bool`

HasSchneiderElectricId returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySchneiderElectricAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySchneiderElectricAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySchneiderElectricAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySchneiderElectricAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisorySchneiderElectricAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisorySchneiderElectricAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisorySchneiderElectricAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisorySchneiderElectricAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisorySchneiderElectricAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisorySchneiderElectricAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisorySchneiderElectricAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisorySchneiderElectricAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


