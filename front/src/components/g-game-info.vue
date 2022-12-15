<template>
    <div className="game-page">
        <h1 className="game-page__title">{{ CURRENT_GAME.title }}</h1>
        <div className="game-page__content shadow p-3 mb-5 bg-body rounded">
            <div className="game-page__left ">
                <iframe :src="CURRENT_GAME.video" width="90%" height="400px" title="YouTube video player"
                    frameBorder="0"
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"></iframe>
            </div>
            <div className="game-page__right">
                <gGameCover v-bind:image="CURRENT_GAME.image" />
                <div class="shadow p-3 mb-8 bg-body rounded">
                    <p>{{ CURRENT_GAME.description }}</p>
                    <p className="secondary-text">Популярные метки для этого продукта:</p>
                    <div className='game-item__genre-container' v-for="genre in CURRENT_GAME.genres" :key="genre">
                        {{ genre }}
                    </div>
                    <div className="game-page__buy-game">
                        <button type="button" class="btn btn-primary btn-lg" @click="ADD_GAME_TO_CART(CURRENT_GAME)">
                            Купить</button>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<script>
import gGameCover from './g-game-cover.vue';

import { mapGetters, mapActions } from 'vuex'


export default {
    components: { gGameCover },

    data() {

    },
    methods: {
        ...mapActions([
            "ADD_GAME_TO_CART"
        ])
    },

    computed: {
        ...mapGetters([
            "CURRENT_GAME",
            "CART"
        ]),
    }

}
</script>

<style>
.game-page {
    padding-top: 60px;
    max-width: 1200px;
    margin: 0 auto;
}

.game-page__content {
    display: flex;
}

.game-page__left {
    width: 70%;
}

.game-page__right {
    width: 30%;
}

.game-page__buy-game {
    display: flex;
    justify-content: center;
    margin-top: 20px;
}

.game-item__g {
    background-color: rgb(17, 14, 15);
    border-radius: 5px;
    padding: 5px;
    margin: 2px;
    justify-content: center;
    display: flex;
}
</style>