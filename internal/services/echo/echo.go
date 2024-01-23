package echo

import (
	"context"
	"time"

	"github.com/namkatcedrickjumtock/go_layered_arch_starter_kit/internal/models"
	"github.com/namkatcedrickjumtock/go_layered_arch_starter_kit/internal/persistence"
)

//go:generate mockgen -source ./echo.go -destination mocks/echo.mock.go -package mocks

// Echoer exaple interface of a service that echos a message with the current time.
type Echoer interface {
	Echo(ctx context.Context, msg string) (*models.EchoResponse, error)
}

// Echoer concrete implementation of Echoer.
type EchoerImpl struct {
	repo persistence.Repository
}

var _ Echoer = &EchoerImpl{}

func NewEchoer(repo persistence.Repository) (*EchoerImpl, error) {
	return &EchoerImpl{
		repo: repo,
	}, nil
}

// Echo implements Echoer.
func (e *EchoerImpl) Echo(ctx context.Context, msg string) (*models.EchoResponse, error) {
	currentTime, err := e.repo.GetTime(ctx)
	if err != nil {
		return nil, err
	}

	return &models.EchoResponse{
		Message:    msg,
		Timestramp: currentTime.Format(time.RFC3339),
	}, nil
}
