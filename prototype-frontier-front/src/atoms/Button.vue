<template>
  <button class="btn" :class="[type, size]">
    {{ text }}
  </button>
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
        return ["primary", "secondary", "success", "warning", "danger"].some(
          v => v === value
        );
      }
    },
    size: {
      type: String,
      default: ""
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
.small {
  font-size: small;
  padding: 0.2rem 0.4rem;
}
.large {
  font-size: large;
}
.btn:disabled {
  opacity: 0.65;
}

.primary {
  color: #fff;
  background-color: #34b5ff;
  border-color: #34b5ff;
}

.secondary {
  color: #fff;
  background-color: #6c757d;
  border-color: #6c757d;
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
