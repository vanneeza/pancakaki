package transactioncontroller

import (
	"fmt"
	"log"
	"net/http"
	"pancakaki/internal/domain/web"
	webtransaction "pancakaki/internal/domain/web/transaction"
	paymentservice "pancakaki/internal/service/payment"
	transactionservice "pancakaki/internal/service/transaction"
	"pancakaki/utils/helper"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
		Message: "the transaction still progress, waiting to payment",
		Data:    TransactionResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"Transaction": webResponse})

}

func (TransactionController *TransactionControllerImpl) CustomerPayment(context *gin.Context) {
	// D:\Programming\go\pancakaki\document\uploads\customer_payment
	var CustomerPayment webtransaction.PaymentCreateRequest
	paymentId, _ := strconv.Atoi(context.Param("id"))
	err := context.ShouldBind(&CustomerPayment)
	helper.InternalServerError(err, context)

	CustomerPayment.Transaction_detail_order_Id = paymentId
	log.Println(CustomerPayment)
	fmt.Scanln()

	// Mengambil file foto yang diupload
	file, err := context.FormFile("photo")
	if err != nil {
		// Handle error jika tidak ada foto yang diupload
		helper.InternalServerError(err, context)
		return
	}

	ext := filepath.Ext(file.Filename)
	currentTime := time.Now()
	formattedDate := currentTime.Format("20060102")
	newFilename := fmt.Sprintf("%s%s%s%s%d%s", formattedDate, "_", "CustomerPayment", "_", CustomerPayment.VirtualAccount, ext)

	// Menyimpan file foto ke direktori customer_payment
	uploadPath := filepath.Join("document/uploads/customer_payment", newFilename)
	err = context.SaveUploadedFile(file, uploadPath)
	if err != nil {
		// Handle error jika gagal menyimpan foto
		helper.InternalServerError(err, context)
		return
	}

	// Lanjutkan dengan pemrosesan transaksi dan respons

	TransactionResponse, err := TransactionController.TransactionService.CustomerPayment(CustomerPayment)
	helper.InternalServerError(err, context)

	webResponse := web.WebResponse{
		Code:    http.StatusCreated,
		Status:  "CREATED",
		Message: "the transaction was completed",
		Data:    TransactionResponse,
	}

	context.JSON(http.StatusCreated, gin.H{"payment": webResponse})

}

type CheckoutData struct {
	ClientSecret string
}

// func (TransactionController *TransactionControllerImpl) CreatePaymentIntent(c *gin.Context) {
// 	checkoutTmpl := template.Must(template.ParseFiles("internal/views/checkout.html"))
// 	var payment entity.Payment
// 	err := c.ShouldBind(&payment)
// 	if err != nil {
// 		// Jika terjadi kesalahan dalam membaca permintaan pengguna
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "error",
// 			"message": "Terjadi kesalahan dalam membaca permintaan pengguna",
// 		})
// 		return
// 	}

// 	stripe.Key = "sk_test_51NBe0ZIEAauHIyesxp61VF2X3dL8Eb0bcDREFn6v0k6uSSSmuptdx85VwJOuvi8OHRCAd0t9cRN8XM5NbdaQk8YK00CbOxnXTX" // Ganti dengan kunci rahasia Stripe Anda

// 	totalAmount := calculateTotalAmount(payment.Price, payment.Qty)

// 	params := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(totalAmount),
// 		Currency: stripe.String("idr"),
// 		PaymentMethodTypes: stripe.StringSlice([]string{
// 			"card",
// 		}),
// 		Params: stripe.Params{
// 			Metadata: map[string]string{
// 				"product_name": payment.Name,
// 			},
// 		},
// 		Description: &payment.Name,
// 	}

// 	intent, err := paymentintent.New(params)
// 	if err != nil {
// 		// Jika terjadi kesalahan dalam pembuatan PaymentIntent
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  "error",
// 			"message": "Terjadi kesalahan dalam pembuatan PaymentIntent",
// 		})
// 		return
// 	}

// 	// p, _ := TransactionController.PaymentService.InsertPayment(&payment)

// 	data := struct {
// 		ClientSecret string
// 		ProductName  string
// 		TotalAmount  int64
// 	}{
// 		ClientSecret: intent.ClientSecret,
// 		ProductName:  payment.Name,
// 		TotalAmount:  totalAmount,
// 	}

// 	err = checkoutTmpl.Execute(c.Writer, data)
// 	if err != nil {
// 		// Jika terjadi kesalahan dalam mengeksekusi template
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  "error",
// 			"message": "Terjadi kesalahan dalam mengeksekusi template",
// 		})
// 		return

// 	}
// }

// // Fungsi untuk menghitung total pembayaran
// func calculateTotalAmount(price, qty int) int64 {
// 	// Lakukan perhitungan total berdasarkan harga produk dan jumlah
// 	// Pastikan Anda melakukan validasi dan konversi tipe data yang sesuai
// 	// Misalnya, jika harga dan jumlah dalam bentuk string, Anda dapat mengubahnya menjadi float64 terlebih dahulu,
// 	// kemudian mengalikan dan mengonversi hasilnya menjadi int64.

// 	// Contoh sederhana:
// 	productPrice := 2000000 // Harga produk dalam rupiah
// 	productQty := 1         // Jumlah produk

// 	totalAmount := int64(productPrice * productQty)

// 	return totalAmount
// }
