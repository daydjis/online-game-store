<template>
    <div>
        <g-header />
    </div>
    <div class="home-page_wrapper shadow p-3 mb-5 bg-body rounded">
        <g-game-catalog />
    </div>
    <div class="home-page_wrapper shadow p-3 mb-5 bg-body rounded">
        <g-cart v-if="CART.length" :itemsInCart="CART.length" :cartItemsInfo="CART" />
    </div>

</template>

<script>

import GGameCatalog from './g-game-catalog.vue'
import gHeader from './g-header.vue'
import gCart from './g-cart.vue'
import { mapGetters } from 'vuex';

export default {
    components: { gHeader, GGameCatalog, gCart },
    computed: {
        ...mapGetters([
            'CART'
        ])
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
</style>