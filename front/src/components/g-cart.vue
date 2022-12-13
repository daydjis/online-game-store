<template>
    <router-view>
        <div class="g-cart">
            <h1>
                Корзина {{ CART.length }} шт.
            </h1>
            <g-cart-item v-for="(item, index) in CART" :key="item._id" :cart_data="item"
                @deleteItemFromCart="deleteGameFromCart(index)" />
        </div>
    </router-view>
</template>

<script>
import GCartItem from './g-cart-item.vue'
import { mapActions, mapGetters } from 'vuex';
export default {
    components: {
        GCartItem
    },
    props: {
        itemsInCart: {
            type: String,
            default() {
                return null
            }
        },

    },
    methods: {
        ...mapActions([
            'DELETE_FROM_CART'
        ]),
        deleteGameFromCart(index) {
            this.DELETE_FROM_CART(index)
        }
    }, computed: {
        ...mapGetters([
            "CART"
        ])
    }
}
</script>

<style>
.badge {
    color: black;
}

.g-cart {
    width: 100%;
    height: 500px;
}
</style>