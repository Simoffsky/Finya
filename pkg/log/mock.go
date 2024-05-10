package log

type MockLogger struct {
}

func (m *MockLogger) Debug(string) {
}

func (m *MockLogger) Info(string) {
}

func (m *MockLogger) Warning(string) {
}

func (m *MockLogger) Error(string) {
}
