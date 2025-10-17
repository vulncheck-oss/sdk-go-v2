# AdvisoryOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**AdvisoryOverrideAnnotation**](AdvisoryOverrideAnnotation.md) |  | [optional] 
**Cve** | Pointer to [**AdvisoryOverrideCVE**](AdvisoryOverrideCVE.md) |  | [optional] 

## Methods

### NewAdvisoryOverride

`func NewAdvisoryOverride() *AdvisoryOverride`

NewAdvisoryOverride instantiates a new AdvisoryOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOverrideWithDefaults

`func NewAdvisoryOverrideWithDefaults() *AdvisoryOverride`

NewAdvisoryOverrideWithDefaults instantiates a new AdvisoryOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *AdvisoryOverride) GetAnnotation() AdvisoryOverrideAnnotation`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *AdvisoryOverride) GetAnnotationOk() (*AdvisoryOverrideAnnotation, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *AdvisoryOverride) SetAnnotation(v AdvisoryOverrideAnnotation)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *AdvisoryOverride) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryOverride) GetCve() AdvisoryOverrideCVE`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOverride) GetCveOk() (*AdvisoryOverrideCVE, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOverride) SetCve(v AdvisoryOverrideCVE)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOverride) HasCve() bool`

HasCve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


