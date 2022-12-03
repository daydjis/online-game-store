import React, { useState } from 'react'
import { useSelector } from 'react-redux'
import "./cart-block.css"
import { AiOutlineShoppingCart } from "react-icons/ai"
import CartMenu from '../cart-menu'
import { calcTotalPrice } from '../utils'

export const CartBlock = () => {
    const [isCartMenuVisible, setCartMenuVisible] = useState(false)
    const items = useSelector(state => state.cart.itemsInCart)
    const totalPrice = calcTotalPrice(items)
    return (
        <div>
            <AiOutlineShoppingCart size={30} className="cart-block__icon" onClick={() => setCartMenuVisible(!isCartMenuVisible)} />
            {totalPrice > 0 ? <span className='cart-block__total-price'>{totalPrice}</span> : null}
            {isCartMenuVisible && <CartMenu items={items} onClick={e => null} />}
        </div>
    )
}
