<script>
import { defineComponent } from 'vue';
import store from '../store';
export default defineComponent({
    name: 'Home',
    setup() {
        return {
        user: store.state.user,
        }
    },
})
</script>
<template>
    <h1>Home</h1>
    <h1>
        User: {{ user }}
    </h1>
</template>