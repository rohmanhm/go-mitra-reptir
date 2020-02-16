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

package example

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rohmanhm/go-mitra-reptir/mitrareptir"
)

// Client mitra reptir
var Client *mitrareptir.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not env file")
	}
	Client = mitrareptir.New(os.Getenv("REPTIR_AUTH_KEY"))
}
