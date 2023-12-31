// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type PipelineFunctionConf struct {
	Conf        PipelineFunctionConfFunctionSpecificConfigs `tfsdk:"conf"`
	Description types.String                                `tfsdk:"description"`
	Disabled    types.Bool                                  `tfsdk:"disabled"`
	Filter      types.String                                `tfsdk:"filter"`
	Final       types.Bool                                  `tfsdk:"final"`
	GroupID     types.String                                `tfsdk:"group_id"`
	ID          types.String                                `tfsdk:"id"`
}
