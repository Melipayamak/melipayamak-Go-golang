package melipayamak

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type SOAPClient struct {
	httpClient *http.Client
	username   string
	password   string
}

func NewSOAPClient(username, password string, httpClient *http.Client) *SOAPClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &SOAPClient{
		httpClient: httpClient,
		username:   username,
		password:   password,
	}
}

const (
	SOAPSendUrl     = "https://api.payamak-panel.com/post/Send.asmx"
	SOAPReceiveUrl  = "https://api.payamak-panel.com/post/receive.asmx"
	SOAPContactsUrl = "https://api.payamak-panel.com/post/contacts.asmx"
	SOAPActionsUrl  = "https://api.payamak-panel.com/post/actions.asmx"
	SOAPScheduleUrl = "https://api.payamak-panel.com/post/schedule.asmx"
	SOAPTicketsUrl  = "https://api.payamak-panel.com/post/tickets.asmx"
	SOAPUsersUrl    = "https://api.payamak-panel.com/post/users.asmx"
	SOAPVoiceUrl    = "https://api.payamak-panel.com/post/voice.asmx"
	SoapNamespace   = "http://tempuri.org/"
)

type SOAPParam struct {
	Name     string
	Value    any
	Children []SOAPParam
}

var validXMLName = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_.-]*$`)

func validateXMLName(name string) error {
	if !validXMLName.MatchString(name) {
		return fmt.Errorf("invalid XML element name: %q", name)
	}

	return nil
}

func SOAPArray[T any](name, itemName string, values []T) SOAPParam {
	param := SOAPParam{Name: name}

	for _, value := range values {
		param.Children = append(param.Children, SOAPParam{
			Name:  itemName,
			Value: value,
		})
	}

	return param
}

func writeSOAPElement(buf *bytes.Buffer, param SOAPParam) error {
	if err := validateXMLName(param.Name); err != nil {
		return err
	}

	buf.WriteByte('<')
	buf.WriteString(param.Name)
	buf.WriteByte('>')

	if len(param.Children) > 0 {
		for _, child := range param.Children {
			if err := writeSOAPElement(buf, child); err != nil {
				return err
			}
		}
	} else if param.Value != nil {
		if err := xml.EscapeText(buf, []byte(fmt.Sprint(param.Value))); err != nil {
			return fmt.Errorf("escaping XML parameter %q: %w", param.Name, err)
		}
	}

	buf.WriteString("</")
	buf.WriteString(param.Name)
	buf.WriteByte('>')

	return nil
}

func buildSOAPEnvelope(method string, params ...SOAPParam) ([]byte, error) {
	var buf bytes.Buffer

	if err := validateXMLName(method); err != nil {
		return nil, err
	}

	buf.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	buf.WriteString(`<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"`)
	buf.WriteString(` xmlns:xsd="http://www.w3.org/2001/XMLSchema"`)
	buf.WriteString(` xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">`)
	buf.WriteString(`<soap:Body>`)
	buf.WriteString(`<` + method + ` xmlns="` + SoapNamespace + `">`)

	for _, param := range params {

		if err := writeSOAPElement(&buf, param); err != nil {
			return nil, err
		}
	}

	buf.WriteString(`</` + method + `>`)
	buf.WriteString(`</soap:Body>`)
	buf.WriteString(`</soap:Envelope>`)

	return buf.Bytes(), nil
}

func (c *SOAPClient) doRequest(
	ctx context.Context,
	soapUrl string,
	soapAction string,
	requestBody []byte,
) ([]byte, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		soapUrl,
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"text/xml; charset=utf-8",
	)
	req.Header.Set("SOAPAction", SoapNamespace+soapAction)

	// Set SOAPAction according to the service contract.

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("SOAP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading SOAP response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"SOAP HTTP error: status=%d, body=%s",
			resp.StatusCode,
			string(body),
		)
	}

	return body, nil
}
