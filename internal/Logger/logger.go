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
package logger

import (
	"os"
	"fmt"
	"log"
	"time"
	"strings"
	"encoding/json"
)
 
var (
	WarningLogger	*log.Logger
	InfoLogger		*log.Logger
	ErrorLogger		*log.Logger
	MetricLogger	*log.Logger
)

type M map[string]log.Logger

func init() {	
	// logs := []string{"info", "warning", "error", "metrics"}
	InfoLogger = log.New(createLogFile("info"), fmt.Sprintf("%s: ", strings.ToUpper("info")), log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(createLogFile("warning"), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(createLogFile("error"), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
func createLogFile(logtype string) *os.File {
	file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", logtype), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func LogData(logtype string, logfile string, ip string, user_id string, request_method string, status_code string, message string, pack string, func_call string) {
	 logdata, _ := json.Marshal(map[string]string{"user-identifier": user_id,
		 "time": fmt.Sprintf("%s", time.Now()),
		 "request-method": request_method,
		 "status-code": status_code,
		 "message": message,
		 "package": pack,
		 "function": func_call,
		 "server": ip})
	LogByType(logtype, string(logdata))
}
func LogByType(logtype string, message string) {
	 fmt.Println(fmt.Sprintf("%s: %s", logtype, message))
	 switch logtype {
		case "INFO":
			InfoLogger.Println(fmt.Sprintf("%s: %s", logtype, message))
		case "WARNING":
			WarningLogger.Println(fmt.Sprintf("%s: %s", logtype, message))
		case "ERROR":
			InfoLogger.Println(fmt.Sprintf("%s: %s", logtype, message))
		default:
			InfoLogger.Println(fmt.Sprintf("%s: %s", logtype, message))
	 }
}
 func Log(message string) {
	LogByType("INFO", message)
}
