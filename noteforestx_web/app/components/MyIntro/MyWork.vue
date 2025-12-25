<script setup lang="ts">
import Ulli from "~/components/Ulli.vue"
import SkillRepresent from "~/components/SkillRepresent.vue"
import { useScrollFadeIn } from "~/composables/useScrollFadeIn"
import { useI18n } from "vue-i18n"

const { t } = useI18n()

// 原有代码：完全不变
useScrollFadeIn({
  selector: '.animate-card-mywork',
  direction: 'up',
  y: 50,
  duration: 1,
  stagger: 0.1,
  start: 'top 85%',
  toggleActions: 'play none none reverse'
})
</script>

<template>
  <div class="flex justify-center">
    <!-- 页面文字基准色 -->
    <div
        class="relative max-w-[1000px] container z-10 px-4 pt-20
             text-slate-800 dark:text-slate-100"
    >

      <!-- 🌟 作品标题 -->
      <div class="mb-6">
        <p class="text-2xl font-bold text-[#3261e4] dark:text-[#4f77e6]">
          {{ t("my_work.title_tag") }}
        </p>
        <p class="text-7xl font-bold mt-5 mb-4">
          {{ t("my_work.main_title") }}
        </p>
        <p class="text-xl opacity-85">
          {{ t("my_work.main_desc") }}
        </p>
      </div>

      <!-- 简介 -->
      <div class="animate-card-mywork mb-6 opacity-85">
        <p>{{ t("my_work.intro") }}</p>
      </div>

      <!-- App Developer -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">
          {{ t("my_work.app_dev_title") }}
        </p>

        <p class="mt-2 mb-2 opacity-85">
          {{ t("my_work.app_dev_desc1") }}
        </p>
        <p class="mt-2 mb-2 opacity-85">
          {{ t("my_work.app_dev_desc2") }}
        </p>

        <Ulli :value="t('my_work.app_dev_item1')" bold />
        <Ulli :value="t('my_work.app_dev_item2')" bold />
        <Ulli :value="t('my_work.app_dev_item3')" bold />
        <Ulli :value="t('my_work.app_dev_item4')" bold />
        <Ulli :value="t('my_work.app_dev_item5')" />
      </div>

      <div class="animate-card-mywork mb-6 opacity-85">
        <p>{{ t("my_work.tool_intro") }}</p>
      </div>

      <!-- 技术栈 -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">
          {{ t("my_work.vue_tech_title") }}
        </p>
        <p class="mt-2 mb-2 opacity-70">
          {{ t("my_work.vue_tech_desc1") }}
        </p>
        <p class="mt-2 mb-2 opacity-85">
          {{ t("my_work.vue_tech_desc2") }}
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">
          {{ t("my_work.backend_tech_title") }}
        </p>
        <p class="mt-2 mb-2 opacity-70">
          {{ t("my_work.backend_tech_desc") }}
        </p>
      </div>

      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">
          {{ t("my_work.ops_title") }}
        </p>
        <p class="mt-2 mb-2 opacity-85">
          {{ t("my_work.ops_desc") }}
        </p>
      </div>

      <!-- Detail and Summary -->
      <div class="animate-card-mywork mb-6">
        <p class="text-3xl font-bold mt-5 mb-4">
          {{ t("my_work.detail_summary_title") }}
        </p>
        <p class="mt-2 mb-2 opacity-85">
          {{ t("my_work.detail_summary_desc") }}
        </p>

        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent :value="t('my_work.usage_example1')" :percent="90" />
          <SkillRepresent :value="t('my_work.usage_example2')" :percent="23" />
        </div>
      </div>

      <!-- Programming Languages -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">
          {{ t("my_work.program_lang_title") }}
        </p>

        <div class="flex flex-wrap gap-2 justify-start">
          <SkillRepresent :value="t('my_work.lang_ts')" :percent="100" />
          <SkillRepresent :value="t('my_work.lang_js')" :percent="90" />
          <SkillRepresent :value="t('my_work.lang_go')" :percent="100" />
          <SkillRepresent :value="t('my_work.lang_java')" :percent="30" />
          <SkillRepresent :value="t('my_work.lang_shell')" :percent="50" />
        </div>
      </div>

      <!-- Technologies -->
      <div class="animate-card-mywork mb-6">
        <p class="text-2xl font-bold mt-5 mb-4">
          {{ t("my_work.tech_tools_title") }}
        </p>

        <div class="flex flex-wrap gap-2 justify-start">
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
        <p class="text-2xl font-bold mt-5 mb-4">
          {{ t("my_work.software_tools_title") }}
        </p>

        <div class="flex flex-wrap gap-2 justify-start">
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