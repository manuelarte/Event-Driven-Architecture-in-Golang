package application

import (
	"context"

	"eda-in-golang/internal/ddd"
)

type DomainEventHandlers interface {
	OnStoreCreated(ctx context.Context, event ddd.Event) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error
	OnProductAdded(ctx context.Context, event ddd.Event) error
	OnProductRemoved(ctx context.Context, event ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = new(ignoreUnimplementedDomainEvents)

func (ignoreUnimplementedDomainEvents) OnStoreCreated(_ context.Context, _ ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationEnabled(_ context.Context, _ ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationDisabled(_ context.Context, _ ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductAdded(_ context.Context, _ ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductRemoved(_ context.Context, _ ddd.Event) error {
	return nil
}
