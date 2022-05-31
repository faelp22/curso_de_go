import {
  mapState,
  mapActions
} from 'vuex'

export default {

  data () {
    return {
      local_product: {
        id: null,
        name: '',
        code: '',
        price: 0
      }
    }
  },

  computed: {
    ...mapState('Products', [
      'product'
    ])
  },

  created () {
    if (this.$route.params.id) {
      (async () => {
        console.log(this.$route.params.id)
        await this.setProductAction({ id: this.$route.params.id })
        this.local_product = JSON.parse(JSON.stringify(this.product))
      })()
    } else {
      this.iniciarProductAction()
      this.local_product = JSON.parse(JSON.stringify(this.product))
    }
  },

  methods: {

    ...mapActions('Products', [
      'iniciarProductAction',
      'setProductAction',
      'setProductsInListAction',
      'deleteProductAction'
    ]),

    async deleteProduct () {
      if (confirm('Deseja realmente deletar este produto?') === true) {
        await this.deleteProductAction()
        this.$router.push({ name: 'ListProducts' })
      }
    },

    async saveProduct () {
      const tmpProduct = JSON.parse(JSON.stringify(this.local_product))
      tmpProduct.price = parseFloat(tmpProduct.price)
      await this.setProductsInListAction(tmpProduct)
      this.$router.push({ name: 'ListProducts' })
    },

    resetProduct () {
      this.iniciarProductAction()
      this.local_product = JSON.parse(JSON.stringify(this.product))
    }

  }
}
