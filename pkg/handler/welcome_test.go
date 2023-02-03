package handler_test

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestHandler_welcome(t *testing.T) {
	testTable := []struct {
		name                 string
		inputBody            string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "OK",
			inputBody:            `{"message":"dummy input body"}`,
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"dummy input body"}`,
		},
		{
			name:                 "ERROR",
			inputBody:            `{"message":"dummy input body"}`,
			expectedStatusCode:   429,
			expectedResponseBody: `{"status": 429,"message":"too many requests"}`,
		},
	}

	for idx, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init
			// Test server
			r := gin.New()
			r.POST("/api", handler.Welcome)

			// Testing
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api",
				bytes.NewBufferString(`{"message":"dummy input body"}`))

			// Make request
			var wg sync.WaitGroup
			ln := 5
			if idx == 1 {
				ln = 1
			}

			for i := 0; i < ln; i++ {
				go func() {
					wg.Add(1)
					r.ServeHTTP(w, req)
					// Assert check if equal
					fmt.Printf("Case %d -> Expected status code: %d -> Handled status code: %d \n", idx, testCase.expectedStatusCode, w.Code)
					assert.Equal(t, testCase.expectedStatusCode, w.Code)
					assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
					wg.Done()
				}()
			}
			wg.Wait()
		})
	}
}
