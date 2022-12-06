import React from 'react'
import { useSelector } from 'react-redux'
import GameBtn from "../../components/game-btn";
import GameCover from '../../components/game-cover';
import "./game-page.css"

const GamePage = () => {

    const game = useSelector((state) => state.game.currentGame)

    console.log("qwe", game);


    return (
        <div className="game-page">
            <h1 className="game-page__title">{game.title}</h1>
            <div className="game-page__content">
                <div className="game-page__left">
                    <iframe
                        width="90%"
                        height="400px"
                        src={game.video}
                        title="YouTube video player"
                        frameBorder="0"
                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                    ></iframe>
                </div>
                <div className="game-page__right">
                    <GameCover image={game.image} />
                    <p>{game.description}</p>
                    <p className="secondary-text">Популярные метки для этого продукта:</p>
                    <div className='game-item__genre-container'>
                        {game.genres.map((genre) => (
                            <div className='game-item__g'>{genre} </div>
                        ))}
                    </div>
                    <div className="game-page__buy-game">
                        <GameBtn game={game} />
                    </div>
                </div>
            </div>
        </div>
    )
}


export default GamePage