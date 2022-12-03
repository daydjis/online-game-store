import React from 'react'
import GameBtn from '../game-btn'
import GameCover from '../game-cover'
import "./game-item.css"

export const GameItem = ({ game }) => {
    return (
        <div className='game-item'>
            <GameCover image={game.image} />
            <div className='game-item__details'>
                <span className='game-item__title'>{game.title}</span>
                <div className='game-item__genre'>
                    {game.genres.map(genre =>
                        <div key={genre} className='game-item__genre-box'>{genre}</div>
                    )}
                </div>
                <div className='game-item__btn'> <GameBtn game={game} /></div>
            </div>
        </div>
    )
}
