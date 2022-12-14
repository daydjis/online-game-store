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
            <form class="d-flex" role="search">
                <input class="form-control me-2" type="search" placeholder="Поиск" aria-label="Поиск">
                <button class="btn btn-outline-success" type="submit">Поиск</button>
            </form>
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


export default {

}
</script>

<style>
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