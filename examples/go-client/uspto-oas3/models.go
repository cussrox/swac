// Code generated by github.com/swaggest/swac v0.1.0, DO NOT EDIT.

package uspto

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

// ComponentsSchemasDataSetList structure is generated from "#/components/schemas/dataSetList".
type ComponentsSchemasDataSetList struct {
	Total                int64                                   `json:"total,omitempty"`
	Apis                 []ComponentsSchemasDataSetListApisItems `json:"apis,omitempty"`
	AdditionalProperties map[string]interface{}                  `json:"-"`               // All unmatched properties
}

type marshalComponentsSchemasDataSetList ComponentsSchemasDataSetList

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemasDataSetList) UnmarshalJSON(data []byte) error {
	ii := marshalComponentsSchemasDataSetList(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"total",
			"apis",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = ComponentsSchemasDataSetList(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemasDataSetList) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSchemasDataSetList(i), i.AdditionalProperties)
}

// ComponentsSchemasDataSetListApisItems structure is generated from "#/components/schemas/dataSetList->apis->items".
type ComponentsSchemasDataSetListApisItems struct {
	APIKey               string                 `json:"apiKey,omitempty"`              // To be used as a dataset parameter value
	APIVersionNumber     string                 `json:"apiVersionNumber,omitempty"`    // To be used as a version parameter value
	APIURL               string                 `json:"apiUrl,omitempty"`              // The URL describing the dataset's fields
	APIDocumentationURL  string                 `json:"apiDocumentationUrl,omitempty"` // A URL to the API console for each API
	AdditionalProperties map[string]interface{} `json:"-"`                             // All unmatched properties
}

type marshalComponentsSchemasDataSetListApisItems ComponentsSchemasDataSetListApisItems

// UnmarshalJSON decodes JSON.
func (i *ComponentsSchemasDataSetListApisItems) UnmarshalJSON(data []byte) error {
	ii := marshalComponentsSchemasDataSetListApisItems(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"apiKey",
			"apiVersionNumber",
			"apiUrl",
			"apiDocumentationUrl",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = ComponentsSchemasDataSetListApisItems(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i ComponentsSchemasDataSetListApisItems) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponentsSchemasDataSetListApisItems(i), i.AdditionalProperties)
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
