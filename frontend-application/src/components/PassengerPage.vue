<template>
    <header>
        <h1>PassengerPage</h1>
    </header>

    <main>
        <form @submit.prevent="submitForm">
            <div v-for="(passenger, index) in passengers" :key="index" class="passenger-form">
                <h3>Passenger {{ index + 1 }}</h3>

                <div style="display: flex; gap: 10px; flex-wrap: wrap;">
                    <div>
                        <label for="title">Title:</label>
                        <select v-model="passenger.title" id="title" required>
                            <option value="" disabled>Select Title</option>
                            <option value="Mr">Mr</option>
                            <option value="Mrs">Mrs</option>
                            <option value="Miss">Miss</option>
                        </select>
                    </div>

                    <div>
                        <label for="firstName">First Name:</label>
                        <input type="text" v-model="passenger.firstName" id="firstName" required />
                    </div>

                    <div>
                        <label for="lastName">Last Name:</label>
                        <input type="text" v-model="passenger.lastName" id="lastName" required />
                    </div>

                    <div>
                        <label for="dob">Date of Birth:</label>
                        <input type="date" v-model="passenger.dateOfBirth" id="dob" required />
                    </div>
                </div>
            </div>

            <button type="submit">Submit</button>
        </form>
    </main>

   
</template>

<script>
import { useFormStore, useFlightStore, usePassengerStore } from "@/scripts/usePiniaStorage";

export default {
    name: "PassengerPage",
    data() {
        return {
            passengers: [], // Array to hold passenger details
        };
    },
    setup() {
        const formStore = useFormStore();
        const flightStore = useFlightStore();
        const passengerStore = usePassengerStore();
        return { formStore, flightStore, passengerStore };
    },
    mounted() {
        const savedFormData = this.formStore.getFormData();
        if (savedFormData) {
            console.log("Loaded form data:", savedFormData);
        } else {
            console.warn("No saved form data found.");
        }

        const numPassengers = savedFormData?.passengers || 1; // Default to 1 if no data
        this.passengers = Array.from({ length: numPassengers }, () => ({
            title: "",
            firstName: "",
            lastName: "",
            dateOfBirth: "",
        }));
    },
    methods: {
        submitForm() {
            // Validate form fields
            const hasEmptyFields = this.passengers.some(
                (passenger) =>
                    !passenger.title || !passenger.firstName || !passenger.lastName || !passenger.dateOfBirth
            );

            if (hasEmptyFields) {
                alert("Please fill all passenger details.");
                return;
            }

            // Save passenger data to the store
            this.passengerStore.passengers = this.passengers;
            console.log("Passengers saved to store:", this.passengerStore.passengers);

            // Navigate to the next page
            this.$router.push({
                path: "/seat-selection-page",
            });

            alert("Passenger details submitted successfully!");
        },
    },
};
</script>

<style scoped>
.passenger-form {
    margin-bottom: 20px;
    border: 1px solid #ddd;
    padding: 10px;
    background-color: white;
    border-radius: 5px;
}

label {
    display: block;
    margin: 5px 0;
    font-weight: bold;
}

input,
select {
    width: 100%;
    padding: 5px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
}

button {
    padding: 10px 20px;
    background-color: #0078d7;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button:hover {
    background-color: #005fa3;
}
</style>
