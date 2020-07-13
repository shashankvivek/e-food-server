package handlers

import (
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGuestCartRemoveItemHandler_WithoutCookie(t *testing.T) {
	expectedPayload := "error with cookie"
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: "random_key", Value: "AScolkZ3ZZZXX"})
	params := &guest.RemoveItemParams{
		HTTPRequest: &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}},
		ProductID:   123,
	}

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	service := NewGuestCartRemoveItemHandler(db)
	actualResponse := service.Handle(*params)
	require.IsType(t, &guest.RemoveItemInternalServerError{}, actualResponse)
	responseOK := actualResponse.(*guest.RemoveItemInternalServerError)
	require.Equal(t, expectedPayload, responseOK.Payload)
}

func TestNewGuestCartRemoveItemHandler_WithCookie(t *testing.T) {
	expectedPayload := &models.SuccessResponse{Success: true}
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: "guest_session", Value: "AScolkZ3ZZZXX"})
	params := &guest.RemoveItemParams{
		HTTPRequest: &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}},
		ProductID:   123,
	}

	db, _, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	service := NewGuestCartRemoveItemHandler(db)
	actualResponse := service.Handle(*params)
	require.IsType(t, &guest.RemoveItemOK{}, actualResponse)
	responseOK := actualResponse.(*guest.RemoveItemOK)
	require.Equal(t, expectedPayload, responseOK.Payload)
}
