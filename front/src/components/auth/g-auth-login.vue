<template>
    <div class="auth-container ">
        <div class="auth-border">
            <form>
                <h2>Логин</h2>
                <div class="form-group">
                    <input v-model="userLogin.login" type="login" class="form-control" id="exampleInputlLogin1"
                        aria-describedby="loginHelp" placeholder="Введите ваш логин">
                    <small id="loginHelp" class="form-text text-muted">Никому не сообщайте свои данные</small>
                </div>
                <div class="form-group">
                    <input v-model="userLogin.password" type="password" class="form-control" id="exampleInputPassword1"
                        placeholder="Пароль">
                </div>


                <router-link to="/auth/register">
                    <a class="nav-link active black" aria-current="page" href="#">Ещё не зарегистрировались?</a>
                </router-link>

                <button class="btn btn-primary" @click="LOGIN_USER(userLogin)">Submit</button>
            </form>
        </div>
    </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
    methods: {
        ...mapActions([
            "LOGIN_USER"
        ])
    }
    ,
    data() {
        return {
            userLogin: {
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