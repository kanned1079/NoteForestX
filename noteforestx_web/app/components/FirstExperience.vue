<script setup lang="ts">
import { ref } from 'vue'
import { useLoop } from '@tresjs/core'

// 立方体和小球引用
const cubeRef = ref()
const sphereRef = ref()

// 动画循环
const { onBeforeRender } = useLoop()
onBeforeRender(({ elapsed }) => {
  if (cubeRef.value) {
    cubeRef.value.rotation.x = elapsed * 0.6
    cubeRef.value.rotation.y = elapsed * 0.4
  }
  if (sphereRef.value) {
    sphereRef.value.position.y = Math.sin(elapsed * 2) // 小球上下浮动
  }
})
</script>

<template>
  <!-- 小型透视相机，右侧可视化使用 -->
  <TresPerspectiveCamera
      :position="[3, 3, 5]"
      :look-at="[0, 0, 0]"
  />

  <!-- 旋转立方体 -->
  <TresMesh ref="cubeRef" :position="[0, 0, 0]">
    <TresBoxGeometry :args="[1, 1, 1]" />
    <TresMeshStandardMaterial color="#ff6b6b" />
  </TresMesh>

  <!-- 浮动小球 -->
  <TresMesh ref="sphereRef" :position="[2, 0, 0]">
    <TresSphereGeometry :args="[0.5, 32, 32]" />
    <TresMeshStandardMaterial color="#1dd1a1" />
  </TresMesh>

  <!-- 灯光 -->
  <TresDirectionalLight :position="[5, 5, 5]" :intensity="1" />
  <TresAmbientLight :intensity="0.3" />

  <!-- 辅助工具，可在开发阶段开启 -->
  <TresAxesHelper :size="2" />
  <TresGridHelper :size="5" :divisions="5" />
</template>

<style scoped>
/* 让 TresCanvas 只占右侧小区域 */
:host {
  width: 200px;
  height: 200px;
}
</style>