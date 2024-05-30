package asteroid_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "ALT-go-challenge/internal/asteroid" 
)

func TestCreateAsteroid(t *testing.T) {
    repo := asteroid.NewRepository()
    handler := asteroid.NewHandler(repo)

    testAsteroid := asteroid.Asteroid{
        ID: "test_id",
        Name: "TestAsteroid",
        Diameter: 100.0, 
        DiscoveryDate: "01-01-2020",
    }
    body, _ := json.Marshal(testAsteroid)

    req, err := http.NewRequest("POST", "/api/v1/asteroides", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler.CreateAsteroid(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    var returnedAsteroid asteroid.Asteroid
    err = json.NewDecoder(rr.Body).Decode(&returnedAsteroid)
    if err != nil {
        t.Fatal(err)
    }

    if returnedAsteroid.Name != testAsteroid.Name {
        t.Errorf("handler returned unexpected body: got %v want %v", returnedAsteroid.Name, testAsteroid.Name)
    }
}

