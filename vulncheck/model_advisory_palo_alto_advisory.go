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

// checks if the AdvisoryPaloAltoAdvisory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryPaloAltoAdvisory{}

// AdvisoryPaloAltoAdvisory struct for AdvisoryPaloAltoAdvisory
type AdvisoryPaloAltoAdvisory struct {
	Affected *string `json:"affected,omitempty"`
	ApplicableVersions *string `json:"applicableVersions,omitempty"`
	AttackComplexity *string `json:"attackComplexity,omitempty"`
	AttackRequirements *string `json:"attackRequirements,omitempty"`
	AttackVector *string `json:"attackVector,omitempty"`
	AvailabilityImpact *string `json:"availabilityImpact,omitempty"`
	ConfidentialityImpact *string `json:"confidentialityImpact,omitempty"`
	Cve []string `json:"cve,omitempty"`
	CvssbaseScore *string `json:"cvssbaseScore,omitempty"`
	DatePublished *string `json:"datePublished,omitempty"`
	DateUpdated *string `json:"dateUpdated,omitempty"`
	DateAdded *string `json:"date_added,omitempty"`
	Id *string `json:"id,omitempty"`
	IntegrityImpact *string `json:"integrityImpact,omitempty"`
	PrivilegesRequired *string `json:"privilegesRequired,omitempty"`
	Problem *string `json:"problem,omitempty"`
	Product *string `json:"product,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Solution *string `json:"solution,omitempty"`
	Title *string `json:"title,omitempty"`
	Unaffected *string `json:"unaffected,omitempty"`
	Url *string `json:"url,omitempty"`
	UserInteraction *string `json:"userInteraction,omitempty"`
	Workaround *string `json:"workaround,omitempty"`
}

// NewAdvisoryPaloAltoAdvisory instantiates a new AdvisoryPaloAltoAdvisory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryPaloAltoAdvisory() *AdvisoryPaloAltoAdvisory {
	this := AdvisoryPaloAltoAdvisory{}
	return &this
}

// NewAdvisoryPaloAltoAdvisoryWithDefaults instantiates a new AdvisoryPaloAltoAdvisory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryPaloAltoAdvisoryWithDefaults() *AdvisoryPaloAltoAdvisory {
	this := AdvisoryPaloAltoAdvisory{}
	return &this
}

// GetAffected returns the Affected field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetAffected() string {
	if o == nil || IsNil(o.Affected) {
		var ret string
		return ret
	}
	return *o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetAffectedOk() (*string, bool) {
	if o == nil || IsNil(o.Affected) {
		return nil, false
	}
	return o.Affected, true
}

// HasAffected returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasAffected() bool {
	if o != nil && !IsNil(o.Affected) {
		return true
	}

	return false
}

// SetAffected gets a reference to the given string and assigns it to the Affected field.
func (o *AdvisoryPaloAltoAdvisory) SetAffected(v string) {
	o.Affected = &v
}

// GetApplicableVersions returns the ApplicableVersions field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetApplicableVersions() string {
	if o == nil || IsNil(o.ApplicableVersions) {
		var ret string
		return ret
	}
	return *o.ApplicableVersions
}

// GetApplicableVersionsOk returns a tuple with the ApplicableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetApplicableVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicableVersions) {
		return nil, false
	}
	return o.ApplicableVersions, true
}

// HasApplicableVersions returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasApplicableVersions() bool {
	if o != nil && !IsNil(o.ApplicableVersions) {
		return true
	}

	return false
}

// SetApplicableVersions gets a reference to the given string and assigns it to the ApplicableVersions field.
func (o *AdvisoryPaloAltoAdvisory) SetApplicableVersions(v string) {
	o.ApplicableVersions = &v
}

// GetAttackComplexity returns the AttackComplexity field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetAttackComplexity() string {
	if o == nil || IsNil(o.AttackComplexity) {
		var ret string
		return ret
	}
	return *o.AttackComplexity
}

// GetAttackComplexityOk returns a tuple with the AttackComplexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetAttackComplexityOk() (*string, bool) {
	if o == nil || IsNil(o.AttackComplexity) {
		return nil, false
	}
	return o.AttackComplexity, true
}

// HasAttackComplexity returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasAttackComplexity() bool {
	if o != nil && !IsNil(o.AttackComplexity) {
		return true
	}

	return false
}

// SetAttackComplexity gets a reference to the given string and assigns it to the AttackComplexity field.
func (o *AdvisoryPaloAltoAdvisory) SetAttackComplexity(v string) {
	o.AttackComplexity = &v
}

// GetAttackRequirements returns the AttackRequirements field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetAttackRequirements() string {
	if o == nil || IsNil(o.AttackRequirements) {
		var ret string
		return ret
	}
	return *o.AttackRequirements
}

// GetAttackRequirementsOk returns a tuple with the AttackRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetAttackRequirementsOk() (*string, bool) {
	if o == nil || IsNil(o.AttackRequirements) {
		return nil, false
	}
	return o.AttackRequirements, true
}

// HasAttackRequirements returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasAttackRequirements() bool {
	if o != nil && !IsNil(o.AttackRequirements) {
		return true
	}

	return false
}

// SetAttackRequirements gets a reference to the given string and assigns it to the AttackRequirements field.
func (o *AdvisoryPaloAltoAdvisory) SetAttackRequirements(v string) {
	o.AttackRequirements = &v
}

// GetAttackVector returns the AttackVector field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetAttackVector() string {
	if o == nil || IsNil(o.AttackVector) {
		var ret string
		return ret
	}
	return *o.AttackVector
}

// GetAttackVectorOk returns a tuple with the AttackVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetAttackVectorOk() (*string, bool) {
	if o == nil || IsNil(o.AttackVector) {
		return nil, false
	}
	return o.AttackVector, true
}

// HasAttackVector returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasAttackVector() bool {
	if o != nil && !IsNil(o.AttackVector) {
		return true
	}

	return false
}

// SetAttackVector gets a reference to the given string and assigns it to the AttackVector field.
func (o *AdvisoryPaloAltoAdvisory) SetAttackVector(v string) {
	o.AttackVector = &v
}

// GetAvailabilityImpact returns the AvailabilityImpact field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetAvailabilityImpact() string {
	if o == nil || IsNil(o.AvailabilityImpact) {
		var ret string
		return ret
	}
	return *o.AvailabilityImpact
}

// GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetAvailabilityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityImpact) {
		return nil, false
	}
	return o.AvailabilityImpact, true
}

// HasAvailabilityImpact returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasAvailabilityImpact() bool {
	if o != nil && !IsNil(o.AvailabilityImpact) {
		return true
	}

	return false
}

// SetAvailabilityImpact gets a reference to the given string and assigns it to the AvailabilityImpact field.
func (o *AdvisoryPaloAltoAdvisory) SetAvailabilityImpact(v string) {
	o.AvailabilityImpact = &v
}

// GetConfidentialityImpact returns the ConfidentialityImpact field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetConfidentialityImpact() string {
	if o == nil || IsNil(o.ConfidentialityImpact) {
		var ret string
		return ret
	}
	return *o.ConfidentialityImpact
}

// GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetConfidentialityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.ConfidentialityImpact) {
		return nil, false
	}
	return o.ConfidentialityImpact, true
}

// HasConfidentialityImpact returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasConfidentialityImpact() bool {
	if o != nil && !IsNil(o.ConfidentialityImpact) {
		return true
	}

	return false
}

// SetConfidentialityImpact gets a reference to the given string and assigns it to the ConfidentialityImpact field.
func (o *AdvisoryPaloAltoAdvisory) SetConfidentialityImpact(v string) {
	o.ConfidentialityImpact = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetCve() []string {
	if o == nil || IsNil(o.Cve) {
		var ret []string
		return ret
	}
	return o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetCveOk() ([]string, bool) {
	if o == nil || IsNil(o.Cve) {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasCve() bool {
	if o != nil && !IsNil(o.Cve) {
		return true
	}

	return false
}

// SetCve gets a reference to the given []string and assigns it to the Cve field.
func (o *AdvisoryPaloAltoAdvisory) SetCve(v []string) {
	o.Cve = v
}

// GetCvssbaseScore returns the CvssbaseScore field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetCvssbaseScore() string {
	if o == nil || IsNil(o.CvssbaseScore) {
		var ret string
		return ret
	}
	return *o.CvssbaseScore
}

// GetCvssbaseScoreOk returns a tuple with the CvssbaseScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetCvssbaseScoreOk() (*string, bool) {
	if o == nil || IsNil(o.CvssbaseScore) {
		return nil, false
	}
	return o.CvssbaseScore, true
}

// HasCvssbaseScore returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasCvssbaseScore() bool {
	if o != nil && !IsNil(o.CvssbaseScore) {
		return true
	}

	return false
}

// SetCvssbaseScore gets a reference to the given string and assigns it to the CvssbaseScore field.
func (o *AdvisoryPaloAltoAdvisory) SetCvssbaseScore(v string) {
	o.CvssbaseScore = &v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetDatePublished() string {
	if o == nil || IsNil(o.DatePublished) {
		var ret string
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetDatePublishedOk() (*string, bool) {
	if o == nil || IsNil(o.DatePublished) {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasDatePublished() bool {
	if o != nil && !IsNil(o.DatePublished) {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given string and assigns it to the DatePublished field.
func (o *AdvisoryPaloAltoAdvisory) SetDatePublished(v string) {
	o.DatePublished = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetDateUpdated() string {
	if o == nil || IsNil(o.DateUpdated) {
		var ret string
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetDateUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.DateUpdated) {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasDateUpdated() bool {
	if o != nil && !IsNil(o.DateUpdated) {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given string and assigns it to the DateUpdated field.
func (o *AdvisoryPaloAltoAdvisory) SetDateUpdated(v string) {
	o.DateUpdated = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetDateAdded() string {
	if o == nil || IsNil(o.DateAdded) {
		var ret string
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetDateAddedOk() (*string, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given string and assigns it to the DateAdded field.
func (o *AdvisoryPaloAltoAdvisory) SetDateAdded(v string) {
	o.DateAdded = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdvisoryPaloAltoAdvisory) SetId(v string) {
	o.Id = &v
}

// GetIntegrityImpact returns the IntegrityImpact field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetIntegrityImpact() string {
	if o == nil || IsNil(o.IntegrityImpact) {
		var ret string
		return ret
	}
	return *o.IntegrityImpact
}

// GetIntegrityImpactOk returns a tuple with the IntegrityImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetIntegrityImpactOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrityImpact) {
		return nil, false
	}
	return o.IntegrityImpact, true
}

// HasIntegrityImpact returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasIntegrityImpact() bool {
	if o != nil && !IsNil(o.IntegrityImpact) {
		return true
	}

	return false
}

// SetIntegrityImpact gets a reference to the given string and assigns it to the IntegrityImpact field.
func (o *AdvisoryPaloAltoAdvisory) SetIntegrityImpact(v string) {
	o.IntegrityImpact = &v
}

// GetPrivilegesRequired returns the PrivilegesRequired field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetPrivilegesRequired() string {
	if o == nil || IsNil(o.PrivilegesRequired) {
		var ret string
		return ret
	}
	return *o.PrivilegesRequired
}

// GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetPrivilegesRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.PrivilegesRequired) {
		return nil, false
	}
	return o.PrivilegesRequired, true
}

// HasPrivilegesRequired returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasPrivilegesRequired() bool {
	if o != nil && !IsNil(o.PrivilegesRequired) {
		return true
	}

	return false
}

// SetPrivilegesRequired gets a reference to the given string and assigns it to the PrivilegesRequired field.
func (o *AdvisoryPaloAltoAdvisory) SetPrivilegesRequired(v string) {
	o.PrivilegesRequired = &v
}

// GetProblem returns the Problem field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetProblem() string {
	if o == nil || IsNil(o.Problem) {
		var ret string
		return ret
	}
	return *o.Problem
}

// GetProblemOk returns a tuple with the Problem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetProblemOk() (*string, bool) {
	if o == nil || IsNil(o.Problem) {
		return nil, false
	}
	return o.Problem, true
}

// HasProblem returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasProblem() bool {
	if o != nil && !IsNil(o.Problem) {
		return true
	}

	return false
}

// SetProblem gets a reference to the given string and assigns it to the Problem field.
func (o *AdvisoryPaloAltoAdvisory) SetProblem(v string) {
	o.Problem = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AdvisoryPaloAltoAdvisory) SetProduct(v string) {
	o.Product = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AdvisoryPaloAltoAdvisory) SetScope(v string) {
	o.Scope = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AdvisoryPaloAltoAdvisory) SetSeverity(v string) {
	o.Severity = &v
}

// GetSolution returns the Solution field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetSolution() string {
	if o == nil || IsNil(o.Solution) {
		var ret string
		return ret
	}
	return *o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetSolutionOk() (*string, bool) {
	if o == nil || IsNil(o.Solution) {
		return nil, false
	}
	return o.Solution, true
}

// HasSolution returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasSolution() bool {
	if o != nil && !IsNil(o.Solution) {
		return true
	}

	return false
}

// SetSolution gets a reference to the given string and assigns it to the Solution field.
func (o *AdvisoryPaloAltoAdvisory) SetSolution(v string) {
	o.Solution = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AdvisoryPaloAltoAdvisory) SetTitle(v string) {
	o.Title = &v
}

// GetUnaffected returns the Unaffected field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetUnaffected() string {
	if o == nil || IsNil(o.Unaffected) {
		var ret string
		return ret
	}
	return *o.Unaffected
}

// GetUnaffectedOk returns a tuple with the Unaffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetUnaffectedOk() (*string, bool) {
	if o == nil || IsNil(o.Unaffected) {
		return nil, false
	}
	return o.Unaffected, true
}

// HasUnaffected returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasUnaffected() bool {
	if o != nil && !IsNil(o.Unaffected) {
		return true
	}

	return false
}

// SetUnaffected gets a reference to the given string and assigns it to the Unaffected field.
func (o *AdvisoryPaloAltoAdvisory) SetUnaffected(v string) {
	o.Unaffected = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvisoryPaloAltoAdvisory) SetUrl(v string) {
	o.Url = &v
}

// GetUserInteraction returns the UserInteraction field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetUserInteraction() string {
	if o == nil || IsNil(o.UserInteraction) {
		var ret string
		return ret
	}
	return *o.UserInteraction
}

// GetUserInteractionOk returns a tuple with the UserInteraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetUserInteractionOk() (*string, bool) {
	if o == nil || IsNil(o.UserInteraction) {
		return nil, false
	}
	return o.UserInteraction, true
}

// HasUserInteraction returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasUserInteraction() bool {
	if o != nil && !IsNil(o.UserInteraction) {
		return true
	}

	return false
}

// SetUserInteraction gets a reference to the given string and assigns it to the UserInteraction field.
func (o *AdvisoryPaloAltoAdvisory) SetUserInteraction(v string) {
	o.UserInteraction = &v
}

// GetWorkaround returns the Workaround field value if set, zero value otherwise.
func (o *AdvisoryPaloAltoAdvisory) GetWorkaround() string {
	if o == nil || IsNil(o.Workaround) {
		var ret string
		return ret
	}
	return *o.Workaround
}

// GetWorkaroundOk returns a tuple with the Workaround field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryPaloAltoAdvisory) GetWorkaroundOk() (*string, bool) {
	if o == nil || IsNil(o.Workaround) {
		return nil, false
	}
	return o.Workaround, true
}

// HasWorkaround returns a boolean if a field has been set.
func (o *AdvisoryPaloAltoAdvisory) HasWorkaround() bool {
	if o != nil && !IsNil(o.Workaround) {
		return true
	}

	return false
}

// SetWorkaround gets a reference to the given string and assigns it to the Workaround field.
func (o *AdvisoryPaloAltoAdvisory) SetWorkaround(v string) {
	o.Workaround = &v
}

func (o AdvisoryPaloAltoAdvisory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryPaloAltoAdvisory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affected) {
		toSerialize["affected"] = o.Affected
	}
	if !IsNil(o.ApplicableVersions) {
		toSerialize["applicableVersions"] = o.ApplicableVersions
	}
	if !IsNil(o.AttackComplexity) {
		toSerialize["attackComplexity"] = o.AttackComplexity
	}
	if !IsNil(o.AttackRequirements) {
		toSerialize["attackRequirements"] = o.AttackRequirements
	}
	if !IsNil(o.AttackVector) {
		toSerialize["attackVector"] = o.AttackVector
	}
	if !IsNil(o.AvailabilityImpact) {
		toSerialize["availabilityImpact"] = o.AvailabilityImpact
	}
	if !IsNil(o.ConfidentialityImpact) {
		toSerialize["confidentialityImpact"] = o.ConfidentialityImpact
	}
	if !IsNil(o.Cve) {
		toSerialize["cve"] = o.Cve
	}
	if !IsNil(o.CvssbaseScore) {
		toSerialize["cvssbaseScore"] = o.CvssbaseScore
	}
	if !IsNil(o.DatePublished) {
		toSerialize["datePublished"] = o.DatePublished
	}
	if !IsNil(o.DateUpdated) {
		toSerialize["dateUpdated"] = o.DateUpdated
	}
	if !IsNil(o.DateAdded) {
		toSerialize["date_added"] = o.DateAdded
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrityImpact) {
		toSerialize["integrityImpact"] = o.IntegrityImpact
	}
	if !IsNil(o.PrivilegesRequired) {
		toSerialize["privilegesRequired"] = o.PrivilegesRequired
	}
	if !IsNil(o.Problem) {
		toSerialize["problem"] = o.Problem
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Solution) {
		toSerialize["solution"] = o.Solution
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Unaffected) {
		toSerialize["unaffected"] = o.Unaffected
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UserInteraction) {
		toSerialize["userInteraction"] = o.UserInteraction
	}
	if !IsNil(o.Workaround) {
		toSerialize["workaround"] = o.Workaround
	}
	return toSerialize, nil
}

type NullableAdvisoryPaloAltoAdvisory struct {
	value *AdvisoryPaloAltoAdvisory
	isSet bool
}

func (v NullableAdvisoryPaloAltoAdvisory) Get() *AdvisoryPaloAltoAdvisory {
	return v.value
}

func (v *NullableAdvisoryPaloAltoAdvisory) Set(val *AdvisoryPaloAltoAdvisory) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryPaloAltoAdvisory) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryPaloAltoAdvisory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryPaloAltoAdvisory(val *AdvisoryPaloAltoAdvisory) *NullableAdvisoryPaloAltoAdvisory {
	return &NullableAdvisoryPaloAltoAdvisory{value: val, isSet: true}
}

func (v NullableAdvisoryPaloAltoAdvisory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryPaloAltoAdvisory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


