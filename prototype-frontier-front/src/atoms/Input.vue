<template>
  <input class="form-control" v-model="state.value" />
</template>

<script lang="ts">
import { defineComponent, watch, reactive } from "vue";
export default defineComponent({
  props: {
    input: { type: String, require: true, default: "" }
  },
  setup(props, context) {
    const state = reactive<{ value: string }>({ value: props.input });
    watch(
      () => state.value,
      value => {
        context.emit("model", value);
      }
    );
    return { state };
  }
});
</script>

<style scoped>
.form-control {
  display: inline-block;
  padding: 0.375rem 0.75rem;
  font-weight: 400;
  border: 1px solid #ced4da;
  border-radius: 0.25rem;
}
</style>
