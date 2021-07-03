import './App.css';
import './Components/WorldMap/WorldMap'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Toolbar from './Components/UI/Toolbar/Toolbar'
import WorldMap from './Components/WorldMap/WorldMap';
import ControlPanel from './Components/ControlsPanel/ControlPanel';

function App() {
  return (
    <div className="App">
      <Toolbar />
      <Container fluid className={["align-items-center", "min-vh-100"].join(" ")}>
        <Row className={["align-self-start"].join(" ")}>
          <Col md={3} className={["align-items-start"].join(" ")}>
            <ControlPanel />
          </Col>
          <Col md={9} className={["align-items-center"].join(" ")}>
            <WorldMap />
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
