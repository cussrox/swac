// Code generated by github.com/swaggest/swac v0.1.0, DO NOT EDIT.

package uber

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// GetProductsRequest is operation request value.
type GetProductsRequest struct {
	// Latitude is a required `latitude` parameter in query.
	// Latitude component of location.
	Latitude  float64
	// Longitude is a required `longitude` parameter in query.
	// Longitude component of location.
	Longitude float64
}

// encode creates *http.Request for request data.
func (request *GetProductsRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestUri := baseURL + "/products"
	query := make(url.Values, 2)
	query.Set("latitude", strconv.FormatFloat(request.Latitude, 'g', -1, 64))
	query.Set("longitude", strconv.FormatFloat(request.Longitude, 'g', -1, 64))
	if len(query) > 0 {
		requestUri += "?" + query.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, requestUri, nil)
	if err != nil {
	    return nil, err
	}
	req = req.WithContext(ctx)
	return req, err
}

// GetProductsResponse is operation response value.
type GetProductsResponse struct {
	StatusCode int
	OK         []Product  // OK is a value of 200 OK response.
	Default    *Error     // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetProductsResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.OK)
	    if err != nil {
	        return err
	    }
	default:
	    err := json.NewDecoder(resp.Body).Decode(&result.Default)
	    if err != nil {
	        return err
	    }
	}
	return nil

}

// GetProducts performs REST operation.
func (c *Client) GetProducts(ctx context.Context, request GetProductsRequest) (GetProductsResponse, error) {
	result := GetProductsResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/products")
	if c.Timeout != 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, c.Timeout)
		defer cancel()
	}
	req, err := request.encode(ctx, c.BaseURL)
	if err != nil {
	    return result, err
	}
	resp, err := c.transport.RoundTrip(req)

	if err != nil {
	    return result, err
	}
	defer resp.Body.Close()
	err = result.decode(resp)
	if err != nil {
	    return result, err
	}
	return result, nil
}
