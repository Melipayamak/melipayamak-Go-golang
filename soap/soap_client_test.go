package melipayamak

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBuildSOAPEnvelope(t *testing.T) {
	body, err := buildSOAPEnvelope(
		"SendSimpleSMS2",
		SOAPParam{Name: "username", Value: "user"},
		SOAPParam{Name: "password", Value: "pass"},
		SOAPParam{Name: "text", Value: "Hello & World"},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
}

func TestSendSimpleSMS2(t *testing.T) {
	const expectedSOAPAction = "http://tempuri.org/SendSimpleSMS2"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check HTTP method.
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}

		// Check endpoint.
		if r.URL.Path != "/Send.asmx" {
			t.Errorf("unexpected endpoint: %s", r.URL.Path)
		}

		// Check Content-Type.
		if got := r.Header.Get("Content-Type"); !strings.Contains(got, "text/xml") {
			t.Errorf("unexpected Content-Type: %s", got)
		}

		// Check SOAPAction.
		if got := r.Header.Get("SOAPAction"); got != expectedSOAPAction {
			t.Errorf("unexpected SOAPAction: %s", got)
		}

		// Read request body.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("reading request body: %v", err)
			return
		}

		xmlBody := string(body)

		// Check SOAP envelope content.
		expectedElements := []string{
			"<SendSimpleSMS2",
			"<username>user</username>",
			"<password>pass</password>",
			"<to>09120000000</to>",
			"<from>5000</from>",
			"<text>Hello &amp; World</text>",
			"<isflash>false</isflash>",
		}

		for _, element := range expectedElements {
			if !strings.Contains(xmlBody, element) {
				t.Errorf("expected XML to contain %q, got:\n%s", element, xmlBody)
			}
		}

		// Simulate successful SOAP response.
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<response>OK</response>"))
	}))
	defer server.Close()

	client := NewSOAPClient("user", "pass", server.Client())
	// client.baseURL = server.URL + "/"

	// Execute the method.
	response, err := client.SendSimpleSMS2(
		context.Background(),
		"09120000000",
		"5000",
		"Hello & World",
		false,
	)

	if err != nil {
		t.Fatalf("SendSimpleSMS2 returned an error: %v", err)
	}

	if !strings.Contains(string(response), "SendSimpleSMS2Result") {
		t.Errorf("unexpected response: %s", response)
	}
}

func TestSendSms(t *testing.T) {
	const expectedSOAPAction = "http://tempuri.org/SendSms"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check HTTP method.
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}

		// Check endpoint.
		if r.URL.Path != "/Send.asmx" {
			t.Errorf("unexpected endpoint: %s", r.URL.Path)
		}

		// Check Content-Type.
		if got := r.Header.Get("Content-Type"); !strings.Contains(got, "text/xml") {
			t.Errorf("unexpected Content-Type: %s", got)
		}

		// Check SOAPAction.
		if got := r.Header.Get("SOAPAction"); got != expectedSOAPAction {
			t.Errorf("unexpected SOAPAction: %s", got)
		}

		// Read request body.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("reading request body: %v", err)
			return
		}

		xmlBody := string(body)

		// Check SOAP envelope content.
		expectedElements := []string{
			"<SendSms",
			"<username>user</username>",
			"<password>pass</password>",
			"<to><string>09120000000</string><string>09130000000</string></to>",
			"<from>5000</from>",
			"<text>Hello &amp; World</text>",
			"<isflash>false</isflash>",
			"<udh></udh>",
		}

		for _, element := range expectedElements {
			if !strings.Contains(xmlBody, element) {
				t.Errorf("expected XML to contain %q, got:\n%s", element, xmlBody)
			}
		}

		// Simulate successful SOAP response.
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<response>OK</response>"))
	}))
	defer server.Close()

	client := NewSOAPClient("user", "pass", server.Client())
	// client.baseURL = server.URL + "/"

	// Execute the method.
	response, err := client.SendSms(
		context.Background(),
		[]string{"09120000000", "09130000000"},
		"5000",
		"Hello & World",
		false,
		"",
	)

	if err != nil {
		t.Fatalf("SendSms returned an error: %v", err)
	}

	if !strings.Contains(string(response), "SendSmsResult") {
		t.Errorf("unexpected response: %s", response)
	}
}
