# AdvisoryOpenJDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**OpenjdkCves** | Pointer to [**[]AdvisoryOpenJDKCVE**](AdvisoryOpenJDKCVE.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOpenJDK

`func NewAdvisoryOpenJDK() *AdvisoryOpenJDK`

NewAdvisoryOpenJDK instantiates a new AdvisoryOpenJDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOpenJDKWithDefaults

`func NewAdvisoryOpenJDKWithDefaults() *AdvisoryOpenJDK`

NewAdvisoryOpenJDKWithDefaults instantiates a new AdvisoryOpenJDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryOpenJDK) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOpenJDK) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOpenJDK) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOpenJDK) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryOpenJDK) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryOpenJDK) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryOpenJDK) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryOpenJDK) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetOpenjdkCves

`func (o *AdvisoryOpenJDK) GetOpenjdkCves() []AdvisoryOpenJDKCVE`

GetOpenjdkCves returns the OpenjdkCves field if non-nil, zero value otherwise.

### GetOpenjdkCvesOk

`func (o *AdvisoryOpenJDK) GetOpenjdkCvesOk() (*[]AdvisoryOpenJDKCVE, bool)`

GetOpenjdkCvesOk returns a tuple with the OpenjdkCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenjdkCves

`func (o *AdvisoryOpenJDK) SetOpenjdkCves(v []AdvisoryOpenJDKCVE)`

SetOpenjdkCves sets OpenjdkCves field to given value.

### HasOpenjdkCves

`func (o *AdvisoryOpenJDK) HasOpenjdkCves() bool`

HasOpenjdkCves returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOpenJDK) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOpenJDK) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOpenJDK) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOpenJDK) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOpenJDK) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOpenJDK) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOpenJDK) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOpenJDK) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryOpenJDK) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryOpenJDK) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryOpenJDK) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryOpenJDK) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


