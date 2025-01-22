package handlers

import (
	"bytes"
	"flight-app/config"
	"flight-app/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrewcharlton/wkhtmltopdf-go"
	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
)

var newBooking struct {
	FlightID         uint                    `json:"flight_id"`
	DetailsClassID   uint                    `json:"details_class_id"`
	BookingReference string                  `json:"booking_reference"`
	TotalAmount      float64                 `json:"total_amount"`
	Passengers       []models.PassengerInput `json:"passengers"`
	Payment          models.PaymentInput     `json:"payment"`
}

func CreateBooking(ctx *gin.Context) {

	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	if err := ctx.ShouldBindJSON(&newBooking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	//added booking data
	booking := models.Booking{
		FlightID:         newBooking.FlightID,
		DetailsClassID:   newBooking.DetailsClassID,
		BookingReference: newBooking.BookingReference,
		TotalAmount:      newBooking.TotalAmount,
	}

	if err := db.Create(&booking).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save booking"})
		return
	}

	// Added Passengers
	for _, p := range newBooking.Passengers {
		passenger := models.Passenger{
			BookingID:   booking.ID,
			Title:       models.Title(p.Title),
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			DateOfBirth: p.DateOfBirth,
			SeatNumber:  p.SeatNumber,
		}
		if err := db.Create(&passenger).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save passenger details"})
			return
		}
	}

	// Step 3: Add Payment
	payment := models.Payment{
		BookingID:     booking.ID,
		Amount:        newBooking.Payment.Amount,
		PaymentMethod: newBooking.Payment.PaymentMethod,
		TransactionID: newBooking.Payment.TransactionID,
	}
	if err := db.Create(&payment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save payment details"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Booking created successfully",
		"booking":    booking,
		"passengers": newBooking.Passengers,
		"payment":    payment,
	})

}

//BOOK-1737444480674-627981
//get all bookings

func GetAllBookings(ctx *gin.Context) {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	var bookings []models.Booking
	if err := db.Preload("Passengers").Preload("Flight").Preload("Payments").Find(&bookings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve bookings"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"bookings": bookings})

}

func GetBookingByReference(ctx *gin.Context) {
	// Retrieve the booking reference from the URL parameter
	bookingReference := ctx.Query("bookingReference")
	if bookingReference == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Booking reference is required"})
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	var booking models.Booking
	if err := db.Preload("Passengers").Preload("Flight").Preload("Payments").Where("booking_reference = ?", bookingReference).First(&booking).Error; err != nil {

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"booking": booking})
}

func GetPdf(ctx *gin.Context) {
	// Retrieve the booking reference from the URL parameter
	bookingReference := ctx.Query("bookingReference")
	if bookingReference == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Booking reference is required"})
		return
	}

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	// Retrieve booking details along with passengers, flight, and payments
	var booking models.Booking
	if err := db.Preload("Passengers").Preload("Flight").Preload("Payments").Where("booking_reference = ?", bookingReference).First(&booking).Error; err != nil {
		log.Fatal("Error fetching booking:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Generate the HTML content for the PDF
	htmlContent := generateHtml(booking)

	// Convert HTML to PDF using wkhtmltopdf
	doc := wkhtmltopdf.NewDocument()
	buf := bytes.NewBufferString(htmlContent)
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading from reader.")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
		return
	}
	doc.AddPages(pg)

	// Set the headers to serve the PDF file as an attachment
	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", `attachment; filename="booking_report.pdf"`)

	// Write the PDF to the response
	err = doc.Write(ctx.Writer)
	if err != nil {
		log.Fatal("Error generating PDF:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating PDF"})
		return
	}
}

func generateHtml(booking models.Booking) string {
	// Define the base HTML structure for the booking report
	basehtml := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Flight Booking Report</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9fafb;
            margin: 0;
            padding: 20px;
        }
        table {
            border-collapse: collapse;
            width: 100%%;
        }
        .main-container {
            width: 800px;
            margin: 0 auto;
        }
        .report-header {
            background: white;
            padding: 24px;
            border-radius: 12px;
            margin-bottom: 24px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .section-title {
            color: #111827;
            font-size: 24px;
            margin: 0 0 24px 0;
            padding: 0;
        }
        .sub-title {
            color: #374151;
            font-size: 18px;
            margin: 0 0 16px 0;
            padding: 0;
        }
        .info-cell {
            background: #f9fafb;
            padding: 16px;
            border-radius: 8px;
            vertical-align: top;
        }
        .info-label {
            color: #6b7280;
            font-size: 14px;
            margin-bottom: 4px;
        }
        .info-value {
            font-size: 16px;
            font-weight: bold;
            color: #111827;
            margin-bottom: 12px;
        }
        .spacer-row {
            height: 24px;
        }
        .divider {
            border-top: 1px solid #e5e7eb;
            margin: 24px 0;
        }
    </style>
</head>
<body>
    <table class="main-container">
        <tr>
            <td>
                <div class="report-header">
                    <h1 class="section-title">Flight Booking Report</h1>

                    <!-- Booking Info -->
                    <table>
                        <tr>
                            <td width="50%%" class="info-cell">
                                <div class="info-label">Booking Reference Number</div>
                                <div class="info-value">%v</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">Flight ID</div>
                                <div class="info-value">%v</div>
                            </td>
                        </tr>
                    </table>

                    <div class="divider"></div>

                    <!-- Flight Details -->
                    <h2 class="sub-title">Flight Details</h2>
                    <table>
                        <tr>
                            <td width="50%%" class="info-cell">
                                <div class="info-label">From</div>
                                <div class="info-value">%v</div>
                                <div class="info-label">Departure Time</div>
                                <div class="info-value">%v</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">To</div>
                                <div class="info-value">%v</div>
                                <div class="info-label">Arrival Time</div>
                                <div class="info-value">%v</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">Price</div>
                                <div class="info-value">Rs %.2v</div>
                            </td>
                        </tr>
                    </table>

                    <div class="divider"></div>

                    <!-- Passenger Details -->
                    <h2 class="sub-title">Passenger Details</h2>`

	// Format the booking details into the HTML content
	html := fmt.Sprintf(basehtml,
		booking.BookingReference,                      // Booking Reference
		booking.FlightID,                              // Flight ID
		booking.Flight.DepartureAirport,               // From (Source)
		formatTime(booking.Flight.ScheduledDeparture), // Departure Time
		booking.Flight.ArrivalAirport,                 // To (Destination)
		formatTime(booking.Flight.ScheduledArrival),   // Arrival Time
		booking.TotalAmount,                           // Total Price (assuming Price is in Rs)
	)

	// Add passenger details dynamically
	for _, passenger := range booking.Passengers {
		passengerHtml := fmt.Sprintf(`
                    <table>
                        <tr>
                            <td class="info-cell" style="margin-bottom: 12px;">
                                <table width="">
                                    <tr>
                                        <td width="50%%">
                                            <div class="info-label">Passenger Name</div>
                                            <div class="info-value">%s %s</div>
                                        </td>
                                        <td width="50%%">
                                            <div class="info-label">Seat Number</div>
                                            <div class="info-value">%v</div>
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                        <tr class="spacer-row"><td></td></tr>`,
			passenger.FirstName,
			passenger.LastName,
			passenger.SeatNumber)
		html += passengerHtml
	}

	// Close all tags for HTML
	html += `
                </div>
            </td>
        </tr>
    </table>
</body>
</html>`

	return html
}

// Helper function to format time
func formatTime(timestamp time.Time) string {
	return timestamp.Format("03:04 PM")
}

func SendEmailWithPdf(ctx *gin.Context) {
	var emailDto models.EmailSendingDTO
	if err := ctx.ShouldBindJSON(&emailDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if emailDto.ReferencNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ReferencNo cannot be empty",
		})
		return
	}

	result, err := SendMail(emailDto.Email, emailDto.ReferencNo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":  result,
		"message": "Email has been sent successfully",
	})
}

func SendMail(email string, referenceNo string) (bool, error) {

	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	result, err := FindByReferenceNumber(referenceNo)
	if err != nil {
		return false, err
	}
	htmlContent := generateHtml(*result)
	pdfBuffer, err := generatePDF(htmlContent)
	if err != nil {
		return false, err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", smtpUser)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Flight Ticket Report")
	msg.SetBody("text/html", "<b>Here is your flight ticket</b>")
	msg.Attach("booking_report.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(pdfBuffer.Bytes())
		return err
	}))

	n := gomail.NewDialer("smtp.gmail.com", 587, smtpUser, smtpPass)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return false, err
	}
	return true, nil
}

func generatePDF(htmlContent string) (*bytes.Buffer, error) {
	doc := wkhtmltopdf.NewDocument()
	buf := bytes.NewBufferString(htmlContent)
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading from reader.")
	}
	doc.AddPages(pg)

	output := &bytes.Buffer{}
	errs := doc.Write(output)
	if errs != nil {
		log.Fatal("Error writing to writer.")
	}
	return output, nil
}
func FindByReferenceNumber(referenceNumber string) (*models.Booking, error) {
	if referenceNumber == "" {
		return nil, fmt.Errorf("ReferencNo cannot be empty")
	}

	db, _ := config.ConnectDB()

	var booking models.Booking
	if err := db.Preload("Passengers").Preload("Flight").Preload("Payments").
		Where("booking_reference = ?", referenceNumber).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}
