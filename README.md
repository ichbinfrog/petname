
# Petname

[![Go Report Card](https://goreportcard.com/badge/ichbinfrog/petname)](https://goreportcard.com/report/github.com/ichbinfrog/petname) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/ichbinfrog/petname/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/ichbinfrog/petname?status.svg)](https://godoc.org/github.com/ichbinfrog/petname) ![Build](https://travis-ci.org/ichbinfrog/petname.svg?branch=master) [![codecov](https://codecov.io/gh/ichbinfrog/petname/branch/master/graph/badge.svg)](https://codecov.io/gh/ichbinfrog/petname)

Petname is a server that generates unique petnames (see [RFC](https://tools.ietf.org/html/rfc1178)) which are pronounceable, sometimes even memorable names consisting of a random combination of adverbs, an adjective, and an animal name.

You can use this server to:
- Create unique memorable names for your container/processes
- Dynamically import more names into the server by adding values to a yaml file
- Dynamically add names to the API by querying the API itself
- (**future**) Distributed lightweight API server to handle unique cluster names
- (**future**) Secure access to the API using Bearer Tokens

## Running

Using the docker image:
```sh
docker run --rm -d -p $PORT:$PORT ichbinfrog/petname:v0.2 serve --port $PORT
```

Using the binary:
```sh
# Generate locally
petname gen <nb>

# Create petname server
petname serve [flags]

flags:
  -p, --port string   Port for serving the API server (default: 8000)
```

## Querying the API

| Endpoint                                            | Description                                                           |
| --------------------------------------------------- | --------------------------------------------------------------------- |
| `/get/default?amount=n`                             | Default endpoint present, returns n petname                           |
| `/api?name={name}&template={template}`              | Create a new API endpoint                                             |
| `/api/{api}/reload`                                 | Clear the used binary tree for the specific API                       |

In order to create an API endpoint, you have to send a GET request with the query parameters:
- `name`: the name of the desired endpoints (will result in the `/get/{name}` API being available)
- `template`: template for the petname generation. Semi Golang text/template with two available variable atm ( `{{ Name }}`, `{{ Adjective }}` and `{{ Adverb }}`).

For example with golang:
```go
// Creates a request to send to the /api endpoint
req, err := http.NewRequest("GET", "http://localhost:8000/api", nil)
if err != nil {
  panic(err)
}

// Creates a http client
client := &http.Client{}

// Adds the v2 endpoint
q := req.URL.Query()
q.Add("name", "v2")
q.Add("template", "{{.Name}}~{{.Name}}~{{.Adjective}}")

// Perform the request
r, reqErr := client.Do(req)
```
Now you can query the `/get/v2` endpoint and it would give you names such as: `gopher~duck~boring`


## Installation

Using the docker image (recommended):
```sh
# If you want to build your own image
docker build . -t petname:v0.2

# If you want to pull directly from dockerhub
docker pull ichbinfrog/petname:v0.2
```

Using the binary:
```sh
# Assuming you've downloaded the latest release
# on the github release tab
chmod +x ./petname:v0.2

# If you want to add it to your path
mv ./petname:v0.2 /usr/local/bin/petname
```
