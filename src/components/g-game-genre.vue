<template>
    <span class="game-genre">{{ game_genre }}</span>
</template>

<script>
export default {
    props: {
        game_genre: {
            type: Array,
            defautl() {
                return {}
            }
        }
    }
}
</script>

<style>
.game-genre {
    background-color: rgb(97, 97, 97);
    color: white;
    padding: 5px;
    margin: 2px;
    border-radius: 5px;
}
</style>