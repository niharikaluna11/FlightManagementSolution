import { defineStore } from 'pinia';

export const useFormStore = defineStore('form', {
    state: () => ({
        form: {} // This is the data that we want to persist
    }),
    actions: {
        setFormData(data) {
            this.form = data;
        },
        getFormData() {
            return this.form;
        }
    }
});
export const useFlightStore = defineStore('flightStore', {
    state: () => ({
        selectedFlight: {}
    }),
    actions: {
        setSelectedFlight(flight) {
            this.selectedFlight = flight;
        },
        getSelectedFlight() {
            return this.selectedFlight;
        }
    }
});


export const usePassengerStore = defineStore("passengerStore", {
    state: () => ({
        passengers: [], // Initialize as an empty array
    }),
    actions: {
        addPassenger(passenger) {
            this.passengers.push(passenger);
        },
    },
});