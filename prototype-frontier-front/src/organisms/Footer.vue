<template>
  <div class="footer">
    <div>
      アクセス数:<span>{{ state.count ? state.count : "" }}</span>
    </div>
    <div id="copylight">created by teshima.fu@gmail.com</div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive } from "vue";
import firebaseService from "../services/firebaseService";

export default defineComponent({
  setup() {
    const fs = firebaseService.getInstance();
    const state = reactive<{ count: number | null }>({
      count: null
    });
    fs.getAccessCount().then(count => (state.count = count));
    return { state };
  }
});
</script>
>
<style>
.footer {
  background-color: black;
  color: white;
  padding: 0;
  width: 100%;
  position: absolute;
  bottom: 0;
}
#copylight {
  font-size: small;
  margin: auto;
}
</style>
