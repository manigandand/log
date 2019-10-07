package log

import "testing"

func TestInitWithoutMultiHandle(t *testing.T) {
	// TODO: write test for logLevel check. If FatalLevel is set other handlers should have iouti.Discard writer
	// currently noway to get the writer for any logger via standard log package.

	Init(InfoLevel, nil)
	if infoLogger.Prefix() != "INFO: " {
		t.Error("Error in infologger initialize")
	}

	if warningLogger.Prefix() != "WARNING: " {
		t.Error("Error in  warninglogger initialize")
	}

	if errorLogger.Prefix() != "ERROR: " {
		t.Error("Error in errorlogger initialize")
	}

	if fatalLogger.Prefix() != "FATAL: " {
		t.Error("Error in errorlogger initialize")
	}
}
