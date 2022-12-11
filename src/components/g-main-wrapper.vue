<template>
    <div>
        <g-header />
    </div>
    <div class="home-page_wrapper shadow p-3 mb-5 bg-body rounded">
        <g-game-catalog />
    </div>
</template>

<script>
import GGameCatalog from './g-game-catalog.vue'
import gHeader from './g-header.vue'
export default {
    components: { gHeader, GGameCatalog },
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
</style>