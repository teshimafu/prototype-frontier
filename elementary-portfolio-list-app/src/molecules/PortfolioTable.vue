<template>
  <table>
    <thead>
      <tr>
        <th>id</th>
        <td>タイトル</td>
        <td>登録者</td>
        <td>作成日</td>
      </tr>
    </thead>
    <tbody>
      <tr
        class="portfolio-link"
        v-for="data in portfolioList"
        :key="data.id"
        v-on:click="onClickTitle(data.id)"
      >
        <th>{{ data.id }}</th>
        <td>{{ data.title }}</td>
        <td>{{ data.author }}</td>
        <td>{{ data.created_at }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Portfolio } from "../models/portfolio2";
import router from "../router";

export default defineComponent({
  name: "PortolioTable",
  props: {
    portfolioList: Array as PropType<Portfolio[]>
  },
  methods: {
    onClickTitle(id: number) {
      router.push({ path: `/detail/${id}` });
    }
  }
});
</script>

<style>
table {
  width: 100%;
  border: 1px solid #ccc;
  border-collapse: collapse;
}

th,
td {
  padding: 5px;
  border: 1px solid #ccc;
}

thead tr {
  color: #fff;
  background-color: gray;
}

tbody tr:nth-of-type(odd) {
  background-color: #efefef;
}

.portfolio-link {
  cursor: pointer;
}
</style>
