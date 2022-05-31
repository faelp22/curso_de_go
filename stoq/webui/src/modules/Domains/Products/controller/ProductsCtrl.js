import {
  mapState,
  mapActions
} from 'vuex';


export default {

  computed: {

    ...mapState('Products', [
      'list_products',
    ]),
  },

  created() {
    this.setListProductsAction();
  },

  methods: {

    ...mapActions('Products', [
      'setListProductsAction',
    ]),

    deleteItem(item) {
      this.list_products = this.list_products.filter((e) => {
        return e.hash !== item.hash;
      });
      this.$forceUpdate();
    },

    addItem() {
      this.list_products.push({
        "hash": "fsdfsdf",
        "path": "http://localhost:5002",
        "name": "fsdfsdf",
        "count": 0,
        "send": false,
        "found": false,
        "jwt_check": false,
        "db_auto": true
      })
      console.log(this.list_products)
      this.$forceUpdate();
    },
  }
}
