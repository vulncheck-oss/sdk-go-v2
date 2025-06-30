# AdvisoryMicrosoftCVRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvrf** | Pointer to [**AdvisoryMSCVRF**](AdvisoryMSCVRF.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**ExploitedList** | Pointer to [**[]AdvisoryITW**](AdvisoryITW.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMicrosoftCVRF

`func NewAdvisoryMicrosoftCVRF() *AdvisoryMicrosoftCVRF`

NewAdvisoryMicrosoftCVRF instantiates a new AdvisoryMicrosoftCVRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMicrosoftCVRFWithDefaults

`func NewAdvisoryMicrosoftCVRFWithDefaults() *AdvisoryMicrosoftCVRF`

NewAdvisoryMicrosoftCVRFWithDefaults instantiates a new AdvisoryMicrosoftCVRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryMicrosoftCVRF) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMicrosoftCVRF) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMicrosoftCVRF) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMicrosoftCVRF) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvrf

`func (o *AdvisoryMicrosoftCVRF) GetCvrf() AdvisoryMSCVRF`

GetCvrf returns the Cvrf field if non-nil, zero value otherwise.

### GetCvrfOk

`func (o *AdvisoryMicrosoftCVRF) GetCvrfOk() (*AdvisoryMSCVRF, bool)`

GetCvrfOk returns a tuple with the Cvrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvrf

`func (o *AdvisoryMicrosoftCVRF) SetCvrf(v AdvisoryMSCVRF)`

SetCvrf sets Cvrf field to given value.

### HasCvrf

`func (o *AdvisoryMicrosoftCVRF) HasCvrf() bool`

HasCvrf returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMicrosoftCVRF) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMicrosoftCVRF) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMicrosoftCVRF) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMicrosoftCVRF) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetExploitedList

`func (o *AdvisoryMicrosoftCVRF) GetExploitedList() []AdvisoryITW`

GetExploitedList returns the ExploitedList field if non-nil, zero value otherwise.

### GetExploitedListOk

`func (o *AdvisoryMicrosoftCVRF) GetExploitedListOk() (*[]AdvisoryITW, bool)`

GetExploitedListOk returns a tuple with the ExploitedList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitedList

`func (o *AdvisoryMicrosoftCVRF) SetExploitedList(v []AdvisoryITW)`

SetExploitedList sets ExploitedList field to given value.

### HasExploitedList

`func (o *AdvisoryMicrosoftCVRF) HasExploitedList() bool`

HasExploitedList returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMicrosoftCVRF) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMicrosoftCVRF) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMicrosoftCVRF) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMicrosoftCVRF) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMicrosoftCVRF) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMicrosoftCVRF) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMicrosoftCVRF) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMicrosoftCVRF) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


