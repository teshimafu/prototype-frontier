<template>
  <div>
    <Loading v-if="state.isLoading" />
    <div v-else-if="state.isEdit">
      <Edit :inputPortfolio="state.portfolio" />
      <Button @click="updatePortfolio" :text="'更新'" />
    </div>
    <div v-else-if="state.portfolio">
      <Preview :inputPortfolio="state.portfolio" />
      <Button v-if="state.isEditable" @click="editPortfolio" :text="'編集'" />
      <Button
        v-if="state.isEditable"
        @click="deletePortfolio"
        :type="'danger'"
        :text="'削除'"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { useRoute, useRouter } from "vue-router";
import Button from "@/atoms/Button.vue";
import Loading from "@/organisms/Loading.vue";
import Preview from "./Detail/Preview.vue";
import Edit from "./Detail/Edit.vue";
import { InputPortfolio, Portfolio } from "../../models/portfolio";
import PortfolioService from "../../services/portolioService.vue";
import FirebaseService from "../../services/firebaseService";

interface Content {
  isEdit: boolean;
  isLoading: boolean;
  isEditable: boolean;
  portfolio: InputPortfolio;
}

const initPortfolio: InputPortfolio = {
  title: "",
  author: "",
  abstruct: "",
  source: "",
  link: "",
  readme: ""
};

export default defineComponent({
  components: {
    Button,
    Loading,
    Edit,
    Preview
  },
  setup(): {
    state: Content;
    editPortfolio?: () => void;
    updatePortfolio?: () => void;
    deletePortfolio?: () => void;
  } {
    const route = useRoute();
    const router = useRouter();
    const state = reactive<Content>({
      isEdit: false,
      isLoading: true,
      isEditable: false,
      portfolio: initPortfolio
    });

    //methods
    const editPortfolio = () => {
      state.isEdit = true;
    };

    const updatePortfolio = () => {
      if ("new" === route.params.id) {
        PortfolioService.postPortfolio(state.portfolio).then(r => {
          state.portfolio = r.data;
          state.isEdit = false;
          state.isLoading = false;
          state.isEditable = true;
          router.replace("/detail/" + r.data.id);
        });
      } else {
        PortfolioService.putPortfolio(
          route.params.id as string,
          state.portfolio as Portfolio
        ).then(r => {
          state.portfolio = r.data;
          state.isEdit = false;
          state.isLoading = false;
        });
      }
      state.isLoading = true;
    };

    const deletePortfolio = () => {
      PortfolioService.deletePortfolio(route.params.id as string).then(() => {
        state.isEdit = false;
        state.isLoading = false;
        router.replace({ name: "PortolioList" });
      });
      state.isLoading = true;
    };

    if ("new" === route.params.id) {
      state.portfolio = JSON.parse(JSON.stringify(initPortfolio));
      state.isEdit = true;
      state.isLoading = false;
      return { state, editPortfolio, updatePortfolio, deletePortfolio };
    }
    const id = Number(route.params.id);
    if (isNaN(id)) {
      router.push({ name: "NotFoundError" });
      return { state };
    }
    PortfolioService.getPortfolio(id)
      .then(r => {
        state.portfolio = r;
        FirebaseService.getInstance()
          .getUID()
          .then(uid => {
            if (uid) {
              state.isEditable =
                !state.portfolio.uid || uid === state.portfolio.uid;
            }
          });
        state.isLoading = false;
      })
      .catch(e => {
        if (e.response?.status == 404) {
          router.push({ name: "NotFoundError" });
        } else {
          router.push({ name: "LoadError" });
        }
      });

    return { state, editPortfolio, updatePortfolio, deletePortfolio };
  },
  beforeRouteUpdate(to, from, next) {
    next();
  }
});
</script>
