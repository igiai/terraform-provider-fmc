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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccFmcAccessControlPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "description", "My access control policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "default_action_send_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "categories.0.name", "cat1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy.test", "rules.0.destination_network_literals.0.value", "10.2.2.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcAccessControlPolicyPrerequisitesConfig + testAccFmcAccessControlPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcAccessControlPolicyPrerequisitesConfig + testAccFmcAccessControlPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_access_control_policy.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccFmcAccessControlPolicyPrerequisitesConfig = `
resource "fmc_network" "this" {
  name   = "mynetwork1"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "this" {
  name = "myhost1"
  ip   = "10.1.1.1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccFmcAccessControlPolicyConfig_minimum() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccFmcAccessControlPolicyConfig_all() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My access control policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = true` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_send_syslog = true` + "\n"
	config += `	categories = [{` + "\n"
	config += `	  name = "cat1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  action = "ALLOW"` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  category_name = "cat1"` + "\n"
	config += `	  enabled = true` + "\n"
	config += `	  source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  destination_network_literals = [{` + "\n"
	config += `		value = "10.2.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  source_network_objects = [{` + "\n"
	config += `		id = fmc_network.this.id` + "\n"
	config += `		type = fmc_network.this.type` + "\n"
	config += `	}]` + "\n"
	config += `	  destination_network_objects = [{` + "\n"
	config += `		id = fmc_host.this.id` + "\n"
	config += `		type = fmc_host.this.type` + "\n"
	config += `	}]` + "\n"
	config += `	  log_begin = false` + "\n"
	config += `	  log_end = false` + "\n"
	config += `	  log_files = false` + "\n"
	config += `	  send_events_to_fmc = false` + "\n"
	config += `	  description = ""` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestNewValidAccessControlPolicy(t *testing.T) {

	steps := []resource.TestStep{{
		Config: `resource fmc_access_control_policy step1 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`}`,
	}, {
		Config: `resource fmc_access_control_policy step2 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`}`,
	}, {
		Config: `resource fmc_access_control_policy step3 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "catd" },` + "\n" +
			`		{ name = "catm", section = "mandatory" }` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`"catm" must be somewhere above category "catd"`),
	}, {
		Config: `resource fmc_access_control_policy step4 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [{ name = "cat1", section = "" }]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`value must be one of`),
	}, {
		Config: `resource fmc_access_control_policy step5 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "mandatory" },` + "\n" +
			`		{ name = "cat2", section = "mandatory" },` + "\n" +
			`		{ name = "cat3", section = "default"   },` + "\n" +
			`		{ name = "cat4", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat1",      name = "r1", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat3",      name = "r3", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat4",      name = "r4", action = "ALLOW"},` + "\n" +
			`		{ section = "default",         name = "r5", action = "ALLOW"},` + "\n" +
			`	]` + "\n" +
			`}`,
	}, {
		Config: `resource fmc_access_control_policy step6 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "mandatory" },` + "\n" +
			`		{ name = "cat2", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",    name = "step6r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",     name = "step6r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step7 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",      name = "r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step8 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" },` + "\n" +
			`		{ name = "cat2" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat2",      name = "r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat1",      name = "r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step9 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1", section = "default"   }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{                              name = "step9r2", action = "ALLOW"},` + "\n" +
			`		{ category_name = "cat1",      name = "step9r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step10 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [` + "\n" +
			`		{                              name = "step10r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "step10r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step11 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [` + "\n" +
			`		{ name = "cat1" }` + "\n" +
			`	]` + "\n" +
			`	rules = [` + "\n" +
			`		{ category_name = "cat1",      name = "step11r2", action = "ALLOW"},` + "\n" +
			`		{ section = "mandatory",       name = "step11r1", action = "ALLOW"}` + "\n" +
			`	]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Rule .* must be somewhere above rule`),
	}, {
		Config: `resource fmc_access_control_policy step12 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [{ name = "cat1", section = "default" }]` + "\n" +
			`	rules = [{ category_name = "cat1", section = "default", name = "r1", action = "ALLOW"}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Cannot use section together with category_name`),
	}, {
		Config: `resource fmc_access_control_policy step13 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [{ category_name = "--Undefined--", name = "r1", action = "ALLOW"}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`value is reserved`),
	}, {
		Config: `resource fmc_access_control_policy step14 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	categories = [{ name = "--Undefined--" }]` + "\n" +
			`	rules = [{ category_name = "--Undefined--", name = "r1", action = "ALLOW"}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`value is reserved`),
	}, {
		Config: `resource fmc_access_control_policy step15 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [{ name = "r1", action = "MONITOR"}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Cannot use log_end=false when action="MONITOR"`),
	}, {
		Config: `resource fmc_access_control_policy step16 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [{ name = "r1", action = "MONITOR", log_end=true, send_events_to_fmc=null}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Cannot use send_events_to_fmc=false when action="MONITOR"`),
	}, {
		Config: `resource fmc_access_control_policy step17 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [{ name = "r1", action = "MONITOR", log_end=true, send_events_to_fmc=true}]` + "\n" +
			`}`,
	}, {
		Config: `resource fmc_access_control_policy step18 {` + "\n" +
			`	name = "pol1"` + "\n" +
			`	default_action = "BLOCK"` + "\n" +
			`	rules = [{ name = "r1", action = "MONITOR", log_begin=true, log_end=true, send_events_to_fmc=true}]` + "\n" +
			`}`,
		ExpectError: regexp.MustCompile(`Cannot use log_begin=true when action="MONITOR"`),
	}}

	resource.Test(t, resource.TestCase{
		// If you see "Step 7/x, expected an error" look above for the name "step7".
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}
