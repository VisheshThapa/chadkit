package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

func StripePaymentHandler(c echo.Context) error {
	stripe.Key = "sk_test_51O2PQzLUswntPTxShPesVmjlf9q4h8uapLcMu19tXVYR6r6YSV26keSfwQbw0Hbq7duN5Knct2i7PsdZskTYPPa300SmnABo4l"
	request := c.Request()
	payload, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}

	var event stripe.Event
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return err
	}

	// // This is your Stripe CLI webhook secret for testing your endpoint locally.
	// endpointSecret := "whsec_f2a7179bfc015bbec13436fef29b300807b2775c192a92ba8bf426b53fc40d92"
	// // Pass the request body and Stripe-Signature header to ConstructEvent, along
	// // with the webhook signing key.
	// event, err := webhook.ConstructEvent(payload, request.Header.Get("Stripe-Signature"),
	// 	endpointSecret)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
	// 	return err
	// }
	//
	switch event.Type {
	case "checkout.session.completed":
		var checkout_session stripe.CheckoutSession

		err = json.Unmarshal(event.Data.Raw, &checkout_session)
		if err != nil {
			return err
		}
		params := &stripe.CheckoutSessionListLineItemsParams{
			Session: stripe.String(checkout_session.ID),
		}
		result := session.ListLineItems(params)
		var li string
		for result.Next() {
			li = result.LineItem().Description
		}

		CreatePaidUser(
			checkout_session.CustomerDetails.Name,
			checkout_session.CustomerDetails.Email,
			checkout_session.PaymentIntent.ID,
			li,
		)
	case "checkout.session.expired":
		// checkout wasn't completed send email reminder
	case "invoice.paid":
		// do something
	case "invoice.payment_failed":
		// revoke access or email about updating payment method
	case "customer.subscription.deleted":
		// revoke access
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)

	}

	return c.JSON(http.StatusOK, event)
}

func CreatePaidUser(name string, email string, paymentID string, membership string) error {
	collection, err := app.Dao().FindCollectionByNameOrId("PaidUsers")
	if err != nil {
		fmt.Println("FindCollectionByNameOrID Error", err)
		return err
	}
	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)
	form.LoadData(map[string]any{
		"name":           name,
		"email":          email,
		"payment_intent": paymentID,
		"membership":     membership,
	})
	if err := form.Submit(); err != nil {
		fmt.Println("form.Submit() eror", err)
		return err
	}
	err = SendProduct(name, email, membership)
	if err != nil {
		fmt.Println("Send Product", err)
		return err
	}

	return err
}
