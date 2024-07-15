package tests

import (
	transportHTTP "MailBeacon/internal/transport/http"
	"MailBeacon/internal/types"
	"MailBeacon/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	token = utils.GenerateRandomToken()
)

// Store
type MockStore struct {
	mock.Mock
}

func (m *MockStore) AddUser(ctx context.Context, user types.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockStore) ConfirmUser(ctx context.Context, user types.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// PubSub
type MockPubSub struct {
	mock.Mock
}

func (m *MockPubSub) Publish(topicId string, payload any) error {
	args := m.Called(topicId, payload)
	return args.Error(0)
}

// Mailer
type MockMailer struct {
	mock.Mock
}

func (m *MockMailer) SendConfirmationEmail(email string, token string) error {
	args := m.Called(email, token)
	return args.Error(0)
}

func (m *MockMailer) SendWelcomeEmail(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

type MockService struct {
	mock.Mock
}

func (m *MockService) SignUp(ctx context.Context, user types.User) error {
	user.Token = token
	args := m.Called(ctx, user)
	return args.Error(0)

}

func (m *MockService) SendConfirmationEmail(ctx context.Context, user types.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockService) ConfirmSubscription(ctx context.Context, user types.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)

}

func (m *MockService) SendWelcomeEmail(ctx context.Context, user types.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)

}

func TestNewsletterSignup(t *testing.T) {
	mockService := new(MockService)
	handler := transportHTTP.Handler{Service: mockService}

	tests := []struct {
		name           string
		email          string
		mockServiceErr error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Kennedy Karoko",
			email:          "test@example.com",
			mockServiceErr: nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "Thank you for signing up! Please check your email to confirm your subscription.",
		},
		// {
		// 	name:           "invalid email",
		// 	email:          "",
		// 	mockServiceErr: nil,
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedBody:   "Invalid email",
		// },
		// {
		// 	name:           "service error",
		// 	email:          "test@example.com",
		// 	mockServiceErr: errors.New("internal error"),
		// 	expectedStatus: http.StatusInternalServerError,
		// 	expectedBody:   "internal error",
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body := bytes.NewBufferString("email=" + tc.email)
			req := httptest.NewRequest("POST", "/api/v1/newsletter/signup", body)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()

			if tc.email != "" {
				user := types.User{Email: tc.email, Token: token}
				mockService.On("SignUp", mock.Anything, user).Return(tc.mockServiceErr)
			}

			handler.NewsletterSignup(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			if tc.expectedBody != "" {
				buffer := new(bytes.Buffer)
				buffer.ReadFrom(res.Body)
				bodyString := buffer.String()
				assert.Contains(t, bodyString, tc.expectedBody)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestSendConfirmationEmail(t *testing.T) {
	mockService := new(MockService)
	handler := transportHTTP.Handler{Service: mockService}

	tests := []struct {
		name           string
		payload        string
		mockServiceErr error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid payload",
			payload:        `{"message": {"data": "eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ=="}}`,
			mockServiceErr: nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "Confirmation email sent!",
		},
		// {
		// 	name:           "invalid payload",
		// 	payload:        `{"message": {"data": "invalid"}}`,
		// 	mockServiceErr: nil,
		// 	expectedStatus: http.StatusBadRequest,
		// 	expectedBody:   "Invalid request body",
		// },
		// {
		// 	name:           "service error",
		// 	payload:        `{"message": {"data": "eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ=="}}`,
		// 	mockServiceErr: errors.New("internal error"),
		// 	expectedStatus: http.StatusInternalServerError,
		// 	expectedBody:   "internal error",
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body := bytes.NewBufferString(tc.payload)
			req := httptest.NewRequest("POST", "/api/v1/newsletter/send-confirmation-email", body)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if tc.payload != "" {
				parsedPayload := types.SendConfirmationEmailRequest{}
				json.Unmarshal([]byte(tc.payload), &parsedPayload)
				user := types.ConvertSendConfirmationEmailRequestToUser(parsedPayload)
				user.Email = "test@example.com"
				mockService.On("SendConfirmationEmail", mock.Anything, user).Return(tc.mockServiceErr)
			}

			handler.SendConfirmationEmail(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			if tc.expectedBody != "" {
				buffer := new(bytes.Buffer)
				buffer.ReadFrom(res.Body)
				bodyString := buffer.String()
				assert.Contains(t, bodyString, tc.expectedBody)
			}

			mockService.AssertExpectations(t)

		})
	}
}

func TestConfirmNewsletterSignup(t *testing.T) {
	mockService := new(MockService)
	handler := transportHTTP.Handler{Service: mockService}

	tests := []struct {
		name           string
		token          string
		email          string
		mockServiceErr error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid token and email",
			token:          token,
			email:          "test@example.com",
			mockServiceErr: nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "You have successfully confirmed your email.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/confirm-email?token="+tc.token+"&email="+tc.email, nil)
			w := httptest.NewRecorder()

			if tc.token != "" && tc.email != "" {
				user := types.User{Email: tc.email, Token: tc.token}
				mockService.On("ConfirmSubscription", mock.Anything, user).Return(tc.mockServiceErr)
			}

			handler.ConfirmNewsletterSignup(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			if tc.expectedBody != "" {
				buffer := new(bytes.Buffer)
				buffer.ReadFrom(res.Body)
				bodyString := buffer.String()
				assert.Contains(t, bodyString, tc.expectedBody)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestSendWelcomeEmail(t *testing.T) {
	mockService := new(MockService)
	handler := transportHTTP.Handler{Service: mockService}

	// encodedEmail := base64.StdEncoding.EncodeToString([]byte("test@example.com"))

	tsts := []struct {
		name           string
		payload        string
		mockServiceErr error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid payload",
			payload:        `{"message": {"data": "eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ=="}}`,
			mockServiceErr: nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome email sent!",
		},
	}

	for _, tc := range tsts {
		t.Run(tc.name, func(t *testing.T) {
			body := bytes.NewBufferString(tc.payload)
			req := httptest.NewRequest("POST", "/api/v1/newsletter/send-welcome-email", body)
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			if tc.payload != "" {

				parsedPayload := types.SendWelcomeEmailRequest{}
				json.Unmarshal([]byte(tc.payload), &parsedPayload)
				user := types.ConvertSendWelcomeEmailRequestToUser(parsedPayload)
				user.Email = "test@example.com"
				mockService.On("SendWelcomeEmail", mock.Anything, user).Return(tc.mockServiceErr)
			}

			handler.SendWelcomeEmail(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tc.expectedStatus, res.StatusCode)
			if tc.expectedBody != "" {
				buffer := new(bytes.Buffer)
				buffer.ReadFrom(res.Body)
				bodyString := buffer.String()
				assert.Contains(t, bodyString, tc.expectedBody)
			}

			mockService.AssertExpectations(t)

		})

	}

}
