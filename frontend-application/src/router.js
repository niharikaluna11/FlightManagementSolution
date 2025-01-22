import { createRouter, createWebHistory } from 'vue-router';
import MainPage from './components/MainPage.vue';
import SearchFlightPage from './components/SearchFlightPage.vue';
import NotFoundPage from './components/NotFoundPage.vue';
import FlightResultPage from './components/FlightResultPage.vue';
import PassengerPage from './components/PassengerPage.vue';
import PaymentPage from './components/PaymentPage.vue';
import SeatSelectionPage from './components/SeatSelectionPage.vue';
import FlightResult2Page from './components/FlightResult2Page.vue';
import NoFlightPage from './components/NoFlightPage.vue';
import BookingResultPage from './components/BookingResultPage.vue';

const routes = [
    {
        path: '/',
        name: 'MainPage',
        component: MainPage,
        meta: { requiresAuth: true, }
    },
    {
        path: '/search-flight',
        name: 'SearchFlight',
        component: SearchFlightPage,
        meta: { requiresAuth: true, }
    },
    {
        path: '/not-found',
        name: 'NotFoundPage',
        component: NotFoundPage,

    },
    {
        path: '/flight-result',
        name: 'FlightResultPage',
        component: FlightResultPage,

    },
    {
        path: '/passenger-page',
        name: 'PassengerPage',
        component: PassengerPage,

    },
    {
        path: '/Payment-page',
        name: 'PaymentPage',
        component: PaymentPage,

    },
    {
        path: '/seat-selection-page',
        name: 'SeatSelectionPage',
        component: SeatSelectionPage,

    },
    {
        path: '/flight-result-two',
        name: 'FlightResult2Page',
        component: FlightResult2Page,

    },
    {
        path: '/no-flights-page',
        name: 'NoFlightPage',
        component: NoFlightPage,

    },
    {
        path: '/booking-confirm',
        name: 'BookingResultPage',
        component: BookingResultPage,

    },
    //    name: "BookingResultPage",
    //seat-selection-page
    //PaymentPage
    { path: '/:pathMatch(.*)*', redirect: '/not-found' }
    //no-flights-page
    ///flight-result-two

];

const router = createRouter({
    history: createWebHistory(),
    routes
});


export default router;
