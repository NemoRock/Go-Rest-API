package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_handleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}

/*type fields struct {
		config *Config
		logger *logrus.Logger
		router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &APIServer{
				config: tt.fields.config,
				logger: tt.fields.logger,
				router: tt.fields.router,
			}
			if got := s.handleHello(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

*/
