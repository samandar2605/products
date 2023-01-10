package v1_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/bxcodec/faker/v4"
// 	"github.com/golang/mock/gomock"
// 	"github.com/samandar2605/products/storage/repo"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateUser(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	strg := mockdb.NewMockStorageI(ctrl)

// 	reqBody := storage.User{
// 		FirstName:   faker.FirstName(),
// 		LastName:    faker.LastName(),
// 		Email:       faker.Email(),
// 		PhoneNumber: faker.Phonenumber(),
// 	}

// 	resp := reqBody
// 	resp.ID = 1

// 	strg.EXPECT().
// 		CreateUser(&reqBody).
// 		Times(1).Return(&resp, nil)

// 	server := NewServer(strg)

// 	rec := httptest.NewRecorder()

// 	payload, err := json.Marshal(reqBody)
// 	assert.NoError(t, err)
// 	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

// 	server.Router.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusCreated, rec.Code)
// }

// func TestCreateProductTable(t *testing.T) {
// 	product := repo.Product{
// 		Id:          1,
// 		Name:        "noutbook",
// 		Sku:         "1 23 456 789",
// 		Description: "It is nice",
// 		Price:       500.99,
// 		Count:       2,
// 		CreatedAt:   time.Now(),
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(strg *mockdb.MockStorageI)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "success case",
// 			buildStubs: func(strg *mockdb.MockStorageI) {
// 				strg.EXPECT().
// 					CreateUser(&user).
// 					Times(1).Return(&user, nil)
// 			},
// 			checkResponse: func(t *testing.T, rec *httptest.ResponseRecorder) {
// 				assert.Equal(t, http.StatusCreated, rec.Code)
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			strg := mockdb.NewMockStorageI(ctrl)
// 			tc.buildStubs(strg)

// 			server := NewServer(strg)

// 			rec := httptest.NewRecorder()

// 			payload, err := json.Marshal(user)
// 			assert.NoError(t, err)
// 			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))

// 			server.Router.ServeHTTP(rec, req)
// 			tc.checkResponse(t, rec)
// 		})
// 	}
// }
