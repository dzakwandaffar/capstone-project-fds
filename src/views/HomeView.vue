<script>
import { reactive } from 'vue'
import UserInfo from '../views/User.vue'
import myAxios from '@/plugin/axios'

const data = reactive({
  name: '',
  saldo: '',
  items: []
})

const info = () => {
  myAxios
    .get('/transaction/get', {
      name: data.name,
      saldo: data.saldo
    })
    .then(
      (res) => {
        if (res.status == 200) {
          data.items = res.data.data
          // console.log(res.data)
        }
        data.snackbar = true
      },
      (err) => {
        data.pesanMutation = 'Mutasi Gagal'
        data.snackbar = true
      }
    )
}

export default {
  components: {
    UserInfo
  },
  data() {
    return {
      user: {
        name: 'Rio Dewanto',
        saldo: 1000000 // Saldo contoh
      }
    }
  },
  methods: {
    logout() {
      // Logika logout
      this.$router.push('/login')
    }
  }
}
</script>

<template>
  <div class="home">
    <UserInfo :user="user" />
    <button @click="logout">Logout</button>
  </div>
</template>

<style scoped>
.home {
  text-align: center;
  margin-top: 50px;
}

button {
  margin-top: 20px;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  background-color: blue;
  border-radius: 4px;
  color: white;
}
</style>
