import React, { useCallback, useState } from 'react'
import { useSelector } from 'react-redux'
import "./cart-block.css"
import { AiOutlineShoppingCart } from "react-icons/ai"
import CartMenu from '../cart-menu'
import { calcTotalPrice } from '../utils'
import ItemsInCart from '../items-in-Cart'
import { useHistory } from 'react-router-dom'


export const CartBlock = () => {
    const history = useHistory()
    const [isCartMenuVisible, setCartMenuVisible] = useState(false)
    const items = useSelector(state => state.cart.itemsInCart)
    const totalPrice = calcTotalPrice(items)

    const handleClick = useCallback(() => {
        setCartMenuVisible(false)
        history.push("/order")
    }, [history])

    return (
        <div>
            <ItemsInCart quantity={items.length} />
            <AiOutlineShoppingCart size={30} className="cart-block__icon" onClick={() => setCartMenuVisible(!isCartMenuVisible)} />
            {totalPrice > 0 ? <span className='cart-block__total-price'>{totalPrice}</span> : null}
            {isCartMenuVisible && <CartMenu items={items} onClick={handleClick} />}
        </div>
    )
}
