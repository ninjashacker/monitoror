package service

import (
	"testing"

	"github.com/monitoror/monitoror/config"

	"github.com/stretchr/testify/assert"
)

// This tests are realy basic and juste check if Server member are not nil

func TestInitEcho(t *testing.T) {
	server := &Server{}
	server.initEcho()

	assert.NotNil(t, server.Echo)
}

func TestInitMiddleware(t *testing.T) {
	conf, _ := config.InitConfig()
	server := &Server{config: conf}
	server.initEcho()
	server.initMiddleware()

	assert.NotNil(t, server.cm)
}

func TestRegister(t *testing.T) {
	register("test", true, func() {
		assert.True(t, true)
	})

	register("test", false, func() {
		assert.True(t, false)
	})
}