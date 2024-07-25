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
	"math"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-fmc/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AccessControlPolicy struct {
	Id                             types.String                    `tfsdk:"id"`
	Domain                         types.String                    `tfsdk:"domain"`
	Name                           types.String                    `tfsdk:"name"`
	Description                    types.String                    `tfsdk:"description"`
	DefaultAction                  types.String                    `tfsdk:"default_action"`
	DefaultActionId                types.String                    `tfsdk:"default_action_id"`
	DefaultActionLogBegin          types.Bool                      `tfsdk:"default_action_log_begin"`
	DefaultActionLogEnd            types.Bool                      `tfsdk:"default_action_log_end"`
	DefaultActionSendEventsToFmc   types.Bool                      `tfsdk:"default_action_send_events_to_fmc"`
	DefaultActionSendSyslog        types.Bool                      `tfsdk:"default_action_send_syslog"`
	DefaultActionSyslogConfigId    types.String                    `tfsdk:"default_action_syslog_config_id"`
	DefaultActionSyslogSeverity    types.String                    `tfsdk:"default_action_syslog_severity"`
	DefaultActionSnmpConfigId      types.String                    `tfsdk:"default_action_snmp_config_id"`
	DefaultActionIntrusionPolicyId types.String                    `tfsdk:"default_action_intrusion_policy_id"`
	Categories                     []AccessControlPolicyCategories `tfsdk:"categories"`
	Rules                          []AccessControlPolicyRules      `tfsdk:"rules"`
}

type AccessControlPolicyCategories struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Section types.String `tfsdk:"section"`
}

type AccessControlPolicyRules struct {
	Id                                 types.String                                                 `tfsdk:"id"`
	Action                             types.String                                                 `tfsdk:"action"`
	Name                               types.String                                                 `tfsdk:"name"`
	CategoryName                       types.String                                                 `tfsdk:"category_name"`
	Section                            types.String                                                 `tfsdk:"section"`
	Enabled                            types.Bool                                                   `tfsdk:"enabled"`
	SourceNetworkLiterals              []AccessControlPolicyRulesSourceNetworkLiterals              `tfsdk:"source_network_literals"`
	DestinationNetworkLiterals         []AccessControlPolicyRulesDestinationNetworkLiterals         `tfsdk:"destination_network_literals"`
	SourceNetworkObjects               []AccessControlPolicyRulesSourceNetworkObjects               `tfsdk:"source_network_objects"`
	DestinationNetworkObjects          []AccessControlPolicyRulesDestinationNetworkObjects          `tfsdk:"destination_network_objects"`
	SourceDynamicObjects               []AccessControlPolicyRulesSourceDynamicObjects               `tfsdk:"source_dynamic_objects"`
	DestinationDynamicObjects          []AccessControlPolicyRulesDestinationDynamicObjects          `tfsdk:"destination_dynamic_objects"`
	SourcePortObjects                  []AccessControlPolicyRulesSourcePortObjects                  `tfsdk:"source_port_objects"`
	DestinationPortObjects             []AccessControlPolicyRulesDestinationPortObjects             `tfsdk:"destination_port_objects"`
	SourceSecurityGroupTagObjects      []AccessControlPolicyRulesSourceSecurityGroupTagObjects      `tfsdk:"source_security_group_tag_objects"`
	DestinationSecurityGroupTagObjects []AccessControlPolicyRulesDestinationSecurityGroupTagObjects `tfsdk:"destination_security_group_tag_objects"`
	SourceZones                        []AccessControlPolicyRulesSourceZones                        `tfsdk:"source_zones"`
	DestinationZones                   []AccessControlPolicyRulesDestinationZones                   `tfsdk:"destination_zones"`
	UrlObjects                         []AccessControlPolicyRulesUrlObjects                         `tfsdk:"url_objects"`
	UrlCategories                      []AccessControlPolicyRulesUrlCategories                      `tfsdk:"url_categories"`
	LogBegin                           types.Bool                                                   `tfsdk:"log_begin"`
	LogEnd                             types.Bool                                                   `tfsdk:"log_end"`
	LogFiles                           types.Bool                                                   `tfsdk:"log_files"`
	SendEventsToFmc                    types.Bool                                                   `tfsdk:"send_events_to_fmc"`
	SendSyslog                         types.Bool                                                   `tfsdk:"send_syslog"`
	SyslogConfigId                     types.String                                                 `tfsdk:"syslog_config_id"`
	SyslogSeverity                     types.String                                                 `tfsdk:"syslog_severity"`
	SnmpConfigId                       types.String                                                 `tfsdk:"snmp_config_id"`
	Description                        types.String                                                 `tfsdk:"description"`
	FilePolicyId                       types.String                                                 `tfsdk:"file_policy_id"`
	IntrusionPolicyId                  types.String                                                 `tfsdk:"intrusion_policy_id"`
}

type AccessControlPolicyRulesSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type AccessControlPolicyRulesDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
}
type AccessControlPolicyRulesSourceNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessControlPolicyRulesDestinationNetworkObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessControlPolicyRulesSourceDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesDestinationDynamicObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesSourceSecurityGroupTagObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessControlPolicyRulesDestinationSecurityGroupTagObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type AccessControlPolicyRulesSourceZones struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesDestinationZones struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesUrlObjects struct {
	Id types.String `tfsdk:"id"`
}
type AccessControlPolicyRulesUrlCategories struct {
	Id         types.String `tfsdk:"id"`
	Reputation types.String `tfsdk:"reputation"`
}

// End of section. //template:end types

func (c AccessControlPolicyCategories) GetSection() string {
	if s := c.Section.ValueString(); s != "" {
		return s
	}

	if c.Section.IsUnknown() {
		panic(c.Section)
	}

	return "default"
}

func (r AccessControlPolicyRules) GetSection() string {
	if s := r.Section.ValueString(); s != "" {
		return s
	}

	if r.CategoryName.ValueString() != "" {
		return ""
	}

	if r.Section.IsUnknown() {
		panic(r.Section)
	}

	return "default"
}

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AccessControlPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AccessControlPolicy) toBody(ctx context.Context, state AccessControlPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.action", data.DefaultAction.ValueString())
	}
	if state.DefaultActionId.ValueString() != "" {
		body, _ = sjson.Set(body, "defaultAction.id", state.DefaultActionId.ValueString())
	}
	if !data.DefaultActionLogBegin.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logBegin", data.DefaultActionLogBegin.ValueBool())
	}
	if !data.DefaultActionLogEnd.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.logEnd", data.DefaultActionLogEnd.ValueBool())
	}
	if !data.DefaultActionSendEventsToFmc.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.sendEventsToFMC", data.DefaultActionSendEventsToFmc.ValueBool())
	}
	if !data.DefaultActionSendSyslog.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.enableSyslog", data.DefaultActionSendSyslog.ValueBool())
	}
	if !data.DefaultActionSyslogConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.syslogConfig.id", data.DefaultActionSyslogConfigId.ValueString())
	}
	if !data.DefaultActionSyslogSeverity.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.syslogSeverity", data.DefaultActionSyslogSeverity.ValueString())
	}
	if !data.DefaultActionSnmpConfigId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.snmpConfig.id", data.DefaultActionSnmpConfigId.ValueString())
	}
	if !data.DefaultActionIntrusionPolicyId.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.intrusionPolicy.id", data.DefaultActionIntrusionPolicyId.ValueString())
	}
	if len(data.Categories) > 0 {
		body, _ = sjson.Set(body, "dummy_categories", []interface{}{})
		for _, item := range data.Categories {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Section.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metadata.section", item.Section.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_categories.-1", itemBody)
		}
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "dummy_rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.CategoryName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metadata.category", item.CategoryName.ValueString())
			}
			if !item.Section.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metadata.section", item.Section.ValueString())
			}
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []interface{}{})
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.literals", []interface{}{})
				for _, childItem := range item.DestinationNetworkLiterals {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", []interface{}{})
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.objects", []interface{}{})
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.SourceDynamicObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceDynamicObjects.objects", []interface{}{})
				for _, childItem := range item.SourceDynamicObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "DynamicObject")
					itemBody, _ = sjson.SetRaw(itemBody, "sourceDynamicObjects.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationDynamicObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationDynamicObjects.objects", []interface{}{})
				for _, childItem := range item.DestinationDynamicObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "DynamicObject")
					itemBody, _ = sjson.SetRaw(itemBody, "destinationDynamicObjects.objects.-1", itemChildBody)
				}
			}
			if len(item.SourcePortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.objects", []interface{}{})
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.objects", []interface{}{})
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects.-1", itemChildBody)
				}
			}
			if len(item.SourceSecurityGroupTagObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceSecurityGroupTags.objects", []interface{}{})
				for _, childItem := range item.SourceSecurityGroupTagObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceSecurityGroupTags.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationSecurityGroupTagObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationSecurityGroupTags.objects", []interface{}{})
				for _, childItem := range item.DestinationSecurityGroupTagObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationSecurityGroupTags.objects.-1", itemChildBody)
				}
			}
			if len(item.SourceZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceZones.objects", []interface{}{})
				for _, childItem := range item.SourceZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "SecurityZone")
					itemBody, _ = sjson.SetRaw(itemBody, "sourceZones.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationZones.objects", []interface{}{})
				for _, childItem := range item.DestinationZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "SecurityZone")
					itemBody, _ = sjson.SetRaw(itemBody, "destinationZones.objects.-1", itemChildBody)
				}
			}
			if len(item.UrlObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "urls.objects", []interface{}{})
				for _, childItem := range item.UrlObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "urls.objects.-1", itemChildBody)
				}
			}
			if len(item.UrlCategories) > 0 {
				itemBody, _ = sjson.Set(itemBody, "urls.urlCategoriesWithReputation", []interface{}{})
				for _, childItem := range item.UrlCategories {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "category.id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "category.type", "AnyNonEmptyString")
					if !childItem.Reputation.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "reputation", childItem.Reputation.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "urls.urlCategoriesWithReputation.-1", itemChildBody)
				}
			}
			if !item.LogBegin.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logBegin", item.LogBegin.ValueBool())
			}
			if !item.LogEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logEnd", item.LogEnd.ValueBool())
			}
			if !item.LogFiles.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logFiles", item.LogFiles.ValueBool())
			}
			if !item.SendEventsToFmc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sendEventsToFMC", item.SendEventsToFmc.ValueBool())
			}
			if !item.SendSyslog.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableSyslog", item.SendSyslog.ValueBool())
			}
			if !item.SyslogConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogConfig.id", item.SyslogConfigId.ValueString())
			}
			if !item.SyslogSeverity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "syslogSeverity", item.SyslogSeverity.ValueString())
			}
			if !item.SnmpConfigId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "snmpConfig.id", item.SnmpConfigId.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.FilePolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filePolicy.id", item.FilePolicyId.ValueString())
			}
			if !item.IntrusionPolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipsPolicy.id", item.IntrusionPolicyId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AccessControlPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("defaultAction.action"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogBegin = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionLogEnd = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendEventsToFmc = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.enableSyslog"); value.Exists() {
		data.DefaultActionSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.DefaultActionSendSyslog = types.BoolValue(false)
	}
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.syslogSeverity"); value.Exists() {
		data.DefaultActionSyslogSeverity = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogSeverity = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.intrusionPolicy.id"); value.Exists() {
		data.DefaultActionIntrusionPolicyId = types.StringValue(value.String())
	} else {
		data.DefaultActionIntrusionPolicyId = types.StringNull()
	}
	if value := res.Get("dummy_categories"); value.Exists() {
		data.Categories = make([]AccessControlPolicyCategories, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessControlPolicyCategories{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.Categories = append(data.Categories, item)
			return true
		})
	}
	if value := res.Get("dummy_rules"); value.Exists() {
		data.Rules = make([]AccessControlPolicyRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AccessControlPolicyRules{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			} else {
				item.Action = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("metadata.category"); cValue.Exists() {
				item.CategoryName = types.StringValue(cValue.String())
			} else {
				item.CategoryName = types.StringNull()
			}
			if cValue := v.Get("metadata.section"); cValue.Exists() {
				item.Section = types.StringValue(cValue.String())
			} else {
				item.Section = types.StringNull()
			}
			if cValue := v.Get("enabled"); cValue.Exists() {
				item.Enabled = types.BoolValue(cValue.Bool())
			} else {
				item.Enabled = types.BoolValue(true)
			}
			if cValue := v.Get("sourceNetworks.literals"); cValue.Exists() {
				item.SourceNetworkLiterals = make([]AccessControlPolicyRulesSourceNetworkLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceNetworkLiterals{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.SourceNetworkLiterals = append(item.SourceNetworkLiterals, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationNetworks.literals"); cValue.Exists() {
				item.DestinationNetworkLiterals = make([]AccessControlPolicyRulesDestinationNetworkLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationNetworkLiterals{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.DestinationNetworkLiterals = append(item.DestinationNetworkLiterals, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceNetworks.objects"); cValue.Exists() {
				item.SourceNetworkObjects = make([]AccessControlPolicyRulesSourceNetworkObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceNetworkObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.SourceNetworkObjects = append(item.SourceNetworkObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationNetworks.objects"); cValue.Exists() {
				item.DestinationNetworkObjects = make([]AccessControlPolicyRulesDestinationNetworkObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationNetworkObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.DestinationNetworkObjects = append(item.DestinationNetworkObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceDynamicObjects.objects"); cValue.Exists() {
				item.SourceDynamicObjects = make([]AccessControlPolicyRulesSourceDynamicObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceDynamicObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.SourceDynamicObjects = append(item.SourceDynamicObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationDynamicObjects.objects"); cValue.Exists() {
				item.DestinationDynamicObjects = make([]AccessControlPolicyRulesDestinationDynamicObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationDynamicObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.DestinationDynamicObjects = append(item.DestinationDynamicObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("sourcePorts.objects"); cValue.Exists() {
				item.SourcePortObjects = make([]AccessControlPolicyRulesSourcePortObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourcePortObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.SourcePortObjects = append(item.SourcePortObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationPorts.objects"); cValue.Exists() {
				item.DestinationPortObjects = make([]AccessControlPolicyRulesDestinationPortObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationPortObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.DestinationPortObjects = append(item.DestinationPortObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceSecurityGroupTags.objects"); cValue.Exists() {
				item.SourceSecurityGroupTagObjects = make([]AccessControlPolicyRulesSourceSecurityGroupTagObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceSecurityGroupTagObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.SourceSecurityGroupTagObjects = append(item.SourceSecurityGroupTagObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationSecurityGroupTags.objects"); cValue.Exists() {
				item.DestinationSecurityGroupTagObjects = make([]AccessControlPolicyRulesDestinationSecurityGroupTagObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationSecurityGroupTagObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.DestinationSecurityGroupTagObjects = append(item.DestinationSecurityGroupTagObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceZones.objects"); cValue.Exists() {
				item.SourceZones = make([]AccessControlPolicyRulesSourceZones, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesSourceZones{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.SourceZones = append(item.SourceZones, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationZones.objects"); cValue.Exists() {
				item.DestinationZones = make([]AccessControlPolicyRulesDestinationZones, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesDestinationZones{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.DestinationZones = append(item.DestinationZones, cItem)
					return true
				})
			}
			if cValue := v.Get("urls.objects"); cValue.Exists() {
				item.UrlObjects = make([]AccessControlPolicyRulesUrlObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesUrlObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.UrlObjects = append(item.UrlObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("urls.urlCategoriesWithReputation"); cValue.Exists() {
				item.UrlCategories = make([]AccessControlPolicyRulesUrlCategories, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := AccessControlPolicyRulesUrlCategories{}
					if ccValue := cv.Get("category.id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("reputation"); ccValue.Exists() {
						cItem.Reputation = types.StringValue(ccValue.String())
					} else {
						cItem.Reputation = types.StringNull()
					}
					item.UrlCategories = append(item.UrlCategories, cItem)
					return true
				})
			}
			if cValue := v.Get("logBegin"); cValue.Exists() {
				item.LogBegin = types.BoolValue(cValue.Bool())
			} else {
				item.LogBegin = types.BoolValue(false)
			}
			if cValue := v.Get("logEnd"); cValue.Exists() {
				item.LogEnd = types.BoolValue(cValue.Bool())
			} else {
				item.LogEnd = types.BoolValue(false)
			}
			if cValue := v.Get("logFiles"); cValue.Exists() {
				item.LogFiles = types.BoolValue(cValue.Bool())
			} else {
				item.LogFiles = types.BoolValue(false)
			}
			if cValue := v.Get("sendEventsToFMC"); cValue.Exists() {
				item.SendEventsToFmc = types.BoolValue(cValue.Bool())
			} else {
				item.SendEventsToFmc = types.BoolValue(false)
			}
			if cValue := v.Get("enableSyslog"); cValue.Exists() {
				item.SendSyslog = types.BoolValue(cValue.Bool())
			} else {
				item.SendSyslog = types.BoolValue(false)
			}
			if cValue := v.Get("syslogConfig.id"); cValue.Exists() {
				item.SyslogConfigId = types.StringValue(cValue.String())
			} else {
				item.SyslogConfigId = types.StringNull()
			}
			if cValue := v.Get("syslogSeverity"); cValue.Exists() {
				item.SyslogSeverity = types.StringValue(cValue.String())
			} else {
				item.SyslogSeverity = types.StringNull()
			}
			if cValue := v.Get("snmpConfig.id"); cValue.Exists() {
				item.SnmpConfigId = types.StringValue(cValue.String())
			} else {
				item.SnmpConfigId = types.StringNull()
			}
			if cValue := v.Get("filePolicy.id"); cValue.Exists() {
				item.FilePolicyId = types.StringValue(cValue.String())
			} else {
				item.FilePolicyId = types.StringNull()
			}
			if cValue := v.Get("ipsPolicy.id"); cValue.Exists() {
				item.IntrusionPolicyId = types.StringValue(cValue.String())
			} else {
				item.IntrusionPolicyId = types.StringNull()
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AccessControlPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("defaultAction.action"); value.Exists() && !data.DefaultAction.IsNull() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("defaultAction.id"); value.Exists() {
		data.DefaultActionId = types.StringValue(value.String())
	} else {
		data.DefaultActionId = types.StringNull()
	}
	if value := res.Get("defaultAction.logBegin"); value.Exists() && !data.DefaultActionLogBegin.IsNull() {
		data.DefaultActionLogBegin = types.BoolValue(value.Bool())
	} else if data.DefaultActionLogBegin.ValueBool() != false {
		data.DefaultActionLogBegin = types.BoolNull()
	}
	if value := res.Get("defaultAction.logEnd"); value.Exists() && !data.DefaultActionLogEnd.IsNull() {
		data.DefaultActionLogEnd = types.BoolValue(value.Bool())
	} else if data.DefaultActionLogEnd.ValueBool() != false {
		data.DefaultActionLogEnd = types.BoolNull()
	}
	if value := res.Get("defaultAction.sendEventsToFMC"); value.Exists() && !data.DefaultActionSendEventsToFmc.IsNull() {
		data.DefaultActionSendEventsToFmc = types.BoolValue(value.Bool())
	} else if data.DefaultActionSendEventsToFmc.ValueBool() != false {
		data.DefaultActionSendEventsToFmc = types.BoolNull()
	}
	if value := res.Get("defaultAction.enableSyslog"); value.Exists() && !data.DefaultActionSendSyslog.IsNull() {
		data.DefaultActionSendSyslog = types.BoolValue(value.Bool())
	} else if data.DefaultActionSendSyslog.ValueBool() != false {
		data.DefaultActionSendSyslog = types.BoolNull()
	}
	if value := res.Get("defaultAction.syslogConfig.id"); value.Exists() && !data.DefaultActionSyslogConfigId.IsNull() {
		data.DefaultActionSyslogConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.syslogSeverity"); value.Exists() && !data.DefaultActionSyslogSeverity.IsNull() {
		data.DefaultActionSyslogSeverity = types.StringValue(value.String())
	} else {
		data.DefaultActionSyslogSeverity = types.StringNull()
	}
	if value := res.Get("defaultAction.snmpConfig.id"); value.Exists() && !data.DefaultActionSnmpConfigId.IsNull() {
		data.DefaultActionSnmpConfigId = types.StringValue(value.String())
	} else {
		data.DefaultActionSnmpConfigId = types.StringNull()
	}
	if value := res.Get("defaultAction.intrusionPolicy.id"); value.Exists() && !data.DefaultActionIntrusionPolicyId.IsNull() {
		data.DefaultActionIntrusionPolicyId = types.StringValue(value.String())
	} else {
		data.DefaultActionIntrusionPolicyId = types.StringNull()
	}
	{
		l := len(res.Get("dummy_categories").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_categories array resizing from %d to %d", len(data.Categories), l))
		for i := len(data.Categories); i < l; i++ {
			data.Categories = append(data.Categories, AccessControlPolicyCategories{})
		}
		if len(data.Categories) > l {
			data.Categories = data.Categories[:l]
		}
	}
	for i := range data.Categories {
		parent := data
		data := parent.Categories[i]
		parentRes := res
		res := parentRes.Get(fmt.Sprintf("dummy_categories.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		parent.Categories[i] = data
	}
	{
		l := len(res.Get("dummy_rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("dummy_rules array resizing from %d to %d", len(data.Rules), l))
		for i := len(data.Rules); i < l; i++ {
			data.Rules = append(data.Rules, AccessControlPolicyRules{})
		}
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := data
		data := parent.Rules[i]
		parentRes := res
		res := parentRes.Get(fmt.Sprintf("dummy_rules.%d", i))
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("metadata.category"); value.Exists() && !data.CategoryName.IsNull() {
			data.CategoryName = types.StringValue(value.String())
		} else {
			data.CategoryName = types.StringNull()
		}
		if value := res.Get("metadata.section"); value.Exists() && !data.Section.IsNull() {
			data.Section = types.StringValue(value.String())
		} else {
			data.Section = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else if data.Enabled.ValueBool() != true {
			data.Enabled = types.BoolNull()
		}
		for i := 0; i < len(data.SourceNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.SourceNetworkLiterals[i].Value.ValueString()}

			parent := data
			data := parent.SourceNetworkLiterals[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourceNetworks.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkLiterals[%d] = %+v",
					i,
					parent.SourceNetworkLiterals[i],
				))
				parent.SourceNetworkLiterals = slices.Delete(parent.SourceNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			parent.SourceNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.DestinationNetworkLiterals[i].Value.ValueString()}

			parent := data
			data := parent.DestinationNetworkLiterals[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationNetworks.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkLiterals[%d] = %+v",
					i,
					parent.DestinationNetworkLiterals[i],
				))
				parent.DestinationNetworkLiterals = slices.Delete(parent.DestinationNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			parent.DestinationNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.SourceNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceNetworkObjects[i].Id.ValueString()}

			parent := data
			data := parent.SourceNetworkObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourceNetworks.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkObjects[%d] = %+v",
					i,
					parent.SourceNetworkObjects[i],
				))
				parent.SourceNetworkObjects = slices.Delete(parent.SourceNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			parent.SourceNetworkObjects[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationNetworkObjects[i].Id.ValueString()}

			parent := data
			data := parent.DestinationNetworkObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationNetworks.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkObjects[%d] = %+v",
					i,
					parent.DestinationNetworkObjects[i],
				))
				parent.DestinationNetworkObjects = slices.Delete(parent.DestinationNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			parent.DestinationNetworkObjects[i] = data
		}
		for i := 0; i < len(data.SourceDynamicObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceDynamicObjects[i].Id.ValueString()}

			parent := data
			data := parent.SourceDynamicObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourceDynamicObjects.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceDynamicObjects[%d] = %+v",
					i,
					parent.SourceDynamicObjects[i],
				))
				parent.SourceDynamicObjects = slices.Delete(parent.SourceDynamicObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.SourceDynamicObjects[i] = data
		}
		for i := 0; i < len(data.DestinationDynamicObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationDynamicObjects[i].Id.ValueString()}

			parent := data
			data := parent.DestinationDynamicObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationDynamicObjects.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationDynamicObjects[%d] = %+v",
					i,
					parent.DestinationDynamicObjects[i],
				))
				parent.DestinationDynamicObjects = slices.Delete(parent.DestinationDynamicObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.DestinationDynamicObjects[i] = data
		}
		for i := 0; i < len(data.SourcePortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourcePortObjects[i].Id.ValueString()}

			parent := data
			data := parent.SourcePortObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourcePorts.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourcePortObjects[%d] = %+v",
					i,
					parent.SourcePortObjects[i],
				))
				parent.SourcePortObjects = slices.Delete(parent.SourcePortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.SourcePortObjects[i] = data
		}
		for i := 0; i < len(data.DestinationPortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationPortObjects[i].Id.ValueString()}

			parent := data
			data := parent.DestinationPortObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationPorts.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortObjects[%d] = %+v",
					i,
					parent.DestinationPortObjects[i],
				))
				parent.DestinationPortObjects = slices.Delete(parent.DestinationPortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.DestinationPortObjects[i] = data
		}
		for i := 0; i < len(data.SourceSecurityGroupTagObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceSecurityGroupTagObjects[i].Id.ValueString()}

			parent := data
			data := parent.SourceSecurityGroupTagObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourceSecurityGroupTags.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceSecurityGroupTagObjects[%d] = %+v",
					i,
					parent.SourceSecurityGroupTagObjects[i],
				))
				parent.SourceSecurityGroupTagObjects = slices.Delete(parent.SourceSecurityGroupTagObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			parent.SourceSecurityGroupTagObjects[i] = data
		}
		for i := 0; i < len(data.DestinationSecurityGroupTagObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationSecurityGroupTagObjects[i].Id.ValueString()}

			parent := data
			data := parent.DestinationSecurityGroupTagObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationSecurityGroupTags.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationSecurityGroupTagObjects[%d] = %+v",
					i,
					parent.DestinationSecurityGroupTagObjects[i],
				))
				parent.DestinationSecurityGroupTagObjects = slices.Delete(parent.DestinationSecurityGroupTagObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			parent.DestinationSecurityGroupTagObjects[i] = data
		}
		for i := 0; i < len(data.SourceZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceZones[i].Id.ValueString()}

			parent := data
			data := parent.SourceZones[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("sourceZones.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceZones[%d] = %+v",
					i,
					parent.SourceZones[i],
				))
				parent.SourceZones = slices.Delete(parent.SourceZones, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.SourceZones[i] = data
		}
		for i := 0; i < len(data.DestinationZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationZones[i].Id.ValueString()}

			parent := data
			data := parent.DestinationZones[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("destinationZones.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationZones[%d] = %+v",
					i,
					parent.DestinationZones[i],
				))
				parent.DestinationZones = slices.Delete(parent.DestinationZones, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.DestinationZones[i] = data
		}
		for i := 0; i < len(data.UrlObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.UrlObjects[i].Id.ValueString()}

			parent := data
			data := parent.UrlObjects[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("urls.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing UrlObjects[%d] = %+v",
					i,
					parent.UrlObjects[i],
				))
				parent.UrlObjects = slices.Delete(parent.UrlObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			parent.UrlObjects[i] = data
		}
		for i := 0; i < len(data.UrlCategories); i++ {
			keys := [...]string{"category.id"}
			keyValues := [...]string{data.UrlCategories[i].Id.ValueString()}

			parent := data
			data := parent.UrlCategories[i]
			parentRes := res
			var res gjson.Result

			parentRes.Get("urls.urlCategoriesWithReputation").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing UrlCategories[%d] = %+v",
					i,
					parent.UrlCategories[i],
				))
				parent.UrlCategories = slices.Delete(parent.UrlCategories, i, i+1)
				i--

				continue
			}
			if value := res.Get("category.id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("reputation"); value.Exists() && !data.Reputation.IsNull() {
				data.Reputation = types.StringValue(value.String())
			} else {
				data.Reputation = types.StringNull()
			}
			parent.UrlCategories[i] = data
		}
		if value := res.Get("logBegin"); value.Exists() && !data.LogBegin.IsNull() {
			data.LogBegin = types.BoolValue(value.Bool())
		} else if data.LogBegin.ValueBool() != false {
			data.LogBegin = types.BoolNull()
		}
		if value := res.Get("logEnd"); value.Exists() && !data.LogEnd.IsNull() {
			data.LogEnd = types.BoolValue(value.Bool())
		} else if data.LogEnd.ValueBool() != false {
			data.LogEnd = types.BoolNull()
		}
		if value := res.Get("logFiles"); value.Exists() && !data.LogFiles.IsNull() {
			data.LogFiles = types.BoolValue(value.Bool())
		} else if data.LogFiles.ValueBool() != false {
			data.LogFiles = types.BoolNull()
		}
		if value := res.Get("sendEventsToFMC"); value.Exists() && !data.SendEventsToFmc.IsNull() {
			data.SendEventsToFmc = types.BoolValue(value.Bool())
		} else if data.SendEventsToFmc.ValueBool() != false {
			data.SendEventsToFmc = types.BoolNull()
		}
		if value := res.Get("enableSyslog"); value.Exists() && !data.SendSyslog.IsNull() {
			data.SendSyslog = types.BoolValue(value.Bool())
		} else if data.SendSyslog.ValueBool() != false {
			data.SendSyslog = types.BoolNull()
		}
		if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogConfigId.IsNull() {
			data.SyslogConfigId = types.StringValue(value.String())
		} else {
			data.SyslogConfigId = types.StringNull()
		}
		if value := res.Get("syslogSeverity"); value.Exists() && !data.SyslogSeverity.IsNull() {
			data.SyslogSeverity = types.StringValue(value.String())
		} else {
			data.SyslogSeverity = types.StringNull()
		}
		if value := res.Get("snmpConfig.id"); value.Exists() && !data.SnmpConfigId.IsNull() {
			data.SnmpConfigId = types.StringValue(value.String())
		} else {
			data.SnmpConfigId = types.StringNull()
		}
		if value := res.Get("filePolicy.id"); value.Exists() && !data.FilePolicyId.IsNull() {
			data.FilePolicyId = types.StringValue(value.String())
		} else {
			data.FilePolicyId = types.StringNull()
		}
		if value := res.Get("ipsPolicy.id"); value.Exists() && !data.IntrusionPolicyId.IsNull() {
			data.IntrusionPolicyId = types.StringValue(value.String())
		} else {
			data.IntrusionPolicyId = types.StringNull()
		}
		parent.Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// adjustFromBody applies a few corrections after an auto-generated fromBody/fromBodyPartial.
func (data *AccessControlPolicy) adjustFromBody(ctx context.Context, res gjson.Result) {
	for i := range data.Rules {
		if data.Rules[i].CategoryName.Equal(types.StringValue("--Undefined--")) {
			data.Rules[i].CategoryName = types.StringNull()
		}

		switch data.Rules[i].CategoryName.ValueString() {
		case "":
			// API has "Mandatory", OAS has "mandatory".
			data.Rules[i].Section = helpers.ToLower(data.Rules[i].Section)
		default:
			if !data.Rules[i].Section.IsUnknown() {
				data.Rules[i].Section = types.StringNull()
			}
		}
	}

	for i := range data.Categories {
		// API has "Mandatory", OAS has "mandatory".
		data.Categories[i].Section = helpers.ToLower(data.Categories[i].Section)
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AccessControlPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.DefaultActionId.IsUnknown() {
		if value := res.Get("defaultAction.id"); value.Exists() {
			data.DefaultActionId = types.StringValue(value.String())
		} else {
			data.DefaultActionId = types.StringNull()
		}
	}
	for i := range data.Categories {
		r := res.Get(fmt.Sprintf("dummy_categories.%d", i))
		if data.Categories[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				data.Categories[i].Id = types.StringValue(value.String())
			} else {
				data.Categories[i].Id = types.StringNull()
			}
		}
	}
	for i := range data.Rules {
		r := res.Get(fmt.Sprintf("dummy_rules.%d", i))
		if data.Rules[i].Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				data.Rules[i].Id = types.StringValue(value.String())
			} else {
				data.Rules[i].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// NewValidAccessControlPolicy validates the terraform Plan and converts it to a new AccessControlPolicy object.
// Does not tolerate unknown values (IsUnknown), primarily because tfplan.Get cannot unmarshal unknown lists to []T
// and `.rules` and `.categories` attributes have type []T.
func NewValidAccessControlPolicy(ctx context.Context, tfplan tfsdk.Plan) (AccessControlPolicy, diag.Diagnostics) {
	var plan AccessControlPolicy
	diags := tfplan.Get(ctx, &plan)
	if diags.HasError() {
		return plan, diags
	}

	// Validate lenghts. TODO(#66): make lenghts unlimited
	if len(plan.Categories) > 1000 {
		diags.AddAttributeError(path.Root("categories"), "Too many categories",
			fmt.Sprintf("%d categories specified, expected at most %d.", len(plan.Categories), 1000))
		return plan, diags
	}

	if len(plan.Rules) > 1000 {
		diags.AddAttributeError(path.Root("rules"), "Too many rules",
			fmt.Sprintf("%d rules specified, expected at most %d.", len(plan.Rules), 1000))
		return plan, diags
	}

	// Validate categories.*.section
	def := types.StringNull()
	insertion := len(plan.Categories)
	for i, cat := range plan.Categories {
		switch {
		case !def.IsNull() && cat.Section.Equal(types.StringValue("mandatory")):
			diags.AddAttributeError(path.Root("categories"), "Wrong order of categories",
				fmt.Sprintf("Category %s must be somewhere above category %s, not below it.\n"+
					"This is because the section=\"mandatory\" categories must precede all other categories.\n",
					cat.Name, def))
			return plan, diags
		case !def.IsNull():
			continue
		case !cat.Section.Equal(types.StringValue("mandatory")):
			def = cat.Name
			insertion = i
		}
	}

	// Validate rules.*.category_name clash with rules.*.section
	for _, rule := range plan.Rules {
		switch {
		case !rule.CategoryName.IsNull() && rule.GetSection() != "":
			diags.AddAttributeError(path.Root("rules"), "Cannot use section together with category_name",
				fmt.Sprintf("Rule %s cannot have both section and category_name specified.", rule.Name))
			return plan, diags
		case rule.CategoryName.Equal(types.StringValue("--Undefined--")):
			diags.AddAttributeError(path.Root("rules"), "Invalid category_name value",
				fmt.Sprintf("Cannot use category_name=%s that value is reserved for internal use.", rule.CategoryName))
			return plan, diags
		}
	}

	// Validate rules.*.action==MONITOR clash
	for _, rule := range plan.Rules {
		switch {
		case rule.Action.ValueString() == "MONITOR" && rule.LogBegin.ValueBool():
			diags.AddAttributeError(path.Root("rules"), "Cannot use log_begin=true when action=\"MONITOR\"",
				fmt.Sprintf("Rule %s has action=\"MONITOR\" so it must use:\n	log_begin=null/false,\n	log_end=true,\n	send_events_to_fmc=true,", rule.Name))
			return plan, diags
		case rule.Action.ValueString() == "MONITOR" && !rule.LogEnd.ValueBool():
			diags.AddAttributeError(path.Root("rules"), "Cannot use log_end=false when action=\"MONITOR\"",
				fmt.Sprintf("Rule %s has action=\"MONITOR\" so it must use:\n	log_begin=null/false,\n	log_end=true,\n	send_events_to_fmc=true,", rule.Name))
			return plan, diags
		case rule.Action.ValueString() == "MONITOR" && !rule.SendEventsToFmc.ValueBool():
			diags.AddAttributeError(path.Root("rules"), "Cannot use send_events_to_fmc=false when action=\"MONITOR\"",
				fmt.Sprintf("Rule %s has action=\"MONITOR\" so it must use:\n	log_begin=null/false,\n	log_end=true,\n	send_events_to_fmc=true,", rule.Name))
			return plan, diags
		}
	}

	// Validate rules.*.category_name relative order and rules.*.section relative order.
	type Node struct {
		sec string
		cat string
	}

	ranks := map[Node]int{}
	cat2sec := map[string]string{}
	count := 0
	for _, cat := range plan.Categories {
		if count == insertion {
			ranks[Node{sec: "mandatory", cat: ""}] = count
			count++
		}

		var node Node
		node.cat = cat.Name.ValueString()
		node.sec = cat.Section.ValueString()
		if node.sec == "" {
			node.sec = "default"
		}
		cat2sec[node.cat] = node.sec
		ranks[node] = count
		count++
	}
	ranks[Node{sec: "default", cat: ""}] = math.MaxInt

	var prev Node
	reached := 0
	for i, rule := range plan.Rules {
		var node Node
		node.cat = rule.CategoryName.ValueString()
		node.sec = cat2sec[node.cat]
		if node.sec == "" {
			node.sec = rule.GetSection()
		}

		// Never return to earlier rank.
		if ranks[node] < reached {
			diags.AddAttributeError(path.Root("rules"), "Wrong order of rules",
				fmt.Sprintf("Rule %s must be somewhere above rule %s, not directly below it.\n"+
					"  - rule %s: category_name=%s  section=%s\n"+
					"  - rule %s: category_name=%s  section=%s\n\n"+
					"That's because rules must be in this order (1->4):\n"+
					"  1. All rules from \"mandatory\" categories, in the order of their categories,\n"+
					"  2. then rules from \"mandatory\" section (uncategorized),\n"+
					"  3. then rules from non-mandatory categories, in the order of their categories,\n"+
					"  4. then rules from non-mandatory section (uncategorized).\n",
					plan.Rules[i].Name, plan.Rules[i-1].Name,
					plan.Rules[i-1].Name, plan.Rules[i-1].CategoryName, prev.sec, plan.Rules[i].Name, plan.Rules[i].CategoryName, node.sec))
			return plan, diags
		}

		reached = ranks[node]
		prev = node
	}

	return plan, diags
}
