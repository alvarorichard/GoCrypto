<template>
    <div>
        <form @submit.prevent="submitForm">
            <div>
                <label for="name">Name:</label>
                <input type="text" id="name" v-model="coin.Name">
            </div>

            <div>
                <label for="symbol">Symbol:</label>
                <input type="text" id="symbol" v-model="coin.Symbol">
            </div>

            <div>
                <label for="price">Price:</label>
                <input type="number" id="price" v-model="coin.Price">
            </div>

            <button type="submit">Create Coin</button>
        </form>
    </div>
    <div>
        <ul>
            <li v-for="coin in coins" :key="coin.ID">
                {{ coin.name }} - {{ coin.symbol }} - {{ coin.price }}
            </li>
        </ul>
    </div>
</template>
  
<script>
import axios from 'axios';
import { ref } from 'vue';
import api from '../api';
export default {
    name: 'Admin',
    data() {
        return {
            coin: {
                Name: '',
                Symbol: '',
                Price: 0
            },
            coins: ref([])
        };
    },
    methods: {
        async submitForm() {
            try {
                const response = await api.post('/api/admin/coin', this.coin, {
                    headers: {
                        Authorization: `Bearer ${localStorage.getItem('token')}`
                    }
                });
                this.coins.push(response.data);
            } catch (error) {
                alert(error);
            }
        }
    },
    mounted() {
        api.get('/api/admin/coin', {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`
            }
        }).then((res) => {
            this.coins = res.data;
        });
    }
};
</script>
  