import React from 'react'
import { InputGroup, FormControl } from 'react-bootstrap';

const ControlPanel = (props) => {
    return (
        <div style={{
            backgroundColor: 'red',
        }}>
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