package transactioncontroller

import (
	"net/http"
	"pancakaki/internal/domain/entity"
	"pancakaki/internal/domain/web"
	webtransaction "pancakaki/internal/domain/web/transaction"
	paymentservice "pancakaki/internal/service/payment"
	transactionservice "pancakaki/internal/service/transaction"
	"pancakaki/utils/helper"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type TransactionControllerImpl struct {
	TransactionService transactionservice.TransactionService
	PaymentService     paymentservice.PaymentService
}

func NewTransactionController(TransactionService transactionservice.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		TransactionService: TransactionService,
	}
}

func (TransactionController *TransactionControllerImpl) MakeOrder(context *gin.Context) {

	var Transaction webtransaction.TransactionOrderCreateRequest

	err := context.ShouldBind(&Transaction)
	helper.InternalServerError(err, context)

	TransactionResponse, err := TransactionController.TransactionService.MakeOrder(Transaction)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "The data has been successfully added",
		Data:    TransactionResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"Transaction": webResponse})

}

type CheckoutData struct {
	ClientSecret string
}

func (TransactionController *TransactionControllerImpl) CreatePaymentIntent(c *gin.Context) {
	checkoutTmpl := template.Must(template.ParseFiles("internal/views/checkout.html"))
	var payment entity.Payment
	err := c.ShouldBind(&payment)
	if err != nil {
		// Jika terjadi kesalahan dalam membaca permintaan pengguna
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan dalam membaca permintaan pengguna",
		})
		return
	}

	stripe.Key = "sk_test_51NBe0ZIEAauHIyesxp61VF2X3dL8Eb0bcDREFn6v0k6uSSSmuptdx85VwJOuvi8OHRCAd0t9cRN8XM5NbdaQk8YK00CbOxnXTX" // Ganti dengan kunci rahasia Stripe Anda

	totalAmount := calculateTotalAmount(payment.Price, payment.Qty)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(totalAmount),
		Currency: stripe.String("usd"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}

	intent, err := paymentintent.New(params)
	if err != nil {
		// Jika terjadi kesalahan dalam pembuatan PaymentIntent
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan dalam pembuatan PaymentIntent",
		})
		return
	}

	// p, _ := TransactionController.PaymentService.InsertPayment(&payment)

	data := struct {
		ClientSecret string
		ProductName  string
		TotalAmount  int64
	}{
		ClientSecret: intent.ClientSecret,
		ProductName:  payment.Name,
		TotalAmount:  totalAmount,
	}

	err = checkoutTmpl.Execute(c.Writer, data)
	if err != nil {
		// Jika terjadi kesalahan dalam mengeksekusi template
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan dalam mengeksekusi template",
		})
		return

	}
}

// Fungsi untuk menghitung total pembayaran
func calculateTotalAmount(price, qty int) int64 {
	// Lakukan perhitungan total berdasarkan harga produk dan jumlah
	// Pastikan Anda melakukan validasi dan konversi tipe data yang sesuai
	// Misalnya, jika harga dan jumlah dalam bentuk string, Anda dapat mengubahnya menjadi float64 terlebih dahulu,
	// kemudian mengalikan dan mengonversi hasilnya menjadi int64.

	// Contoh sederhana:
	productPrice := 2000000 // Harga produk dalam rupiah
	productQty := 1         // Jumlah produk

	totalAmount := int64(productPrice * productQty)

	return totalAmount
}
