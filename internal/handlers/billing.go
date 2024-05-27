package handlers

import (
	"math"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
	"quanta/internal/single"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/stripe/stripe-go/v78"
)

const (
	intervalMonthly = "monthly"
	intervalAnnual  = "annual"
)

func checkout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(string)
	redirect := chi.URLParam(r, "redirect")

	customer, err := model.GetOrCreateCustomerId(r.Context(), user)
	if err != nil {
		errorInternalResponse(w, err)
		return
	}

	var req struct {
		Interval string
		Amount   int
	}
	if err := jsonRequest(w, r, &req); err != nil {
		return
	}
	req.Interval = strings.ToLower(req.Interval)

	amount, price := calculateAmountAndPrice(req.Interval, req.Amount)
	params := checkoutSessionParams(customer, redirect, req.Interval, amount, price)
	sess, err := single.Stripe.CheckoutSessions.New(params)
	if err != nil {
		errorBadRequestResponse(w)
	}

	single.Inertia.Location(w, r, sess.URL)
}

func checkoutSessionParams(customer string, redirect string, interval string, amount int64, price float64) *stripe.CheckoutSessionParams {
	params := &stripe.CheckoutSessionParams{
		Customer:   stripe.String(customer),
		SuccessURL: stripe.String(globals.AppURL + redirect),
		CancelURL:  stripe.String(globals.AppURL + "/shop"),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyGBP)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Credits"),
					},
					UnitAmountDecimal: stripe.Float64(math.Round(price*1e12) / 1e12),
				},
				Quantity: stripe.Int64(amount),
			},
		},
	}

	if interval == intervalMonthly || interval == intervalAnnual {
		planInterval := stripe.String(string(stripe.PlanIntervalMonth))
		if interval == intervalAnnual {
			planInterval = stripe.String(string(stripe.PlanIntervalYear))
		}
		params.Mode = stripe.String(string(stripe.CheckoutSessionModeSubscription))
		params.LineItems[0].PriceData.Recurring = &stripe.CheckoutSessionLineItemPriceDataRecurringParams{
			Interval: planInterval,
		}
	}

	return params
}

func calculateAmountAndPrice(period string, amount int) (int64, float64) {
	amount = roundUpToMinimumAmount(period, amount)
	bonus := calculateBonusCredits(amount)
	finalPrice := calculatePrice(period, amount)
	finalAmount := int64(amount + bonus)
	return finalAmount, finalPrice
}

func roundUpToMinimumAmount(interval string, amount int) int {
	min := 100
	step := 100

	if interval == intervalAnnual {
		min = 1000
		step = 200
	}

	amount = int(math.Max(float64(min), float64(amount)))
	amount = int(math.Ceil(float64(amount)/float64(step))) * step
	return amount
}

func calculatePrice(interval string, amount int) float64 {
	pricePer := 0.05

	switch interval {
	case intervalMonthly:
		pricePer *= 0.9
	case intervalAnnual:
		pricePer *= 0.8
	}

	return pricePer * float64(amount)
}

func calculateBonusCredits(amount int) int {
	return int(math.Floor(float64(amount)/500.0) * 100)
}
