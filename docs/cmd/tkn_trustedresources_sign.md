## tkn trustedresources sign

Sign Tekton Task/Pipeline

### Usage

```
tkn trustedresources sign
```

### Synopsis

Sign Tekton Task/Pipeline

### Options

```
      --allow-missing-template-keys   If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats. (default true)
  -f, --file-name string              Skip verifying the payload'signature
  -h, --help                          help for sign
  -K, --key-file string               Key file
  -d, --kind string                   Skip verifying the payload'signature
  -m, --kms-key string                Skip verifying the payload'signature
  -o, --output string                 Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.
  -r, --resource-file string          Skip verifying the payload'signature
      --show-managed-fields           If true, keep the managedFields when printing objects in JSON or YAML format.
  -t, --target-dir string             Skip verifying the payload'signature
      --template string               Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].
```

### Options inherited from parent commands

```
  -c, --context string      name of the kubeconfig context to use (default: kubectl config current-context)
  -k, --kubeconfig string   kubectl config file (default: $HOME/.kube/config)
  -n, --namespace string    namespace to use (default: from $KUBECONFIG)
  -C, --no-color            disable coloring (default: false)
```

### SEE ALSO

* [tkn trustedresources](tkn_trustedresources.md)	 - Sign and verify Tekton Resources

