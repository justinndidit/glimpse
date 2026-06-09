package service

import (
	"context"

	"github.com/Adedunmol/glimpse/internal/repository"
	"github.com/Adedunmol/glimpse/internal/server"
	"github.com/clerk/clerk-sdk-go/v2"
)

// internal representation of various clerk Events
const (
	userCreatedEvent = "user.created"
	userDeletedEvent = "user.deleted"
	userUpdatedEvent = "user.updated"
)

type ClerkService struct {
	server   *server.Server
	userRepo repository.UserRepository
}

type ClerkEventPayload struct {
	Data       any    `json:"data"`
	Object     string `json:"object"`
	EventType  string `json:"type"`
	TimeStamp  string `json:"timestamp"`
	InstanceID string `json:"instance_id"`
}

func NewClerkService(srv *server.Server, ur repository.UserRepository) *ClerkService {
	return &ClerkService{
		server:   srv,
		userRepo: ur,
	}
}

func (c *ClerkService) HandleClerkEvents(ctx context.Context, payload ClerkEventPayload) error {

	switch payload.EventType {
	case userCreatedEvent:
		c.server.Logger.Info().Msg("clerk event: user created event")

		userData := payload.Data.(*clerk.User)

		_, err := c.userRepo.CreateUser(ctx, userData.EmailAddresses[0].EmailAddress, userData.ID)
		if err != nil {
			c.server.Logger.Error().Err(err).Msg("failed to create user")
			return err
		}
	case userDeletedEvent:
		c.server.Logger.Info().Msg("clerk event: user deleted event")

		userData := payload.Data.(*clerk.User)

		err := c.userRepo.DeleteUser(ctx, userData.ID)
		if err != nil {
			c.server.Logger.Error().Err(err).Msg("failed to delete user")
			return err
		}
	case userUpdatedEvent:
		c.server.Logger.Info().Msg("clerk event: user updated event")

	default:
		c.server.Logger.Warn().Str("event_type", payload.EventType).Msg("unhandled event sent by clerk")
	}
	return nil
}
