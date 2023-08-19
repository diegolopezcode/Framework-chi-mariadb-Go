package configs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigSuccess(t *testing.T) {
	c := require.New(t)
	port, _ := Config("PORT")
	host, _ := Config("host")
	c.Equal("8080", port)
	c.Equal("localhost", host)
}

func TestConfigError(t *testing.T) {
	c := require.New(t)
	variable, _ := Config("Dummy")
	c.Equal("", variable)
}

// func TestConfigFileError(t *testing.T) {
// 	c := require.New(t)

// 	// Mocking the loading of a non-existent .env file for testing
// 	err := godotenv.Load("../nonexistent.env")
// 	if err == nil {
// 		t.Fatalf("Expected error, got nil")
// 	}

// 	variable, err := Config("Dummy")

// 	c.Equal("", variable)
// 	c.NotNil(err)
// }
