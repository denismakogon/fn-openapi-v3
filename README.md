OpenAPI v3 spec generator tool for Fn applications
==================================================

Idea
----
The Fn application is nothing but set of function where each has its own HTTP route for execution.

What if there's a way to build client binding for the particular serverless application?

This library and tool are designed to provide necessary API to generate OpenAPI v3.0.0 specification using Function spec language.


Function spec language
----------------------

This tool relies on improved Swagger API 2.0 plus additional inline referencing features 
that are missing in both Swagger 2.0 and OpenAPI 3.0 that are allowing developers to structure their application in more modular way.


Function spec example
---------------------

```yaml
version: 0.0.1

application:
  name: app
  config:
    c: 123
    10: "helloworld"

description: Functions spec that describes Fn-powered serverless application

functions:
  createUser:
    handler: handler.create
    events:
      - http:
          method: post
          fn: ${file(models/func.yml):first}
          documentation:
            summary: Create User
            description: Creates a user and then sends a generated password email
            requestBody:
              schema: ${file(models/PutDocumentRequest.json)}
            parameters:
              - name: username
                description: The username for a user to create
                required: true
                in: path
                schema:
                  type: string
                  pattern: "^[-a-z0-9_]+$"
              - name: membershipType
                description: The user's Membership Type
                required: true
                in: query
                schema:
                  type: string
                  enum:
                    - premium
                    - standard
            responses:
              200:
                description: create a user
                content:
                  application/json:
                    schema: ${file(models/PutDocumentResponse.json)}
              500:
                description: error
                content:
                  application/json:
                    schema: ${file(models/ErrorResponse.json)}
```
This sample you can find [here](examples/fn.yml)

using the following code:
```go
package main

import (
	"fmt"
	"github.com/denismakogon/fn-openapi/models"
	"io/ioutil"
	"os"
)

func main() {

	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fn models.Fn

	err = fn.Unmarshal(yamlFile, os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var oai models.OpenAPISpec
	err = oai.FromFnSpec("http://localhost:8080", &fn)
	err = oai.Marshal(os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
```

it is possible to turn Function spec into valid OpenAPI v3 specification.
To confirm that spec is valid use the following command:
```bash
docker run --rm -i -v `pwd`:/go fnproject/openapiv3-validator:0.0.1 /go/examples/openapi.yml
```
