package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func setenvs(t *testing.T, args ...string) {
	if len(args)%2 != 0 {
		t.Error("count of args for setenvs should be even")
		t.FailNow()
	}

	for i := 0; i < len(args); i += 2 {
		err := os.Setenv(args[i], args[i+1])
		require.NoErrorf(t, err, "cannot set env variable %q with value %q", args[i], args[i+1])
	}
}

func TestCollect(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("without optional envs", func(t *testing.T) {
			setenvs(
				t,
				envDBDriver, "postgres",
				envDBURI, "postgres://",
			)

			cfg, err := Collect()
			require.NoError(t, err)

			assert.Equal(t, "0.0.0.0:4000", cfg.Address)
			assert.Equal(t, time.Minute, cfg.StopTimeout)
			assert.Equal(t, "info", cfg.LogLevel)
			assert.Equal(t, "postgres", cfg.DBDriver)
			assert.Equal(t, "postgres://", cfg.DBURI)
		})

		t.Run("with all envs", func(t *testing.T) {
			setenvs(
				t,
				envKeyIP, "192.0.2.1",
				envKeyPort, "9999",
				envKeyStopTimeout, "30s",
				envLogLevel, "off",
				envDBDriver, "postgres",
				envDBURI, "postgres://",
			)

			cfg, err := Collect()
			require.NoError(t, err)

			assert.Equal(t, "192.0.2.1:9999", cfg.Address)
			assert.Equal(t, 30*time.Second, cfg.StopTimeout)
			assert.Equal(t, "off", cfg.LogLevel)
			assert.Equal(t, "postgres", cfg.DBDriver)
			assert.Equal(t, "postgres://", cfg.DBURI)
		})
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("invalid timeout", func(t *testing.T) {
			setenvs(
				t,
				envKeyStopTimeout, "invalid",
			)

			_, err := Collect()
			require.EqualError(t, err, `invalid stop timeout format: time: invalid duration "invalid"`)
		})

		t.Run("without DB driver", func(t *testing.T) {
			_, err := Collect()
			require.EqualError(t, err, `invalid stop timeout format: time: invalid duration "invalid"`)
		})

		t.Run("without DB URI", func(t *testing.T) {
			setenvs(
				t,
				envDBDriver, "postgres",
			)

			_, err := Collect()
			require.EqualError(t, err, `invalid stop timeout format: time: invalid duration "invalid"`)
		})
	})
}
