<template>
    <nav class="navbar navbar-dark bg-dark">
        <div class="container-md">
            <a class="navbar-brand">Game Store</a>
            <a class="nav-link active" aria-current="page" href="#">Главная</a>
            <a class="nav-link active" aria-current="page" href="#">Добавить</a>
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

.nav-link {
    color: white;
}

.container-fluid {
    width: 64%;
}
</style>