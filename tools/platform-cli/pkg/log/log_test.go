package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestConfigureLogger(t *testing.T) {
	tests := []struct {
		name          string
		logLevel      string
		expectedErr   bool
		expectedLevel logrus.Level
	}{
		{
			name:          "Valid Trace",
			logLevel:      "trace",
			expectedErr:   false,
			expectedLevel: logrus.TraceLevel,
		},
		{
			name:          "Valid Debug",
			logLevel:      "debug",
			expectedErr:   false,
			expectedLevel: logrus.DebugLevel,
		},
		{
			name:          "Valid Info",
			logLevel:      "info",
			expectedErr:   false,
			expectedLevel: logrus.InfoLevel,
		},
		{
			name:          "Valid Warn",
			logLevel:      "warn",
			expectedErr:   false,
			expectedLevel: logrus.WarnLevel,
		},
		{
			name:          "Valid Error",
			logLevel:      "error",
			expectedErr:   false,
			expectedLevel: logrus.ErrorLevel,
		},
		{
			name:          "Valid Fatal",
			logLevel:      "fatal",
			expectedErr:   false,
			expectedLevel: logrus.FatalLevel,
		},
		{
			name:          "Invalid Level",
			logLevel:      "invalid",
			expectedErr:   true,
			expectedLevel: logrus.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ConfigureLogger(tt.logLevel)

			if tt.expectedErr && err == nil {
				t.Errorf("Expected an error, but got nil")
			}

			if !tt.expectedErr && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}

			if logrus.GetLevel() != tt.expectedLevel {
				t.Errorf("Expected log level %v, but got %v", tt.expectedLevel, logrus.GetLevel())
			}
		})
	}
}
