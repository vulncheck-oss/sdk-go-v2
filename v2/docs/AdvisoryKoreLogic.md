# AdvisoryKoreLogic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProduct** | Pointer to **string** |  | [optional] 
**AffectedVendor** | Pointer to **string** |  | [optional] 
**AffectedVersion** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryKoreLogic

`func NewAdvisoryKoreLogic() *AdvisoryKoreLogic`

NewAdvisoryKoreLogic instantiates a new AdvisoryKoreLogic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryKoreLogicWithDefaults

`func NewAdvisoryKoreLogicWithDefaults() *AdvisoryKoreLogic`

NewAdvisoryKoreLogicWithDefaults instantiates a new AdvisoryKoreLogic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProduct

`func (o *AdvisoryKoreLogic) GetAffectedProduct() string`

GetAffectedProduct returns the AffectedProduct field if non-nil, zero value otherwise.

### GetAffectedProductOk

`func (o *AdvisoryKoreLogic) GetAffectedProductOk() (*string, bool)`

GetAffectedProductOk returns a tuple with the AffectedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProduct

`func (o *AdvisoryKoreLogic) SetAffectedProduct(v string)`

SetAffectedProduct sets AffectedProduct field to given value.

### HasAffectedProduct

`func (o *AdvisoryKoreLogic) HasAffectedProduct() bool`

HasAffectedProduct returns a boolean if a field has been set.

### GetAffectedVendor

`func (o *AdvisoryKoreLogic) GetAffectedVendor() string`

GetAffectedVendor returns the AffectedVendor field if non-nil, zero value otherwise.

### GetAffectedVendorOk

`func (o *AdvisoryKoreLogic) GetAffectedVendorOk() (*string, bool)`

GetAffectedVendorOk returns a tuple with the AffectedVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVendor

`func (o *AdvisoryKoreLogic) SetAffectedVendor(v string)`

SetAffectedVendor sets AffectedVendor field to given value.

### HasAffectedVendor

`func (o *AdvisoryKoreLogic) HasAffectedVendor() bool`

HasAffectedVendor returns a boolean if a field has been set.

### GetAffectedVersion

`func (o *AdvisoryKoreLogic) GetAffectedVersion() string`

GetAffectedVersion returns the AffectedVersion field if non-nil, zero value otherwise.

### GetAffectedVersionOk

`func (o *AdvisoryKoreLogic) GetAffectedVersionOk() (*string, bool)`

GetAffectedVersionOk returns a tuple with the AffectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersion

`func (o *AdvisoryKoreLogic) SetAffectedVersion(v string)`

SetAffectedVersion sets AffectedVersion field to given value.

### HasAffectedVersion

`func (o *AdvisoryKoreLogic) HasAffectedVersion() bool`

HasAffectedVersion returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryKoreLogic) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryKoreLogic) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryKoreLogic) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryKoreLogic) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryKoreLogic) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryKoreLogic) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryKoreLogic) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryKoreLogic) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryKoreLogic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryKoreLogic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryKoreLogic) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryKoreLogic) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryKoreLogic) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryKoreLogic) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryKoreLogic) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryKoreLogic) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryKoreLogic) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryKoreLogic) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryKoreLogic) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryKoreLogic) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryKoreLogic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryKoreLogic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryKoreLogic) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryKoreLogic) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


