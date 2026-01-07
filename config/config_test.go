package config

import (
	"strings"
	"testing"
	"time"
)

func TestQuotesCSVPath_ContainsDate(t *testing.T) {
	// Act
	path := QuotesCSVPath()

	// Assert
	today := time.Now().Format("2006-01-02")
	if !strings.Contains(path, today) {
		t.Errorf("expected path to contain %q, got %q", today, path)
	}
}

func TestQuotesCSVPath_StartsWithTmp(t *testing.T) {
	// Act
	path := QuotesCSVPath()

	// Assert
	if !strings.HasPrefix(path, "/tmp/") {
		t.Errorf("expected path to start with /tmp/, got %q", path)
	}
}

func TestQuotesCSVPath_EndsWithCSV(t *testing.T) {
	// Act
	path := QuotesCSVPath()

	// Assert
	if !strings.HasSuffix(path, ".csv") {
		t.Errorf("expected path to end with .csv, got %q", path)
	}
}

func TestTesouroDiretoAPIUrl_IsValid(t *testing.T) {
	// Act
	url := TesouroDiretoAPIUrl()

	// Assert
	if !strings.HasPrefix(url, "https://") {
		t.Errorf("expected URL to start with https://, got %q", url)
	}

	if !strings.Contains(url, "tesourotransparente.gov.br") {
		t.Errorf("expected URL to contain tesourotransparente.gov.br, got %q", url)
	}
}

