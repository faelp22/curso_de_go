<template>
  <div class="q-pa-md row items-start q-gutter-md">

    <div class="row no-wrap items-center q-mt-md q-pa-sm bg-grey-3 rounded-borders" style="width: 100%">
      <div class="text-h4">Lista de Produtos</div>
      <q-space />
      <q-btn flat align="right" icon="add" color="positive" @click="addItem">Adicionar</q-btn>
    </div>

    <q-card v-for="item, idx in list_products" :key="idx" class="my-card" flat bordered>
      <q-card-section class="q-pt-xs">
        <q-btn flat round icon="category" />
        <div class="text-h6 q-mt-sm q-mb-xs">{{item[0].name}}</div>
        Validar JWT [{{ (item[0].jwt_check) ? "Sim" : "Não" }}] <br>
        Total de Requisições [{{item[0].count}}] <br>
      </q-card-section>
      <q-separator />
      <q-card-actions align="right">
        <q-btn flat icon="link" color="primary">Acessar</q-btn>
        <q-btn flat icon="delete_forever" color="negative" @click="deleteItem(item)" >Remover</q-btn>
      </q-card-actions>
    </q-card>

  </div>
</template>

<script>
  import ProductsCtrl from "../controller/ProductsCtrl";
  export default ProductsCtrl;
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 350px

</style>
