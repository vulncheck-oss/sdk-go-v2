# AdvisoryGoVulnRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]AdvisoryGoEvent**](AdvisoryGoEvent.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGoVulnRanges

`func NewAdvisoryGoVulnRanges() *AdvisoryGoVulnRanges`

NewAdvisoryGoVulnRanges instantiates a new AdvisoryGoVulnRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGoVulnRangesWithDefaults

`func NewAdvisoryGoVulnRangesWithDefaults() *AdvisoryGoVulnRanges`

NewAdvisoryGoVulnRangesWithDefaults instantiates a new AdvisoryGoVulnRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AdvisoryGoVulnRanges) GetEvents() []AdvisoryGoEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AdvisoryGoVulnRanges) GetEventsOk() (*[]AdvisoryGoEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AdvisoryGoVulnRanges) SetEvents(v []AdvisoryGoEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AdvisoryGoVulnRanges) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryGoVulnRanges) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryGoVulnRanges) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryGoVulnRanges) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryGoVulnRanges) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


