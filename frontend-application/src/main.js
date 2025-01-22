

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPersist from 'pinia-plugin-persistedstate';
import App from './App.vue';
import router from './router';

const pinia = createPinia();


pinia.use(piniaPersist);

// const app = createApp(App);
// app.use(createPinia());
// app.use(router);
// app.mount('#app');

createApp(App)
    .use(pinia)
    .use(router)
    .mount('#app'); 
