<template>
    <div class="game-cover" :style="{ backgroundImage: `url(${image})` }"></div>
</template>

<script>
export default {
    props: {
        image: String,
    },
}
</script>

<style>
.game-cover {
    height: 180px;
    width: 100%;
    background-size: cover;
    background-position: center;
    border-radius: 20px;
}
</style>