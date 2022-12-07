import React, { useState } from 'react'
import "./new-game-page.css"
import axios from 'axios'


export const NewGamePage = () => {

    const [title, setTitle] = useState("")
    const [description, setDescription] = useState("")
    const [price, setPrice] = useState("")
    const [genres, setGenres] = useState("")
    const [video, setVideo] = useState("")
    const [image, setImage] = useState("")
    const headers = {
        "Accept": "application/json",
        "Access-Control-Allow-Origin": "*",
        "X-Requested-With": "XMLHttpRequest",
        "Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,OPTIONS",
        "Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"
    }


    const handleCreateGame = () => {
        axios.get('http://localhost:5000/api/games/new', headers, {
            title: title,
            genres: [genres],
            price: Number(price),
            video: video,
            image: image,
            description: description
        })
            .then((res) => console.log('Response', res))

    }


    return (
        <div className='new-game-page__container'>
            <input type="text" placeholder='title' className='new-game-page__form' onChange={(e) => setTitle(e.target.value)} />
            <input type="text" placeholder='description' className='new-game-page__form' onChange={(e) => setDescription(e.target.value)} />
            <input type="text" placeholder='price' className='new-game-page__form' onChange={(e) => setPrice(e.target.value)} />
            <input type="text" placeholder='Genres ' className='new-game-page__form' onChange={(e) => setGenres(e.target.value)} />
            <input type="text" placeholder='Video  ' className='new-game-page__form' onChange={(e) => setVideo(e.target.value)} />
            <input type="text" placeholder='image ' className='new-game-page__form' onChange={(e) => setImage(e.target.value)} />
            <button className='btn btn--primory ' onClick={handleCreateGame}>создать</button>
        </div>
    )
}
