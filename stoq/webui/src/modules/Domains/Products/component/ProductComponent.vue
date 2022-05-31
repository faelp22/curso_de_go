<template>
    <div class="row q-pa-md items-start q-gutter-md justify-center full-height full-width text-center">

    <q-card class="my-card center" >
      <q-form class="q-gutter-md">
      <div class="row no-wrap items-center q-mt-md q-pa-sm bg-grey-3 rounded-borders">
        <div v-if="local_product.id" class="text-h4">Editar Produto</div>
        <div v-else class="text-h4">Cadastrar Produto</div>
        <q-space />
      </div>
      <q-input
        filled
        v-model="local_product.name"
        label="Nome do produto"
      />

      <q-input
        filled
        v-model="local_product.code"
        label="Código"
      />

      <q-input
        filled
        type="number"
        v-model="local_product.price"
        label="Preço"
      />
      <q-separator />
      <q-card-actions>
        <q-btn flat icon="reply" color="info" :to="{name: 'ListProducts'}" > Voltar</q-btn>
        <q-btn v-if="!local_product.id" flat icon="cleaning_services" type="submit" color="warning" @click="resetProduct()" > Limpar</q-btn>
        <q-btn v-if="local_product.id" flat icon="delete_forever" type="reset" color="negative" @click="deleteProduct(item)" > Remover</q-btn>
        <q-btn flat icon="send" type="submit" color="positive" @click="saveProduct()" > Salvar</q-btn>
      </q-card-actions>
    </q-form>
    </q-card>

  </div>
</template>

<script>
import ProductCtrl from '../controller/ProductCtrl'
export default ProductCtrl
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 750px
</style>
