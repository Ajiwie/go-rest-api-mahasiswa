package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Ajiwie/go-rest-api-mahasiswa/config"
	"github.com/Ajiwie/go-rest-api-mahasiswa/helper"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/controller"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/middleware"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/model"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/repository"
	"github.com/Ajiwie/go-rest-api-mahasiswa/internal/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	dsn := "root:Mysqlajiwicaksono2005@tcp(localhost:3306)/database_mahasiswa_dosen?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepository, db, validate)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaService)
	router := config.NewRouter(mahasiswaController)

	return middleware.NewAuthMiddleware(router)
}

func truncateMahasiswa(db *sql.DB) {
	db.Exec("TRUNCATE mahasiswa")
}

func TestCreateMahasiswaSucces(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"nama" : "Aji" , "nim" :"Ada" , "jurusan" : "Informatika"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/mahasiswa", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Aji", responseBody["data"].(map[string]interface{})["nama"])
	assert.Equal(t, "Ada", responseBody["data"].(map[string]interface{})["nim"])
	assert.Equal(t, "Informatika", responseBody["data"].(map[string]interface{})["jurusan"])

}

func TestCreateMahasiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"nama" : "" , "nim" :"" , "jurusan" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/mahasiswa", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateMahasiswaSucces(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)

	tx, _ := db.Begin()
	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswa := mahasiswaRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama: "Galih",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"nama" : "Aji"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/mahasiswa/"+strconv.Itoa(mahasiswa.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	code, ok := responseBody["code"].(float64)
	if !ok {
		t.Fatalf("Response body tidak memiliki field 'code': %v", responseBody)
	}
	assert.Equal(t, 200, int(code))
	assert.Equal(t, "OK", responseBody["status"])
	data, ok := responseBody["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Response body tidak memiliki field 'data': %v", responseBody)
	}
	assert.Equal(t, mahasiswa.Id, int(data["id"].(float64)))
	assert.Equal(t, "Aji", data["nama"])

}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewMahasiswaRepository()
	category := categoryRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"nama" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/mahasiswa/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewMahasiswaRepository()
	mahasiswa := categoryRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama:    "Ajiewi",
		NIM:     "1111",
		Jurusan: "Informatika",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/mahasiswa/"+strconv.Itoa(mahasiswa.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	data, ok := responseBody["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("Response body tidak memiliki field 'data': %v", responseBody)
	}
	assert.Equal(t, mahasiswa.Id, int(data["id"].(float64)))
	assert.Equal(t, mahasiswa.Nama, data["nama"])
	assert.Equal(t, mahasiswa.NIM, data["nim"])
	assert.Equal(t, mahasiswa.Jurusan, data["jurusan"])

}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/mahasiswa/400", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	code, exists := responseBody["code"]
	if !exists {
		t.Fatal("Field 'code' is missing in response")
	}

	codeFloat, ok := code.(float64)
	if !ok {
		t.Fatal("Field 'code' is not a float64")
	}
	assert.Equal(t, 404, int(codeFloat))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteMahasiswaSucces(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)

	tx, _ := db.Begin()
	mahasiswaRepository := repository.NewMahasiswaRepository()
	mahasiswa := mahasiswaRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama: "Aji",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/mahasiswa/"+strconv.Itoa(mahasiswa.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	code, ok := responseBody["code"].(float64)
	if !ok {
		t.Fatalf("Response body tidak memiliki field 'code': %v", responseBody)
	}
	assert.Equal(t, 200, int(code))
	assert.Equal(t, "OK", responseBody["status"])

}

func TestDeleteMahasiswaFailed(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/mahasiswa/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	code, ok := responseBody["code"].(float64)
	if !ok {
		t.Fatal("Field 'code' is missing in response")
	}
	assert.Equal(t, 404, int(code))

	status, ok := responseBody["status"].(string)
	if !ok {
		t.Fatal("Field 'status' is missing in response")
	}
	assert.Equal(t, "NOT FOUND", status)
}

func TestListMahasiswaSuccess(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewMahasiswaRepository()
	mahasiswa1 := categoryRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama:    "Ajiewi",
		NIM:     "1111",
		Jurusan: "Informatika",
	})
	mahasiswa2 := categoryRepository.Create(context.Background(), tx, model.Mahasiswa{
		Nama:    "Ajie",
		NIM:     "2222",
		Jurusan: "Multimedia",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/mahasiswa", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(body))

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var mahasiswaa = responseBody["data"].([]interface{})

	mahasiswaResponse1 := mahasiswaa[0].(map[string]interface{})
	mahasiswaResponse2 := mahasiswaa[1].(map[string]interface{})

	assert.Equal(t, mahasiswa1.Id, int(mahasiswaResponse1["id"].(float64)))
	assert.Equal(t, mahasiswa1.Nama, mahasiswaResponse1["nama"])
	assert.Equal(t, mahasiswa1.NIM, mahasiswaResponse1["nim"])
	assert.Equal(t, mahasiswa1.Jurusan, mahasiswaResponse1["jurusan"])

	assert.Equal(t, mahasiswa2.Id, int(mahasiswaResponse2["id"].(float64)))
	assert.Equal(t, mahasiswa2.Nama, mahasiswaResponse2["nama"])
	assert.Equal(t, mahasiswa2.NIM, mahasiswaResponse2["nim"])
	assert.Equal(t, mahasiswa2.Jurusan, mahasiswaResponse2["jurusan"])

}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateMahasiswa(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/mahasiswa", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
