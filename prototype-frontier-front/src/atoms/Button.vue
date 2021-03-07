<template>
  <button class="btn" :class="type">{{ text }}</button>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  props: {
    text: { type: String, default: "ボタン" },
    type: {
      type: String,
      default: "primary",
      validator: function(value: string) {
        return ["primary", "success", "warning", "danger"].some(
          v => v === value
        );
      }
    }
  }
});
</script>
<style scoped>
.btn {
  cursor: pointer;
  display: inline-block;
  font-weight: 400;
  text-align: center;
  font-size: 1rem;
  border: 1px solid transparent;
  padding: 0.375rem 0.75rem;
  border-radius: 0.25rem;
}
.btn:disabled {
  opacity: 0.65;
}

.primary {
  color: #fff;
  background-color: #34b5ff;
  border-color: #34b5ff;
}

.success {
  color: #fff;
  background-color: #28a745;
  border-color: #28a745;
}

.danger {
  color: #fff;
  background-color: #dc3545;
  border-color: #dc3545;
}
</style>
