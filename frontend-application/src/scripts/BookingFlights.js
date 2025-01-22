import axios from "axios";

const BASE_URL = "http://localhost:8090/";


export const CreateBooking = async (bookingData) => {
    const url = BASE_URL + "create-booking";

    // Ensure the passengers' data is correctly mapped to snake_case before sending the request
    const formattedData = {
        flight_id: bookingData.flight_id,
        details_class_id: bookingData.details_class_id,
        booking_reference: bookingData.booking_reference,
        total_amount: bookingData.total_amount,
        passengers: bookingData.passengers.map((passenger) => ({
            title: passenger.title,  // These should be in the correct structure
            first_name: passenger.first_name,  // Convert to snake_case
            last_name: passenger.last_name,  // Convert to snake_case
            date_of_birth: passenger.date_of_birth,  // Convert to snake_case
            seat_number: passenger.seat_number,  // Convert to snake_case
        })),
        payment: {
            amount: bookingData.payment.amount,
            payment_method: bookingData.payment.payment_method,
            transaction_id: bookingData.payment.transaction_id,
        },
    };

    try {
        const response = await axios.post(url, formattedData);
        console.log("Response:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error creating booking:", err);
        throw err;
    }
};
