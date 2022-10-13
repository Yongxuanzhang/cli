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

package task

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/tektoncd/cli/pkg/cli"
	"github.com/tektoncd/cli/pkg/trustedresources"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	cliopts "k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/yaml"
)

type verifyOptions struct {
	keyfile   string
	kmsKey    string
}

func verifyCommand(p cli.Params) *cobra.Command {
	opts := &verifyOptions{}
	f := cliopts.NewPrintFlags("trustedresources")

	c := &cobra.Command{
		Use:   "verify",
		Short: "Verify Tekton Task/Pipeline",
		Annotations: map[string]string{
			"commandType": "main",
		},
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			s := &cli.Stream{
				Out: cmd.OutOrStdout(),
				Err: cmd.OutOrStderr(),
			}
			b, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("error reading file: %v", err)
				return err
			}

			crd := &v1beta1.Task{}
			if err := yaml.Unmarshal(b, &crd); err != nil {
				log.Fatalf("error unmarshalling Task/Pipeline: %v", err)
				return err
			}

			if err := trustedresources.Verify(crd, opts.keyfile, opts.kmsKey); err != nil {
				log.Fatalf("error signing Task/Pipeline: %v", err)
				return err
			}
			fmt.Fprintf(s.Out, "Task %s passes verification \n", args[0])
			return nil
		},
	}
	f.AddFlags(c)
	c.Flags().StringVarP(&opts.keyfile, "key-file", "K", "", "Key file")
	c.Flags().StringVarP(&opts.kmsKey, "kms-key", "m", "", "KMS key url")
	return c
}
