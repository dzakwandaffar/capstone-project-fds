<template>
  <v-card>
    <main>
      <HomeView class="user-info">
        <img alt="user logo" class="logo" src="@/assets/user-icon.jpg" width="120" height="120" />
        <h2>Selamat Datang, {{ user.name }} di Bank Celengan</h2>
        <h2>Informasi Saldo Anda Saat ini: {{ user.saldo }}</h2>
      </HomeView>
    </main>
  </v-card>
</template>

<script>
import HomeView from '../views/HomeView.vue'
export default {
  props: {
    user: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
.user-info {
  text-align: center;
  margin-top: 20px;
}
</style>
