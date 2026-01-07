package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload_Success(t *testing.T) {
	// Arrange: cria servidor mock
	expectedContent := "col1,col2\nval1,val2"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedContent))
	}))
	defer server.Close()

	// Cria arquivo temporário
	tmpDir := t.TempDir()
	filepath := filepath.Join(tmpDir, "test.csv")

	// Act
	err := Download(server.URL, filepath)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	if string(content) != expectedContent {
		t.Errorf("expected %q, got %q", expectedContent, string(content))
	}
}

func TestDownload_ServerError(t *testing.T) {
	// Arrange: servidor retorna 500
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	filepath := filepath.Join(tmpDir, "test.csv")

	// Act
	err := Download(server.URL, filepath)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDownload_NotFound(t *testing.T) {
	// Arrange: servidor retorna 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	filepath := filepath.Join(tmpDir, "test.csv")

	// Act
	err := Download(server.URL, filepath)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDownload_InvalidURL(t *testing.T) {
	tmpDir := t.TempDir()
	filepath := filepath.Join(tmpDir, "test.csv")

	// Act
	err := Download("http://invalid-url-that-does-not-exist.local", filepath)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDownload_InvalidFilepath(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("content"))
	}))
	defer server.Close()

	// Act: tenta criar arquivo em diretório que não existe
	err := Download(server.URL, "/nonexistent/directory/file.csv")

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

