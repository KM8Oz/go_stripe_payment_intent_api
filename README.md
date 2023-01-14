# go_stripe_payment_intent_api

This is a Go library for interacting with the Stripe Payment Intent API. With this library, you can easily create, retrieve, update, and cancel Payment Intents.

## Installation

To install the library, simply use `go get`:

`go get github.com/KM8Oz/go_stripe_payment_intent_api`


## Usage

First, import the library into your Go project:

```go
import "github.com/KM8Oz/go_stripe_payment_intent_api"


Before using the library, you'll need to set your Stripe API key. You can do this by calling stripe.Key = "YOUR_SECRET_KEY".
## Creating a Payment Intent

To create a Payment Intent, you can use the Create function. This function takes a CreateParams struct as its only argument, which can be used to set various options for the Payment Intent.


```
params := &stripe.PaymentIntentCreateParams{
  Amount:   stripe.Int64(1000),
  Currency: stripe.String(string(stripe.CurrencyUSD)),
}

intent, _ := paymentintent.Create(params)
```

