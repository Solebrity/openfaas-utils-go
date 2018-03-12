# OpenFaaS Utils for Go
OpenFaas utils is a tiny package for Golang-based OpenFaas functions that help with some common use cases.  This repo was created mainly because each function by default is in its own package, and dep does not support local depencencies.

## Secrets
OpenFaas mounts secrets as volumes to your pods in Kubernetes.  In order to read the value of the secrets, you must open and read from `/run/secrets/{YOUR_SECRET_KEY}`.  We've written a helper function to do this for you.

```go
package function

import ofutils "github.com/Solebrity/openfaas-utils-go"

func Handle(req []byte) string {
    value, err := ofutils.GetSecretValue("my-secret-key")
    if err != nil {
        return err.Error()
    }
    return value
}
```

More to come...