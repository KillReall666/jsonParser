package update

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"KillReall666/jsonParser.git/internal/model"
)

type MockStorage struct {
	Repo map[string]model.PortData
}

func (s *MockStorage) SaveJSONData(name string, port model.PortData) {
	s.Repo[name] = port
}

func (s *MockStorage) UpdateJSONData(name string, port model.PortData) {
	s.Repo[name] = port
}

func (s *MockStorage) IsDataInStorage(port string) bool {
	_, ok := s.Repo[port]
	return ok
}

func (s *MockStorage) PrintData() {
	//fmt.Println(s.Repo)
}

func TestHandler_Update(t *testing.T) {
	type testCase struct {
		name         string
		input        string
		expectedCode int
		expectedBody string
	}

	tests := []testCase{
		{
			name: "Valid input data",
			input: `{
    "ZWUTAA": {
    "name": "Mutare",
    "city": "Mutare",
    "country": "Zimbabwe",
    "alias": [],
    "regions": [],
    "coordinates": [
      32.650351,
      -18.9757714
    ],
    "province": "Manicaland",
    "timezone": "Africa/Harare",
    "unlocs": [
      "ZWUTA"
    ]
  }
}`,
			expectedCode: http.StatusOK,
			expectedBody: `New data added successfully`,
		},
		{
			name: "Invalid input data",
			input: `{
    "ZWUTAA": {
    "name": "Mutare",
    "city": "Mutare",
    "country": "Zimbabwe",
    "alias": [],
    "regions": [],`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "json format invalid",
		},
		{
			name:         "Empty input data",
			input:        ``,
			expectedCode: http.StatusBadRequest,
			expectedBody: "json format invalid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := &MockStorage{
				Repo: make(map[string]model.PortData),
			}

			updateBody := strings.NewReader(tc.input)
			req, err := http.NewRequest("POST", "/update", updateBody)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			h := Handler{store}

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.Update(w, r)
			})

			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tc.expectedCode {
				t.Errorf("handler returned wrong status code. got: %v want: %v", status, tc.expectedCode)
			}

			if rr.Body.String() != tc.expectedBody {
				t.Errorf("handler returned unexpected body. got: %v want: %v", rr.Body.String(), tc.expectedBody)
				fmt.Println(rr.Body.String() == tc.expectedBody)
				fmt.Println("body:", rr.Body.String())
				fmt.Println("body:", tc.expectedBody)
			}
		})
	}
}
