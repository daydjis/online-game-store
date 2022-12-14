<template>
    <router-view>
        <div class="home-page_wrapper-cart-container">
            <router-link to="/cart">
                <button class="home-page_wrapper-cart-icon">В корзине: {{ CART.length }}</button>
            </router-link>
        </div>
        <div class="home-page_wrapper shadow p-3 mb-5 bg-body rounded">
            <g-game-catalog />
        </div>
    </router-view>
</template>

<script>
import { mapGetters } from 'vuex'

import GGameCatalog from './g-game-catalog.vue'


export default {
    components: { GGameCatalog },
    computed: {
        ...mapGetters([
            "CART"
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
    min-height: 1000px;
    margin: 0 auto;
}

.home-page_wrapper-cart-container {
    position: absolute;
    bottom: 800px;
    right: 180px;
}

.home-page_wrapper-cart-icon {
    z-index: 2;
    margin: 0 auto;
    border-radius: 5px;
    padding: 10px;
    background-color: white;
}
</style>