package otp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type GenerateOTPRequest struct {
	Otp_key string `json:"otp_key,omitempty"`
}
type ValidateOTPRequest struct {
	Otp_key        string `json:"otp_key,omitempty"`
	Otp_user_value string `json:"user_otp_value,omitempty"`
}

type GenerateOTPResponse struct {
	OTP string `json:"otp,omitempty"`
	Err string `json:"err,omitempty"`
}

func makeGenerateOTPEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateOTPRequest)
		otp_value, err := svc.GenerateOTP(ctx, req.Otp_key)
		if err != nil {
			return GenerateOTPResponse{"", err.Error()}, err
		}
		return GenerateOTPResponse{otp_value, ""}, err

	}
}

func makeValidateOTPEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ValidateOTPRequest)
		status, err := svc.ValidateOTP(ctx, req.Otp_key, req.Otp_user_value)
		if err != nil {
			return GenerateOTPResponse{"", err.Error()}, err
		}
		return GenerateOTPResponse{status, ""}, err

	}
}
