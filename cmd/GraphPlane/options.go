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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)
func setupGrpcServerOptions() []grpc.ServerOption {
	// This is where you can setup custom options for the grpc server
	// https://godoc.org/google.golang.org/grpc#ServerOption
	opts := []grpc.ServerOption{grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000"))}
	return opts
}
func setupServeMuxOptions() []runtime.ServeMuxOption {
	// https://godoc.org/github.com/grpc-ecosystem/grpc-gateway/runtime#ServeMuxOption
	return nil
}
