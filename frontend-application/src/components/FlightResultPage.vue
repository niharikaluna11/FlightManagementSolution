<template>
    <header>
        <button @click="goBack()">Back</button>
    </header>
    <main class="main">
        <template v-if="Array.isArray(flights) && flights.length > 0">
            <div v-for="(flight, index) in flights" :key="flight.id" class="flight-card">
                <div class="flight-details">
                    <div class="flight-airline">
                        <span>{{ flight.airline }} {{ flight.flightNumber }}</span>
                    </div>

                    <div class="flight-time">
                        <div class="time">{{ flight.scheduledDeparture }}</div>
                        <div class="arrival-city">{{ flight.departureAirport }}</div>
                    </div>

                    <div class="flight-time">
                        <div class="time">{{ flight.scheduledArrival }}</div>
                        <div class="arrival-city">{{ flight.arrivalAirport }}</div>
                    </div>

                    <div class="flight-price">
                        <div class="amount">₹{{ flight.selectedFarePrice }}</div>
                        <button class="view-prices-btn" @click="selectFlight(flight)"
                            :disabled="!selectedFare[flight.id]">
                            CONTINUE
                        </button>
                    </div>
                </div>

                <div class="flight-details-footer">
                    <a href="#" @click.prevent="toggleAccordion(index)">View Flight Details</a>
                </div>

                <div v-if="activeAccordion === index" class="accordion-details">
                    <div>
                        <label>
                            <input type="radio" name="fareType" :value="'Saver'" v-model="selectedFare[flight.id]"
                                @change="changePrice(flight, 'Saver')" />
                            Saver (₹{{ flight.saverPrice }})
                        </label>
                    </div>
                    <div>
                        <label>
                            <input type="radio" name="fareType" :value="'Flexi'" v-model="selectedFare[flight.id]"
                                @change="changePrice(flight, 'Flexi')" />
                            Flexi (₹{{ flight.flexiPrice }})
                        </label>
                    </div>
                </div>
            </div>
        </template>
        <template v-else>
            <p>No flights available. Please modify your search.</p>
        </template>
    </main>
   ]
</template>

<script>
import { DisplayFlights } from "@/scripts/GetFlights";
import { useFormStore } from '@/scripts/usePiniaStorage';
import { useFlightStore } from '@/scripts/usePiniaStorage'; // import the flight store

export default {
    name: "FlightResultPage",
    data() {
        return {
            flights: [],
            activeAccordion: null,
            selectedFare: {},
            form: {},
        };
    },
    setup() {
        const formStore = useFormStore();
        const flightStore = useFlightStore(); // Access flight store
        return { formStore, flightStore };
    },
    methods: {
        calculateFares(basePrice) {
            return {
                saverPrice: basePrice,
                flexiPrice: basePrice * 1.8,
            };
        },

        validateForm() {
            if (!this.form.source || !this.form.destination || !this.form.departureDate) {
                alert("Source, destination, and departure date are required.");
                return false;
            }
            return true;
        },

        async fetchFlights() {
            if (!this.validateForm()) return;
            try {
                const response = await DisplayFlights(
                    this.form.source,
                    this.form.destination,
                    this.form.departureDate
                );
                if (!response || !response.flights || response.flights.length === 0) {
                    alert("No flights available for the selected route and date.");
                    this.$router.push("/no-flights-page");
                    return;
                }
                this.flights = response.flights.map(this.mapFlightData);
            } catch (error) {
                console.error("Error fetching flights:", error);
                alert("Unable to load flights.");
            }
        },

        toggleAccordion(index) {
            this.activeAccordion = this.activeAccordion === index ? null : index;
        },

        mapFlightData(flight) {
            const { saverPrice, flexiPrice } = this.calculateFares(flight.BasePrice);
            if (!this.selectedFare[flight.ID]) {
                this.selectedFare[flight.ID] = 'Saver'; // Default to Saver
            }
            return {
                id: flight.ID,
                flightNumber: flight.FlightNumber,
                airline: flight.AirlineName,
                departureAirport: flight.DepartureAirport.Code,
                arrivalAirport: flight.ArrivalAirport.Name,
                scheduledDeparture: flight.ScheduledDeparture,
                scheduledArrival: flight.ScheduledArrival,
                basePrice: flight.BasePrice,
                saverPrice,
                flexiPrice,
                selectedFarePrice: saverPrice, // Default selected price is Saver
            };
        },

        changePrice(flight, fareType) {
            flight.selectedFarePrice = fareType === 'Flexi' ? flight.flexiPrice : flight.saverPrice;
        },

        selectFlight(flight) {
            const selectedFareType = this.selectedFare[flight.id] || 'Saver';
            const selectedFarePrice = selectedFareType === 'Saver' ? flight.saverPrice : flight.flexiPrice;

            this.flightStore.setSelectedFlight({
                flightId: flight.id,
                selectedFareType,
                selectedFarePrice,
                detailsClassID: flight.detailsClassID,
                totalAmount: flight.totalAmount,
                ...flight,
            });

            alert(`Selected ${flight.flightNumber} with ${selectedFareType} fare`);

            this.$router.push({
                path: "/passenger-page",
                query: {
                    flightId: flight.id,
                    fareType: selectedFareType,
                    selectedFarePrice,
                },
            });
        },

        goBack() {
            window.location.href = "/search-flight";
        },
    },
    mounted() {
        const savedFormData = this.formStore.getFormData();
        this.form = savedFormData;

        this.fetchFlights();
    },
};
</script>



<style scoped>
.flight-card {
    display: flex;
    flex-direction: column;
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    padding: 16px;
    margin: 10px 0;
    background-color: #ffffff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.flight-card-header {
    font-size: 0.9rem;
    color: #ff5722;
    font-weight: bold;
    margin-bottom: 10px;
}

.flight-details {
    display: flex;
    justify-content: space-between;
    padding: 10px 0;
}

.flight-price {
    text-align: right;
}

.view-prices-btn {
    padding: 8px 16px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.flight-details-footer {
    text-align: center;
    margin-top: 10px;
}

.accordion-details {
    display: flex;
    gap: 20px;
    /* Space between radio buttons */
    flex-direction: row;
    justify-content: start;
    /* Align to the left */
    margin-top: 10px;
}

.accordion-details div {
    display: flex;
    align-items: center;
    gap: 8px;
    /* Space between radio button and text */
}

.accordion-details label {
    display: flex;
    align-items: center;
    cursor: pointer;
}

.accordion-details input[type="radio"] {
    margin-right: 8px;
    /* Space between radio button and label text */
}
</style>
