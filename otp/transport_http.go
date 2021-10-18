package otp

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc Service, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
	GenerateOTPHandler := httptransport.NewServer(
		makeGenerateOTPEndpoint(svc),
		DecodeGenerateOTPRequest,
		encodeResponse,
		options...,
	)

	ValidateOTPHandler := httptransport.NewServer(
		makeValidateOTPEndpoint(svc),
		DecodeValidateOTPRequest,
		encodeResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path("/generateOTP").Handler(GenerateOTPHandler)
	r.Methods("POST").Path("/validateOTP").Handler(ValidateOTPHandler)
	return r
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrExpiredOTP:
		return http.StatusGone
	case ErrUmatchingOTP:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}

}

func DecodeGenerateOTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GenerateOTPRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(reqBody, &request); err != nil {
		return nil, err
	}
	if request.Otp_key == "" {
		return nil, errors.New("otp_key not present in request body")
	}
	return request, nil
}

func DecodeValidateOTPRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request ValidateOTPRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(reqBody, &request); err != nil {
		return nil, err
	}

	if request.Otp_key == "" {
		return nil, errors.New("otp_key not present in request body")
	}

	if request.Otp_user_value == "" {
		return nil, errors.New("user_otp_value not present in request body")
	}

	return request, nil
}

//converts the struct returned by the endpoint to a json response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
