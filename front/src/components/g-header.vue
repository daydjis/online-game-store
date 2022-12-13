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
        </div>
    </nav>
</template>

<script>


export default {

}
</script>

<style>
.navbar {
    margin-bottom: 20px;
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