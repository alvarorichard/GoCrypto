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
        <th>Transação</th>
        <th>Data</th>
        <th>Tipo</th>
        <th>Valor</th>
        <th>Status</th>
    </tr>
    <tr>
        <td>Sexo</td>
        <td>{{ tx.time }}</td>
        <td>{{ tx.coin }}</td>
        <td>{{ tx.ammount }}</td>
        <td class="text-red-600 flex items-center gap-2">
            <span v-if="tipo == 'sell'" class="w-[10px] h-[10px] rounded-full bg-red-600">Venda</span>
            <span v-if="tipo == 'buy'" class="w-[10px] h-[10px] rounded-full bg-green-600">Compra</span>
        </td>
    </tr>
</template>