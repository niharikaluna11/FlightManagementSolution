<template>
  <div class="booking-confirm">
    <header>
      <h1>Booking Confirmation</h1>
    </header>

    <main>
      <div class="booking-info">
        <h2>Booking Reference: {{ bookingReference }}</h2>
        <p>Your booking has been successfully created. Please find your booking details by:</p>
        <!-- Additional booking information can go here -->

        <div class="actions">
          <button @click="downloadPDF">Download PDF</button>
        </div>
      </div>

      <div class="email-section">
        <input type="email" v-model="email" placeholder="Enter your email" class="email-input" />
        <button @click="handleSendEmail">Send Email</button>
      </div>

    </main>
  </div>
</template>

<script>
import { GetPdf, sendEmail } from "@/scripts/GetPdf"; // Import GetPdf function

export default {
  name: "BookingResultPage",
  data() {
    return {
      bookingReference: this.$route.query.bookingReference || "", // Get booking reference from query
      email: "", // To store the user's email input
    };
  },
  methods: {
    async handleSendEmail() {
      if (!this.email) {
        alert("Please enter a valid email address.");
        return;
      }

      const emailDetails = {
        Email: this.email,
        ReferencNo: this.bookingReference,
      };

      try {
        const response = await sendEmail(emailDetails); // Call sendEmail function
        alert("Email sent successfully!");
        console.log("Response:", response);
      } catch (err) {
        alert("Failed to send email. Please try again.");
        console.error("Error:", err);
      }
    },
    downloadPDF() {
      console.log("Downloading PDF for booking reference:", this.bookingReference);

      GetPdf(this.bookingReference)
        .then((pdfBlob) => {
          const url = window.URL.createObjectURL(pdfBlob);

          const link = document.createElement("a");
          link.href = url;
          link.setAttribute("download", "booking_ticket.pdf");
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);

          window.URL.revokeObjectURL(url);
        })
        .catch((err) => {
          console.error("Error downloading PDF:", err);
        });
    },
  },
};
</script>

<style scoped>
.booking-confirm {
  text-align: center;
  margin: 20px;
}

.booking-info {
  margin-bottom: 20px;
}

.booking-info h2 {
  font-size: 24px;
  font-weight: bold;
  margin: 10px 0;
}

.email-section {
  margin: 20px 0;
}

.email-input {
  padding: 10px;
  width: 250px;
  margin-right: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.actions button,
.email-section button {
  padding: 10px 20px;
  margin: 10px;
  background-color: #0078d7;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.actions button:hover,
.email-section button:hover {
  background-color: #005fa3;
}
</style>
