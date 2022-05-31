import config from '../config'
import Http from '../modules/Http'

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
      const url = config.apiPath + 'api/v1/user/login'
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
