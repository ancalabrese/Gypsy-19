import React from 'react'
import { InputGroup, FormControl } from 'react-bootstrap';
import Style from '../../index.scss'

const ControlPanel = (props) => {
    return (
        <div className="primary-dark-element">
            <InputGroup className="mb-3">
                <InputGroup.Checkbox aria-label="Checkbox for following text input" />
                {/* <FormControl aria-label="Text input with checkbox" /> */}
            </InputGroup>
            <InputGroup>
                <InputGroup.Radio aria-label="Radio button for following text input" />
                <FormControl aria-label="Text input with radio button" />
            </InputGroup>
            Control Panel placeholder
        </div>
    )
}

export default ControlPanel;