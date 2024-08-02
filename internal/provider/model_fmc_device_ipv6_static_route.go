// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceIPv6StaticRoute struct {
	Id                   types.String                               `tfsdk:"id"`
	Domain               types.String                               `tfsdk:"domain"`
	DeviceId             types.String                               `tfsdk:"device_id"`
	InterfaceLogicalName types.String                               `tfsdk:"interface_logical_name"`
	InterfaceId          types.String                               `tfsdk:"interface_id"`
	DestinationNetworks  []DeviceIPv6StaticRouteDestinationNetworks `tfsdk:"destination_networks"`
	MetricValue          types.Int64                                `tfsdk:"metric_value"`
	GatewayObjectId      types.String                               `tfsdk:"gateway_object_id"`
	GatewayLiteral       types.String                               `tfsdk:"gateway_literal"`
	IsTunneled           types.Bool                                 `tfsdk:"is_tunneled"`
}

type DeviceIPv6StaticRouteDestinationNetworks struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceIPv6StaticRoute) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/ipv6staticroutes", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceIPv6StaticRoute) toBody(ctx context.Context, state DeviceIPv6StaticRoute) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.InterfaceLogicalName.IsNull() {
		body, _ = sjson.Set(body, "interfaceName", data.InterfaceLogicalName.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "links.parent", data.InterfaceId.ValueString())
	}
	if len(data.DestinationNetworks) > 0 {
		body, _ = sjson.Set(body, "selectedNetworks", []interface{}{})
		for _, item := range data.DestinationNetworks {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "selectedNetworks.-1", itemBody)
		}
	}
	if !data.MetricValue.IsNull() {
		body, _ = sjson.Set(body, "metricValue", data.MetricValue.ValueInt64())
	}
	if !data.GatewayObjectId.IsNull() {
		body, _ = sjson.Set(body, "gateway.object.id", data.GatewayObjectId.ValueString())
	}
	if !data.GatewayLiteral.IsNull() {
		body, _ = sjson.Set(body, "gateway.literal.value", data.GatewayLiteral.ValueString())
	}
	if !data.IsTunneled.IsNull() {
		body, _ = sjson.Set(body, "isTunneled", data.IsTunneled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceIPv6StaticRoute) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("interfaceName"); value.Exists() {
		data.InterfaceLogicalName = types.StringValue(value.String())
	} else {
		data.InterfaceLogicalName = types.StringNull()
	}
	if value := res.Get("selectedNetworks"); value.Exists() {
		data.DestinationNetworks = make([]DeviceIPv6StaticRouteDestinationNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DeviceIPv6StaticRouteDestinationNetworks{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.DestinationNetworks = append(data.DestinationNetworks, item)
			return true
		})
	}
	if value := res.Get("metricValue"); value.Exists() {
		data.MetricValue = types.Int64Value(value.Int())
	} else {
		data.MetricValue = types.Int64Null()
	}
	if value := res.Get("gateway.object.id"); value.Exists() {
		data.GatewayObjectId = types.StringValue(value.String())
	} else {
		data.GatewayObjectId = types.StringNull()
	}
	if value := res.Get("gateway.literal.value"); value.Exists() {
		data.GatewayLiteral = types.StringValue(value.String())
	} else {
		data.GatewayLiteral = types.StringNull()
	}
	if value := res.Get("isTunneled"); value.Exists() {
		data.IsTunneled = types.BoolValue(value.Bool())
	} else {
		data.IsTunneled = types.BoolValue(false)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceIPv6StaticRoute) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("interfaceName"); value.Exists() && !data.InterfaceLogicalName.IsNull() {
		data.InterfaceLogicalName = types.StringValue(value.String())
	} else {
		data.InterfaceLogicalName = types.StringNull()
	}
	for i := 0; i < len(data.DestinationNetworks); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationNetworks[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationNetworks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("selectedNetworks").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworks[%d] = %+v",
				i,
				(*parent).DestinationNetworks[i],
			))
			(*parent).DestinationNetworks = slices.Delete((*parent).DestinationNetworks, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationNetworks[i] = data
	}
	if value := res.Get("metricValue"); value.Exists() && !data.MetricValue.IsNull() {
		data.MetricValue = types.Int64Value(value.Int())
	} else {
		data.MetricValue = types.Int64Null()
	}
	if value := res.Get("gateway.object.id"); value.Exists() && !data.GatewayObjectId.IsNull() {
		data.GatewayObjectId = types.StringValue(value.String())
	} else {
		data.GatewayObjectId = types.StringNull()
	}
	if value := res.Get("gateway.literal.value"); value.Exists() && !data.GatewayLiteral.IsNull() {
		data.GatewayLiteral = types.StringValue(value.String())
	} else {
		data.GatewayLiteral = types.StringNull()
	}
	if value := res.Get("isTunneled"); value.Exists() && !data.IsTunneled.IsNull() {
		data.IsTunneled = types.BoolValue(value.Bool())
	} else if data.IsTunneled.ValueBool() != false {
		data.IsTunneled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceIPv6StaticRoute) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
