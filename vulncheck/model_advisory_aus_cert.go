/*
VulnCheck API

Version 3 of the VulnCheck API

API version: 3.0
Contact: support@vulncheck.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vulncheck

import (
	"encoding/json"
)

// checks if the AdvisoryAusCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryAusCert{}

// AdvisoryAusCert struct for AdvisoryAusCert
type AdvisoryAusCert struct {
	Body *string `json:"body,omitempty"`
	BulletinId *string `json:"bulletinId,omitempty"`
	Cve []string `json:"cve,omitempty"`
	Cvss *string `json:"cvss,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Link *string `json:"link,omitempty"`
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	Product *string `json:"product,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Resolution *string `json:"resolution,omitempty"`
}

// NewAdvisoryAusCert instantiates a new AdvisoryAusCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryAusCert() *AdvisoryAusCert {
	this := AdvisoryAusCert{}
	return &this
}

// NewAdvisoryAusCertWithDefaults instantiates a new AdvisoryAusCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryAusCertWithDefaults() *AdvisoryAusCert {
	this := AdvisoryAusCert{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *AdvisoryAusCert) SetBody(v string) {
	o.Body = &v
}

// GetBulletinId returns the BulletinId field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetBulletinId() string {
	if o == nil || IsNil(o.BulletinId) {
		var ret string
		return ret
	}
	return *o.BulletinId
}

// GetBulletinIdOk returns a tuple with the BulletinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetBulletinIdOk() (*string, bool) {
	if o == nil || IsNil(o.BulletinId) {
		return nil, false
	}
	return o.BulletinId, true
}

// HasBulletinId returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasBulletinId() bool {
	if o != nil && !IsNil(o.BulletinId) {
		return true
	}

	return false
}

// SetBulletinId gets a reference to the given string and assigns it to the BulletinId field.
func (o *AdvisoryAusCert) SetBulletinId(v string) {
	o.BulletinId = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryAusCert) SetCve(v []string) {
	o.Cve = v
}

// GetCvss returns the Cvss field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetCvss() string {
	if o == nil || IsNil(o.Cvss) {
		var ret string
		return ret
	}
	return *o.Cvss
}

// GetCvssOk returns a tuple with the Cvss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetCvssOk() (*string, bool) {
	if o == nil || IsNil(o.Cvss) {
		return nil, false
	}
	return o.Cvss, true
}

// HasCvss returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasCvss() bool {
	if o != nil && !IsNil(o.Cvss) {
		return true
	}

	return false
}

// SetCvss gets a reference to the given string and assigns it to the Cvss field.
func (o *AdvisoryAusCert) SetCvss(v string) {
	o.Cvss = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryAusCert) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *AdvisoryAusCert) SetLink(v string) {
	o.Link = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetOperatingSystem() string {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetOperatingSystemOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *AdvisoryAusCert) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AdvisoryAusCert) SetProduct(v string) {
	o.Product = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetPublisher() string {
	if o == nil || IsNil(o.Publisher) {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetPublisherOk() (*string, bool) {
	if o == nil || IsNil(o.Publisher) {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasPublisher() bool {
	if o != nil && !IsNil(o.Publisher) {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *AdvisoryAusCert) SetPublisher(v string) {
	o.Publisher = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *AdvisoryAusCert) GetResolution() string {
	if o == nil || IsNil(o.Resolution) {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryAusCert) GetResolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *AdvisoryAusCert) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *AdvisoryAusCert) SetResolution(v string) {
	o.Resolution = &v
}

func (o AdvisoryAusCert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryAusCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.BulletinId) {
		toSerialize["bulletinId"] = o.BulletinId
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.Cvss) {
		toSerialize["cvss"] = o.Cvss
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["operatingSystem"] = o.OperatingSystem
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Publisher) {
		toSerialize["publisher"] = o.Publisher
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	return toSerialize, nil
}

type NullableAdvisoryAusCert struct {
	value *AdvisoryAusCert
	isSet bool
}

func (v NullableAdvisoryAusCert) Get() *AdvisoryAusCert {
	return v.value
}

func (v *NullableAdvisoryAusCert) Set(val *AdvisoryAusCert) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryAusCert) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryAusCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryAusCert(val *AdvisoryAusCert) *NullableAdvisoryAusCert {
	return &NullableAdvisoryAusCert{value: val, isSet: true}
}

func (v NullableAdvisoryAusCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryAusCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


