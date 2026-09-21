# PurlUnprocessedPurl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purl** | Pointer to **string** | The purl exactly as submitted. | [optional] 
**Reason** | Pointer to **string** | Why this purl was not looked up. One of: \&quot;unsupported_type\&quot; (a valid purl for an ecosystem VulnCheck does not index), \&quot;unparseable\&quot; (not a valid purl), \&quot;unsupported_distro\&quot; (a distro-scoped purl whose distro qualifier is missing or unrecognised, e.g. pkg:deb/debian/curl with no distro&#x3D;). Treat this as an open set: further values may be added. | [optional] 

## Methods

### NewPurlUnprocessedPurl

`func NewPurlUnprocessedPurl() *PurlUnprocessedPurl`

NewPurlUnprocessedPurl instantiates a new PurlUnprocessedPurl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurlUnprocessedPurlWithDefaults

`func NewPurlUnprocessedPurlWithDefaults() *PurlUnprocessedPurl`

NewPurlUnprocessedPurlWithDefaults instantiates a new PurlUnprocessedPurl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurl

`func (o *PurlUnprocessedPurl) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *PurlUnprocessedPurl) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *PurlUnprocessedPurl) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *PurlUnprocessedPurl) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetReason

`func (o *PurlUnprocessedPurl) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PurlUnprocessedPurl) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PurlUnprocessedPurl) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PurlUnprocessedPurl) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


