// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

type RoutesRouteInput struct {
	// Short description of this Route
	Description *string `json:"description,omitempty"`
	// Whether this routing rule is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// If toggled to Yes, you can use a JavaScript expression that evaluates to the name of the Output below
	EnableOutputExpression *bool `json:"enableOutputExpression,omitempty"`
	// Expression (JS) to select data to route
	Filter *string `json:"filter,omitempty"`
	// Flag to control whether the event gets consumed by this Route (Final), or cloned into it
	Final            *bool       `json:"final,omitempty"`
	Name             string      `json:"name"`
	Output           interface{} `json:"output,omitempty"`
	OutputExpression interface{} `json:"outputExpression,omitempty"`
	// Pipeline to send the matching data to
	Pipeline string `json:"pipeline"`

	AdditionalProperties interface{} `json:"-"`
}
type _RoutesRouteInput RoutesRouteInput

func (c *RoutesRouteInput) UnmarshalJSON(bs []byte) error {
	data := _RoutesRouteInput{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RoutesRouteInput(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "description")
	delete(additionalFields, "disabled")
	delete(additionalFields, "enableOutputExpression")
	delete(additionalFields, "filter")
	delete(additionalFields, "final")
	delete(additionalFields, "name")
	delete(additionalFields, "output")
	delete(additionalFields, "outputExpression")
	delete(additionalFields, "pipeline")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RoutesRouteInput) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RoutesRouteInput(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type RoutesRoute struct {
	// Short description of this Route
	Description *string `json:"description,omitempty"`
	// Whether this routing rule is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// If toggled to Yes, you can use a JavaScript expression that evaluates to the name of the Output below
	EnableOutputExpression *bool `json:"enableOutputExpression,omitempty"`
	// Expression (JS) to select data to route
	Filter *string `json:"filter,omitempty"`
	// Flag to control whether the event gets consumed by this Route (Final), or cloned into it
	Final            *bool       `json:"final,omitempty"`
	ID               *string     `json:"id,omitempty"`
	Name             string      `json:"name"`
	Output           interface{} `json:"output,omitempty"`
	OutputExpression interface{} `json:"outputExpression,omitempty"`
	// Pipeline to send the matching data to
	Pipeline string `json:"pipeline"`

	AdditionalProperties interface{} `json:"-"`
}
type _RoutesRoute RoutesRoute

func (c *RoutesRoute) UnmarshalJSON(bs []byte) error {
	data := _RoutesRoute{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RoutesRoute(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "description")
	delete(additionalFields, "disabled")
	delete(additionalFields, "enableOutputExpression")
	delete(additionalFields, "filter")
	delete(additionalFields, "final")
	delete(additionalFields, "id")
	delete(additionalFields, "name")
	delete(additionalFields, "output")
	delete(additionalFields, "outputExpression")
	delete(additionalFields, "pipeline")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RoutesRoute) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RoutesRoute(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}
