<template>
    <main>
        <header>
            <h1>Payment Page</h1>
        </header>

        <div class="passenger-form">
            <form @submit.prevent="processPayment">
                <!-- Credit Card Number -->
                <label for="cc-number">Credit Card Number</label>
                <input type="text" id="cc-number" v-model="creditCardNumber" maxlength="16"
                    placeholder="XXXX-XXXX-XXXX-XXXX" />
                <span v-if="errors.ccNumber" class="error">{{ errors.ccNumber }}</span>

                <!-- CVV -->
                <label for="cvv">CVV</label>
                <input type="text" id="cvv" v-model="cvv" maxlength="3" placeholder="Enter 3-digit CVV" />
                <span v-if="errors.cvv" class="error">{{ errors.cvv }}</span>

                <!-- Expiry Date -->
                <label for="expiry-date">Expiry Date</label>
                <div class="expiry-container">
                    <select v-model="expiryMonth" id="expiry-month">
                        <option value="" disabled>Month</option>
                        <option v-for="month in 12" :key="month" :value="month">
                            {{ month < 10 ? '0' + month : month }} </option>
                    </select>
                    <select v-model="expiryYear" id="expiry-year">
                        <option value="" disabled>Year</option>
                        <option v-for="year in years" :key="year" :value="year">
                            {{ year }}
                        </option>
                    </select>
                </div>
                <span v-if="errors.expiryDate" class="error">{{ errors.expiryDate }}</span>

                <!-- Submit Button -->
                <button type="submit">Submit Payment</button>
            </form>
        </div>

        <div v-if="errors.apiError" class="error">
            <p>{{ errors.apiError }}</p>
        </div>
    </main>
    
</template>

<script>
import { CreateBooking } from "@/scripts/BookingFlights";
import { useFormStore, useFlightStore, usePassengerStore } from "@/scripts/usePiniaStorage";

import { useRouter } from "vue-router";


export default {
    name: "PaymentPage",
    setup() {
        const formStore = useFormStore();
        const flightStore = useFlightStore();
        const passengerStore = usePassengerStore();
        const router = useRouter();



        return { formStore, flightStore, passengerStore, router };
    },
    data() {
        return {
            creditCardNumber: "",
            cvv: "",
            expiryMonth: "",
            expiryYear: "",
            errors: {}, // Store validation errors
        };
    },
    computed: {
        years() {
            const currentYear = new Date().getFullYear();
            return Array.from({ length: 10 }, (_, i) => currentYear + i);
        },
    },
    methods: {
        validateForm() {
            this.errors = {}; // Reset errors

            // Validate Credit Card Number
            const ccNumberRegex = /^[0-9]{16}$/;
            if (!ccNumberRegex.test(this.creditCardNumber)) {
                this.errors.ccNumber = "Invalid Credit Card Number (16 digits required).";
            }

            // Validate CVV
            const cvvRegex = /^[0-9]{3}$/;
            if (!cvvRegex.test(this.cvv)) {
                this.errors.cvv = "Invalid CVV (3 digits required).";
            }

            // Validate Expiry Date
            if (!this.expiryMonth || !this.expiryYear) {
                this.errors.expiryDate = "Please select a valid expiry date.";
            } else {
                const today = new Date();
                const selectedDate = new Date(this.expiryYear, this.expiryMonth - 1);
                if (selectedDate < today) {
                    this.errors.expiryDate = "The expiry date must be in the future.";
                }
            }

            return Object.keys(this.errors).length === 0;
        },
        //    async  processPayment() {
        //         if (!this.validateForm()) {
        //             return;
        //         }

        //         if (!Array.isArray(this.passengerStore.passengers) || this.passengerStore.passengers.length === 0) {
        //             this.errors.passengerData = "No passengers found. Please add passenger details.";
        //             return;
        //         }

        //         const bookingData = {
        //             flight_id: this.flightStore.flightID,
        //             details_class_id: this.formStore.detailsClassID,
        //             booking_reference: `BOOK-${Math.floor(Math.random() * 100000)}`,
        //             total_amount: this.formStore.totalAmount,
        //             passengers: this.passengerStore.passengers.map((passenger) => ({
        //                 title: passenger.title,
        //                 first_name: passenger.firstName,
        //                 last_name: passenger.lastName,
        //                 date_of_birth: passenger.dateOfBirth,
        //                 seat_number: passenger.seatNumber,
        //             })),
        //             payment: {
        //                 amount: this.formStore.totalAmount,
        //                 payment_method: "Credit Card",
        //                 transaction_id: `TRANS-${Math.floor(Math.random() * 1000000)}`,
        //             },
        //         };

        //         try {
        //             const response = await axios.post(BASE_URL + "create-booking", bookingData);
        //             console.log("Booking created successfully:", response.data);

        //             this.router.push({
        //                 path: "/booking-confirm",
        //                 query: { bookingReference: response.data.booking.booking_reference },
        //             });
        //         } catch (error) {
        //             console.error("Failed to create booking:", error);
        //             this.errors.apiError = "Failed to create booking. Please try again.";
        //         }
        //     }
        //     ,
        async processPayment() {
            if (!this.validateForm()) {
                return;
            }

            if (!Array.isArray(this.passengerStore.passengers) || this.passengerStore.passengers.length === 0) {
                this.errors.passengerData = "No passengers found. Please add passenger details.";
                return;
            }
            const generateUniqueBookingReference = () => {
                const timestamp = Date.now();  // Use current timestamp for uniqueness
                const randomNum = Math.floor(Math.random() * 1000000);  // Random number for added complexity
                const uniqueRef = `BOOK-${timestamp}-${randomNum}`;  // Combine timestamp and random number
                return uniqueRef;
            };
            const generateUniquetransaction = () => {
                const timestamp = Date.now();  // Use current timestamp for uniqueness
                const randomNum = Math.floor(Math.random() * 1000000);  // Random number for added complexity
                const uniqueRef = `TXN-${timestamp}-${randomNum}`;  // Combine timestamp and random number
                return uniqueRef;
            };



            const bookingData = {
                flight_id: this.flightStore.selectedFlight.flightId, // Use flightId from flightStore
                details_class_id: this.flightStore.selectedFlight.selectedFareType === "Saver" ? 1 :
                    this.flightStore.selectedFlight.selectedFareType === "Flexi" ? 2 : null,
                booking_reference: generateUniqueBookingReference(),
                total_amount: this.flightStore.selectedFlight.selectedFarePrice, // Use 'selectedFarePrice' as the total amount
                passengers: this.passengerStore.passengers.map((passenger) => ({
                    title: passenger.title, // Map title
                    first_name: passenger.firstName, // Map firstName to first_name
                    last_name: passenger.lastName, // Map lastName to last_name
                    date_of_birth: passenger.dateOfBirth, // Map dateOfBirth
                    seat_number: passenger.seatNumber || null, // Add seat_number if available or default to null
                })),
                payment: {
                    amount: this.flightStore.selectedFlight.selectedFarePrice, // Use the same amount as the total fare
                    payment_method: "Credit Card", // Hardcoded payment method
                    transaction_id: generateUniquetransaction(), // Generate a unique transaction ID
                },
            };


            console.log(bookingData);
            // Call the API to create the booking
            const response = await CreateBooking(bookingData);
            console.log("Booking created successfully:", response);

            if (response && response.booking) {
                // Redirect to the booking confirmation page with the correct booking reference
                this.router.push({
                    path: "/booking-confirm",
                    query: { bookingReference: response.booking.BookingReference },
                });
            } else {
                console.error("Booking data is missing in the response.");
                this.errors.apiError = "Booking data is missing. Please try again.";
            }

        }
        ,

    },
};
</script>



<style scoped>
.passenger-form {
    margin-bottom: 20px;
    border: 1px solid #ddd;
    padding: 20px;
    background-color: white;
    border-radius: 5px;
    max-width: 400px;
    margin: auto;
}

label {
    display: block;
    margin: 5px 0;
}

input,
select {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
    font-size: 14px;
}

button {
    padding: 10px 20px;
    background-color: #0078d7;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    width: 100%;
}

button:hover {
    background-color: #005fa3;
}

.error {
    color: red;
    font-size: 12px;
    margin-bottom: 10px;
}

.expiry-container {
    display: flex;
    gap: 10px;
}

.confirmation {
    text-align: center;
    margin: 20px auto;
}
</style>
