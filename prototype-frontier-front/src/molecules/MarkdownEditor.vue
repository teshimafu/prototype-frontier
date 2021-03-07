<template>
  <form>
    <table class="markdown-editor">
      <thead>
        <tr>
          <th>
            編集
          </th>
          <th>
            プレビュー
          </th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>
            <span class="editor">
              <textarea v-model="text"></textarea>
            </span>
          </td>
          <td>
            <div class="preview">
              <MarkdownPreview :text="text" />
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </form>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import MarkdownPreview from "@/atoms/MarkdownPreview.vue";

export default defineComponent({
  components: { MarkdownPreview },
  props: {
    markdownText: { type: String, required: true }
  },
  data() {
    return { text: this.markdownText };
  },
  watch: {
    text: function() {
      this.$emit("returnText", this.text);
    }
  }
});
</script>

<style scoped>
.markdown-editor {
  height: 50vh;
}
td {
  width: 50%;
}
.editor {
  overflow: scroll;
}
.editor > textarea {
  resize: none;
  width: 100%;
  height: 100%;
}

.preview {
  height: 100%;
  overflow: scroll;
}
.preview > span {
  width: 100%;
  height: 100%;
  vertical-align: top;
  text-align: left;
}
</style>
