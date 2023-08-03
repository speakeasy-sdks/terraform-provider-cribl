// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"Cribl/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *PolicyRuleResourceModel) ToCreateSDKType() *shared.PolicyRule {
	var args []string = nil
	for _, argsItem := range r.Args {
		args = append(args, argsItem.ValueString())
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	id := r.ID.ValueString()
	var template []string = nil
	for _, templateItem := range r.Template {
		template = append(template, templateItem.ValueString())
	}
	out := shared.PolicyRule{
		Args:        args,
		Description: description,
		ID:          id,
		Template:    template,
	}
	return &out
}

func (r *PolicyRuleResourceModel) ToGetSDKType() *shared.PolicyRule {
	out := r.ToCreateSDKType()
	return out
}

func (r *PolicyRuleResourceModel) ToUpdateSDKType() *shared.PolicyRule {
	out := r.ToCreateSDKType()
	return out
}

func (r *PolicyRuleResourceModel) ToDeleteSDKType() *shared.PolicyRule {
	out := r.ToCreateSDKType()
	return out
}

func (r *PolicyRuleResourceModel) RefreshFromGetResponse(resp *shared.Error) {
	if resp.Message != nil {
		r.Message = types.StringValue(*resp.Message)
	} else {
		r.Message = types.StringNull()
	}
}

func (r *PolicyRuleResourceModel) RefreshFromCreateResponse(resp *shared.Error) {
	r.RefreshFromGetResponse(resp)
}

func (r *PolicyRuleResourceModel) RefreshFromUpdateResponse(resp *shared.Error) {
	r.RefreshFromGetResponse(resp)
}
