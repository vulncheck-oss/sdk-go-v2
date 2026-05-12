# SearchResponseDataOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe** | Pointer to **string** |  | [optional] 
**CpeStruct** | Pointer to [**ApiCPE**](ApiCPE.md) |  | [optional] 
**Cves** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchResponseDataOut

`func NewSearchResponseDataOut() *SearchResponseDataOut`

NewSearchResponseDataOut instantiates a new SearchResponseDataOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseDataOutWithDefaults

`func NewSearchResponseDataOutWithDefaults() *SearchResponseDataOut`

NewSearchResponseDataOutWithDefaults instantiates a new SearchResponseDataOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe

`func (o *SearchResponseDataOut) GetCpe() string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *SearchResponseDataOut) GetCpeOk() (*string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *SearchResponseDataOut) SetCpe(v string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *SearchResponseDataOut) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCpeStruct

`func (o *SearchResponseDataOut) GetCpeStruct() ApiCPE`

GetCpeStruct returns the CpeStruct field if non-nil, zero value otherwise.

### GetCpeStructOk

`func (o *SearchResponseDataOut) GetCpeStructOk() (*ApiCPE, bool)`

GetCpeStructOk returns a tuple with the CpeStruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeStruct

`func (o *SearchResponseDataOut) SetCpeStruct(v ApiCPE)`

SetCpeStruct sets CpeStruct field to given value.

### HasCpeStruct

`func (o *SearchResponseDataOut) HasCpeStruct() bool`

HasCpeStruct returns a boolean if a field has been set.

### GetCves

`func (o *SearchResponseDataOut) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *SearchResponseDataOut) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *SearchResponseDataOut) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *SearchResponseDataOut) HasCves() bool`

HasCves returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


