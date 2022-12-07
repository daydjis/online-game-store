import React, { useState } from 'react'
import "./new-game-page.css"



export const NewGamePage = () => {

    const [title, setTitle] = useState("")
    const [description, setDescription] = useState("")
    const [price, setPrice] = useState("")
    const [genres, setGenres] = useState("")
    const [video, setVideo] = useState("")
    const [image, setImage] = useState("")



    const handleCreateGame = async (data) => {
        const responce = await fetch("http://localhost:5000/api/games/new", {
            method: "POST",
            body: JSON.stringify(data = {
                title: title,
                genres: [genres],
                price: Number(price),
                video: video,
                image: image,
                description: description
            })
        })
        if (!responce.ok) {
            throw new Error(`Ошибка по адресу ${"http://localhost:5000/api/games/new"}, статус ошибки ${responce}`)
        }
        return await responce.json()
    }


    return (
        <div className='new-game-page__container'>
            <div className='new-game-page__container-items'>
                <h2 className='new-game-page__title'>Добавить новую игру в базу данных</h2>
                <input type="text" placeholder='Название игры' className='new-game-page__form' onChange={(e) => setTitle(e.target.value)} />
                <input type="text" placeholder='Описание игры' className='new-game-page__form' onChange={(e) => setDescription(e.target.value)} />
                <input type="text" placeholder='Цена в (руб)' className='new-game-page__form' onChange={(e) => setPrice(e.target.value)} />
                <input type="text" placeholder='Жанры ' className='new-game-page__form' onChange={(e) => setGenres(e.target.value)} />
                <input type="text" placeholder='Сыллка на трейлер  ' className='new-game-page__form' onChange={(e) => setVideo(e.target.value)} />
                <input type="text" placeholder='Картинка (base64)' className='new-game-page__form' onChange={(e) => setImage(e.target.value)} />
                <button className='btn btn--primory ' onClick={handleCreateGame}>создать</button>
            </div>
        </div>
    )
}
