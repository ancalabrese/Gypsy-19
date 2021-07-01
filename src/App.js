import logo from './logo.svg';
import './App.css';
import './Components/WorldMap/WorldMap'
import WorldMap from './Components/WorldMap/WorldMap';
import ControlPanel from './Components/ControlsPanel/ControlPanel';

function App() {
  return (
    <div className="App">
        <WorldMap />
        <ControlPanel />
        {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}
    </div>
  );
}

export default App;
