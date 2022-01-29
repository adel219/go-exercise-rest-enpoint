package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostDroneLocation(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(handleDroneLocation))
	defer s.Close()

	endpoint := s.URL + "/drone/location"
	// Happy flow
	{
		var req = DroneLocationRequest{
			SectorID: 1.0,
			X:        123.12,
			Y:        456.56,
			Z:        789.89,
			Vel:      20.0,
		}
		reqJSON, err := json.Marshal(req)
		assert.NoError(t, err)

		r, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(reqJSON))
		assert.NoError(t, err)

		resp, err := s.Client().Do(r)
		assert.NoError(t, err)

		var dresp DroneLocationResponse
		err = json.NewDecoder(resp.Body).Decode(&dresp)
		assert.NoError(t, err)

		assert.Equal(t, resp.StatusCode, http.StatusOK)
		assert.Equal(t, 1389.57, dresp.Loc)
	}

	// Invalid method
	{
		r, err := http.NewRequest(http.MethodGet, endpoint, strings.NewReader("{}"))
		assert.NoError(t, err)

		resp, err := s.Client().Do(r)
		assert.NoError(t, err)

		assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)
	}

	// Invalid body format
	{
		r, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader("{invalid json}"))
		assert.NoError(t, err)

		resp, err := s.Client().Do(r)
		assert.NoError(t, err)

		b, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		assert.Equal(t, "invalid character 'i' looking for beginning of object key string\n", string(b))
	}

	// Invalid sector ID
	{
		r, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader("{}"))
		assert.NoError(t, err)

		resp, err := s.Client().Do(r)
		assert.NoError(t, err)

		b, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
		assert.Equal(t, "sectorId must be greater than zero\n", string(b))
	}
}
