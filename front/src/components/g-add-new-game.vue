<template>
    <div class="home-page_wrapper">
        <div>
            <h1>Новая игра</h1>
            <form class="add-new-game__menu">
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Ведите название игры">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Описание картинки">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Ссылка на трейлер">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Жанры">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Описание">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Изображение">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" id="formGroupExampleInput" placeholder="Цена">
                </div>
            </form>
            <button type="button" class="btn btn-outline-primary" @click="POST_NEW_GAME">Добавить</button>
        </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
    computed: {
        ...mapGetters([
            "NEW_GAME"
        ]),

    },
    methods: {
        ...mapActions([
            'CREATE_NEW_GAME',
            'POST_NEW_GAME'
        ])
    },

    data() {

    }

}
</script>

<style>
.home-page_wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    max-width: 1200px;
    margin: 0 auto;
}

.btn {
    margin-top: 20px;
    max-width: 100px;
}

.form-group {
    width: 400px;
}

.add-new-game__menu {
    width: 900px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    height: 400px;
}
</style>