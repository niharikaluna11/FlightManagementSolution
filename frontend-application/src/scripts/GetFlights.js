import axios from "axios";

const BASE_URL = "http://localhost:8090/";

export const DisplayFlights = async (sourceId, destId, doj) => {
    //source_id=1&dest_id=2&date_of_journey=2025-01-14"
    const url = `${BASE_URL}search-flights-oneway?source_id=${sourceId}&dest_id=${destId}&date_of_journey=${doj}`;
    try {
        const response = await axios.get(url);
        console.log("Response:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error fetching flights:", err);
        throw err;
    }
};


export const DisplayFlights2 = async (sourceId, destId, doj, dor) => {
    //"/search-flights-twoway?source_id=1&dest_id=2&date_of_journey=2025-01-15&date_of_return=2025-01-16"
    const url = `${BASE_URL}search-flights-twoway?source_id=${sourceId}&dest_id=${destId}&date_of_journey=${doj}&date_of_return=${dor}`;
    try {
        const response = await axios.get(url);
        console.log("Response:", response.data);
        return response.data;
    } catch (err) {
        console.error("Error fetching flights:", err);
        throw err;
    }
};

