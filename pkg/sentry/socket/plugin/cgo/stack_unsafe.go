// Copyright 2023 The gVisor Authors.
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

// Package cgo provides interfaces definition to interact with third-party
// network stack. It also implements CGO wrappers to handle Golang arguments
// to CGO and CGO return values to Golang.
//
// Third-party external network stack will implement interfaces defined in this
// package in order to be used by gVisor.
package cgo

/*
int plugin_initstack(int argc, char **argv);
int plugin_preinitstack(int argc, char **argv);
int plugin_postinitstack(int argc, char **argv);
*/
import "C"

// InitStack implements CGO wrapper for plugin_initstack.
func InitStack(argv []string) int {
	//TODO: implement cgo wrapper
	return 0
}

// PreInitStack implements CGO wrapper for plugin_preinitstack.
func PreInitStack(argv []string) int {
	//TODO: implement cgo wrapper
	return 0
}

// PostInitStack implements CGO wrapper for plugin_postinitstack.
func PostInitStack(argv []string) int {
	//TODO: implement cgo wrapper
	return 0
}
