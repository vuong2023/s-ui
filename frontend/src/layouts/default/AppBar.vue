<template>
  <v-app-bar :elevation="5">
    <v-icon v-if="isMobile" icon="mdi-menu" @click="$emit('toggleDrawer')" />
    <v-app-bar-title :text="$t(<string>$router.currentRoute.value.name)" class="align-center text-center " />
    <v-btn prepend-icon="mdi-content-save" v-if="stateChange" :text="$t('actions.save')" @click="saveChanges"></v-btn>
    <v-icon icon="mdi-theme-light-dark" @click="toggleTheme()" style="margin: 0 10px;"></v-icon>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref,watch } from "vue"
import { useTheme } from "vuetify"
import { FindDiff } from "@/plugins/utils"
import Data from "@/store/modules/data"

defineProps(['isMobile'])

const theme = useTheme()
const darkMode = ref(localStorage.getItem('theme') == "dark")

const store = Data()

const toggleTheme = () => {
  darkMode.value = !darkMode.value
  theme.global.name.value = darkMode.value ? "dark" : "light"
  localStorage.setItem('theme', theme.global.name.value)
}

const saveChanges = () => {
  store.pushData()
}

const oldData = computed((): any => {
  return {config: store.oldData.config, clients: store.oldData.clients}
})

const newData = computed((): any => {
  return {config: store.config, clients: store.clients}
})

const stateChange = computed((): any => {
  return !FindDiff.deepCompare(newData.value,oldData.value)
})
</script>
