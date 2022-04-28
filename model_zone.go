/*
DNS API

## Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](/docs/getstarted).  Please note that any zone or record updates might conflict with active services. In such cases, the DNS records belonging to the conflicting services will be deactivated.  ## Support Support questions may be posted in English: <a href='/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET. <h2> <details>   <summary>Release notes</summary>   <ul>     <li>Version 1.0.0 Exposed CRUD operations for customer zone.</li>     <li>Version 1.0.1 Added response body for UPDATE and CREATE record operations.</li>   </ul> </details> </h2> 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// Zone struct for Zone
type Zone struct {
	// The zone name.
	Name *string `json:"name,omitempty"`
	// The zone id.
	Id *string `json:"id,omitempty"`
	Type *ZoneTypes `json:"type,omitempty"`
}

// NewZone instantiates a new Zone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZone() *Zone {
	this := Zone{}
	return &this
}

// NewZoneWithDefaults instantiates a new Zone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneWithDefaults() *Zone {
	this := Zone{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Zone) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Zone) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Zone) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Zone) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Zone) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Zone) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Zone) GetType() ZoneTypes {
	if o == nil || o.Type == nil {
		var ret ZoneTypes
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetTypeOk() (*ZoneTypes, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Zone) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ZoneTypes and assigns it to the Type field.
func (o *Zone) SetType(v ZoneTypes) {
	o.Type = &v
}

func (o Zone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableZone struct {
	value *Zone
	isSet bool
}

func (v NullableZone) Get() *Zone {
	return v.value
}

func (v *NullableZone) Set(val *Zone) {
	v.value = val
	v.isSet = true
}

func (v NullableZone) IsSet() bool {
	return v.isSet
}

func (v *NullableZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZone(val *Zone) *NullableZone {
	return &NullableZone{value: val, isSet: true}
}

func (v NullableZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


