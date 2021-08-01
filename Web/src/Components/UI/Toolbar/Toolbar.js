import React, { useState, useEffect } from 'react'
import PropTypes from 'prop-types'
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'
import NavDropdown from 'react-bootstrap/NavDropdown'
import Style from './Toolbar.module.css'

const Toolbar = (props) => {
  const [showShadow, displayShadow] = useState(props.fixedShadow);

  return (
    <Navbar className={Style.Toolbar} expand="md" sticky="top">
      <Navbar.Brand href="#home">GYPSY-19</Navbar.Brand>
      {/* <Navbar.Toggle aria-controls="basic-navbar-nav" /> */}
      {/* <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="mr-auto">
            <Nav.Link href="#home">Home</Nav.Link>
            <Nav.Link href="#link">Link</Nav.Link>
            <NavDropdown title="Dropdown" id="basic-nav-dropdown">
              <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
              <NavDropdown.Item href="#action/3.2">Another action</NavDropdown.Item>
              <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
              <NavDropdown.Divider />
              <NavDropdown.Item href="#action/3.4">Separated link</NavDropdown.Item>
            </NavDropdown>
          </Nav>
        </Navbar.Collapse> */}
    </Navbar>
  );
}

Toolbar.propTypes = {
  fixedShadow: PropTypes.bool,
  navigationItems: PropTypes.arrayOf(Nav.Link)
};

Toolbar.defaultProps = {
  fixedShadow: false,
  navigationItems: []
};

export default Toolbar;