<template>
  <div class="transfer-form">
    <h2>Transfer</h2>
    <form @submit.prevent="submitTransfer">
      <div class="form-group">
        <label for="sender">Pengirim</label>
        <input type="text" id="sender" v-model="transfer.sender" required />
      </div>
      <div class="form-group">
        <label for="receiver">Penerima</label>
        <input type="text" id="receiver" v-model="transfer.receiver" required />
      </div>
      <div class="form-group">
        <label for="amount">Jumlah</label>
        <input type="number" id="amount" v-model="transfer.amount" required />
      </div>
      <button type="submit">Kirim</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      transfer: {
        sender: '',
        receiver: '',
        amount: ''
      }
    };
  },
  methods: {
    async submitTransfer() {
      try {
        const response = await axios.post('http://localhost:3000/api/transfers', this.transfer);
        alert('Transfer berhasil: ' + response.data.message);
        this.transfer.sender = '';
        this.transfer.receiver = '';
        this.transfer.amount = '';
      } catch (error) {
        console.error('Terjadi kesalahan:', error);
        alert('Transfer gagal');
      }
    }
  }
};
</script>

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
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}
</style>
