<template>
    <div>
        <h1>
            {{ text }}
        </h1>
        <slot />
        <slot name="part1_header" :param1="data.param_from_slot"></slot>
    </div>
</template>
<script setup>
import { reactive } from 'vue';

const data = reactive({
    param_from_slot: 'Hii from part1 header to u<3'
})

defineProps({
    text: {
        type: String,
        default: ''
    }
})
</script>