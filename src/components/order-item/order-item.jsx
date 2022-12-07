import React from 'react'
import { useDispatch } from 'react-redux'
import GameCover from '../game-cover'
import "./order-item.css"
import { RiDeleteBin7Fill } from "react-icons/ri"
import { deleteItemFromCart } from '../../redux/cart/reducer'

export const OrderItem = ({ game }) => {
    const dispatch = useDispatch()

    const handleDelete = () => {
        dispatch(deleteItemFromCart(game.id))
    }

    return (
        <div className='order-item'>
            <div className='order-item__cover'>
                <GameCover image={game.image} />
            </div>
            <div className='order-item__title'>
                <span>{game.title}</span>
            </div>
            <div className='order-item__price'>
                <span>{game.price} руб.</span>
                <RiDeleteBin7Fill size={30} className="cart-item__delete-icon" onClick={handleDelete} />
            </div>
        </div>
    )
}
