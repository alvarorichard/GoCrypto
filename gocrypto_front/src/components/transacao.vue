<script>
import { defineComponent } from "vue";
export default defineComponent({
    name: "Trasacao",
    props: {
        tx: {
            time: String,
            ammount: Number,
            coin: String,
            operation: String,
        }
    },
    setup(props) {
        console.log(props)
    }
});
</script>

<template>
    <tr>
        <td class="flex items-center gap-2">{{ tx.id }}</td>
        <td class="flex items-center gap-2">>{{ tx.time }}</td>
        <td class="flex items-center gap-2">>{{ tx.coin }}</td>
        <td class="flex items-center gap-2">>{{ tx.amount }}</td>
        <td class="flex items-center gap-2">
            <span v-if="tx.operation == 'buy'" class="text-green-600">{{ tx.operation }}</span>
            <span v-if="tx.operation == 'sell'" class="text-red-600">{{ tx.operation }}</span>
        </td>
    </tr>
</template>