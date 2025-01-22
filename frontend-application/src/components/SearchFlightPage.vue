<template>
    <header>

        <h1> Search For Flights </h1>
    </header>
    <main class="main">
        <div class="content">
            <form @submit.prevent="searchFlight">
                <div class="radio-group-container">
                    <div class="radio-group">
                        <label>
                            <input type="radio" :checked="form.triptype === 'one-way'" value="one-way"
                                @change="updateForm('triptype', $event.target.value)" />
                            One Way
                        </label>
                        <label>
                            <input type="radio" :checked="form.triptype === 'round-trip'" value="round-trip"
                                @change="updateForm('triptype', $event.target.value)" />
                            Round Trip
                        </label>
                    </div>
                </div>
                <div class="form-grp">
                    <div class="form-group" title="Select the starting city (From)">

                        <select :value="form.source" id="source" class="form-control wide-select" required
                            @change="updateForm('source', $event.target.value)">
                            <option value="" disabled selected>Select the starting city (From)</option>
                            <option v-for="city in cities" :key="city.id" :value="city.id">
                                {{ city.name }}
                            </option>
                        </select>
                    </div>
                    <div class="form-group" title="Select the destination city (To)">

                        <select :value="form.destination" id="destination" class="form-control wide-select" required
                            :disabled="!destinations.length" @change="updateForm('destination', $event.target.value)">
                            <option value="" disabled selected>Select the destination city (To)</option>
                            <option v-for="city in destinations" :key="city.id" :value="city.id">
                                {{ city.name }}
                            </option>
                        </select>
                    </div>
                    <div class="form-group" title="Choose the Date of Journey (DOJ)">
                        <input type="date" :value="form.departureDate" id="departure" class="form-control" :min="today"
                            placeholder="Date of Journey (DOJ)" required
                            @input="updateForm('departureDate', $event.target.value)" />
                    </div>

                    <div class="form-group">

                        <input type="date" :value="form.returnDate" id="return" class="form-control"
                            :min="form.departureDate || today" @input="updateForm('returnDate', $event.target.value)"
                            @click="handleReturnDateClick" />
                    </div>
                    <div class="form-group" title="Select the number of passengers">
                        <select :value="form.passengers" id="passengers" class="form-control" required
                            @change="updateForm('passengers', $event.target.value)">
                            <option value="">Select passenger</option>
                            <option v-for="num in 10" :key="num" :value="num">
                                {{ num }}
                            </option>
                        </select>
                    </div>
                </div>
                <div class="search-btn-container">
                    <button @click="goBack()" class="search-btn">Back</button>
                    <button type="submit" class="search-btn">SEARCH</button>
                </div>
            </form>
        </div>
    </main>
   
</template>


<script>
import { DisplayAirportBySource, DisplayAirports } from '@/scripts/GetallAirports';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { useFormStore } from '@/scripts/usePiniaStorage';

export default {
    name: "SearchFlightPage",
    setup() {
        const formStore = useFormStore();
        return { formStore };
    },
    data() {
        return {
            
            form: this.formStore.getFormData(), // Retrieve form data from Pinia store
            cities: [],
            destinations: [],
            today: new Date().toISOString().split("T")[0],

        };
    },
    methods: {
        goBack() {
            window.location.href = "/";
        },
        async fetchAirports() {
            try {
                const response = await DisplayAirports();
                this.cities = response.airports.map((airport) => ({
                    id: airport.ID,
                    name: airport.City,
                }));
            } catch (error) {
                alert("Unable to load airports. Please try again later.");
            }
        },
        async fetchDestinations() {
            if (!this.form.source) return;
            try {
                const response = await DisplayAirportBySource(this.form.source);
                this.destinations = response.routes.map((route) => ({
                    id: route.Destination.ID,
                    name: route.Destination.City,
                }));
            } catch (error) {
                alert("Unable to load destinations. Please try again later.");
            }
        },
        toggleReturnDate() {
            if (this.form.triptype === 'one-way') {
                this.form.returnDate = '';
            }
        },
        handleReturnDateClick() {
            if (this.form.triptype === 'one-way') {
                this.form.triptype = 'round-trip';
            }
        },
        updateForm(field, value) {
            this.form[field] = value;
            if (field === 'source') {
                this.fetchDestinations();
            }
            if (field === 'triptype' && value === 'one-way') {
                this.form.returnDate = '';
            }

            // Save form data to Pinia store
            this.saveFormData();
        },
        saveFormData() {
            // Set form data into Pinia store (replace localStorage)
            this.formStore.setFormData(this.form);
        },
        searchFlight() {
            if (this.form.source === this.form.destination) {
                alert("Source and destination cannot be the same.");
                return;
            }
            if (this.form.triptype === "round-trip" && !this.form.returnDate) {
                alert("Please select a return date for a round-trip.");
                return;
            }

            this.saveFormData(); // Save form data to Pinia store

            if (this.form.triptype === "one-way") {
                toast.success("Redirecting to Flight Result Page", {
                    position: toast.POSITION.TOP_RIGHT,
                    autoClose: 1000,
                });
                setTimeout(() => {
                    this.$router.push("/flight-result");
                }, 1000);
            } else {
                toast.success("Redirecting to Flight Result Page", {
                    position: toast.POSITION.TOP_RIGHT,
                    autoClose: 1000,
                });
                setTimeout(() => {
                    this.$router.push("/flight-result-two");
                }, 1000);
            }
        },
    },
    mounted() {
        this.fetchAirports();
        this.form.triptype = 'one-way';
      
    },
};
</script>



<style scoped>
header h1 {
    text-align: center;
    font-size: 2rem;
    color: var(--c-text-primary, #333);


}


.main {

    width: 1500px;
    justify-content: center;
    align-items: center;
    margin-left: 50px;
    margin-right: 50px;
    text-align: center;
    background-color: #f5f5f5;
}

.wide-select {
    width: 100%;
    min-width: 450px;
    max-width: 500px;
    padding: 10px;
    font-size: 16px;
}

.content {
    padding: 20px;
    height: 160px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.form-grp {
    display: flex;
    gap: 45px;
}

.form-group {
    margin-bottom: 20px;
}


.form-group select,
.form-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 4px;
}

.special-fare {
    display: flex;
    gap: 10px;
    flex-wrap: wrap;
}

.special-fare button {
    flex: 1;
    padding: 10px;
    background: #f3f3f3;
    border: 1px solid #ccc;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.special-fare button.active,
.special-fare button:hover {
    background: #0078d7;
    color: #fff;
}



.radio-group-container {
    text-align: center;
    margin-top: 0px;
    margin-bottom: 20px;
}

.radio-group {
    display: inline-flex;
    gap: 15px;
    padding: 10px;
    background-color: #f3f3f3;
    border: 1px solid #ccc;
    border-radius: 10px;
}

.search-btn-container {
    text-align: center;
    margin-bottom: 20px;
    display: flex;
    gap: 20px;
    margin-left: 450px;
}

.search-btn {

    width: 30%;
    padding: 15px;
    background-color: var(--c-text-primary);
    color: #fff;
    border: none;
    border-radius: 10px;
    font-size: 18px;

    cursor: pointer;
    display: block;



}
</style>
