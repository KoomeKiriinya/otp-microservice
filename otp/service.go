package otp

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrExpiredOTP = errors.New("OTP expired")
var ErrUmatchingOTP = errors.New("OTP does not Match")

type Service interface {
	GenerateOTP(ctx context.Context, otp_key string) (string, error)
	ValidateOTP(ctx context.Context, otp_key, otp_user_value string) (string, error)
}

type service struct {
	Redis *redis.Client
}

// service required a Redis connection
func NewService(db int) *service {
	return &service{
		Redis: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_DSN"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       db,
		}),
	}
}

// function generate OTP and store it in Redis

func (s *service) GenerateOTP(ctx context.Context, otp_key string) (string, error) {
	otp_value := generateOTP()
	err := s.Redis.Set(ctx, otp_key, otp_value, time.Minute*5).Err()
	if err != nil {
		return "", err
	}
	return otp_value, nil
}

// gets otp stored in Redis and Validates
func (s *service) ValidateOTP(ctx context.Context, otp_key, otp_user_value string) (string, error) {

	otp_stored_value, err := s.Redis.Get(ctx, otp_key).Result()
	if err != nil {
		return "", ErrExpiredOTP
	}
	if otp_stored_value != otp_user_value {
		return "", ErrUmatchingOTP
	} else {
		return "OTP Matched", nil
	}
}

func generateOTP() string {
	// omitted few confusing characters
	charSet := "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"
	pass := randomStringGenerator(charSet, 4) + "-" + randomStringGenerator(charSet, 4)
	return pass
}

func randomStringGenerator(charSet string, codeLength int32) string {
	code := ""
	charSetLength := int32(len(charSet))
	for i := int32(0); i < codeLength; i++ {
		index := randomNumber(0, charSetLength)
		code += string(charSet[index])
	}
	return code
}

func randomNumber(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + int32(rand.Intn(int(max-min)))

}
