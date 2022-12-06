import React, { useEffect, useState } from 'react'
import './home-page.css'
import { GameItem } from '../../components/game-item/game-item'
import axios from 'axios'



const HomePage = () => {
  const [newGames, setGames] = useState([])
  useEffect(() => {
    const getGames = async () => {
      try {
        const response = await axios.get('http://localhost:5000/api/games');
        setGames(response.data)
        console.log(response);
      } catch (error) {
        console.error(error);
      }
    }
    getGames()
  }, [])
  console.log("Наш объект", newGames);
  return <div className="home-page">{newGames.map(game => <GameItem game={game} key={game.id} />)
  }
  </div >
}

export default HomePage
