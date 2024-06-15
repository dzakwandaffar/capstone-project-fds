<script setup>
import { inject, onMounted, reactive } from 'vue'
import { ref } from 'vue'
import router from '@/router/index'
import { VNumberInput } from 'vuetify/labs/VNumberInput'

const dataBilling = reactive({
  CheckBiller: [],
  AccountID: '',
  BillerID: '',
  Biller: {
    Amount: 0
  }
})

const myAxios = inject('myAxios')

const bil = () => {
  myAxios.get('/transaction/biller').then(
    (res) => {
      dataBilling.CheckBiller = res.data.Data
    },
    (err) => {}
  )
}

const generateTagihan = () => {
  myAxios.get('/transaction/biller/' + dataBilling.BillerID + '/' + dataBilling.AccountID).then(
    (res) => {
      // dataBilling.ListBiller = res.data.data
      if (res.status == 200) {
        dataBilling.Biller = res.data.Biller
      }
    },
    (err) => {}
  )
}

const open = ref(['Users'])

onMounted(() => {
  bil()
})
</script>

<script>
export default {
  data: () => ({
    open: ['Users']
  })
}
</script>

<template>
  <v-card variant="outlined" class="pa-5">
    <div class="container">
      <div>
        <label>ID Account Billing</label>
        <v-text-field type="text" v-model="dataBilling.AccountID" />
      </div>
      <div>
        <label>Tujuan Billing</label>
        <v-select
          label="Select"
          v-model="dataBilling.BillerID"
          :items="dataBilling.CheckBiller"
          :item-title="'Name'"
          :item-value="'BillerID'"
        ></v-select>
      </div>
    </div>

    <v-container>
      <div class="d-flex">
        <div class="mr-2">
          <v-btn @click="bil"> Transfer </v-btn>
        </div>

        <v-snackbar v-model="dataBilling.snackbar">
          {{ dataBilling.pesanBil }}

          <template v-slot:actions>
            <v-btn color="blue" variant="text" @click="dataBilling.snackbar = false"> Close </v-btn>
          </template>
        </v-snackbar>
        <v-dialog max-width="500">
          <template v-slot:activator="{ props: activatorProps }">
            <v-btn v-bind="activatorProps" text="Cek Tagihan"></v-btn>
          </template>

          <template v-slot:default="{ isActive }" @click="generateTagihan">
            <v-card title="Tagihan anda saat ini : ">
              <v-card-text>
                <v-btn @click="generateTagihan"> Cek Tagihan </v-btn>

                {{ dataBilling.Biller.Amount }}
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>

                <v-btn text="Close" @click="isActive.value = false"></v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </v-dialog>
        <!-- <div class="ml-2">
            <v-btn @click="generateTagihan">   
              Cek Tagihan
            </v-btn>
          </div> -->
        <!-- {{ dataBilling.Biller.Amount }} -->
      </div>
    </v-container>
  </v-card>
</template>
