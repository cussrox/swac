// Code generated by github.com/swaggest/swac v0.0.2, DO NOT EDIT.

package uber

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

// Product structure is generated from "#/definitions/Product".
type Product struct {
	ID                   string                 `json:"product_id,omitempty"`   // Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco will have a different product_id than uberX in Los Angeles.
	Description          string                 `json:"description,omitempty"`  // Description of product.
	DisplayName          string                 `json:"display_name,omitempty"` // Display name of product.
	Capacity             string                 `json:"capacity,omitempty"`     // Capacity of product. For example, 4 people.
	Image                string                 `json:"image,omitempty"`        // Image URL representing the product.
	AdditionalProperties map[string]interface{} `json:"-"`                      // All unmatched properties
}

type marshalProduct Product

// UnmarshalJSON decodes JSON.
func (i *Product) UnmarshalJSON(data []byte) error {
	ii := marshalProduct(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"product_id",
			"description",
			"display_name",
			"capacity",
			"image",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = Product(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i Product) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProduct(i), i.AdditionalProperties)
}

// Error structure is generated from "#/definitions/Error".
type Error struct {
	Code                 int64                  `json:"code,omitempty"`
	Message              string                 `json:"message,omitempty"`
	Fields               string                 `json:"fields,omitempty"`
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
			"fields",
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

// PriceEstimate structure is generated from "#/definitions/PriceEstimate".
type PriceEstimate struct {
	ProductID            string                 `json:"product_id,omitempty"`       // Unique identifier representing a specific product for a given latitude & longitude. For example, uberX in San Francisco will have a different product_id than uberX in Los Angeles
	CurrencyCode         string                 `json:"currency_code,omitempty"`    // [ISO 4217](http://en.wikipedia.org/wiki/ISO_4217) currency code.
	DisplayName          string                 `json:"display_name,omitempty"`     // Display name of product.
	Estimate             string                 `json:"estimate,omitempty"`         // Formatted string of estimate in local currency of the start location. Estimate could be a range, a single number (flat rate) or "Metered" for TAXI.
	LowEstimate          float64                `json:"low_estimate,omitempty"`     // Lower bound of the estimated price.
	HighEstimate         float64                `json:"high_estimate,omitempty"`    // Upper bound of the estimated price.
	SurgeMultiplier      float64                `json:"surge_multiplier,omitempty"` // Expected surge multiplier. Surge is active if surge_multiplier is greater than 1. Price estimate already factors in the surge multiplier.
	AdditionalProperties map[string]interface{} `json:"-"`                          // All unmatched properties
}

type marshalPriceEstimate PriceEstimate

// UnmarshalJSON decodes JSON.
func (i *PriceEstimate) UnmarshalJSON(data []byte) error {
	ii := marshalPriceEstimate(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"product_id",
			"currency_code",
			"display_name",
			"estimate",
			"low_estimate",
			"high_estimate",
			"surge_multiplier",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = PriceEstimate(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i PriceEstimate) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPriceEstimate(i), i.AdditionalProperties)
}

// Profile structure is generated from "#/definitions/Profile".
type Profile struct {
	FirstName            string                 `json:"first_name,omitempty"` // First name of the Uber user.
	LastName             string                 `json:"last_name,omitempty"`  // Last name of the Uber user.
	Email                string                 `json:"email,omitempty"`      // Email address of the Uber user
	Picture              string                 `json:"picture,omitempty"`    // Image URL of the Uber user.
	PromoCode            string                 `json:"promo_code,omitempty"` // Promo code of the Uber user.
	AdditionalProperties map[string]interface{} `json:"-"`                    // All unmatched properties
}

type marshalProfile Profile

// UnmarshalJSON decodes JSON.
func (i *Profile) UnmarshalJSON(data []byte) error {
	ii := marshalProfile(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"first_name",
			"last_name",
			"email",
			"picture",
			"promo_code",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = Profile(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i Profile) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalProfile(i), i.AdditionalProperties)
}

// Activities structure is generated from "#/definitions/Activities".
type Activities struct {
	Offset               int64                  `json:"offset,omitempty"`  // Position in pagination.
	Limit                int64                  `json:"limit,omitempty"`   // Number of items to retrieve (100 max).
	Count                int64                  `json:"count,omitempty"`   // Total number of items available.
	History              []Activity             `json:"history,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`                 // All unmatched properties
}

type marshalActivities Activities

// UnmarshalJSON decodes JSON.
func (i *Activities) UnmarshalJSON(data []byte) error {
	ii := marshalActivities(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"offset",
			"limit",
			"count",
			"history",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = Activities(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i Activities) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalActivities(i), i.AdditionalProperties)
}

// Activity structure is generated from "#/definitions/Activity".
type Activity struct {
	Uuid                 string                 `json:"uuid,omitempty"` // Unique identifier for the activity
	AdditionalProperties map[string]interface{} `json:"-"`              // All unmatched properties
}

type marshalActivity Activity

// UnmarshalJSON decodes JSON.
func (i *Activity) UnmarshalJSON(data []byte) error {
	ii := marshalActivity(*i)

	err := unionMap{
		mustUnmarshal: []interface{}{&ii},
		ignoreKeys: []string{
			"uuid",
		},
		additionalProperties: &ii.AdditionalProperties,
		jsonData: data,
	}.unmarshal()
	if err != nil {
		return err
	}
	*i = Activity(ii)
	return err
}

// MarshalJSON encodes JSON.
func (i Activity) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalActivity(i), i.AdditionalProperties)
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
