import React from 'react'
import Button from '../button'
import { calcTotalPrice } from '../utils'
import "./cart-menu.css"
import { RiDeleteBin7Fill } from "react-icons/ri"
import { useDispatch } from 'react-redux'
import { deleteItemFromCart } from '../../redux/cart/reducer'


export const CartMenu = ({ items, onClick }) => {

    const dispatch = useDispatch()

    const handleDelete = () => {
        dispatch(deleteItemFromCart())
        console.log("work");
    }

    return (items.length > 0 ?
        (<div className='cart-menu'>
            <div className='cart-menu__games-list'>
                {items.length > 0 ? items.map((game) => <div key={game.title} className="cart-menu__games-list-onCart">
                    <div>{game.title}</div>
                    <div className='cart-menu__games-list-onCart-price'>{game.price} руб </div>
                    <RiDeleteBin7Fill size={20} cursor="pointer" onClick={handleDelete} />
                </div>
                ) : "Пусто"}

            </div>
            {items.length > 0 ? (
                <div className='cart-menu__arrange'>
                    <div className='cart-menu__total-price'>
                        <span>Итого:</span>
                        <span>{calcTotalPrice(items)} руб.</span>
                    </div>
                    <Button type="primory" size='s' onClick={onClick}>Оформить заказ</Button>
                </div>) : null
            }
        </div>) : null
    )
}
