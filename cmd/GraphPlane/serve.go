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
	//  "os"
	 "fmt"
	 "net"
	 "damuth.nick/GraphPlane/internal/Logger"
	//  "google.golang.org/grpc"
	//  "github.com/cockroachdb/cmux"
	 "github.com/zenazn/goji/bind"
	 "github.com/zenazn/goji/graceful"
	 // "github.com/golang/protobuf/proto"
	//  "github.com/grpc-ecosystem/grpc-gateway/runtime"
 )
 var loaded_port interface{}
 var gracefullyStopped bool = false
 
 type ServiceHandlers struct {
 }
 
 func ServeAndWait(port int, configs *map[string]interface{}, listen *net.Listener) {
	logger.LogByType("INFO", "Attempting to Server & Wait gRPC & HTTP Services")
	 if ((*configs)["server"] != nil) {
		 loaded_port = (((*configs)["server"]).(map[string]interface{}) )["port"]
		 if (loaded_port != nil) {
			 port = int(loaded_port.(float64))
		 }
	 }
	 startServer(port, listen)
	 if (gracefullyStopped == false) {
		 go ServeAndWait(port, configs, listen)
	 }
 }
 func startServer(port int, listen *net.Listener) {
	var err error = nil
	*listen, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.LogByType("ERROR", fmt.Sprintf("failed to listen: %v", err))
	}
	
	// Here we handle some signals so that we can stop the server when running.
	graceful.HandleSignals()
	// Handle gracefully stopping the server on signals.
	bind.Ready()
	graceful.PreHook(func() {
		gracefullyStopped = true
		logger.Log("Server received signal, gracefully stopping.")
	})
	graceful.PostHook(func() {
		logger.Log("Server stopped")
	})
	// closeEps(eps)
	// cmux starts all the servers for us when we call Serve() (grpcS and httpS)
	logger.LogByType("INFO", fmt.Sprintf("listening and serving (multiplexed) on: %d", port))
	// err = multiplexer.Serve()
 }
