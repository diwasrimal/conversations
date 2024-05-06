package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/diwasrimal/gochat/backend/types"

	"github.com/google/uuid"
)

func TestRegisterPost(t *testing.T) {
	failingPayloads := []types.Json{
		{
			"fname": "Diwas",
			"lname": "Rimal",
		},
		{
			"fname":    "Diwas",
			"lname":    "Rimal",
			"username": "   ",
			"password": "yes",
		},
		{
			"fname":    "Diwas",
			"lname":    "  ",
			"username": "drimal",
			"password": "yes",
		},
		{
			"fname":    "Diwas",
			"lname":    "Rimal",
			"username": "drimal",
			"password": "yesyesyesyesyesyesyesyesyyesyesyesyesyesyesesyesyesyesyesyesyesyesyesyesyes", // more than 72
		},
	}

	passingPayloads := []types.Json{
		{
			"fname":    "test:Diwas",
			"lname":    "test:Rimal",
			"username": "test:" + uuid.New().String(), // username should not conflict with existing one in db
			"password": "test:yes",
		},
	}

	url := "http://localhost:3030/api/register"

	t.Log("Checking failing cases")
	for i, payload := range failingPayloads {
		encoded, _ := json.Marshal(payload)
		body := bytes.NewBuffer(encoded)
		resp, err := http.Post(url, "application/json", body)
		if err != nil {
			t.Errorf("Error making post request: %v\n", err)
			continue
		}
		respContent, _ := io.ReadAll(resp.Body)
		t.Logf("case: %v, resp status: %v, resp body: %s", i, resp.Status, respContent)
		if resp.StatusCode < 400 {
			t.Fail()
		}
	}

	t.Log("Checking passing cases")
	for i, payload := range passingPayloads {
		encoded, _ := json.Marshal(payload)
		body := bytes.NewBuffer(encoded)
		resp, err := http.Post(url, "application/json", body)
		if err != nil {
			t.Errorf("Error making post request: %v\n", err)
			continue
		}
		respContent, _ := io.ReadAll(resp.Body)
		t.Logf("case: %v, resp status: %v, resp body: %s", i, resp.Status, respContent)
		if resp.StatusCode >= 400 {
			t.Fail()
		}
	}
}
