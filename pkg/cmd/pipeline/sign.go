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

type signOptions struct {
	keyfile    string
	kmsKey     string
	targetFile string
}

func signCommand(p cli.Params) *cobra.Command {
	opts := &signOptions{}
	f := cliopts.NewPrintFlags("trustedresources")

	c := &cobra.Command{
		Use:   "sign",
		Short: "Sign Tekton Task/Pipeline",
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

			crd := &v1beta1.Pipeline{}
			if err := yaml.Unmarshal(b, &crd); err != nil {
				return fmt.Errorf("error unmarshalling Task: %v", err)
			}

			// Sign the task and save to file
			if err := trustedresources.Sign(crd, opts.keyfile, opts.kmsKey, opts.targetFile); err != nil {
				return fmt.Errorf("error signing Task: %v", err)
			}
			fmt.Fprintf(s.Out, "Task %s is signed successfully \n", args[0])
			return nil
		},

	}
	f.AddFlags(c)
	c.Flags().StringVarP(&opts.keyfile, "key-file", "K", "", "Key file")
	c.Flags().StringVarP(&opts.kmsKey, "kms-key", "m", "", "Skip verifying the payload'signature")
	c.Flags().StringVarP(&opts.targetFile, "file-name", "f", "", "Skip verifying the payload'signature")

	return c
}

func (s *signOptions) Run(args []string) error {
	tsBuf, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatalf("error reading file: %v", err)
		return err
	}

	crd := &v1beta1.Pipeline{}
	if err := yaml.Unmarshal(tsBuf, &crd); err != nil {
		log.Fatalf("error unmarshalling Task: %v", err)
		return err
	}

	// Sign the task and write to target file
	if err := trustedresources.Sign(crd, s.keyfile, s.kmsKey, s.targetFile); err != nil {
		log.Fatalf("error signing Task: %v", err)
		return err
	}

	return nil
}
