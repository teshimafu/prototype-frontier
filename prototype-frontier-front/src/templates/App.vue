<template>
  <div>
    <Header />
    <MainContent />
    <Footer />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Header from "@/organisms/Header.vue";
import Footer from "@/organisms/Footer.vue";
import MainContent from "@/organisms/MainContent.vue";
export default defineComponent({
  components: {
    Header,
    Footer,
    MainContent
  }
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  min-height: calc(100vh);
  position: relative;
  padding-top: 70px;
  padding-bottom: calc(1em);
  box-sizing: border-box;
}
</style>
