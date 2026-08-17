package fuel

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost")

	if client == nil {
		t.Fatal("expected client, got nil")
	}

	if client.baseURL != "http://localhost" {
		t.Errorf(
			"expected baseURL %q, got %q",
			"http://localhost",
			client.baseURL,
		)
	}

	if client.httpClient == nil {
		t.Fatal("expected httpClient, got nil")
	}
}

func TestClient_Get_Success(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf(
					"expected GET, got %s",
					r.Method,
				)
			}

			query := r.URL.Query()

			if query.Get("mode") != "municipality" {
				t.Errorf(
					"expected mode=municipality, got %q",
					query.Get("mode"),
				)
			}

			if query.Get("fuel") != string(Gasoline) {
				t.Errorf(
					"expected fuel=%q, got %q",
					Gasoline,
					query.Get("fuel"),
				)
			}

			if query.Get("municipality") != "Salvador" {
				t.Errorf(
					"expected municipality=Salvador, got %q",
					query.Get("municipality"),
				)
			}

			if query.Get("state") != "BA" {
				t.Errorf(
					"expected state=BA, got %q",
					query.Get("state"),
				)
			}

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			w.WriteHeader(http.StatusOK)

			_, _ = w.Write([]byte(`{
				"municipio": "Salvador",
				"preco_medio": 6.29
			}`))
		}),
	)

	defer server.Close()

	client := NewClient(server.URL)

	price, err := client.Get(
		context.Background(),
		Gasoline,
		"Salvador",
		"BA",
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if price.Fuel != Gasoline {
		t.Errorf(
			"expected fuel %q, got %q",
			Gasoline,
			price.Fuel,
		)
	}

	if price.Price != 6.29 {
		t.Errorf(
			"expected price 6.29, got %.2f",
			price.Price,
		)
	}

	if price.City != "Salvador" {
		t.Errorf(
			"expected city Salvador, got %q",
			price.City,
		)
	}

	if price.State != "BA" {
		t.Errorf(
			"expected state BA, got %q",
			price.State,
		)
	}

	if price.Source != "ANP" {
		t.Errorf(
			"expected source ANP, got %q",
			price.Source,
		)
	}
}

func TestClient_Get_HTTPError(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)
		}),
	)

	defer server.Close()

	client := NewClient(server.URL)

	_, err := client.Get(
		context.Background(),
		Gasoline,
		"Salvador",
		"BA",
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestClient_Get_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			_, _ = w.Write([]byte(`invalid-json`))
		}),
	)

	defer server.Close()

	client := NewClient(server.URL)

	_, err := client.Get(
		context.Background(),
		Gasoline,
		"Salvador",
		"BA",
	)

	if err == nil {
		t.Fatal("expected JSON decoding error, got nil")
	}
}

func TestClient_Get_ContextCanceled(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			<-r.Context().Done()
		}),
	)

	defer server.Close()

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	cancel()

	client := NewClient(server.URL)

	_, err := client.Get(
		ctx,
		Gasoline,
		"Salvador",
		"BA",
	)

	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}
}
