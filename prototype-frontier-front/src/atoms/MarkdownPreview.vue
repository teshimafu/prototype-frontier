<template>
  <span v-html="toMarkdown"></span>
</template>
<script lang="ts">
import marked from "marked";
import { defineComponent } from "vue";

export default defineComponent({
  //   props: {
  //     text: String
  //   },
  //   computed: {
  //     toMarkdown() {
  //       return marked(this.text ?? "");
  //     }
  //   }
});
</script>
