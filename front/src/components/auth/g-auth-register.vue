<template>
    <div class="auth-container ">
        <div class="auth-border">
            <form>
                <h2>Регистрация</h2>
                <div class="form-group">
                    <input v-model="userNewLogin.login" type="login" class="form-control" id="exampleInputlLogin1"
                        aria-describedby="loginHelp" placeholder="Введите ваш логин">
                    <small id="loginHelp" class="form-text text-muted">Никому не сообщайте свои данные</small>
                </div>
                <div class="form-group">
                    <input v-model="userNewLogin.password" type="password" class="form-control"
                        id="exampleInputPassword1" placeholder="Пароль">
                </div>


                <router-link to="/auth/login">
                    <a class="nav-link active black" aria-current="page" href="#">уже есть аккаунт?</a>
                </router-link>

                <button v-if="!LOADER" class="btn btn-primary" @click="REGISTER_USER(userNewLogin)">Отправить</button>
                <div class="d-flex justify-content-center" v-else>
                    <div class="spinner-border" role="status">
                        <span class="sr-only"></span>
                    </div>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
    methods: {
        ...mapActions([
            "REGISTER_USER"
        ])
    },
    computed: {
        ...mapGetters([
            "LOADER"
        ])
    }
    ,
    data() {
        return {
            userNewLogin: {
                login: '',
                password: ''
            }
        }
    }
}
</script>

<style scoped>
.black {
    color: black
}

.form-control {
    width: 400px;
    margin-bottom: 20px;
}

.form-group {
    margin-bottom: 20px;
}

.auth-border {
    margin-top: 40px;
    box-shadow: 1px 2px 4px black;
    border-radius: 5px;
    padding: 20px;
}

.auth-container {
    display: flex;
    justify-content: center;
    align-items: center;
    max-width: 1200px;
    margin: 0 auto;
}
</style>