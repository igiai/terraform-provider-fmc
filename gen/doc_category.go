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

//go:build ignore

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath = "./gen/definitions/"
)

type YamlConfig struct {
	Name         string `yaml:"name"`
	DocCategory  string `yaml:"doc_category"`
	NoDataSource bool   `yaml:"no_data_source"`
	NoResource   bool   `yaml:"no_resource"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

var extraDocs = map[string]string{}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

func main() {
	files, _ := os.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(files))

	// Load configs
	for i, filename := range files {
		yamlFile, err := os.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	// Update doc category
	for i := range configs {
		for _, path := range docPaths {
			filename := path + SnakeCase(configs[i].Name) + ".md"
			content, err := os.ReadFile(filename)
			if err != nil {
				if configs[i].NoDataSource && path == "./docs/data-sources/" ||
					configs[i].NoResource && path == "./docs/resources/" {
					continue
				}
				log.Fatalf("Error opening documentation: %v", err)
			}

			cat := configs[i].DocCategory
			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

			os.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Update extra doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := os.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}
}
