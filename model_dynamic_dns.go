/*
DNS API

## Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](https://developer.hosting.ionos.com/docs/getstarted).  Please note that any zone or record updates might conflict with active services. In such cases, the DNS records belonging to the conflicting services will be deactivated.  ## Support Support questions may be posted in English: <a href='https://developer.hosting.ionos.com/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET. <h2> <details>   <summary>Release notes</summary>   <ul>     <li>Version 1.0.0 Exposed CRUD operations for customer zone.</li>     <li>Version 1.0.1 Added response body for UPDATE and CREATE record operations.</li>   </ul> </details> </h2> 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// DynamicDns struct for DynamicDns
type DynamicDns struct {
	// DynDns configuration identifier.
	BulkId *string `json:"bulkId,omitempty"`
	// Use the url with GET to update the ips of (sub)domains. Query parameters: ipv4, ipv6.
	UpdateUrl *string `json:"updateUrl,omitempty"`
	Domains []string `json:"domains,omitempty"`
	// Dynamic Dns description.
	Description *string `json:"description,omitempty"`
}

// NewDynamicDns instantiates a new DynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicDns() *DynamicDns {
	this := DynamicDns{}
	return &this
}

// NewDynamicDnsWithDefaults instantiates a new DynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicDnsWithDefaults() *DynamicDns {
	this := DynamicDns{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *DynamicDns) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicDns) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *DynamicDns) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *DynamicDns) SetBulkId(v string) {
	o.BulkId = &v
}

// GetUpdateUrl returns the UpdateUrl field value if set, zero value otherwise.
func (o *DynamicDns) GetUpdateUrl() string {
	if o == nil || o.UpdateUrl == nil {
		var ret string
		return ret
	}
	return *o.UpdateUrl
}

// GetUpdateUrlOk returns a tuple with the UpdateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicDns) GetUpdateUrlOk() (*string, bool) {
	if o == nil || o.UpdateUrl == nil {
		return nil, false
	}
	return o.UpdateUrl, true
}

// HasUpdateUrl returns a boolean if a field has been set.
func (o *DynamicDns) HasUpdateUrl() bool {
	if o != nil && o.UpdateUrl != nil {
		return true
	}

	return false
}

// SetUpdateUrl gets a reference to the given string and assigns it to the UpdateUrl field.
func (o *DynamicDns) SetUpdateUrl(v string) {
	o.UpdateUrl = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *DynamicDns) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicDns) GetDomainsOk() ([]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *DynamicDns) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *DynamicDns) SetDomains(v []string) {
	o.Domains = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DynamicDns) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicDns) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DynamicDns) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DynamicDns) SetDescription(v string) {
	o.Description = &v
}

func (o DynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.UpdateUrl != nil {
		toSerialize["updateUrl"] = o.UpdateUrl
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableDynamicDns struct {
	value *DynamicDns
	isSet bool
}

func (v NullableDynamicDns) Get() *DynamicDns {
	return v.value
}

func (v *NullableDynamicDns) Set(val *DynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicDns(val *DynamicDns) *NullableDynamicDns {
	return &NullableDynamicDns{value: val, isSet: true}
}

func (v NullableDynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


