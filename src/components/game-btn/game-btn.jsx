import React from 'react'
import Button from '../button'
import "./game-btn.css"
import { useDispatch, useSelector } from 'react-redux'
import { deleteItemFromCart, setItemInCart } from '../../redux/cart/reducer'

export const GameBtn = ({ game }) => {
    const dispatch = useDispatch()
    const items = useSelector(state => state.cart.itemsInCart)
    const isIteminCart = items.some(item => item.id === game.id)

    const handleClick = (e) => {
        e.stopPropagation()
        if (isIteminCart) {
            dispatch(deleteItemFromCart(game.id))
        } else { dispatch(setItemInCart(game)) }

    }
    return (
        <div className='game-btn'>
            <span className='game-btn__price'>{game.price} руб.</span>
            <Button type={isIteminCart ? 'secondary' : 'primory'} onClick={handleClick}>{isIteminCart ? 'Убрать из корзины' : "В корзину"}</Button>
        </div>
    )
}
