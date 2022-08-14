package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initCache(t *testing.T, execute func(t *testing.T, c *cache)) {
	c := newCache()

	execute(t, c)
}

func TestOperationsProcess(t *testing.T) {
	procAdd := func(t *testing.T, c *cache) {
		cfg := cfgProc{
			id:          generatorProcID(),
			etaMilisecs: 10,
			c:           c,
		}

		p := newProcess(cfg)

		time.Sleep(200 * time.Millisecond)

		state, errGet := c.getProcessState(p.id)
		require.NoError(t, errGet, "get process")
		require.NotNil(t, state)
		require.Equal(t, p.procState, *state)
		assert.True(t, c.content[p.id].isReady)

		retryState, errRetry := c.getProcessState(p.id)
		require.NoError(t, errRetry, "retry get process")
		require.Equal(t, p.procState, *retryState)

		time.Sleep(200 * time.Millisecond)
		require.Equal(t, 0, len(c.registry), "registry not cleaned")
	}

	initCache(t, procAdd)
}
