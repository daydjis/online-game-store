<template>
    <div class="v-cart-item_conainer shadow p-3 mb-5 bg-body rounded">
        <div class="v-cart-item_image">
            <g-game-cover :image="cart_data.image" />
        </div>
        <div class="v-cart-item__quantity">
            <p>Кол-во</p>
            {{ cart_data.quantity }}
        </div>
        <div class="v-cart-item__title">{{ cart_data.title }}</div>
        <div class="v-cart-item__price">{{ cart_data.price }} руб</div>
        <button type="button" class="btn btn-outline-info" @click="deleteItemFromCart">Удалить</button>
    </div>
</template>

<script>

import gGameCover from './g-game-cover.vue'
export default {
    components: { gGameCover },
    props: {
        cart_data: {
            type: Object,
            default() {
                return
            }
        }
    },
    mounted() {
        const cart_data = this.cart_data;
        cart_data.quantity = 1;
    },
    methods: {
        deleteItemFromCart() {
            this.$emit('deleteItemFromCart')
        },
    }
}
</script>

<style>
.v-cart-item_conainer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    border-radius: 25px;
    padding: 20px;
}

.v-cart-item_image {
    width: 400px;
    height: 200px;
}

.v-cart-item__price {
    font-size: large;
}

.v-cart-item__title {
    font-size: large;
}
</style>