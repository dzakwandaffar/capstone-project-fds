<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/Header.vue'
import { reactive } from 'vue'
import { useAuthStore } from './stores/auth';
import Transfer from './views/Transfer.vue';

const auth = useAuthStore()

const data = reactive({
  variable1: "test<3",
  counter: 1,
})

const tambahCounter = () => {
  data.counter++
}

</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <Header :text="'HIIIIII'">
        <button @click="tambahCounter">Tambah Counter</button>
        coba coba {{  data.counter }}
        <p>Test 123</p>
        <template #part1_header="{ param1 }">
          <h2>{{ param1 }}</h2>
        </template>
      </Header>
      {{  data.variable1 }}
      <HelloWorld msg="You did it!" />
      
      <nav>
        <RouterLink v-if="auth.isLoggedIn" to="/">Home</RouterLink>
        <RouterLink v-if="auth.isLoggedIn" to="/about">About</RouterLink>
        <RouterLink v-if="!auth.isLoggedIn" to="/login">Login</RouterLink>
        <RouterLink v-if="auth.isLoggedIn" to="/transaction">Transfer</RouterLink>
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
  background-color: plum;
  color: white;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
  border: 5px solid ;
  border-radius: 50px;
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
