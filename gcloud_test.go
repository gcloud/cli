// GCloud - Go Packages for Cloud Services.
// Copyright (c) 2013 Garrett Woodworth (https://github.com/gwoo).

package main

import (
	"encoding/json"
	//"fmt"
	"regexp"
	"testing"

	p "github.com/gcloud/providers"
)

type MockServer struct {
	id   string
	name string
}

func (m *MockServer) Id() string {
	return m.id
}
func (m *MockServer) Name() string {
	return m.name
}
func (m *MockServer) State() string {
	return "running"
}
func (m *MockServer) Ips(t string) []string {
	return []string{}
}
func (m *MockServer) Size() string {
	return ""
}
func (m *MockServer) Image() string {
	return ""
}
func (m *MockServer) String() string {
	b, _ := m.MarshalJSON()
	return string(b)
}
func (m *MockServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Map{
		"id": m.Id(), "name": m.Name(),
	})
}

type MockServers struct{}

func (s *MockServers) NewServer(m p.Map) p.Server {
	return &MockServer{name: "My Server", id: "616fb98f-46ca-475e-917e-2563e5a8cd19"}
}
func (s *MockServers) List() ([]p.Server, error) {
	results := make([]p.Server, 0)
	r := s.NewServer(nil)
	return append(results, r), nil
}
func (s *MockServers) Show(id string) (p.Server, error) {
	r := s.NewServer(nil)
	return r, nil
}
func (s *MockServers) Create(n interface{}) (p.Server, error) {
	r := s.NewServer(nil)
	return r, nil
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
	if len(matches) > 0 && matches[0] == "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		return
	}
	t.Error("List Servers failed.")
}

func Test_ShowMockServers(t *testing.T) {
	result, _ := DoServers("mock", "show")
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(result), -1)
	if len(matches) > 0 && matches[0] == "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		return
	}
	t.Error("Show Servers failed.")
}

func Test_CreateMockServers(t *testing.T) {
	result, _ := DoServers("mock", "add")
	re := regexp.MustCompile("([A-z0-9]{8}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{4}-[A-z0-9]{12})")
	matches := re.FindAllString(string(result), -1)
	if len(matches) > 0 && matches[0] == "616fb98f-46ca-475e-917e-2563e5a8cd19" {
		return
	}
	t.Error("Create Servers failed.")
}

func Test_DestroyMockServers(t *testing.T) {
	result, _ := DoServers("mock", "destroy")
	if string(result) == "true" {
		return
	}
	t.Error("Destroy Servers failed.")
}
