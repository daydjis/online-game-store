<template>
    <div class="game-page">
        <h1 class="game-page__title">{{ GAME_ID.title }}</h1>
        <div class="game-page__content shadow p-3 mb-5 bg-body rounded">
            <div class="game-page__left ">
                <iframe :src="GAME_ID.video" width="90%" height="400px" title="YouTube video player" frameBorder="0"
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"></iframe>
            </div>
            <div class="game-page__right">
                <gGameCover v-bind:image="GAME_ID.image" />
                <div class="game-info__content shadow p-3 mb-8 bg-body rounded">
                    <div>
                        <p>{{ GAME_ID.description }}</p>
                    </div>

                    <div>
                        <p class="secondary-text">Популярные метки для этого продукта:</p>
                    </div>
                    <div>
                        <div class='game-item__g' v-for="genre in GAME_ID.genres" :key="genre">
                            {{ genre }}
                        </div>
                    </div>
                    <div class="game-page__buy-game">
                        <button type="button" class="btn btn-primary btn-lg" @click="ADD_GAME_TO_CART(GAME_ID)">
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
            "ADD_GAME_TO_CART",
            "GET_CURRENT_GAME"
        ])
    },

    computed: {
        ...mapGetters([
            "CURRENT_GAME",
            "CART",
            "GAME_ID"
        ]),
    },
    mounted() {
        console.log(this.$route.params.title)
        this.GET_CURRENT_GAME(this.$route.params.title)
    }

}
</script>

<style>
.game-info__content {
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
    align-content: flex-start;
}



p {
    display: flex;
    justify-content: flex-start;
    text-align: left;
}

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
    color: white;
}
</style>