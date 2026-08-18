# ApiNVD20SsvcDecisionTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionPoints** | Pointer to [**[]ApiNVD20SsvcDecisionPoint**](ApiNVD20SsvcDecisionPoint.md) |  | [optional] 
**DecisionsTable** | Pointer to **[][]int32** | DecisionsTable is an array of objects with no declared properties, so there is nothing to model. In practice each entry maps decision point labels to the chosen option, plus the resulting outcome. | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20SsvcDecisionTree

`func NewApiNVD20SsvcDecisionTree() *ApiNVD20SsvcDecisionTree`

NewApiNVD20SsvcDecisionTree instantiates a new ApiNVD20SsvcDecisionTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20SsvcDecisionTreeWithDefaults

`func NewApiNVD20SsvcDecisionTreeWithDefaults() *ApiNVD20SsvcDecisionTree`

NewApiNVD20SsvcDecisionTreeWithDefaults instantiates a new ApiNVD20SsvcDecisionTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionPoints

`func (o *ApiNVD20SsvcDecisionTree) GetDecisionPoints() []ApiNVD20SsvcDecisionPoint`

GetDecisionPoints returns the DecisionPoints field if non-nil, zero value otherwise.

### GetDecisionPointsOk

`func (o *ApiNVD20SsvcDecisionTree) GetDecisionPointsOk() (*[]ApiNVD20SsvcDecisionPoint, bool)`

GetDecisionPointsOk returns a tuple with the DecisionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionPoints

`func (o *ApiNVD20SsvcDecisionTree) SetDecisionPoints(v []ApiNVD20SsvcDecisionPoint)`

SetDecisionPoints sets DecisionPoints field to given value.

### HasDecisionPoints

`func (o *ApiNVD20SsvcDecisionTree) HasDecisionPoints() bool`

HasDecisionPoints returns a boolean if a field has been set.

### GetDecisionsTable

`func (o *ApiNVD20SsvcDecisionTree) GetDecisionsTable() [][]int32`

GetDecisionsTable returns the DecisionsTable field if non-nil, zero value otherwise.

### GetDecisionsTableOk

`func (o *ApiNVD20SsvcDecisionTree) GetDecisionsTableOk() (*[][]int32, bool)`

GetDecisionsTableOk returns a tuple with the DecisionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionsTable

`func (o *ApiNVD20SsvcDecisionTree) SetDecisionsTable(v [][]int32)`

SetDecisionsTable sets DecisionsTable field to given value.

### HasDecisionsTable

`func (o *ApiNVD20SsvcDecisionTree) HasDecisionsTable() bool`

HasDecisionsTable returns a boolean if a field has been set.

### GetLang

`func (o *ApiNVD20SsvcDecisionTree) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ApiNVD20SsvcDecisionTree) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ApiNVD20SsvcDecisionTree) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ApiNVD20SsvcDecisionTree) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetRoles

`func (o *ApiNVD20SsvcDecisionTree) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiNVD20SsvcDecisionTree) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiNVD20SsvcDecisionTree) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiNVD20SsvcDecisionTree) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTitle

`func (o *ApiNVD20SsvcDecisionTree) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiNVD20SsvcDecisionTree) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiNVD20SsvcDecisionTree) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiNVD20SsvcDecisionTree) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *ApiNVD20SsvcDecisionTree) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiNVD20SsvcDecisionTree) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiNVD20SsvcDecisionTree) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiNVD20SsvcDecisionTree) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


