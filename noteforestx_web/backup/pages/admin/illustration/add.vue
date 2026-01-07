<script lang="ts" setup>
const num = ref<number>(0)
</script>

<template>
  <div>
    <h1>{{ num }}</h1>
    <button @click="num++">+</button>
    <button @click="num--">-</button>
  </div>
</template>