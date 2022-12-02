import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import HomePage from './pages/home-page'
import Header from './components/header'

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Switch>
          <Route path="/">
            <HomePage />
          </Route>
        </Switch>
      </div>
    </Router>
  )
}

export default App
