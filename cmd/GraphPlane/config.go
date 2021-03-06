/*
* Written in 2021 by Nicholas S. Damuth
* V.1.0
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package main

import (
	// "os"
	"fmt"
	"net"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"damuth.nick/GraphPlane/internal/Logger"
)

func LoadConfigs(viper *viper.Viper, listen *net.Listener) {
	logger.Log("Loading GraphPlane Configs.")
	initialConfigFile := "configuration"

	viper.AddConfigPath("Configs/")
	viper.SetConfigName(initialConfigFile)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logger.LogByType("ERROR", "Loading GraphPlane Configs.")
	}

	logger.Log(fmt.Sprintf("Successfully loaded Configs from %s", initialConfigFile))
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Log(fmt.Sprintf("Successfully loaded Configs from %s", initialConfigFile))
	})
}