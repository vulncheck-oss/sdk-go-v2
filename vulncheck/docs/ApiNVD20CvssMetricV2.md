# ApiNVD20CvssMetricV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcInsufInfo** | Pointer to **bool** |  | [optional] 
**BaseSeverity** | Pointer to **string** |  | [optional] 
**CvssData** | Pointer to [**ApiNVD20CvssDataV2**](ApiNVD20CvssDataV2.md) |  | [optional] 
**ExploitabilityScore** | Pointer to **float32** |  | [optional] 
**ImpactScore** | Pointer to **float32** |  | [optional] 
**ObtainAllPrivilege** | Pointer to **bool** |  | [optional] 
**ObtainOtherPrivilege** | Pointer to **bool** |  | [optional] 
**ObtainUserPrivilege** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserInteractionRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiNVD20CvssMetricV2

`func NewApiNVD20CvssMetricV2() *ApiNVD20CvssMetricV2`

NewApiNVD20CvssMetricV2 instantiates a new ApiNVD20CvssMetricV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CvssMetricV2WithDefaults

`func NewApiNVD20CvssMetricV2WithDefaults() *ApiNVD20CvssMetricV2`

NewApiNVD20CvssMetricV2WithDefaults instantiates a new ApiNVD20CvssMetricV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcInsufInfo

`func (o *ApiNVD20CvssMetricV2) GetAcInsufInfo() bool`

GetAcInsufInfo returns the AcInsufInfo field if non-nil, zero value otherwise.

### GetAcInsufInfoOk

`func (o *ApiNVD20CvssMetricV2) GetAcInsufInfoOk() (*bool, bool)`

GetAcInsufInfoOk returns a tuple with the AcInsufInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcInsufInfo

`func (o *ApiNVD20CvssMetricV2) SetAcInsufInfo(v bool)`

SetAcInsufInfo sets AcInsufInfo field to given value.

### HasAcInsufInfo

`func (o *ApiNVD20CvssMetricV2) HasAcInsufInfo() bool`

HasAcInsufInfo returns a boolean if a field has been set.

### GetBaseSeverity

`func (o *ApiNVD20CvssMetricV2) GetBaseSeverity() string`

GetBaseSeverity returns the BaseSeverity field if non-nil, zero value otherwise.

### GetBaseSeverityOk

`func (o *ApiNVD20CvssMetricV2) GetBaseSeverityOk() (*string, bool)`

GetBaseSeverityOk returns a tuple with the BaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSeverity

`func (o *ApiNVD20CvssMetricV2) SetBaseSeverity(v string)`

SetBaseSeverity sets BaseSeverity field to given value.

### HasBaseSeverity

`func (o *ApiNVD20CvssMetricV2) HasBaseSeverity() bool`

HasBaseSeverity returns a boolean if a field has been set.

### GetCvssData

`func (o *ApiNVD20CvssMetricV2) GetCvssData() ApiNVD20CvssDataV2`

GetCvssData returns the CvssData field if non-nil, zero value otherwise.

### GetCvssDataOk

`func (o *ApiNVD20CvssMetricV2) GetCvssDataOk() (*ApiNVD20CvssDataV2, bool)`

GetCvssDataOk returns a tuple with the CvssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssData

`func (o *ApiNVD20CvssMetricV2) SetCvssData(v ApiNVD20CvssDataV2)`

SetCvssData sets CvssData field to given value.

### HasCvssData

`func (o *ApiNVD20CvssMetricV2) HasCvssData() bool`

HasCvssData returns a boolean if a field has been set.

### GetExploitabilityScore

`func (o *ApiNVD20CvssMetricV2) GetExploitabilityScore() float32`

GetExploitabilityScore returns the ExploitabilityScore field if non-nil, zero value otherwise.

### GetExploitabilityScoreOk

`func (o *ApiNVD20CvssMetricV2) GetExploitabilityScoreOk() (*float32, bool)`

GetExploitabilityScoreOk returns a tuple with the ExploitabilityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityScore

`func (o *ApiNVD20CvssMetricV2) SetExploitabilityScore(v float32)`

SetExploitabilityScore sets ExploitabilityScore field to given value.

### HasExploitabilityScore

`func (o *ApiNVD20CvssMetricV2) HasExploitabilityScore() bool`

HasExploitabilityScore returns a boolean if a field has been set.

### GetImpactScore

`func (o *ApiNVD20CvssMetricV2) GetImpactScore() float32`

GetImpactScore returns the ImpactScore field if non-nil, zero value otherwise.

### GetImpactScoreOk

`func (o *ApiNVD20CvssMetricV2) GetImpactScoreOk() (*float32, bool)`

GetImpactScoreOk returns a tuple with the ImpactScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactScore

`func (o *ApiNVD20CvssMetricV2) SetImpactScore(v float32)`

SetImpactScore sets ImpactScore field to given value.

### HasImpactScore

`func (o *ApiNVD20CvssMetricV2) HasImpactScore() bool`

HasImpactScore returns a boolean if a field has been set.

### GetObtainAllPrivilege

`func (o *ApiNVD20CvssMetricV2) GetObtainAllPrivilege() bool`

GetObtainAllPrivilege returns the ObtainAllPrivilege field if non-nil, zero value otherwise.

### GetObtainAllPrivilegeOk

`func (o *ApiNVD20CvssMetricV2) GetObtainAllPrivilegeOk() (*bool, bool)`

GetObtainAllPrivilegeOk returns a tuple with the ObtainAllPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainAllPrivilege

`func (o *ApiNVD20CvssMetricV2) SetObtainAllPrivilege(v bool)`

SetObtainAllPrivilege sets ObtainAllPrivilege field to given value.

### HasObtainAllPrivilege

`func (o *ApiNVD20CvssMetricV2) HasObtainAllPrivilege() bool`

HasObtainAllPrivilege returns a boolean if a field has been set.

### GetObtainOtherPrivilege

`func (o *ApiNVD20CvssMetricV2) GetObtainOtherPrivilege() bool`

GetObtainOtherPrivilege returns the ObtainOtherPrivilege field if non-nil, zero value otherwise.

### GetObtainOtherPrivilegeOk

`func (o *ApiNVD20CvssMetricV2) GetObtainOtherPrivilegeOk() (*bool, bool)`

GetObtainOtherPrivilegeOk returns a tuple with the ObtainOtherPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainOtherPrivilege

`func (o *ApiNVD20CvssMetricV2) SetObtainOtherPrivilege(v bool)`

SetObtainOtherPrivilege sets ObtainOtherPrivilege field to given value.

### HasObtainOtherPrivilege

`func (o *ApiNVD20CvssMetricV2) HasObtainOtherPrivilege() bool`

HasObtainOtherPrivilege returns a boolean if a field has been set.

### GetObtainUserPrivilege

`func (o *ApiNVD20CvssMetricV2) GetObtainUserPrivilege() bool`

GetObtainUserPrivilege returns the ObtainUserPrivilege field if non-nil, zero value otherwise.

### GetObtainUserPrivilegeOk

`func (o *ApiNVD20CvssMetricV2) GetObtainUserPrivilegeOk() (*bool, bool)`

GetObtainUserPrivilegeOk returns a tuple with the ObtainUserPrivilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainUserPrivilege

`func (o *ApiNVD20CvssMetricV2) SetObtainUserPrivilege(v bool)`

SetObtainUserPrivilege sets ObtainUserPrivilege field to given value.

### HasObtainUserPrivilege

`func (o *ApiNVD20CvssMetricV2) HasObtainUserPrivilege() bool`

HasObtainUserPrivilege returns a boolean if a field has been set.

### GetSource

`func (o *ApiNVD20CvssMetricV2) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiNVD20CvssMetricV2) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiNVD20CvssMetricV2) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiNVD20CvssMetricV2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *ApiNVD20CvssMetricV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiNVD20CvssMetricV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiNVD20CvssMetricV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiNVD20CvssMetricV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserInteractionRequired

`func (o *ApiNVD20CvssMetricV2) GetUserInteractionRequired() bool`

GetUserInteractionRequired returns the UserInteractionRequired field if non-nil, zero value otherwise.

### GetUserInteractionRequiredOk

`func (o *ApiNVD20CvssMetricV2) GetUserInteractionRequiredOk() (*bool, bool)`

GetUserInteractionRequiredOk returns a tuple with the UserInteractionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteractionRequired

`func (o *ApiNVD20CvssMetricV2) SetUserInteractionRequired(v bool)`

SetUserInteractionRequired sets UserInteractionRequired field to given value.

### HasUserInteractionRequired

`func (o *ApiNVD20CvssMetricV2) HasUserInteractionRequired() bool`

HasUserInteractionRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


