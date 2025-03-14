// Copyright © 2025 Weald Technology Trading.
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

package blocktrail

import (
	"context"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	if os.Getenv("ETHDO_TEST_CONNECTION") == "" {
		t.Skip("ETHDO_TEST_CONNECTION not configured; cannot run tests")
	}

	tests := []struct {
		name string
		vars map[string]interface{}
		err  string
	}{
		{
			name: "NoBlock",
			vars: map[string]interface{}{
				"timeout":    "60s",
				"connection": os.Getenv("ETHDO_TEST_CONNECTION"),
				"blockid":    "invalid",
			},
			err: "failed to obtain beacon block: failed to request signed beacon block: GET failed with status 400: {\"code\":400,\"message\":\"BAD_REQUEST: Unsupported endpoint version: v2\",\"stacktraces\":[]}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viper.Reset()

			for k, v := range test.vars {
				viper.Set(k, v)
			}
			cmd, err := newCommand(context.Background())
			require.NoError(t, err)
			err = cmd.process(context.Background())
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
