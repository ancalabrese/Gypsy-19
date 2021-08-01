import './App.css';
import './Components/WorldMap/WorldMap'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Toolbar from './Components/UI/Toolbar/Toolbar'
import WorldMap from './Components/WorldMap/WorldMap';
import SettingPanel from './Components/ControlsPanel/SettingPanel';

function App() {
  return (
    <div className="App">
      <Toolbar />
      <Container fluid className={["align-items-center", "min-vh-100"].join(" ")}>
        <Row className={["align-self-start"].join(" ")}>
          <Col xs={{ span: 12, order: 2, }} md={{ span: 3, order: 1, }} className={["align-items-start"].join(" ")}>
            <SettingPanel />
          </Col>
          <Col xs={{ span: 12, order: 1, }} md={{ span: 9, order: 2 }} className={["align-items-center"].join(" ")}>
            <WorldMap />
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
