// Copyright 2025.
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

package services

import (
	"reflect"
	"strings"
	"testing"
)

func TestTmuxInstallScriptHasNoSingleQuotes(t *testing.T) {
	// The script is wrapped in sh -c '...', so a single quote would break it.
	if strings.Contains(tmuxInstallScript, "'") {
		t.Fatalf("tmuxInstallScript must not contain single quotes: %s", tmuxInstallScript)
	}
}

func TestInteractiveSSHArgs(t *testing.T) {
	tests := []struct {
		name  string
		env   string
		extra []string
		want  []string
	}{
		{"tmux default", "", nil, []string{"-t", "host", tmuxRemoteCommand}},
		{"tmux with forwards", "1", []string{"-L", "8080:localhost:80"}, []string{"-L", "8080:localhost:80", "-t", "host", tmuxRemoteCommand}},
		{"tmux disabled", "0", nil, []string{"host"}},
		{"tmux disabled off", "OFF", []string{"-D", "1080"}, []string{"-D", "1080", "host"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("LAZYSSH_TMUX", tt.env)
			got := interactiveSSHArgs(tt.extra, "host")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interactiveSSHArgs() = %q, want %q", got, tt.want)
			}
		})
	}
}
