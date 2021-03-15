<template>
  <div class="header">
    <h1>
      <router-link id="title" to="/">Prototype Frontier</router-link>
    </h1>
    <nav v-if="!state.isLoading" class="nav">
      <ul>
        <li v-if="state.user">
          <span class="menu">
            {{ state.user.displayName }}
            <i class="arrow down"></i>
          </span>
          <ul>
            <li>
              <span class="submenu" @click="signout">サインアウト</span>
            </li>
          </ul>
        </li>
        <li v-else @click="signin">
          <Button
            :id="'signin'"
            :text="'サインイン'"
            :type="'secondary'"
            :size="'small'"
          />
        </li>
      </ul>
    </nav>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive } from "vue";
import firebase from "firebase/app";
import firebaseService from "../services/firebaseService";
import Button from "@/atoms/Button.vue";

export default defineComponent({
  components: {
    Button,
  },
  setup() {
    const fs = firebaseService.getInstance();
    const state = reactive<{ user: firebase.User | null; isLoading: boolean }>({
      user: null,
      isLoading: true,
    });
    //methods
    const signin = () => {
      fs.login();
    };

    const signout = () => {
      state.isLoading = true;
      fs.logout().then(() => {
        window.location.reload();
      });
    };
    fs.getUser()
      .then(user => {
        state.user = user;
        state.isLoading = false;
      })
      .catch(() => {
        state.isLoading = false;
      });
    return { state, signin, signout };
  },
});
</script>

<style>
.header {
  background-color: #d7c447;
  height: 70px;
  width: 100%;
  position: absolute;
  top: 0;
  position: fixed;
  display: flex;
  align-items: center;
}

#title {
  padding: 10px;
  margin: 0;
  font-weight: bold;
  font-style: oblique;
  text-decoration: none;
  color: black;
}

.nav {
  margin: 1rem;
  margin-left: auto;
  padding: 0;
  display: flex;
  cursor: pointer;
  font-weight: bold;
}

.menu {
  font-size: medium;
}
.signin {
  padding: 0;
}
.submenu {
  font-size: x-small;
}
.arrow {
  border: solid black;
  border-width: 0 3px 3px 0;
  display: inline-block;
  padding: 3px;
}
.down {
  transform: rotate(45deg);
  -webkit-transform: rotate(45deg);
}

.nav li {
  position: relative;
  list-style: none;
}
.nav li span {
  display: block;
  text-align: center;
  padding: 0.2rem 0.5rem;
  text-decoration: none;
  box-sizing: border-box;
  background-color: white;
  border-radius: 0.375rem;
}
.nav li ul {
  top: 1.7rem;
  left: -40px;
  position: absolute;
}
.nav .submenu:hover {
  background-color: yellow;
}

.nav li ul li {
  overflow: hidden;
  height: 0;
  font-size: 11px;
}
.nav li:hover > ul > li {
  overflow: visible;
  height: 1.5rem;
  width: 6rem;
}
</style>
