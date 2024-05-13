package log

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(string) {
}

func (m *MockLogger) Info(string) {
}

func (m *MockLogger) Warning(string) {
}

func (m *MockLogger) Error(string) {
}
