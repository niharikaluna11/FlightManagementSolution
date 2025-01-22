import axios from "axios";

const BASE_URL = "http://localhost:8090/";

export const DisplayAirports = async () => {
    const url = BASE_URL + "get-airports";
    try {
        const response = await axios.get(url);

        console.log("Response:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error fetching airports:", err);
        throw err;
    }
};

export const DisplayAirportBySource = async (source_id) => {
    const url = `${BASE_URL}get-routes-by-src?source_id=${source_id}`;
    try {
        const response = await axios.get(url);

        console.log("Response:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error fetching routes:", err);
        throw err;
    }
};
