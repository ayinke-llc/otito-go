# Go API client for otito

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://otito.dev](https://otito.dev)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/ayinke-llc/otito-go"

        // to authenticate
	ctx := context.WithValue(context.Background(), otito.ContextAPIKeys, map[string]otito.APIKey{
		"ApiKeyAuth": {
			Key:    "sk_xyz",
			Prefix: "Bearer",
		},
	})


	cfg := otito.NewConfiguration()

	client := otito.NewAPIClient(cfg)
```

