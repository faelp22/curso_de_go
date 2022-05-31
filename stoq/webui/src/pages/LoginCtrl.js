import config from "../config";
import Http from "../modules/Http";

export default {
    data() {
        return {
            username: null,
            password: null,
            isPwd: true,
        };
    },

    created() {
        this.limpar();
    },

    methods: {

        async login() {
            let headers = {}
            let user_credentials = {"username": this.username, "password": this.password}
            let url = config.apiPath + "api/v1/user/login";
            Http.postData(url, user_credentials, headers).then((response) => {
                localStorage.setItem("autenticated", JSON.stringify(response.data));
                this.$router.push({ name: 'ListProducts' });
            }).catch((error) => {
                console.log("Login", error);
                this.alertError();
            });
        },

        limpar() {
            this.username = "";
            this.password = "";
        },

        alertError() {
            $('#erroralert').fadeIn();
            setTimeout(() => {
                $('#erroralert').fadeOut();
            }, 3000);
        }
    }
};
