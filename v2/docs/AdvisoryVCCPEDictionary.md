# AdvisoryVCCPEDictionary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCPE** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryVCCPEDictionary

`func NewAdvisoryVCCPEDictionary() *AdvisoryVCCPEDictionary`

NewAdvisoryVCCPEDictionary instantiates a new AdvisoryVCCPEDictionary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVCCPEDictionaryWithDefaults

`func NewAdvisoryVCCPEDictionaryWithDefaults() *AdvisoryVCCPEDictionary`

NewAdvisoryVCCPEDictionaryWithDefaults instantiates a new AdvisoryVCCPEDictionary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCPE

`func (o *AdvisoryVCCPEDictionary) GetBaseCPE() string`

GetBaseCPE returns the BaseCPE field if non-nil, zero value otherwise.

### GetBaseCPEOk

`func (o *AdvisoryVCCPEDictionary) GetBaseCPEOk() (*string, bool)`

GetBaseCPEOk returns a tuple with the BaseCPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCPE

`func (o *AdvisoryVCCPEDictionary) SetBaseCPE(v string)`

SetBaseCPE sets BaseCPE field to given value.

### HasBaseCPE

`func (o *AdvisoryVCCPEDictionary) HasBaseCPE() bool`

HasBaseCPE returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryVCCPEDictionary) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryVCCPEDictionary) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryVCCPEDictionary) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryVCCPEDictionary) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


