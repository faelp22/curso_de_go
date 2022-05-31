import {
  mapState,
  mapActions
} from 'vuex'

export default {

  computed: {
    ...mapState('Products', [
      'list_products'
    ])
  },

  created () {
    this.setListProductsAction()
  },

  methods: {

    ...mapActions('Products', [
      'iniciarProductAction',
      'setProductAction',
      'setListProductsAction',
      'deleteProductAction'
    ]),

    async deleteProduct (product) {
      await this.setProductAction({ product })

      if (confirm('Deseja realmente deletar este produto?') === true) {
        await this.deleteProductAction()
        await this.setListProductsAction()
      } else {
        this.iniciarProductAction()
      }
    }

  }
}
