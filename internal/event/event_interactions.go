package event

import (
	"context"
	"fmt"
	"snorp/internal/models"
)

func InteractionHandler(ctx context.Context, interaction models.Interaction) {
	switch interaction.Type {

	default:
		fmt.Printf("%+v\n", interaction)
	}
}
