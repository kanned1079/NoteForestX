<script setup lang="ts">
import { ref, watch, onMounted, nextTick, defineProps, defineEmits } from 'vue'
import { gsap } from 'gsap'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  duration: { type: Number, default: 0.4 }, // 秒
  easing: { type: String, default: 'power2.out' },
  appear: { type: Boolean, default: false },
  keepDisplay: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue'])
const containerRef = ref<HTMLElement | null>(null)

const doEnter = async () => {
  const el = containerRef.value
  if (!el) return

  el.style.display = ''
  el.style.overflow = 'hidden'
  const targetHeight = el.scrollHeight
  gsap.fromTo(
      el,
      { height: 0, opacity: 0 },
      {
        height: targetHeight,
        opacity: 1,
        ease: props.easing,
        duration: props.duration,
        onComplete: () => {
          el.style.height = ''
          el.style.overflow = ''
        }
      }
  )
}

const doLeave = async () => {
  const el = containerRef.value
  if (!el) return

  el.style.overflow = 'hidden'
  const currentHeight = el.scrollHeight
  gsap.fromTo(
      el,
      { height: currentHeight, opacity: 1 },
      {
        height: 0,
        opacity: 0,
        ease: props.easing,
        duration: props.duration,
        onComplete: () => {
          if (!props.keepDisplay) el.style.display = 'none'
          el.style.height = ''
          el.style.overflow = ''
        }
      }
  )
}

watch(
    () => props.modelValue,
    async (val) => {
      await nextTick()
      if (val) doEnter()
      else doLeave()
    }
)

onMounted(async () => {
  const el = containerRef.value
  if (!el) return

  if (props.modelValue) {
    el.style.display = ''
    if (props.appear) {
      await nextTick()
      doEnter()
    }
  } else {
    el.style.display = props.keepDisplay ? '' : 'none'
  }
})
</script>

<template>
  <div ref="containerRef" class="overflow-hidden">
    <slot />
  </div>
</template>