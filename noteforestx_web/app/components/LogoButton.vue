<script setup lang="ts">
import {Icon} from "@vicons/utils";
import {TicketSharp} from "@vicons/ionicons5";
const props = defineProps<{
  click?: () => void | null
}>()
</script>

<template>
  <Button
      size="medium"
      label="NoteForest X"
      severity="secondary"
      variant="text"
      @click="props.click?props.click():null"
  >
    <template #icon>
      <Icon><TicketSharp/></Icon>
    </template>

  </Button>
</template>

<style scoped>

</style>