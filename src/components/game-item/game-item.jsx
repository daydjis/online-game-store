import React from 'react'
import { useDispatch } from 'react-redux'
import GameBtn from '../game-btn'
import GameCover from '../game-cover'
import "./game-item.css"
import { setCurrentGame } from "../../redux/games/reducer"
import { useHistory } from 'react-router'

export const GameItem = ({ game }) => {
    const history = useHistory()
    const dispatch = useDispatch()

    const handleClick = () => {
        dispatch(setCurrentGame(game))
        history.push(`/app/${game.title}`)
    }

    return (
        <div className='game-item' onClick={handleClick}>
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
