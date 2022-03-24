/*
DNS API

## Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](/docs/getstarted).  Please note that any zone or record updates might conflict with active services. In such cases, the DNS records belonging to the conflicting services will be deactivated.  ## Support Support questions may be posted in English: <a href='/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET. <h2> <details>   <summary>Release notes</summary>    Vesion 1.0.0   Exposed CRUD operations for customer zone. </details> </h2> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// ZoneTypes Represents the possible zone types.
type ZoneTypes string

// List of zoneTypes
const (
	NATIVE ZoneTypes = "NATIVE"
	SLAVE ZoneTypes = "SLAVE"
)

// All allowed values of ZoneTypes enum
var AllowedZoneTypesEnumValues = []ZoneTypes{
	"NATIVE",
	"SLAVE",
}

func (v *ZoneTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZoneTypes(value)
	for _, existing := range AllowedZoneTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZoneTypes", value)
}

// NewZoneTypesFromValue returns a pointer to a valid ZoneTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneTypesFromValue(v string) (*ZoneTypes, error) {
	ev := ZoneTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZoneTypes: valid values are %v", v, AllowedZoneTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneTypes) IsValid() bool {
	for _, existing := range AllowedZoneTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to zoneTypes value
func (v ZoneTypes) Ptr() *ZoneTypes {
	return &v
}

type NullableZoneTypes struct {
	value *ZoneTypes
	isSet bool
}

func (v NullableZoneTypes) Get() *ZoneTypes {
	return v.value
}

func (v *NullableZoneTypes) Set(val *ZoneTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTypes(val *ZoneTypes) *NullableZoneTypes {
	return &NullableZoneTypes{value: val, isSet: true}
}

func (v NullableZoneTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
