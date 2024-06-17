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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceFmcAccessControlPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "description", "My access control policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_send_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "categories.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "categories.0.name", "cat1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAccessControlPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcAccessControlPolicyConfig() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My access control policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = true` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_send_syslog = true` + "\n"
	config += `	categories = [{` + "\n"
	config += `	  id = ""` + "\n"
	config += `	  name = "cat1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  id = ""` + "\n"
	config += `	  action = "ALLOW"` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  category_name = "cat1"` + "\n"
	config += `	  enabled = true` + "\n"
	config += `	  source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_control_policy" "test" {
			id = fmc_access_control_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
