import React from 'react'
import "./cart-block.css"
import { AiOutlineShoppingCart } from "react-icons/ai"

export const CartBlock = () => {
    return (
        <div>
            <AiOutlineShoppingCart size={30} className="cart-block__icon" />
            <span className='cart-block__total-price'>2100 руб</span>
        </div>
    )
}
