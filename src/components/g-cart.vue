<template>
    <div class="g-cart">
        <h1>
            Корзина {{ itemsInCart }} шт.
        </h1>
        <g-cart-item v-for="(item, index) in cartItemsInfo" :key="item._id" :cart_data="item"
            @deleteItemFromCart="deleteGameFromCart(index)" />
    </div>
</template>

<script>
import GCartItem from './g-cart-item.vue'
import { mapActions } from 'vuex';
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
        cartItemsInfo: {
            type: Object,
            default() {
                return {}
            }
        }
    },
    methods: {
        ...mapActions([
            'DELETE_FROM_CART'
        ]),
        deleteGameFromCart(index) {
            this.DELETE_FROM_CART(index)
        }
    }
}
</script>

<style>
.badge {
    color: black;
}

.g-cart {
    width: 100%;
}
</style>