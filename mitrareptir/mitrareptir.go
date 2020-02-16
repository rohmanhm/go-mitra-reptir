// Copyright 2020 Rohman Habib M
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package mitrareptir provides request API to mitra reptir.
package mitrareptir

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// Config mitra reptir.
type Config struct {
	BaseURL string
	APIKey  string
}

var client = resty.New()

// New create config for mitra reptir.
func New(apiKey string) *Config {
	return &Config{
		BaseURL: "https://api.reptir.com/v1",
		APIKey:  apiKey,
	}
}

func (c *Config) generateEndpoint(path string) string {
	return fmt.Sprintf("%s/%s", c.BaseURL, path)
}

func (c *Config) getuAuthorizationHeader() (string, string) {
	return "Authorization", fmt.Sprintf("Bearer %s", c.APIKey)
}
