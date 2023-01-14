package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type PaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Type     string `json:"type"`
}

func main() {
	stripe.Key = os.Getenv("SECRET_KEY")

	http.HandleFunc("/create_payment", createPayment)

	fmt.Println("Server started on port 3000")
	http.ListenAndServe(":3000", nil)
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Check the token type
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Extract the token value
	token := parts[1]
	if token != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// The rest of the createPayment function remains the same
	var payment PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the Payment Intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(payment.Amount),
		Currency: stripe.String(payment.Currency),
		PaymentMethodTypes: stripe.StringSlice([]string{
			payment.Type,
		}),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(intent)
}
