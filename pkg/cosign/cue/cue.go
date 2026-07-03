//
// Copyright 2021 The Sigstore Authors.
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

package cue

import (
	"crypto/fips140"
	"fmt"
	"os"
	"strings"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	cuejson "cuelang.org/go/encoding/json"
)

// RHTAS FIPS - DO NOT REMOVE
// ========================================
func checkCUEFIPSImports(entrypoints []string) error {
	for _, path := range entrypoints {
		data, err := os.ReadFile(path) // #nosec G304
		if err != nil {
			return fmt.Errorf("cannot read CUE policy %q for FIPS validation: %w", path, err)
		}
		src := string(data)
		if strings.Contains(src, `"crypto/md5"`) || strings.Contains(src, `"crypto/sha1"`) {
			return fmt.Errorf("CUE policy %q imports crypto/md5 or crypto/sha1 which are not available in FIPS 140-3 mode", path)
		}
	}
	return nil
}

// ========================================

func ValidateJSON(jsonBody []byte, entrypoints []string) error {
	// RHTAS FIPS - DO NOT REMOVE
	// ========================================
	if fips140.Enabled() {
		if err := checkCUEFIPSImports(entrypoints); err != nil {
			return err
		}
	}
	// ========================================

	ctx := cuecontext.New()
	bis := load.Instances(entrypoints, nil)

	for _, bi := range bis {
		if bi.Err != nil {
			return bi.Err
		}

		value := ctx.BuildInstance(bi)
		if value.Err() != nil {
			return value.Err()
		}

		err := cuejson.Validate(jsonBody, value)
		if err != nil {
			return err
		}
	}

	return nil
}
