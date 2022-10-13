// Copyright © 2022 The Tekton Authors.
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

package pipeline

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	cosignsignature "github.com/sigstore/cosign/pkg/signature"
	"github.com/tektoncd/cli/pkg/test"
	"github.com/tektoncd/cli/pkg/trustedresources"
)

func TestSign(t *testing.T) {
	ctx := context.Background()
	p := &test.Params{}

	task := Command(p)

	os.Setenv("COSIGN_PASSWORD", "1234")
	tmpDir := t.TempDir()
	targetFile := filepath.Join(tmpDir, "signed.yaml")
	out, err := test.ExecuteCommand(task, "sign", "testdata/pipeline.yaml", "-K", "testdata/cosign.key", "-f", targetFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "Task testdata/pipeline.yaml is signed successfully \n"
	test.AssertOutput(t, expected, out)

	// verify the signed task
	verifier, err := cosignsignature.LoadPublicKey(ctx, "testdata/cosign.pub")
	if err != nil {
		t.Errorf("error getting verifier from key file: %v", err)
	}

	signed, err := ioutil.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("error reading file: %v", err)
	}

	target, signature, err := trustedresources.UnmarshalCRD(signed, "Pipeline")
	if err := trustedresources.VerifyInterface(target, verifier, signature); err != nil {
		t.Fatalf("VerifyInterface get error: %v", err)
	}

}
