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
        <q-input v-model="password" outlined type="password" label="Senha">
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
    import LoginCtrl from "./LoginCtrl";
    export default LoginCtrl;
</script>
