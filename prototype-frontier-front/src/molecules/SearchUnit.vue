<template>
  <span class="search-unit">
    <Input class="input" :input="state.value" @model="model" />
    <Button :text="'検索'" :type="'success'" :size="'small'" @click="search" />
    <Button
      :text="'リセット'"
      :type="'secondary'"
      :size="'small'"
      @click="reset"
    />
  </span>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import Input from "@/atoms/Input.vue";
import Button from "@/atoms/Button.vue";

export default defineComponent({
  components: { Input, Button },
  props: {
    input: { type: String, require: true, default: "" }
  },
  setup(props, context) {
    const state = reactive<{ value: string }>({ value: props.input });
    const model = (text: string) => {
      state.value = text;
    };
    const search = () => {
      context.emit("search", state.value);
    };
    const reset = () => {
      state.value = "";
      context.emit("search", state.value);
    };
    return { state, model, search, reset };
  }
});
</script>

<style scoped>
.input {
  width: 20rem;
  margin-right: auto;
}
</style>
