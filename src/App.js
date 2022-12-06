import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import HomePage from './pages/home-page'
import GamePage from './pages/game-page'
import Header from './components/header'
import { Provider } from 'react-redux'
import { store } from './redux'

function App() {
  return (
    <Provider store={store}>
      <Router>
        <div className="App">
          <Header />
          <Switch>
            <Route exact path="/" component={HomePage} />
            <Route exact path="/app/:title" component={GamePage} />
          </Switch>
        </div>
      </Router>
    </Provider>
  )
}

export default App
