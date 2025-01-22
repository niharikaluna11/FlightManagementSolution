<template>
    <main>
        <header>
            <div class="search-btn-container">
                <button @click="goBack()" class="search-btn">Back</button>
                <button @click="continuenext()" class="search-btn">Continue</button>
            </div>
            <h1>Flight Seat Booking System</h1>
        </header>
        <div id="app" class="seat-container">


            <!-- Airplane Front -->
            <div class="seat-container-box-in">

            </div>

            <!-- Seat Rows -->
            <div class="seat-container-box">
                <div v-for="row in seatRows" :key="row.rowNumber" class="seat-row">
                    <!-- Render first 3 columns (A, B, C) -->
                    <div v-for="seat in row.seats.slice(0, 3)" :key="seat.id"
                        :class="['seat', seat.status, { selected: seat.selected }]" @click="toggleSeatSelection(seat)">
                        {{ seat.row + seat.column }}
                    </div>

                    <!-- Gap between the first and second group of columns -->
                    <div class="seat-gap"></div>

                    <!-- Render last 3 columns (D, E, F) -->
                    <div v-for="seat in row.seats.slice(3, 6)" :key="seat.id"
                        :class="['seat', seat.status, { selected: seat.selected }]" @click="toggleSeatSelection(seat)">
                        {{ seat.row + seat.column }}
                    </div>
                </div>
            </div>

            <div class="seat-container-box-out">

            </div>

        </div>

    </main>

    
</template>

<script>
import { useFormStore, useFlightStore, usePassengerStore } from '@/scripts/usePiniaStorage';

export default {
    name: "SeatSelectionPage",
    data() {
        return {
            seatRows: this.generateSeatMap(),
        };
    },
    setup() {
        const formStore = useFormStore();
        const flightStore = useFlightStore();
        const passengerStore = usePassengerStore();
        return { formStore, flightStore, passengerStore };
    },
    methods: {
        // Generate 30 rows with 6 columns
        generateSeatMap() {
            const rows = [];
            for (let i = 1; i <= 30; i++) {
                const seats = [];
                for (let j = 1; j <= 6; j++) {
                    seats.push({
                        id: `${i}-${j}`,
                        row: i,
                        column: String.fromCharCode(64 + j), // Columns as A, B, C...
                        status: "available", // Default to available (logic for unavailable to be added later)
                        selected: false,
                    });
                }
                rows.push({ rowNumber: i, seats });
            }
            return rows;
        },
        toggleSeatSelection(seat) {
            const maxSeats = parseInt(this.formStore.form.passengers, 10); // Get the number of passengers
            const selectedSeats = this.getSelectedSeatsCount();

            // Allow toggling only if seat is available and selection count is within the limit
            if (seat.status === "available") {
                if (!seat.selected && selectedSeats < maxSeats) {
                    seat.selected = true; // Select seat
                } else if (seat.selected) {
                    seat.selected = false; // Deselect seat
                }
            }
        },
        getSelectedSeatsCount() {
            // Calculate the number of seats currently selected
            return this.seatRows.reduce(
                (total, row) => total + row.seats.filter(seat => seat.selected).length,
                0
            );
        },
        goBack() {
            this.$router.push({
                path: '/search-flight',

            });
            // window.location.href = "/search-flight";
        },
        continuenext() {
            const selectedSeats = this.getSelectedSeatsCount();
            const requiredSeats = parseInt(this.formStore.form.passengers, 10);

            if (selectedSeats === requiredSeats) {
                this.$router.push({
                    path: '/Payment-page',

                });

            } else {
                alert(`Please select ${requiredSeats - selectedSeats} more seat(s) to continue.`);
            }
        },
    },

};
</script>

<style scoped>
/* Seat Container */
.seat-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 20px;
}

/* Airplane Nose */
.seat-container-box-in {
    width: 60%;
    height: 100px;
    background-color: #ffffff;
    border-radius: 50% 50% 0 0;
    /* Rounded top */
    display: flex;
    justify-content: center;
    align-items: center;
    /* margin-bottom: 20px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); */
    position: relative;
}

.seat-container-box-out {
    width: 60%;
    height: 100px;
    background-color: #ffffff;
    border-radius: 0 0 50% 50%;
    /* Rounded top */
    display: flex;
    justify-content: center;
    align-items: center;
    /* margin-bottom: 20px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); */
    position: relative;
}



/* Seat Rows */
.seat-container-box {
    display: flex;
    flex-direction: column;
    /* border: 1px solid #e0e0e0; */

    /* Rounded bottom */
    /* padding: 16px; */
    width: 60%;
    background-color: #ffffff;
    /* box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); */
}

.seat-row {
    display: flex;
    margin: 5px 0;
    align-items: center;
    justify-content: center;
}

/* Seat Styles */
.seat {
    width: 40px;
    height: 40px;
    margin: 5px;
    text-align: center;
    line-height: 35px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
}

.available {
    background-color: #007bff;
    color: white;
}

.unavailable {
    background-color: gray;
    color: white;
    cursor: not-allowed;
}

.selected {
    background-color: green;
    color: white;
}

/* Gap between column groups */
.seat-gap {
    width: 50px;
    /* Adjust gap size here */
}

@media (max-width: 768px) {

    .seat-container-box,
    .seat-container-box-in {
        width: 90%;
    }

    .seat {
        width: 30px;
        height: 30px;
        line-height: 25px;
        font-size: 12px;
    }

    .seat-gap {
        width: 30px;
    }
}
</style>
