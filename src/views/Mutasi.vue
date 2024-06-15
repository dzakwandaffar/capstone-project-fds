<script setup>
import { reactive, inject, onMounted } from 'vue'
import { VNumberInput } from 'vuetify/labs/VNumberInput'
const data = reactive({
  MutationDateFrom: '',
  MutationDateUntil: '',
  snackbar: false,
  pesanMutation: '',
  items: []
})

const myAxios = inject('myAxios')

const submit = () => {
  console.log('submit clicked', data)

  myAxios
    .get('/transaction/get', {
      MutationDateFrom: data.MutationDateFrom,
      MutationDateUntil: data.MutationDateUntil
    })
    .then(
      (res) => {
        if (res.status == 200) {
          data.pesanMutation = 'Mutasi berhasil'
          // data.MutationDateFrom = res.data.data
          // data.MutationDateUntil = res.data.data
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

  onMounted(() => {
    submit()
  })
}
</script>

<template>
  <v-card class="pa-3">
    <div class="container">
      <div>
        <label>Mutation Date From</label>
        <v-text-field type="date" v-model="data.MutationDateFrom" />
      </div>
      <div>
        <label>Mutation Date Until</label>
        <v-text-field type="date" v-model="data.MutationDateUntil" />
      </div>
      <v-btn @click="submit"> Check Mutation </v-btn>
    </div>
    <v-data-table :items="data.items"></v-data-table>
    <v-snackbar v-model="data.snackbar">
      {{ data.pesanMutation }}
      <template v-slot:actions>
        <v-btn color="plum" variant="text" @click="data.snackbar = false"> Close </v-btn>
      </template>
    </v-snackbar>
  </v-card>
</template>

<style scoped>
.transfer-form {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}

button {
  padding: 10px 15px;
  background-color: rgb(7, 16, 144);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}
</style>
