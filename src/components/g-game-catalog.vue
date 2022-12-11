<template>
    <div class="m">
        <h1>Каталог</h1>
        <div class="game-catalog">
            <g-game-item v-for="game in GAMES" :key="game._id" :game_data="game" />
        </div>
    </div>
</template>

<script>
import gGameItem from './g-game-item.vue'
import { mapActions, mapGetters } from 'vuex'
export default {
    components: { gGameItem },
    data() {
        return
    },
    computed: {
        ...mapGetters([
            "GAMES"
        ])
    },
    methods: {
        ...mapActions([
            "GET_GAMES_FROM_API"
        ]),
    },
    mounted() {
        this.GET_GAMES_FROM_API()
    }

}
</script>

<style>
.game-catalog {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
</style>