<template>
  <AppTemplate />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import AppTemplate from "@/templates/App.vue";

export default defineComponent({
  name: "App",
  components: {
    AppTemplate
  }
});
</script>
