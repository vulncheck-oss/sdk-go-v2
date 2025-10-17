# AdvisoryZDI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to **[]string** |  | [optional] 
**CvssScore** | Pointer to **string** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**CvssVersion** | Pointer to **string** |  | [optional] 
**Discoverers** | Pointer to **[]string** |  | [optional] 
**FilterIdsDv** | Pointer to **[]string** |  | [optional] 
**LastUpdatedAt** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]AdvisoryZDIProduct**](AdvisoryZDIProduct.md) |  | [optional] 
**PublicAdvisory** | Pointer to **string** |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**Responses** | Pointer to [**[]AdvisoryZDIResponse**](AdvisoryZDIResponse.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ZdiCan** | Pointer to **string** |  | [optional] 
**ZdiPublic** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryZDI

`func NewAdvisoryZDI() *AdvisoryZDI`

NewAdvisoryZDI instantiates a new AdvisoryZDI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryZDIWithDefaults

`func NewAdvisoryZDIWithDefaults() *AdvisoryZDI`

NewAdvisoryZDIWithDefaults instantiates a new AdvisoryZDI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *AdvisoryZDI) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *AdvisoryZDI) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *AdvisoryZDI) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *AdvisoryZDI) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryZDI) GetCvssScore() string`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryZDI) GetCvssScoreOk() (*string, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryZDI) SetCvssScore(v string)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryZDI) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryZDI) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryZDI) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryZDI) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryZDI) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCvssVersion

`func (o *AdvisoryZDI) GetCvssVersion() string`

GetCvssVersion returns the CvssVersion field if non-nil, zero value otherwise.

### GetCvssVersionOk

`func (o *AdvisoryZDI) GetCvssVersionOk() (*string, bool)`

GetCvssVersionOk returns a tuple with the CvssVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVersion

`func (o *AdvisoryZDI) SetCvssVersion(v string)`

SetCvssVersion sets CvssVersion field to given value.

### HasCvssVersion

`func (o *AdvisoryZDI) HasCvssVersion() bool`

HasCvssVersion returns a boolean if a field has been set.

### GetDiscoverers

`func (o *AdvisoryZDI) GetDiscoverers() []string`

GetDiscoverers returns the Discoverers field if non-nil, zero value otherwise.

### GetDiscoverersOk

`func (o *AdvisoryZDI) GetDiscoverersOk() (*[]string, bool)`

GetDiscoverersOk returns a tuple with the Discoverers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverers

`func (o *AdvisoryZDI) SetDiscoverers(v []string)`

SetDiscoverers sets Discoverers field to given value.

### HasDiscoverers

`func (o *AdvisoryZDI) HasDiscoverers() bool`

HasDiscoverers returns a boolean if a field has been set.

### GetFilterIdsDv

`func (o *AdvisoryZDI) GetFilterIdsDv() []string`

GetFilterIdsDv returns the FilterIdsDv field if non-nil, zero value otherwise.

### GetFilterIdsDvOk

`func (o *AdvisoryZDI) GetFilterIdsDvOk() (*[]string, bool)`

GetFilterIdsDvOk returns a tuple with the FilterIdsDv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIdsDv

`func (o *AdvisoryZDI) SetFilterIdsDv(v []string)`

SetFilterIdsDv sets FilterIdsDv field to given value.

### HasFilterIdsDv

`func (o *AdvisoryZDI) HasFilterIdsDv() bool`

HasFilterIdsDv returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *AdvisoryZDI) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *AdvisoryZDI) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *AdvisoryZDI) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *AdvisoryZDI) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisoryZDI) GetProducts() []AdvisoryZDIProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisoryZDI) GetProductsOk() (*[]AdvisoryZDIProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisoryZDI) SetProducts(v []AdvisoryZDIProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisoryZDI) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPublicAdvisory

`func (o *AdvisoryZDI) GetPublicAdvisory() string`

GetPublicAdvisory returns the PublicAdvisory field if non-nil, zero value otherwise.

### GetPublicAdvisoryOk

`func (o *AdvisoryZDI) GetPublicAdvisoryOk() (*string, bool)`

GetPublicAdvisoryOk returns a tuple with the PublicAdvisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAdvisory

`func (o *AdvisoryZDI) SetPublicAdvisory(v string)`

SetPublicAdvisory sets PublicAdvisory field to given value.

### HasPublicAdvisory

`func (o *AdvisoryZDI) HasPublicAdvisory() bool`

HasPublicAdvisory returns a boolean if a field has been set.

### GetPublishedDate

`func (o *AdvisoryZDI) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *AdvisoryZDI) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *AdvisoryZDI) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *AdvisoryZDI) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetResponses

`func (o *AdvisoryZDI) GetResponses() []AdvisoryZDIResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *AdvisoryZDI) GetResponsesOk() (*[]AdvisoryZDIResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *AdvisoryZDI) SetResponses(v []AdvisoryZDIResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *AdvisoryZDI) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryZDI) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryZDI) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryZDI) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryZDI) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetZdiCan

`func (o *AdvisoryZDI) GetZdiCan() string`

GetZdiCan returns the ZdiCan field if non-nil, zero value otherwise.

### GetZdiCanOk

`func (o *AdvisoryZDI) GetZdiCanOk() (*string, bool)`

GetZdiCanOk returns a tuple with the ZdiCan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZdiCan

`func (o *AdvisoryZDI) SetZdiCan(v string)`

SetZdiCan sets ZdiCan field to given value.

### HasZdiCan

`func (o *AdvisoryZDI) HasZdiCan() bool`

HasZdiCan returns a boolean if a field has been set.

### GetZdiPublic

`func (o *AdvisoryZDI) GetZdiPublic() string`

GetZdiPublic returns the ZdiPublic field if non-nil, zero value otherwise.

### GetZdiPublicOk

`func (o *AdvisoryZDI) GetZdiPublicOk() (*string, bool)`

GetZdiPublicOk returns a tuple with the ZdiPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZdiPublic

`func (o *AdvisoryZDI) SetZdiPublic(v string)`

SetZdiPublic sets ZdiPublic field to given value.

### HasZdiPublic

`func (o *AdvisoryZDI) HasZdiPublic() bool`

HasZdiPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


