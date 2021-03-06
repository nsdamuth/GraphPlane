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
	"os"
	"net"
	"github.com/spf13/viper"
	"damuth.nick/GraphPlane/internal/Logger"
)

var port = 4000
// var configs map[string]interface{}
var listen net.Listener
// var configs viper.New()

func main() {
	var configs = viper.New()
	logger.LogByType("INFO", "Starting GraphPlane Server")
	if (os.Args != nil) {
		if (stringInSlice("test", os.Args)) {

		}
		if (stringInSlice("green", os.Args)) {

		}
	}
	/////////////////////////////////////////
	// Configuration Loading & Monitoring
	go LoadConfigs(configs, &listen)
	for (configs.Get("server") == nil) { }
	
	/////////////////////////////////////////
	// Begin threading DB

	/////////////////////////////////////////
	// Begin threading Kafka Services

	/////////////////////////////////////////
	// Start up gRPC and Web Services, to includ UID
	ServeAndWait(port, configs, &listen)
	logger.LogByType("INFO", "Completing GraphPlane Server")
}
