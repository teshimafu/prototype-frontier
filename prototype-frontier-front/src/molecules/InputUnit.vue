<template>
  <span class="input-unit">
    <label class="title">{{ title }}:</label>
    <Input class="input" :input="state.value" @model="model" />
  </span>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import Input from "@/atoms/Input.vue";

export default defineComponent({
  components: { Input },
  props: {
    title: { type: String, default: "" },
    input: { type: String, require: true, default: "" }
  },
  setup(props, context) {
    const state = reactive<{ value: string }>({ value: props.input });
    const model = (text: string) => {
      context.emit("model", text);
    };
    return { state, model };
  }
});
</script>

<style scoped>
.title {
  width: 20%;
  margin-right: 1.5rem;
  float: left;
  text-align: right;
}
.input {
  width: 70%;
  margin-right: auto;
}
</style>
