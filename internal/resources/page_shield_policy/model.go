// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageShieldPolicyModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id"`
	Action      types.String `tfsdk:"action" json:"action"`
	Description types.String `tfsdk:"description" json:"description"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled"`
	Expression  types.String `tfsdk:"expression" json:"expression"`
	Value       types.String `tfsdk:"value" json:"value"`
}