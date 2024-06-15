<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/Header.vue'
import { reactive } from 'vue'
import { useAuthStore } from './stores/auth'
import Transfer from './views/Transfer.vue'

const auth = useAuthStore()

const data = reactive({
  variable: 'test'
})
</script>

<template>
  <header>
    <img alt="Bank logo" class="logo" src="@/assets/Bankceleng.jpeg" width="155" height="200" />

    <div class="wrapper">
      <Header :text="'Selamat Datang di'">
        <template #part1_header="{}"> </template>
      </Header>
      <HelloWorld msg="Bank Celengan" />

      <nav>
        <RouterLink v-if="auth.isLoggedIn" to="/home">Home</RouterLink>
        <RouterLink v-if="!auth.isLoggedIn" to="/login">Login</RouterLink>
        <!-- <RouterLink v-if="auth.isLoggedIn" to="/transaction">Transfer</RouterLink> -->
        <RouterLink v-if="auth.isLoggedIn" to="/mutation">Mutation</RouterLink>
        <RouterLink v-if="auth.isLoggedIn" to="/billing">Pembayaran Billing</RouterLink>
      </nav>
    </div>
  </header>

  <div>
    <RouterView />
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
  background-color: white;
  color: black;
  text-align: center;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: white;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
