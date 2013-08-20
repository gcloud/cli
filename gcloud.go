// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).
package main

import (
	"encoding/json"
	"errors"
	"flag"
	c "github.com/gcloud/compute"
	p "github.com/gcloud/compute/providers"
	_ "github.com/gcloud/compute/providers/vbox"
)

var provider = flag.String("p", "vbox", "Provider to use for actions.")
var name = flag.String("name", "", "Name of the object.")

func DoServers(provider string, method string) ([]byte, error) {
	s := &c.Servers{Provider: provider}
	switch method {
	case "list":
		r, err := s.List()
		return respond(r, err)
	case "show":
		r, err := s.Show(*name)
		return respond(r, err)
	case "add":
		r, err := s.Create(&p.Server{Name: *name})
		return respond(r, err)
	case "destroy":
		r, err := s.Destroy(*name)
		return respond(r, err)
	}
	return nil, errors.New("Missing method for Servers.")
}

func main() {
	flag.Parse()
	method := flag.Arg(0)
	pkg := flag.Arg(1)
	switch pkg {
	case "servers":
		result, err := DoServers(*provider, method)
		if err != nil {
			println(string(err.Error()))
			return
		}
		println(string(result))
	}
}

func respond(r interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	switch response := r.(type) {
	case nil:
		return nil, err
	case []byte:
		return response, err
	case bool:
		return []byte(`true`), err
	case string:
		return []byte(response), err
	default:
		return json.Marshal(r)
	}
	return nil, err
}
