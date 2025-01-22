import axios from "axios";

const BASE_URL = "http://localhost:8090/";

export const GetPdf = async (bookingReference) => {
    const url = `${BASE_URL}get-bookingpdf?bookingReference=${bookingReference}`;

    try {
        const response = await axios.get(url, { responseType: "blob" });
        console.log("Response:", response);
        return response.data;
    } catch (err) {
        console.error("Error getting booking PDF:", err);
        throw err;
    }
};


// Function to send email
export const sendEmail = async (emailDetails) => {
    const url = `${BASE_URL}send-email`;

    try {
        const response = await axios.post(url, emailDetails);
        console.log("Email sent successfully:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error sending email:", err.response?.data || err.message);
        throw err;
    }
};