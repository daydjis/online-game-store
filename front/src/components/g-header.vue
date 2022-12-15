<template>
    <nav class="navbar navbar-dark bg-dark">
        <div class="container-md">
            <router-link to="/">
                <a class="navbar-brand">Game Store</a>
            </router-link>
            <router-link to="/">
                <a class="nav-link active" aria-current="page" href="#">Главная</a>
            </router-link>
            <router-link to="/create/game">
                <a class="nav-link active" aria-current="page" href="#">Добавить</a>
            </router-link>
            <div>
                <router-link to="/cart">
                    <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor"
                        class="bi bi-cart" viewBox="0 0 16 16">
                        <path
                            d="M0 1.5A.5.5 0 0 1 .5 1H2a.5.5 0 0 1 .485.379L2.89 3H14.5a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 13 12H4a.5.5 0 0 1-.491-.408L2.01 3.607 1.61 2H.5a.5.5 0 0 1-.5-.5zM3.102 4l1.313 7h8.17l1.313-7H3.102zM5 12a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm7 0a2 2 0 1 0 0 4 2 2 0 0 0 0-4zm-7 1a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm7 0a1 1 0 1 1 0 2 1 1 0 0 1 0-2z" />
                    </svg>
                    <div class="cart-items">{{ CART.length }}</div>
                </router-link>
            </div>
            <router-link to="/auth/login">
                <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor"
                    class="bi bi-person-circle" viewBox="0 0 16 16">
                    <path d="M11 6a3 3 0 1 1-6 0 3 3 0 0 1 6 0z" />
                    <path fill-rule="evenodd"
                        d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8zm8-7a7 7 0 0 0-5.468 11.37C3.242 11.226 4.805 10 8 10s4.757 1.225 5.468 2.37A7 7 0 0 0 8 1z" />
                </svg>
            </router-link>
        </div>
    </nav>
</template>

<script>
import { mapGetters } from 'vuex'


export default {
    computed: {
        ...mapGetters([
            "CART"
        ])
    }
}
</script>

<style>
.cart-items {
    background-color: brown;
    border-radius: 100px;
    height: 18px;
    width: 18px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 15px;
    color: white;
    position: absolute;
    top: 3px;
}

a {
    text-decoration: none;
}

a:hover {
    color: black;
    transition: 150ms;
}

.nav-link {
    color: white;
}

.container-fluid {
    width: 64%;
}
</style>