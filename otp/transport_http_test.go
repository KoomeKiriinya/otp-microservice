package otp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

var otp_test_value_3 string

func TestHTTPGenerateOTP(t *testing.T) {
	var logger log.Logger
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)
	svc := NewService(1)
	r := NewHttpServer(svc, logger)
	srv := httptest.NewServer(r)
	body := `{"otp_key":"3"}`
	req, _ := http.NewRequest("POST", srv.URL+"/generateOTP", strings.NewReader(body))
	resp, _ := http.DefaultClient.Do(req)
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read respose body: "+err.Error())
	}
	otp_res := GenerateOTPResponse{}
	json.Unmarshal(resp_body, &otp_res)
	otp_test_value_3 = otp_res.OTP
	assert.Equal(t, resp.StatusCode, 200)
}

func TestHTTPValidateOTP(t *testing.T) {
	var logger log.Logger
	logger = log.With(logger, "listen", "8081", "caller", log.DefaultCaller)
	svc := NewService(1)
	r := NewHttpServer(svc, logger)
	srv := httptest.NewServer(r)
	body := `{"otp_key":"3","user_otp_value":"`+ otp_test_value_3 +`"}`
	req, _ := http.NewRequest("POST", srv.URL+"/validateOTP", strings.NewReader(body))
	resp, _ := http.DefaultClient.Do(req)
	assert.Equal(t, resp.StatusCode, 200)

}