# AdvisorySAAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Threats** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**WarningDate** | Pointer to **string** |  | [optional] 
**WarningNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySAAdvisory

`func NewAdvisorySAAdvisory() *AdvisorySAAdvisory`

NewAdvisorySAAdvisory instantiates a new AdvisorySAAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySAAdvisoryWithDefaults

`func NewAdvisorySAAdvisoryWithDefaults() *AdvisorySAAdvisory`

NewAdvisorySAAdvisoryWithDefaults instantiates a new AdvisorySAAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisorySAAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySAAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySAAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySAAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySAAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySAAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySAAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySAAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisorySAAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisorySAAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisorySAAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisorySAAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLink

`func (o *AdvisorySAAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisorySAAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisorySAAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisorySAAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisorySAAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisorySAAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisorySAAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisorySAAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisorySAAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisorySAAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisorySAAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisorySAAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetThreats

`func (o *AdvisorySAAdvisory) GetThreats() string`

GetThreats returns the Threats field if non-nil, zero value otherwise.

### GetThreatsOk

`func (o *AdvisorySAAdvisory) GetThreatsOk() (*string, bool)`

GetThreatsOk returns a tuple with the Threats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreats

`func (o *AdvisorySAAdvisory) SetThreats(v string)`

SetThreats sets Threats field to given value.

### HasThreats

`func (o *AdvisorySAAdvisory) HasThreats() bool`

HasThreats returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisorySAAdvisory) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisorySAAdvisory) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisorySAAdvisory) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisorySAAdvisory) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetWarningDate

`func (o *AdvisorySAAdvisory) GetWarningDate() string`

GetWarningDate returns the WarningDate field if non-nil, zero value otherwise.

### GetWarningDateOk

`func (o *AdvisorySAAdvisory) GetWarningDateOk() (*string, bool)`

GetWarningDateOk returns a tuple with the WarningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDate

`func (o *AdvisorySAAdvisory) SetWarningDate(v string)`

SetWarningDate sets WarningDate field to given value.

### HasWarningDate

`func (o *AdvisorySAAdvisory) HasWarningDate() bool`

HasWarningDate returns a boolean if a field has been set.

### GetWarningNumber

`func (o *AdvisorySAAdvisory) GetWarningNumber() string`

GetWarningNumber returns the WarningNumber field if non-nil, zero value otherwise.

### GetWarningNumberOk

`func (o *AdvisorySAAdvisory) GetWarningNumberOk() (*string, bool)`

GetWarningNumberOk returns a tuple with the WarningNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningNumber

`func (o *AdvisorySAAdvisory) SetWarningNumber(v string)`

SetWarningNumber sets WarningNumber field to given value.

### HasWarningNumber

`func (o *AdvisorySAAdvisory) HasWarningNumber() bool`

HasWarningNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


