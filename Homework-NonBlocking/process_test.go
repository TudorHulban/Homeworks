package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProcess(t *testing.T) {
	newProc := func(t *testing.T, c *cache) {
		p := newProcess(cfgProc{
			id:          generatorProcID(),
			etaMilisecs: 500,
			c:           c,
		})

		require.True(t, p.isReady)
	}

	initCache(t, newProc)
}
