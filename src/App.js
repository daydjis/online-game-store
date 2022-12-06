import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import HomePage from './pages/home-page'
import GamePage from './pages/game-page'
import Header from './components/header'
import NewGamePage from './pages/new-game-page'

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Switch>
          <Route exact path="/games/new">
            <NewGamePage />
          </Route>
          <Route exact path="/order">
            <GamePage />
          </Route>
          <Route exact path="/app/:title">
            <GamePage />
          </Route>
          <Route exact path="/">
            <HomePage />
          </Route>
        </Switch>
      </div>
    </Router>
  )
}

export default App
