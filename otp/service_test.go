package otp_test

import (
	"context"
	"github.com/otpservice/otp"
	"testing"
	"github.com/stretchr/testify/assert"
)

var otp_test_value string
func TestGenerateOTP(t *testing.T){
	svc := otp.NewService(1)
	t.Run("Generate OTP",func(t *testing.T){
		otp, err := svc.GenerateOTP(context.Background(),"3")
		otp_test_value = otp
		assert.Nil(t,err)
		assert.Equal(t,len(otp),9)
	})
}

func TestValidateOTP(t *testing.T){
	svc := otp.NewService(1)
	t.Run("Validate OTP",func(t *testing.T){
		res, err := svc.ValidateOTP(context.Background(),"3",otp_test_value)
		assert.Nil(t,err)
		assert.Equal(t,res,"OTP Matched")
	})
}