package logger

import (
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_stringToLvl(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		testCases := map[string]log.Lvl{
			"debug":             log.DEBUG,
			"DEbUg":             log.DEBUG,
			"   \ndebug  \t\t ": log.DEBUG,
			"info":              log.INFO,
			"warn":              log.WARN,
			"error":             log.ERROR,
			"off":               log.OFF,
		}

		for given, expected := range testCases {
			t.Run(given, func(t *testing.T) {
				actual, err := stringToLvl(given)
				require.NoError(t, err)
				require.Equal(t, expected, actual)
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		_, err := stringToLvl("unknown")
		require.EqualError(t, err, `unknown log level "unknown"`)
	})
}
