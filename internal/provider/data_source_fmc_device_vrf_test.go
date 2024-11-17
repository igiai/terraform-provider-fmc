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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceVRF(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vrf.test", "name", "VRF_A"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vrf.test", "description", "My VRF object"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceVRFPrerequisitesConfig + testAccDataSourceFmcDeviceVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceVRFPrerequisitesConfig + testAccNamedDataSourceFmcDeviceVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceVRFPrerequisitesConfig = `

variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceVRFConfig() string {
	config := `resource "fmc_device_vrf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "VRF_A"` + "\n"
	config += `	description = "My VRF object"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vrf" "test" {
			id = fmc_device_vrf.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceVRFConfig() string {
	config := `resource "fmc_device_vrf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "VRF_A"` + "\n"
	config += `	description = "My VRF object"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vrf" "test" {
			name = fmc_device_vrf.test.name
			device_id = var.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig