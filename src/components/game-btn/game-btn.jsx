import React from 'react'
import "./game-btn.css"

export const GameBtn = ({ game }) => {
    return (
        <div className='game-btn'>
            <span className='game-btn__price'>{game.price} руб.</span>
        </div>
    )
}
