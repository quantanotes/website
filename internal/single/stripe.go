package single

import (
	"quanta/internal/globals"

	"github.com/stripe/stripe-go/v78/client"
)

var (
	Stripe = &client.API{}
)

func init() {
	Stripe.Init(globals.StripeSecretKey, nil)
}
