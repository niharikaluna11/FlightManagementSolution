<template>
    <header>
        <button @click="goBack()">Back</button>
        <button @click="continueToPassenger()">Continue</button>
    </header>
    <main class="main">
        <template
            v-if="Array.isArray(outboundFlights) && outboundFlights.length > 0 || Array.isArray(returnFlights) && returnFlights.length > 0">
            <div class="flights-container">
                <!-- Outbound Flights Section -->
                <div class="flights-row">
                    <h2 class="flights-heading">Outbound Flights</h2>


                    <div v-if="outboundFlights.length > 0">

                        <div class="flights-list">
                            <div class="flight-card" v-for="(flight) in outboundFlights" :key="flight.id">
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
                                        <div class="amount">₹{{ flight.basePrice }}</div>
                                        <button class="view-prices-btn" @click="selectFlight(flight, 'outbound')"
                                            :disabled="!selectedFare[flight.id]">
                                            Select
                                        </button>
                                    </div>
                                </div>
                                <!-- Fare Selection Accordion -->
                                <div v-if="activeAccordion === 'outbound'">
                                    <div class="accordion-details">
                                        <button @click="selectFare(flight.id, 'Saver')">Saver (₹{{ flight.saverPrice
                                            }})</button>
                                        <button @click="selectFare(flight.id, 'Flexi')">Flexi (₹{{ flight.flexiPrice
                                            }})</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div v-else class="no-flight-box">
                        <p>No outbound flights available for the selected date.</p>
                    </div>
                </div>

                <!-- Return Flights Section -->
                <div class="flights-row">
                    <h2 class="flights-heading">Return Flights</h2>
                    <!-- <div v-if="Array.isArray(returnFlights) && returnFlights.length > 0">
                        <div class="flights-list">
                            <div class="flight-card" v-for="(flight) in returnFlights" :key="flight.id">
                              
                            </div>
                        </div>
                    </div>
                    <div v-else class="no-flight-box">
                        <p>No return flights available for the selected date.</p>
                    </div> -->

                    <div v-if="returnFlights.length > 0">
                        <div class="flights-list">
                            <div class="flight-card" v-for="(flight) in returnFlights" :key="flight.id">
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
                                        <div class="amount">₹{{ flight.basePrice }}</div>
                                        <button class="view-prices-btn" @click="selectFlight(flight, 'return')"
                                            :disabled="!selectedFare[flight.id]">
                                            Select
                                        </button>
                                    </div>
                                </div>
                                <!-- Fare Selection Accordion -->
                                <div v-if="activeAccordion === 'return'">
                                    <div class="accordion-details">
                                        <button @click="selectFare(flight.id, 'Saver')">Saver (₹{{ flight.saverPrice
                                            }})</button>
                                        <button @click="selectFare(flight.id, 'Flexi')">Flexi (₹{{ flight.flexiPrice
                                            }})</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div v-else class="no-flight-box">
                        <p>No return flights available for the selected date.</p>
                    </div>
                </div>
            </div>
            <!-- Continue Button -->
            <div v-if="selectedOutboundFlight && selectedReturnFlight" class="continue-btn-container">
                <button class="continue-btn" @click="continueBooking">
                    Continue Booking
                </button>
            </div>
        </template>

        <template v-else>
            <p>No flights available. Please modify your search.</p>
        </template>
    </main>

</template>

<style scoped>
.header {
    margin-top: 10px;
    margin-bottom: 20px;
}

.flights-container {
    display: flex;
    flex-direction: row;
    gap: 40px;
    margin-left: 20px;
    margin-right: 20px;
}

.flights-row {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 100%;
}

.flights-heading {
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 10px;
}

.flights-list {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
}

.flight-card {
    display: flex;
    flex-direction: column;
    border: 1px solid #e0e0e0;
    border-radius: 8px;
    padding: 16px;
    width: 100%;
    background-color: #ffffff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.flight-details {
    display: flex;
    flex-direction: column;
    gap: 8px;
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

.accordion-details {
    padding: 10px;
    border-top: 1px solid #e0e0e0;
}

.no-flight-box {
    padding: 20px;
    background-color: #f8d7da;
    border: 1px solid #f5c6cb;
    border-radius: 8px;
    color: #721c24;
    text-align: center;
    font-weight: bold;
}

.continue-btn-container {
    text-align: center;
    margin-top: 20px;
}

.continue-btn {
    padding: 10px 20px;
    background-color: #28a745;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}
</style>
<script>
import { DisplayFlights2 } from "@/scripts/GetFlights";
import { useFormStore } from '@/scripts/usePiniaStorage';

export default {
    name: "FlightResult2Page",
    data() {
        return {
            outboundFlights: [],
            returnFlights: [],
            selectedFare: {},
            selectedOutboundFlight: null,
            selectedReturnFlight: null,
            activeAccordion: '',
            form: {},
        };
    },
    setup() {
        const formStore = useFormStore();
        return { formStore };
    },
    methods: {
        // Method for selecting fare type
        selectFare(flightId, fareType) {
            this.selectedFare[flightId] = fareType;
        },

        // Method for selecting outbound or return flight
        selectFlight(flight, type) {
            if (type === 'outbound') {
                this.selectedOutboundFlight = flight;
            } else {
                this.selectedReturnFlight = flight;
            }
        },

        // Method to continue booking after selecting both flights
        continueBooking() {
            if (this.outboundFlights.length === 0) {
                alert("No outbound flights available. Booking cannot proceed.");
                return;
            }

            if (!this.selectedOutboundFlight) {
                alert("Please select an outbound flight.");
                return;
            }

            if (this.form.returnDate && this.returnFlights.length > 0 && !this.selectedReturnFlight) {
                alert("Please select a return flight.");
                return;
            }

            // Proceed with booking
            alert(`Booking details:
    Outbound: ${this.selectedOutboundFlight.flightNumber} - Fare: ${this.selectedFare[this.selectedOutboundFlight.id] || "Not Selected"},
    Return: ${this.selectedReturnFlight ? `${this.selectedReturnFlight.flightNumber} - Fare: ${this.selectedFare[this.selectedReturnFlight.id] || "Not Selected"}` : "N/A"}`);

            this.$router.push({
                path: "/passenger-page",
                query: {
                    outboundFlightId: this.selectedOutboundFlight.id,
                    returnFlightId: this.selectedReturnFlight ? this.selectedReturnFlight.id : null,
                    outboundFareType: this.selectedFare[this.selectedOutboundFlight.id],
                    returnFareType: this.selectedReturnFlight ? this.selectedFare[this.selectedReturnFlight.id] : null,
                },
            });
        },

        toggleAccordion(accordion) {
            this.activeAccordion = this.activeAccordion === accordion ? null : accordion;
        },
        async fetchFlights() {
            if (!this.validateForm()) return;

            try {
                console.log("Fetching flights with form data:", this.form);

                const response = await DisplayFlights2(
                    this.form.source,
                    this.form.destination,
                    this.form.departureDate,
                    this.form.returnDate
                );

                console.log("Fetched flight response:", response);

                // Handle outbound flights
                if (Array.isArray(response.outbound_flights)) {
                    this.outboundFlights = response.outbound_flights.map(this.mapFlightData);
                } else if (typeof response.outbound_flights === "string") {
                    this.outboundFlights = [];
                    console.warn(response.outbound_flights); // Log the message for debugging
                } else {
                    this.outboundFlights = [];
                    console.warn("Unexpected format for outbound flights data.");
                    this.$router.push({
                        path: "/no-flights-page",

                    });
                }

                // Handle return flights
                if (Array.isArray(response.return_flights)) {
                    this.returnFlights = response.return_flights.map(this.mapFlightData);
                } else if (typeof response.return_flights === "string") {
                    this.returnFlights = [];
                    console.warn(response.return_flights); // Log the message for debugging
                } else {
                    this.returnFlights = [];
                    console.warn("Unexpected format for return flights data.");
                }

            } catch (error) {
                console.error("Error fetching flights:", error);
                alert("Unable to load flights. Please try again later.");
            }
        }
        ,
        validateForm() {
            const { source, destination, departureDate } = this.form;

            if (!source || !destination || !departureDate) {
                alert("Source, destination, and departure date are required.");
                return false;
            }

            return true;
        },

        mapFlightData(flight) {
            const { saverPrice, flexiPrice } = this.calculateFares(flight.BasePrice);
            return {
                id: flight.ID,
                flightNumber: flight.FlightNumber,
                airline: flight.AirlineName || "Unknown Airline",
                departureAirport: flight.DepartureAirport?.Name || "Unknown Airport",
                arrivalAirport: flight.ArrivalAirport?.Name || "Unknown Airport",
                scheduledDeparture: flight.ScheduledDeparture,
                scheduledArrival: flight.ScheduledArrival,
                basePrice: flight.BasePrice,
                saverPrice,
                flexiPrice,
            };
        },

        calculateFares(basePrice) {
            return {
                saverPrice: basePrice,
                flexiPrice: basePrice * 1.8,
            };
        },

        goBack() {
            this.$router.push("/search-flight");
        },

        continueToPassenger() {
            if (this.selectedOutboundFlight && this.selectedReturnFlight) {
                this.$router.push("/passenger-page");
            } else {
                alert("Please select both outbound and return flights before continuing.");
            }
        },
    },

    mounted() {
        const savedFormData = this.formStore.getFormData();
        this.form = (savedFormData);

        this.fetchFlights();
    },
};
</script>
