package service

import (
	"context"
	"fmt"
)

type GreetingService struct {
}

func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

func (g *GreetingService) SayHello(_ context.Context, u string) (string, error) {
	return fmt.Sprintf("hello, %s", u), nil
}
