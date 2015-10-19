package v2

import "testing"

type Service struct{}

type Args struct {
	S string
}

func (s *Service) NoArgsRets() {
}

func (s *Service) Args(str string, i int, args *Args) {
}

func (s *Service) Rets() (string, error) {
	return "", nil
}

func (s *Service) InvalidRets1() (error, string) {
	return nil, ""
}

func (s *Service) InvalidRets2() (string, string) {
	return "", ""
}

func (s *Service) InvalidRets3() (string, string, error) {
	return "", "", nil
}

func TestServerRegister(t *testing.T) {
	server := NewServer()
	service := new(Service)

	if err := server.RegisterName("calc", service); err != nil {
		t.Fatalf("%v", err)
	}

	if err := server.RegisterName("calc", service); err == nil {
		t.Fatal("Second time registering the same service should raise an error")
	}

	if len(server.serviceMap) != 1 {
		t.Fatalf("Expected 1 service entry but got %d", server.serviceMap)
	}

	svc, ok := server.serviceMap["calc"]
	if !ok {
		t.Fatalf("Expected service calc to be registered")
	}

	if len(svc.callbacks) != 3 {
		t.Fatalf("Expected 3 callbacks for service 'calc', got %d", len(svc.callbacks))
	}
}
