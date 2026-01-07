package tesouro

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCSVUrl_Success(t *testing.T) {
	// Arrange: servidor retorna JSON válido com recurso CSV
	jsonResponse := `{
		"result": {
			"resources": [
				{"format": "PDF", "url": "http://example.com/file.pdf"},
				{"format": "CSV", "url": "http://example.com/file.csv"}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	// Act
	url, err := GetCSVUrl(server.URL)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := "http://example.com/file.csv"
	if url != expected {
		t.Errorf("expected %q, got %q", expected, url)
	}
}

func TestGetCSVUrl_NoCSVResource(t *testing.T) {
	// Arrange: JSON sem recurso CSV
	jsonResponse := `{
		"result": {
			"resources": [
				{"format": "PDF", "url": "http://example.com/file.pdf"},
				{"format": "XLS", "url": "http://example.com/file.xls"}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	// Act
	_, err := GetCSVUrl(server.URL)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetCSVUrl_EmptyResources(t *testing.T) {
	// Arrange: JSON com resources vazio
	jsonResponse := `{
		"result": {
			"resources": []
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	// Act
	_, err := GetCSVUrl(server.URL)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetCSVUrl_ServerError(t *testing.T) {
	// Arrange: servidor retorna 500
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Act
	_, err := GetCSVUrl(server.URL)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetCSVUrl_InvalidJSON(t *testing.T) {
	// Arrange: servidor retorna JSON inválido
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	// Act
	_, err := GetCSVUrl(server.URL)

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetCSVUrl_InvalidURL(t *testing.T) {
	// Act
	_, err := GetCSVUrl("http://invalid-url-that-does-not-exist.local")

	// Assert
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetCSVUrl_FirstCSVIsReturned(t *testing.T) {
	// Arrange: múltiplos CSVs, deve retornar o primeiro
	jsonResponse := `{
		"result": {
			"resources": [
				{"format": "CSV", "url": "http://example.com/first.csv"},
				{"format": "CSV", "url": "http://example.com/second.csv"}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	// Act
	url, err := GetCSVUrl(server.URL)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := "http://example.com/first.csv"
	if url != expected {
		t.Errorf("expected %q, got %q", expected, url)
	}
}

