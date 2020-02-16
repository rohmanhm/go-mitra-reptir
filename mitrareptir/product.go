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

package mitrareptir

import (
	"fmt"

	"github.com/rohmanhm/go-mitra-reptir/response"
)

// GetProduct detail of single product.
func (c *Config) GetProduct(code string) (*response.Product, error) {
	resp, err := client.R().
		SetBody(map[string]interface{}{"code": code}).
		SetHeader(c.getuAuthorizationHeader()).
		SetResult(&response.Product{}).
		Post(c.getProductEndpoint())

	if err != nil {
		return nil, fmt.Errorf("could not get product: %s", err)
	}

	return resp.Result().(*response.Product), nil
}

func (c *Config) getProductEndpoint() string {
	return c.generateEndpoint("product")
}
