// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ScriptResourceModel) ToCreateSDKType() *shared.ScriptLibEntry {
	var args []string = nil
	for _, argsItem := range r.Args {
		args = append(args, argsItem.ValueString())
	}
	command := r.Command.ValueString()
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	env := make(map[string]string)
	// Warning. This is a map, but the source tf var is not a map. This might indicate a bug.
	id := r.ID.ValueString()
	var additionalProperties interface{}
	if !r.AdditionalProperties.IsUnknown() && !r.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(r.AdditionalProperties.ValueString()), &additionalProperties)
	}
	out := shared.ScriptLibEntry{
		Args:                 args,
		Command:              command,
		Description:          description,
		Env:                  env,
		ID:                   id,
		AdditionalProperties: additionalProperties,
	}
	return &out
}

func (r *ScriptResourceModel) ToGetSDKType() *shared.ScriptLibEntry {
	out := r.ToCreateSDKType()
	return out
}

func (r *ScriptResourceModel) ToUpdateSDKType() *shared.ScriptLibEntry {
	out := r.ToCreateSDKType()
	return out
}

func (r *ScriptResourceModel) ToDeleteSDKType() *shared.ScriptLibEntry {
	out := r.ToCreateSDKType()
	return out
}

func (r *ScriptResourceModel) RefreshFromGetResponse(resp *shared.ScriptLibEntry) {
	r.Args = nil
	for _, v := range resp.Args {
		r.Args = append(r.Args, types.StringValue(v))
	}
	r.Command = types.StringValue(resp.Command)
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	if r.Env == nil && len(resp.Env) > 0 {
		r.Env = make(map[string]types.String)
		for key, value := range resp.Env {
			r.Env[key] = types.StringValue(value)
		}
	}
	r.ID = types.StringValue(resp.ID)
	if r.AdditionalProperties.IsUnknown() {
		if resp.AdditionalProperties == nil {
			r.AdditionalProperties = types.StringNull()
		} else {
			additionalPropertiesResult, _ := json.Marshal(resp.AdditionalProperties)
			r.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
		}
	}
}

func (r *ScriptResourceModel) RefreshFromCreateResponse(resp *shared.ScriptLibEntry) {
	r.RefreshFromGetResponse(resp)
}

func (r *ScriptResourceModel) RefreshFromUpdateResponse(resp *shared.ScriptLibEntry) {
	r.RefreshFromGetResponse(resp)
}
