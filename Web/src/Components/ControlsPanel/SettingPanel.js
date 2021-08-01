import React, { useState } from 'react'
import { InputGroup, FormControl, Form, Button, Card } from 'react-bootstrap';

const SettingPanel = (props) => {

    const [isVaxed, updateVaxStatus] = useState(false);

    const onVaxStatusChangeHandler = (e) => {
        console.log(e.target.checked)
        updateVaxStatus(e.target.checked)
    }

    return (
        <Card className={["default-card", "primary-dark-element"].join(" ")}>
            <Card.Body>
                <Form className={["primary-dark-element", "default-rounded-borders"].join(" ")}>
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
            </Card.Body>
        </Card>
    )
}

export default SettingPanel;