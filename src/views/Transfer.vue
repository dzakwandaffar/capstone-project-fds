<script setup>
import { reactive, inject } from 'vue'
import { VNumberInput } from 'vuetify/labs/VNumberInput'
const data = reactive({
  AccountID: '',
  BankID: '',
  Amount: 0,
  TransactionDate: '',
  snackbar: false,
  pesanTransfer: ''
})

const myAxios = inject('myAxios')

const submit = () => {
  console.log('submit clicked', data)

  myAxios
    .post('/transaction/transfer-bank', {
      AccountID: data.AccountID,
      BankID: data.BankID,
      Amount: data.Amount
    })
    .then(
      (res) => {
        if (res.status == 200) {
          data.pesanTransfer = 'Transaksi Berhasil'
        }
        data.snackbar = true
      },
      (err) => {
        data.pesanTransfer = 'Transaksi Gagal'
        data.snackbar = true
      }
    )
}
</script>

<template>
  <v-card variant="tonal" class="pa-3">
    <div class="container">
      <div>
        <label>Account ID</label>
        <v-text-field type="text" v-model="data.AccountID" />
      </div>
      <div>
        <label>Bank ID</label>
        <v-text-field type="text" v-model="data.BankID" />
      </div>
      <div>
        <label>Amount</label>
        <v-number-input v-model="data.Amount" />
      </div>
      <v-btn @click="submit"> Button </v-btn>
    </div>
    <v-snackbar v-model="data.snackbar">
      {{ data.pesanTransfer }}
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
  background-color: plum;
  color: black;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}
</style>
