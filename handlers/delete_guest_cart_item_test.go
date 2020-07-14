package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/guest"
	"e-food/pkg/dao"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CartHandlerMock struct{}

func (c *CartHandlerMock) RemoveItemFromGuestCart(db *sql.DB, productId int64, sessionId string) error {
	return nil
}

func (c *CartHandlerMock) GetGuestCart(db *sql.DB, sessionId string) (models.CartPreview, error) {
	return nil, nil
}
func (c *CartHandlerMock) AddItemToGuestCart(db *sql.DB, prodHandler dao.ProductHandler, sessionId string, totalQty, productId int64) (*models.CartSuccessResponse, error) {
	return nil, nil
}
func (c *CartHandlerMock) DeleteExistingGuestCartItemIfAny(db *sql.DB, sessionId string, productId int64) error {
	return nil
}

func (c *CartHandlerMock) InsertItemInGuestCart(db *sql.DB, totalQty, productId int64, sessionId string) error {
	return nil
}
func (c *CartHandlerMock) EmptyGuestCartItem(db *sql.DB, sessionId string) error {
	return nil
}

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
	guestCartHandlerMock := CartHandlerMock{}
	service := NewGuestCartRemoveItemHandler(db, &guestCartHandlerMock)
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
	guestCartHandlerMock := CartHandlerMock{}
	service := NewGuestCartRemoveItemHandler(db, &guestCartHandlerMock)
	actualResponse := service.Handle(*params)
	require.IsType(t, &guest.RemoveItemOK{}, actualResponse)
	responseOK := actualResponse.(*guest.RemoveItemOK)
	require.Equal(t, expectedPayload, responseOK.Payload)
}
