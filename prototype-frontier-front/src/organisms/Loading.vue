<template>
  <div>{{ state.message }}</div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";

export default defineComponent({
  data() {
    const state = reactive<{ message: string; intervalHandler: number }>({
      message: "now loading.",
      intervalHandler: setInterval(() => {
        state.message += ".";
      }, 1000)
    });
    return { state };
  },
  beforeUnmount() {
    clearInterval(this.state.intervalHandler);
  }
});
</script>
