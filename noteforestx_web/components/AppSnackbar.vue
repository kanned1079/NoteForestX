<!-- AppSnackbar.vue -->
<template>
  <v-snackbar
      v-model="visible"
      :color="color"
      :timeout="timeout"
      :location="location"
      :multi-line="true"
  >
    {{ message }}
  </v-snackbar>
</template>

<script setup>
import { ref, defineExpose } from 'vue'

const visible = ref(false)
const message = ref('')
const color = ref('info')
const timeout = ref(3000)
const location = ref('top')

function show(msg, msgColor = 'info', duration = 3000, position = 'top') {
  message.value = msg
  color.value = msgColor
  timeout.value = duration
  location.value = position
  visible.value = true
}

defineExpose({ show })
</script>