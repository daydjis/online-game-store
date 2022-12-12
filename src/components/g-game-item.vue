<template>
    <div class="game-item shadow p-3 mb-5 bg-body rounded">
        <h1>{{ game_data.title }}</h1>
        <g-game-cover v-bind:image="game_data.image" />
        <div class="game-genres">
            <g-game-genre v-for="genre in game_data.genres" :key="genre" :game_genre="genre" />
        </div>
        <span class="game-price">{{ game_data.price }} руб.</span>
        <button type="button" class="btn btn-outline-success" @click="addToCart">Купить</button>
    </div>
</template>

<script>
import gGameGenre from './g-game-genre.vue';
import gGameCover from './g-game-cover.vue';
export default {
    components: { gGameGenre, gGameCover },
    data() {
        return
    },
    props: {
        game_data: {
            type: Object,
            defautl() {
                return {}
            }
        }
    },
    methods: {
        addToCart() {
            this.$emit("addToCart", this.game_data)
        }
    }
}
</script>

<style scoped>
.game-genres {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.game-item {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;

    height: 400px;
    width: 320px;
    margin: 15px;
    padding: 20px;
    border-radius: 20px;
}
</style>