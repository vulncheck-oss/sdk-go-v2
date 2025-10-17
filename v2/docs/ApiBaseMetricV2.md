# ApiBaseMetricV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcInsufInfo** | Pointer to **bool** |  | [optional] 
**CvssV2** | Pointer to [**ApiCVSSV2**](ApiCVSSV2.md) |  | [optional] 
**ExploitabilityScore** | Pointer to **float32** |  | [optional] 
**ImpactScore** | Pointer to **float32** |  | [optional] 
**ObtainAllPrivilege** | Pointer to **bool** |  | [optional] 
**ObtainOtherPrivilege** | Pointer to **bool** |  | [optional] 
**ObtainUserPrivilege** | Pointer to **bool** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**UserInteractionRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiBaseMetricV2

`func NewApiBaseMetricV2() *ApiBaseMetricV2`

NewApiBaseMetricV2 instantiates a new ApiBaseMetricV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBaseMetricV2WithDefaults

`func NewApiBaseMetricV2WithDefaults() *ApiBaseMetricV2`

NewApiBaseMetricV2WithDefaults instantiates a new ApiBaseMetricV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcInsufInfo

`func (o *ApiBaseMetricV2) GetAcInsufInfo() bool`

GetAcInsufInfo returns the AcInsufInfo field if non-nil, zero value otherwise.

### GetAcInsufInfoOk

`func (o *ApiBaseMetricV2) GetAcInsufInfoOk() (*bool, bool)`

GetAcInsufInfoOk returns a tuple with the AcInsufInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcInsufInfo

`func (o *ApiBaseMetricV2) SetAcInsufInfo(v bool)`

SetAcInsufInfo sets AcInsufInfo field to given value.

### HasAcInsufInfo

`func (o *ApiBaseMetricV2) HasAcInsufInfo() bool`

HasAcInsufInfo returns a boolean if a field has been set.

### GetCvssV2

`func (o *ApiBaseMetricV2) GetCvssV2() ApiCVSSV2`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *ApiBaseMetricV2) GetCvssV2Ok() (*ApiCVSSV2, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *ApiBaseMetricV2) SetCvssV2(v ApiCVSSV2)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *ApiBaseMetricV2) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetExploitabilityScore

`func (o *ApiBaseMetricV2) GetExploitabilityScore() float32`

GetExploitabilityScore returns the ExploitabilityScore field if non-nil, zero value otherwise.

### GetExploitabilityScoreOk

`func (o *ApiBaseMetricV2) GetExploitabilityScoreOk() (*float32, bool)`

GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityScore

`func (o *ApiBaseMetricV2) SetExploitabilityScore(v float32)`

SetExploitabilityScore sets ExploitabilityScore field to given value.

### HasExploitabilityScore

`func (o *ApiBaseMetricV2) HasExploitabilityScore() bool`

HasExploitabilityScore returns a boolean if a field has been set.

### GetImpactScore

`func (o *ApiBaseMetricV2) GetImpactScore() float32`

GetImpactScore returns the ImpactScore field if non-nil, zero value otherwise.

### GetImpactScoreOk

`func (o *ApiBaseMetricV2) GetImpactScoreOk() (*float32, bool)`

GetImpactScoreOk returns a tuple with the ImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactScore

`func (o *ApiBaseMetricV2) SetImpactScore(v float32)`

SetImpactScore sets ImpactScore field to given value.

### HasImpactScore

`func (o *ApiBaseMetricV2) HasImpactScore() bool`

HasImpactScore returns a boolean if a field has been set.

### GetObtainAllPrivilege

`func (o *ApiBaseMetricV2) GetObtainAllPrivilege() bool`

GetObtainAllPrivilege returns the ObtainAllPrivilege field if non-nil, zero value otherwise.

### GetObtainAllPrivilegeOk

`func (o *ApiBaseMetricV2) GetObtainAllPrivilegeOk() (*bool, bool)`

GetObtainAllPrivilegeOk returns a tuple with the ObtainAllPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainAllPrivilege

`func (o *ApiBaseMetricV2) SetObtainAllPrivilege(v bool)`

SetObtainAllPrivilege sets ObtainAllPrivilege field to given value.

### HasObtainAllPrivilege

`func (o *ApiBaseMetricV2) HasObtainAllPrivilege() bool`

HasObtainAllPrivilege returns a boolean if a field has been set.

### GetObtainOtherPrivilege

`func (o *ApiBaseMetricV2) GetObtainOtherPrivilege() bool`

GetObtainOtherPrivilege returns the ObtainOtherPrivilege field if non-nil, zero value otherwise.

### GetObtainOtherPrivilegeOk

`func (o *ApiBaseMetricV2) GetObtainOtherPrivilegeOk() (*bool, bool)`

GetObtainOtherPrivilegeOk returns a tuple with the ObtainOtherPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainOtherPrivilege

`func (o *ApiBaseMetricV2) SetObtainOtherPrivilege(v bool)`

SetObtainOtherPrivilege sets ObtainOtherPrivilege field to given value.

### HasObtainOtherPrivilege

`func (o *ApiBaseMetricV2) HasObtainOtherPrivilege() bool`

HasObtainOtherPrivilege returns a boolean if a field has been set.

### GetObtainUserPrivilege

`func (o *ApiBaseMetricV2) GetObtainUserPrivilege() bool`

GetObtainUserPrivilege returns the ObtainUserPrivilege field if non-nil, zero value otherwise.

### GetObtainUserPrivilegeOk

`func (o *ApiBaseMetricV2) GetObtainUserPrivilegeOk() (*bool, bool)`

GetObtainUserPrivilegeOk returns a tuple with the ObtainUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainUserPrivilege

`func (o *ApiBaseMetricV2) SetObtainUserPrivilege(v bool)`

SetObtainUserPrivilege sets ObtainUserPrivilege field to given value.

### HasObtainUserPrivilege

`func (o *ApiBaseMetricV2) HasObtainUserPrivilege() bool`

HasObtainUserPrivilege returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiBaseMetricV2) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiBaseMetricV2) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiBaseMetricV2) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiBaseMetricV2) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetUserInteractionRequired

`func (o *ApiBaseMetricV2) GetUserInteractionRequired() bool`

GetUserInteractionRequired returns the UserInteractionRequired field if non-nil, zero value otherwise.

### GetUserInteractionRequiredOk

`func (o *ApiBaseMetricV2) GetUserInteractionRequiredOk() (*bool, bool)`

GetUserInteractionRequiredOk returns a tuple with the UserInteractionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteractionRequired

`func (o *ApiBaseMetricV2) SetUserInteractionRequired(v bool)`

SetUserInteractionRequired sets UserInteractionRequired field to given value.

### HasUserInteractionRequired

`func (o *ApiBaseMetricV2) HasUserInteractionRequired() bool`

HasUserInteractionRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


