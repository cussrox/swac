// Code generated by github.com/swaggest/swac v0.1.0, DO NOT EDIT.

package petstore

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

// NewPet structure is generated from "#/definitions/NewPet".
type NewPet struct {
	Name                 string                 `json:"name,omitempty"`
	Tag                  string                 `json:"tag,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties
}

type marshalNewPet NewPet

// UnmarshalJSON decodes JSON.
func (i *NewPet) UnmarshalJSON(data []byte) error {
	ii := marshalNewPet(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"name",
			"tag",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = NewPet(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i NewPet) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalNewPet(i), i.AdditionalProperties)
}

// PetAllOf1 structure is generated from "#/definitions/Pet/allOf/1".
type PetAllOf1 struct {
	ID                   int64                  `json:"id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`            // All unmatched properties
}

type marshalPetAllOf1 PetAllOf1

// UnmarshalJSON decodes JSON.
func (i *PetAllOf1) UnmarshalJSON(data []byte) error {
	ii := marshalPetAllOf1(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"id",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = PetAllOf1(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i PetAllOf1) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPetAllOf1(i), i.AdditionalProperties)
}

// Pet structure is generated from "#/definitions/Pet".
type Pet struct {
	NewPet *NewPet    `json:"-"`
	AllOf1 *PetAllOf1 `json:"-"`
}

type marshalPet Pet

// UnmarshalJSON decodes JSON.
func (i *Pet) UnmarshalJSON(data []byte) error {

	err := unionMap{
		mustUnmarshal: []interface{}{&i.NewPet, &i.AllOf1},
		jsonData: data,
	}.unmarshal()

	return err
}

// MarshalJSON encodes JSON.
func (i Pet) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPet(i), i.NewPet, i.AllOf1)
}

// Error structure is generated from "#/definitions/Error".
type Error struct {
	Code                 int64                  `json:"code,omitempty"`
	Message              string                 `json:"message,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                 // All unmatched properties
}

type marshalError Error

// UnmarshalJSON decodes JSON.
func (i *Error) UnmarshalJSON(data []byte) error {
	ii := marshalError(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"code",
			"message",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = Error(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i Error) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalError(i), i.AdditionalProperties)
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := make([]byte, 1, 100)
	result[0] = '{'
	isObject := true
	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		if string(j) == "{}" {
			continue
		}
		if string(j) == "null" {
			continue
		}
		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false
				continue
			}
			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}
		result = append(result, j[1:]...)
	}
	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}
type unionMap struct {
	mustUnmarshal        []interface{}
	ignoreKeys           []string
	additionalProperties interface{}
	jsonData             []byte
}

func (u unionMap) unmarshal() error {
	for _, item := range u.mustUnmarshal {
		// unmarshal to struct
		err := json.Unmarshal(u.jsonData, item)
		if err != nil {
			return err
		}
	}

	if u.additionalProperties == nil {
		return nil
	}

	// unmarshal to a generic map
	var m map[string]*json.RawMessage
	err := json.Unmarshal(u.jsonData, &m)
	if err != nil {
		return err
	}
	// removing ignored keys (defined in struct)
	for _, i := range u.ignoreKeys {
		delete(m, i)
	}
	// returning early on empty map
	if len(m) == 0 {
		return nil
	}

	// Returning early on empty map.
	if len(m) == 0 {
		return nil
	}
	if u.additionalProperties != nil {
		return u.unmarshalAdditionalProperties(m)
	}
	return nil
}
func (u unionMap) unmarshalAdditionalProperties(m map[string]*json.RawMessage) error {
	var err error
	subMap := make([]byte, 1, 100)
	subMap[0] = '{'

	// Iterating map and filling additional properties.
	for key, val := range m {
		keyEscaped := `"` + strings.Replace(key, `"`, `\"`, -1) + `":`
		if len(subMap) != 1 {
			subMap = append(subMap[:len(subMap)-1], ',')
		}
		subMap = append(subMap, []byte(keyEscaped)...)
		if val != nil {
			subMap = append(subMap, []byte(*val)...)
		} else {
			subMap = append(subMap, []byte("null")...)
		}
		subMap = append(subMap, '}')
	}

	if len(subMap) > 1 {
		err = json.Unmarshal(subMap, u.additionalProperties)
		if err != nil {
			return err
		}
	}
	return nil
}
