# AdvisoryAdvisoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bugzilla** | Pointer to [**AdvisoryBugzilla**](AdvisoryBugzilla.md) |  | [optional] 
**Cve** | Pointer to [**AdvisoryOvalCVE**](AdvisoryOvalCVE.md) |  | [optional] 
**Issued** | Pointer to [**AdvisoryIssued**](AdvisoryIssued.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to [**AdvisoryUpdated**](AdvisoryUpdated.md) |  | [optional] 

## Methods

### NewAdvisoryAdvisoryDetails

`func NewAdvisoryAdvisoryDetails() *AdvisoryAdvisoryDetails`

NewAdvisoryAdvisoryDetails instantiates a new AdvisoryAdvisoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAdvisoryDetailsWithDefaults

`func NewAdvisoryAdvisoryDetailsWithDefaults() *AdvisoryAdvisoryDetails`

NewAdvisoryAdvisoryDetailsWithDefaults instantiates a new AdvisoryAdvisoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBugzilla

`func (o *AdvisoryAdvisoryDetails) GetBugzilla() AdvisoryBugzilla`

GetBugzilla returns the Bugzilla field if non-nil, zero value otherwise.

### GetBugzillaOk

`func (o *AdvisoryAdvisoryDetails) GetBugzillaOk() (*AdvisoryBugzilla, bool)`

GetBugzillaOk returns a tuple with the Bugzilla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzilla

`func (o *AdvisoryAdvisoryDetails) SetBugzilla(v AdvisoryBugzilla)`

SetBugzilla sets Bugzilla field to given value.

### HasBugzilla

`func (o *AdvisoryAdvisoryDetails) HasBugzilla() bool`

HasBugzilla returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAdvisoryDetails) GetCve() AdvisoryOvalCVE`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAdvisoryDetails) GetCveOk() (*AdvisoryOvalCVE, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAdvisoryDetails) SetCve(v AdvisoryOvalCVE)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAdvisoryDetails) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetIssued

`func (o *AdvisoryAdvisoryDetails) GetIssued() AdvisoryIssued`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *AdvisoryAdvisoryDetails) GetIssuedOk() (*AdvisoryIssued, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *AdvisoryAdvisoryDetails) SetIssued(v AdvisoryIssued)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *AdvisoryAdvisoryDetails) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryAdvisoryDetails) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryAdvisoryDetails) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryAdvisoryDetails) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryAdvisoryDetails) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetUpdated

`func (o *AdvisoryAdvisoryDetails) GetUpdated() AdvisoryUpdated`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AdvisoryAdvisoryDetails) GetUpdatedOk() (*AdvisoryUpdated, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AdvisoryAdvisoryDetails) SetUpdated(v AdvisoryUpdated)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AdvisoryAdvisoryDetails) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


