<script setup lang="ts">
import Ulli from "~/components/Ulli.vue"
import SkillRepresent from "~/components/SkillRepresent.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
// 仅新增：导入useI18n（不修改其他导入和逻辑）
import { useI18n } from "vue-i18n"

// 仅新增：获取t函数（不修改其他代码）
const { t } = useI18n()

// 原有代码：完全不变
useScrollFadeIn({
  selector: '.animate-card-mywork',
  direction: 'up',      // 默认其实可以不写
  y: 50,
  duration: 1,
  stagger: 0.1,
  start: 'top 85%',
  toggleActions: 'play none none reverse'
})
</script>

<template>
  <div class="flex justify-center">
    <div class="relative max-w-[1000px] container z-10 px-4 pt-20 text-slate-800 dark:text-slate-100">

      <!--      <Button variant="text" class="mt-4 mb-4" label="技能 & 工具" icon="pi pi-wrench" />-->

      <!--      <p class="mb-4 underline">轉到 <span class="font-bold">技能 & 工具</span> <span class="pi pi-wrench"></span> </p>-->

      <!-- 🌟 作品标题 -->
      <div class="mb-6">
        <!-- 仅替换中文为t函数调用，其他不变 -->
        <p class="text-2xl font-bold text-[#3261e4] dark:text-[#4f77e6]">{{ t("my_work.title_tag") }}</p>
        <p class="text-7xl font-bold mt-5 mb-4">{{ t("my_work.main_title") }}</p>
        <p class="text-xl">{{ t("my_work.main_desc") }}</p>
      </div>

      <!--      <Divider class="mt-6 mb-6" />-->

      <!-- 简介 -->
      <div class="animate-card-mywork mb-6">
        <p>{{ t("my_work.intro") }}</p>
      </div>

      <!-- App Developer -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">{{ t("my_work.app_dev_title") }}</p>
        <p class="mt-2 mb-2">{{ t("my_work.app_dev_desc1") }}</p>
        <p class="mt-2 mb-2">{{ t("my_work.app_dev_desc2") }}</p>

        <!-- 仅替换value为t函数调用，组件属性和结构不变 -->
        <Ulli :value="t('my_work.app_dev_item1')" bold />
        <Ulli :value="t('my_work.app_dev_item2')" bold />
        <Ulli :value="t('my_work.app_dev_item3')" bold />
        <Ulli :value="t('my_work.app_dev_item4')" bold />
        <Ulli :value="t('my_work.app_dev_item5')" />
      </div>

      <div class="animate-card-mywork mb-6">
        <p>{{ t("my_work.tool_intro") }}</p>
      </div>

      <!-- 技術棧 -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">{{ t("my_work.vue_tech_title") }}</p>
        <p class="mt-2 mb-2 text-slate-800 dark:text-slate-400">
          {{ t("my_work.vue_tech_desc1") }}
        </p>
        <p class="mt-2 mb-2">
          {{ t("my_work.vue_tech_desc2") }}
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">{{ t("my_work.backend_tech_title") }}</p>
        <p class="mt-2 mb-2 text-slate-800 dark:text-slate-400">
          {{ t("my_work.backend_tech_desc") }}
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">{{ t("my_work.ops_title") }}</p>
        <p class="mt-2 mb-2">
          {{ t("my_work.ops_desc") }}
        </p>
      </div>

      <!--      <Divider class="mt-10 mb-10 animate-card" />-->

      <!-- Detail and Summary -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">{{ t("my_work.detail_summary_title") }}</p>
        <p class="mt-2 mb-2">{{ t("my_work.detail_summary_desc") }}</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <!-- 仅替换value为t函数调用，其他不变 -->
          <SkillRepresent :value="t('my_work.usage_example1')" :percent="90" />
          <SkillRepresent :value="t('my_work.usage_example2')" :percent="23" />
        </div>
      </div>

      <!-- Programming Languages -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">{{ t("my_work.program_lang_title") }}</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <!-- 仅替换value为t函数调用，其他不变 -->
          <SkillRepresent :value="t('my_work.lang_ts')" :percent="100" />
          <SkillRepresent :value="t('my_work.lang_js')" :percent="90" />
          <SkillRepresent :value="t('my_work.lang_go')" :percent="100" />
          <SkillRepresent :value="t('my_work.lang_java')" :percent="30" />
          <SkillRepresent :value="t('my_work.lang_shell')" :percent="50" />
        </div>
      </div>

      <!-- Technologies -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">{{ t("my_work.tech_tools_title") }}</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <!-- 仅替换value为t函数调用，其他不变 -->
          <SkillRepresent :value="t('my_work.tool_tailwind')" :percent="90" />
          <SkillRepresent :value="t('my_work.tool_naive')" :percent="50" />
          <SkillRepresent :value="t('my_work.tool_prime')" :percent="60" />
          <SkillRepresent :value="t('my_work.tool_material')" :percent="30" />
          <SkillRepresent :value="t('my_work.tool_vue')" :percent="90" />
          <SkillRepresent :value="t('my_work.tool_nuxt')" :percent="80" />
          <SkillRepresent :value="t('my_work.tool_electron')" :percent="10" />
          <SkillRepresent :value="t('my_work.tool_pinia')" :percent="90" />
          <SkillRepresent :value="t('my_work.tool_react')" :percent="30" />
          <SkillRepresent :value="t('my_work.tool_gin')" :percent="100" />
          <SkillRepresent :value="t('my_work.tool_gorm')" :percent="100" />
          <SkillRepresent :value="t('my_work.tool_grpc')" :percent="20" />
          <SkillRepresent :value="t('my_work.tool_redis')" :percent="80" />
          <SkillRepresent :value="t('my_work.tool_mysql')" :percent="80" />
          <SkillRepresent :value="t('my_work.tool_mongo')" :percent="10" />
          <SkillRepresent :value="t('my_work.tool_docker')" :percent="95" />
          <SkillRepresent :value="t('my_work.tool_nginx')" :percent="80" />
        </div>
      </div>

      <!-- Softwares and Tools -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">{{ t("my_work.software_tools_title") }}</p>
        <div class="flex flex-wrap gap-2 justify-start">
          <!-- 仅替换value为t函数调用，其他不变 -->
          <SkillRepresent :value="t('my_work.software_idea')" :percent="95" />
          <SkillRepresent :value="t('my_work.software_vscode')" :percent="80" />
          <SkillRepresent :value="t('my_work.software_vim')" :percent="40" />
          <SkillRepresent :value="t('my_work.software_git')" :percent="80" />
          <SkillRepresent :value="t('my_work.software_mysqlworkbench')" :percent="70" />
          <SkillRepresent :value="t('my_work.software_ps')" :percent="80" />
          <SkillRepresent :value="t('my_work.software_ai')" :percent="30" />
        </div>
      </div>

    </div>
  </div>
</template>


<style scoped>

</style>