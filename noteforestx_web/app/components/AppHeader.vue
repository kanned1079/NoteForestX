<script lang="ts" setup>
import {ref} from "vue";

const selectedCountry = ref();
const countries = ref([
  {name: '🇨🇳 简体中文(中國)', code: 'cn'},
  {name: '🇯🇵 日本語', code: 'jp'},
  {name: '🇪🇸 Spain', code: 'es'},
  {name: '🇺🇸 English(US)', code: 'en'}
]);

const onClickLogoButton = () => {
  console.log("pressed")
}

const menu = ref();
const items = ref([
  {
    label: "🇨🇳 简体中文(中國)"
  },
  {
    label: "🇯🇵 日本語"
  },
  {
    label: "🇪🇸 Spain"
  },
  {
    label: "🇺🇸 English(US)"
  },
]);

const toggle = (event) => {
  menu.value.toggle(event);
};

function toggleDarkMode() {
  document.documentElement.classList.toggle('my-app-dark');
}

</script>

<template>
  <div class="p-3 flex-row flex justify-between items-center">
    <div class="ml-3">
      <Button
          class="ml-3"
          size="small"
          icon="pi pi-home"
          label="NoteForest X"
          severity="secondary" variant="primary"
      >
      </Button>
    </div>
    <div class="mr-3">



        <Button size="small"  type="button" label="Toggle" @click="toggle" aria-haspopup="true" aria-controls="overlay_tmenu" />
        <TieredMenu ref="menu" id="overlay_tmenu" :model="items" popup />



      <Button
          class="ml-3"
          size="small"
          icon="pi pi-sun"
          label="Sun"
          severity="secondary" variant="primary"
          @click="toggleDarkMode"
      >
      </Button>

    </div>
  </div>

</template>