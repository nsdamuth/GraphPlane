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
	"fmt"
	"net"
	"strings"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"github.com/fsnotify/fsnotify"

	"damuth.nick/GraphPlane/internal/Logger"
)

var old_configs map[string]interface{}
var byteValue []byte

func LoadConfigs(configs *map[string]interface{}, listen *net.Listener) {
	logger.Log("Loading GraphPlane Configs.")
	folderName := "Configs/"
	initialConfigFile := "configuration.json"
	if (FileExists(fmt.Sprintf("%s%s", folderName, initialConfigFile))) {
		byteValue = loadFile(fmt.Sprintf("%s%s", folderName, initialConfigFile))
		processConfigs(byteValue, configs, &old_configs, initialConfigFile, listen)
	} else {
		fmt.Println(fmt.Sprintf("No such File : %s%s", folderName, initialConfigFile))
	// 	logger.LogByType("ERROR", "Missing Configurations File")
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.LogByType("ERROR", "Loading GraphPlane Configs.")
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fileExtension := filepath.Ext(event.Name)
				if fileExtension == ".json" {
					// var result map[string]interface{}
					filename := strings.ReplaceAll(event.Name, folderName, "")
					logger.LogByType("EVENT", fmt.Sprintf("File Changed : %#v", filename))
					byteValue = loadFile(event.Name)
					processConfigs(byteValue, configs, &old_configs, event.Name, listen)
				}
				// watch for errors
			case err := <-watcher.Errors:
				logger.LogByType("ERROR", fmt.Sprintf("ERROR", err))
			}
		}
	}()

	if err := watcher.Add(fmt.Sprintf("./%s", folderName)); err != nil {
		fmt.Println("ERROR", err)
	}
	<-done

	logger.LogByType("INFO", "Starting Config Monitoring.")
}
func loadFile(event string) []byte {
	jsonFile, err := os.Open(fmt.Sprintf("./%s", event))
	if err != nil {
		logger.LogByType("ERROR", fmt.Sprintf("./%s", err))
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
func processConfigs(byteValue []byte, configs *map[string]interface{}, old_configs *map[string]interface{}, fileName string, listen *net.Listener) {
	json.Unmarshal([]byte(byteValue), &configs)
	if (old_configs != nil) {
		if (reflect.DeepEqual(*old_configs, *configs)) {
			logger.Log(fmt.Sprintf("Successfully loaded Configs from %s", fileName))
			(*listen).Close()
		}
	} else {
		old_configs = configs
	}
	*old_configs = CopyMap(*configs)
}
