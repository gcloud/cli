// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	//"fmt"
	//c "github.com/gcloud/compute"
	p "github.com/gcloud/compute/providers"
	"regexp"
	"testing"
)

type MockServers struct{}

// List servers available on the account.
func (s *MockServers) List() ([]byte, error) {
	return []byte(`[{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}]`), nil
}
func (s *MockServers) Show(id string) ([]byte, error) {
	return []byte(`{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (s *MockServers) Create(n *p.Server) ([]byte, error) {
	return []byte(`{"Name":"My Server","Id": "616fb98f-46ca-475e-917e-2563e5a8cd19"}`), nil
}
func (s *MockServers) Destroy(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Reboot(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Start(id string) (bool, error) {
	return true, nil
}
func (s *MockServers) Stop(id string) (bool, error) {
	return true, nil
}

func init() {
	p.RegisterServers("mock", &MockServers{})
}

func Test_ListMockServers(t *testing.T) {
	result, _ := DoServers("mock", "list")
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(result), -1)
	if matches[0] != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("List Servers failed.")
	}
}

func Test_ShowMockServers(t *testing.T) {
	result, _ := DoServers("mock", "show")
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(result), -1)
	if matches[0] != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Show Servers failed.")
	}
}

func Test_CreateMockServers(t *testing.T) {
	result, _ := DoServers("mock", "add")
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(result), -1)
	if matches[0] != "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		t.Error("Create Servers failed.")
	}
}

func Test_DestroyMockServers(t *testing.T) {
	result, _ := DoServers("mock", "destroy")
	if string(result) != "true" {
		t.Error("Destroy Servers failed.")
	}
}
