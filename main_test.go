package pocketworks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitialSetupAndStart(t *testing.T) {
   app, err := InitialSetup()
   assert.NoError(t, err)
   assert.NotNil(t, app)

   // Start PocketBase in goroutine
   err = app.Start()
   assert.NoError(t, err)

   // Give it time to start up
   time.Sleep(200 * time.Millisecond)
}

func TestSendEmail(t *testing.T) {
   app, _ := InitialSetup()

   err := app.SendEmail("test@example.com", "Test Subject", "<h1>Test Body</h1>")
   assert.NoError(t, err)
}
