import React, { useState } from 'react'
import { InputGroup, FormControl, Form, Button } from 'react-bootstrap';

const SettingPanel = (props) => {

    const [isVaxed, updateVaxStatus] = useState(false);

    const onVaxStatusChangeHandler = (e) => {
        console.log(e.target.checked)
        updateVaxStatus(e.target.checked)
    }

    return (
        <Form className="primary-dark-element">
            <Form.Check
                type="switch"
                id="vaxxed"
                label="I'm fully vaxxed"
                onChange={onVaxStatusChangeHandler} />
            <Form.Check
                type="radio"
                id="quarantine-0"
                name="quarantine"
                value={0}
                label="I Cannot quarantine"
                disabled={isVaxed}
            />
            <Form.Check
                type="radio"
                id="quarantine-1"
                name="quarantine"
                value={1}
                label="Could quarantine for up to 5 days"
                disabled={isVaxed}
            />
            <Form.Check
                type="radio"
                id="quarantine-2"
                name="quarantine"
                value={2}
                label="Could quarantine for up to 10 days"
                disabled={isVaxed}
            />
            <Button variant="outline-secondary">Apply</Button>{" "}
            <Button variant="outline-primary">Reset</Button>
        </Form>
    )
}

export default SettingPanel;