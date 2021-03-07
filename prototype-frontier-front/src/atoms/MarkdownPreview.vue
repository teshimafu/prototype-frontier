<template>
  <span class="preview" v-html="toMarkdown"></span>
</template>
<script lang="ts">
import marked from "marked";
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    text: { type: String, required: true }
  },
  computed: {
    toMarkdown(): string {
      return marked(this.text);
    }
  }
});
</script>
