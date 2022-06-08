<template>
  <q-page class="flex flex-center" style="background-image: linear-gradient(to right, #ff6d0f, #8222ff);">
    <q-card style="width: 300px; height: auto; display: flex; flex-direction: column; align-items: center;">

      <q-img src="../assets/logo.png" ratio="1" width="150px" style="margin-top: 15px;"/>

      <q-card-section class="q-pt-none">
        <div class="text-h5">Login</div>
      </q-card-section>

      <q-card-section class="q-pt-none full-width">
        <q-input outlined v-model="username" label="Usuário" />
      </q-card-section>

      <q-card-section class="q-pt-none full-width">
        <q-input v-model="password" outlined :type="isPwd ? 'password' : 'text'" label="Senha">
        <template v-slot:append>
          <q-icon
            :name="isPwd ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="isPwd = !isPwd"
          />
        </template>
      </q-input>
      </q-card-section>

      <q-card-section>
        <q-btn @click="login" color="primary" icon="login" label="Entrar" />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import config from '../config'
import Http from '../services/Http'

export default {
  data () {
    return {
      username: null,
      password: null,
      isPwd: true
    }
  },

  created () {
    this.limpar()
  },

  methods: {

    async login () {
      const headers = {}
      const userCredentials = { username: this.username, password: this.password }
      const url = config.apiPath + 'api/v1/usuario/token'
      Http.postData(url, userCredentials, headers).then((response) => {
        localStorage.setItem('autenticated', JSON.stringify(response.data))
        this.$router.push({ name: 'ListProducts' })
      }).catch((error) => {
        console.log('Login', error)
      })
    },

    limpar () {
      this.username = ''
      this.password = ''
    }
  }
}

</script>
