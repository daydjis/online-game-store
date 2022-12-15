<template>
    <router-view>
        <div class="g-cart__container">
            <div class="g-cart">
                <h1>
                    Корзина {{ CART.length }} шт.
                </h1>
                <g-cart-item v-for="(item, index) in CART" :key="item._id" :cart_data="item"
                    @deleteItemFromCart="deleteGameFromCart(index)" />
            </div>

            <button class="g-cart__pay-btn btn-success">
                <h1>Всего к оплате: {{ totalPrice }} руб.</h1>
            </button>
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
    data() {
        const totalPrice = 0
        return totalPrice
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
        totalPrice() {
            let result = []

            if (this.CART.length) {
                for (let item of this.CART) {
                    result.push(item.price * item.quantity)
                }
                result = result.reduce(function (sum, el) {
                    return sum + el
                })
                return result
            } else {
                return "0"
            }


        },
        ...mapGetters([
            "CART"
        ])
    }
}
</script>

<style>
.g-cart__container {
    max-width: 1200px;
    min-height: 1000px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-content: center;
}

.badge {
    color: black;
}

.g-cart__pay-btn {
    position: sticky;
    position: -webkit-sticky;
    bottom: 1px;
    height: 100px;
    background-color: rgb(40, 174, 100);
    color: white;
    border: none;
    border-radius: 5px;
    transition: 1s;
}

.g-cart__pay-btn:hover {
    position: sticky;
    height: 100px;
    background-color: rgb(255, 255, 255);
    color: rgb(40, 174, 100);
    border: solid rgb(40, 174, 100) 1px;
    border-radius: 5px;
    transition: 1s;
}
</style>