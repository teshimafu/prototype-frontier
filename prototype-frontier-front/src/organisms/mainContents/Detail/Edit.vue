<template>
  <div v-if="!state.isLoading">
    <ul class="editor-box">
      <li>
        <InputUnit
          :title="'タイトル'"
          :input="state.portfolio.title"
          @model="setInputValue('title', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'作成者'"
          :input="state.portfolio.author"
          @model="setInputValue('author', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'成果物リンク'"
          :input="state.portfolio.link"
          @model="setInputValue('link', $event)"
        />
      </li>
      <li>
        <InputUnit
          :title="'ソースコード'"
          :input="state.portfolio.source"
          @model="setInputValue('source', $event)"
        />
      </li>
    </ul>
    <div>
      <div>詳細説明</div>
      <MarkdownEditor
        :markdownText="state.portfolio.readme"
        @returnText="setReadme"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, reactive } from "vue";
import { useRouter } from "vue-router";
import InputUnit from "@/molecules/InputUnit.vue";
import MarkdownEditor from "@/molecules/MarkdownEditor.vue";
import { InputPortfolio } from "../../../models/portfolio";
import FirebaseService from "../../../services/firebaseService";
export default defineComponent({
  components: {
    InputUnit,
    MarkdownEditor
  },
  props: {
    inputPortfolio: Object as PropType<InputPortfolio>
  },
  setup(props) {
    const router = useRouter();
    const fs = FirebaseService.getInstance();
    if (!props.inputPortfolio) {
      router.push({ name: "NotFoundError" });
      return {};
    }
    const state = reactive<{
      portfolio: InputPortfolio;
      isLoading: boolean;
    }>({ portfolio: props.inputPortfolio, isLoading: true });

    //methods
    const setReadme = function(text: string) {
      state.portfolio.readme = text;
    };
    const setInputValue = function(
      name: keyof Omit<InputPortfolio, "id" | "created_at">,
      text: string
    ) {
      state.portfolio[name] = text;
    };

    fs.getDisplayName().then(displayName => {
      if (!displayName) {
        router.push({ name: "LoadError" });
        return;
      }
      if (!state.portfolio.author) {
        state.portfolio.author = displayName;
      }
      state.isLoading = false;
    });
    return { state, setReadme, setInputValue };
  }
});
</script>

<style scoped>
ul li {
  list-style: none;
  margin: 0.2rem;
}
.editor-box {
  width: 80%;
  margin: auto;
}
</style>
