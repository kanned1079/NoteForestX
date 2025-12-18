<script lang="ts" setup>
import {useScrollFadeIn} from "~/composables/useScrollFadeIn";

useScrollFadeIn({
  selector: '.animated-card-day-sign',
  direction: 'left',
  x: 100,
  stagger: 0,
  duration: 0.5,
  start: 'top 90%'
})

</script>

<template>
  <i class="pi pi-sun text-xl opacity-80 animated-card-day-sign"></i>
</template>

<style scoped>

</style>