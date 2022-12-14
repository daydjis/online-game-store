<template>
    <div class="auth-container">
        <div>
            <h1>Новая игра</h1>
            <form class="add-new-game__menu">
                <div class="form-group">
                    <input v-model="newGameForm.title" type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Ведите название игры">
                </div>
                <div class="form-group">
                    <input v-model="newGameForm.imageDescription" type="text" class="form-control"
                        id="formGroupExampleInput" placeholder="Описание картинки">
                </div>
                <div class="form-group">
                    <input v-model="newGameForm.video" type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Ссылка на трейлер">
                </div>
                <div class="form-group">
                    <input v-model="newGameForm.genres" type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Жанры">
                </div>
                <div class="form-group">
                    <input v-model="newGameForm.description" type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Описание">
                </div>
                <div class="form-group">
                    <input v-model="newGameForm.image" type="text" class="form-control" id="formGroupExampleInput"
                        placeholder="Изображение">
                </div>
                <div class="form-group">
                    <input v-model.number="newGameForm.price" type="number" class="form-control"
                        id="formGroupExampleInput" placeholder="Цена">

                </div>
            </form>
            <button type="button" class="btn btn-outline-primary" @click="POST_NEW_GAME(newGameForm)">Добавить</button>
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
        return {
            newGameForm: {
                title: '',
                description: '',
                price: "",
                genres: [""],
                video: '',
                imageDescription: '',
                image: '',
            },
        }
    },
    mounted() {

    }
}
</script>

<style>
.btn-outline-primary {
    margin-top: 20px;
    max-width: 100px;
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