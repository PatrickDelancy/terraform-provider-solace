package main

import (
	"fmt"

	"github.com/ExalDraen/semp-client/models"
)

// formatError returns a useful string representation of the SempError provided
func formatError(err *models.SempError) string {
	return fmt.Sprintf("[%v]: %v - %v", *err.Code, *err.Description, *err.Status)
}
