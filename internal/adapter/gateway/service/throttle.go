package service

import (
	"context"
	"encoding/json"

	"go-app/internal/domain/gateway"
	"go-app/internal/infrastructure/constant"
	"go-app/pkg/errors"
	"go-app/pkg/utils"
)

// throttleService is a struct that represent the throttle's service
type throttleService struct {
	cm gateway.Cache
}

type throttleData struct {
	Count int `json:"count"`
}

// NewThrottleService will create new an throttleService object representation of domain.ThrottleService interface
func NewThrottleService(cm gateway.Cache) gateway.ThrottleService {
	return &throttleService{
		cm: cm,
	}
}

// Blocked is a function to check if the request login is blocked or not
func (svc *throttleService) Blocked(ctx context.Context, key, ip string) (bool, error) {
	hash := utils.MD5Hash(key + ip)
	if b, err := svc.cm.Get(ctx, hash); err == nil {
		data := throttleData{}
		if err := json.Unmarshal(b, &data); err != nil {
			return false, errors.ErrBadRequest.Wrap(err)
		}
		if data.Count >= constant.MaxLoginAttempt {
			return true, nil
		}
	}

	return false, nil
}

// Incr is a function to increment the request login
func (svc *throttleService) Incr(ctx context.Context, key, ip string) error {
	hash := utils.MD5Hash(key + ip)
	data := throttleData{
		Count: 1,
	}
	// Check if the key exist
	if b, err := svc.cm.Get(ctx, hash); err == nil {
		if err := json.Unmarshal(b, &data); err != nil {
			return errors.ErrBadRequest.Wrap(err)
		}
		data.Count++
	}
	// marshal the data
	bytes, err := json.Marshal(&data)
	if err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}
	// increment the value
	if err := svc.cm.Set(ctx, hash, bytes, constant.ThrottleBlockExpireDuration); err != nil {
		return errors.Throw(err)
	}

	return nil
}

// Clear is a function to clear the request login
func (svc *throttleService) Clear(ctx context.Context, key, ip string) error {
	hash := utils.MD5Hash(key + ip)
	if err := svc.cm.Del(ctx, hash); err != nil {
		return errors.Throw(err)
	}

	return nil
}
