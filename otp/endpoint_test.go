package otp

import (
	"context"
	"testing"
)

var otp_test_value_2 string
func TestMakeGenerateOTPEndpoint(t *testing.T) {
	svc := NewService(1)
	endpoint := makeGenerateOTPEndpoint(svc)
	t.Run("Generate OTP Endpoint", func(t *testing.T) {
		req := GenerateOTPRequest{
			Otp_key: "4",
		}
		res, err := endpoint(context.Background(), req)
		otp_test_value_2 = res.(GenerateOTPResponse).OTP
		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}
	})
}
func TestValidateOTPEndpoint(t *testing.T) {
	svc := NewService(1)
	endpoint := makeValidateOTPEndpoint(svc)
	t.Run("Validate OTP Endpoint", func(t *testing.T) {
		req := ValidateOTPRequest{
			Otp_key: "4",
			Otp_user_value: otp_test_value_2,
		}
		_, err := endpoint(context.Background(), req)
		if err != nil {
			t.Errorf("expected %v received %v", nil, err)
		}
	})
}
