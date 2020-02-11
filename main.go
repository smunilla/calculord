/*
Copyright © 2020 Samuel Munilla <smunilla@redhat.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/smunilla/calculord/cmd"
)

// version will be populated by the Makefile, read from VERSION file of the source code.
var version = ""

// gitCommit will be the hash that the binary was built from and will be populated by the Makefile
var gitCommit = ""

func main() {
	injectVersion()
	cmd.Execute()
}

func injectVersion() {
	if version != "" {
		cmd.Version = version
	}
	if gitCommit != "" {
		cmd.GitCommit = gitCommit
	}
}
