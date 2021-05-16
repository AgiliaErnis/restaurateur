import React, { useState, useContext} from 'react'
import { Button } from '../button/Button';
import PasswordStrengthMeter from '../forms/signup/PasswordStrengthMeter'
import { UserContext } from '../../UserContext';

function ChangePassword() {
    const [input, setInput] = useState("")
    const goodPassword = useContext(UserContext)

    const handleChange = e => {
        var userInput = e.target.value;
        setInput(userInput);
    }

    return (
        <form className='form' noValidate>
            <div className="change-password">
                <label className="old-password">
                    Current Password
                </label>
                <input
                    className='form-input'
                    type='password'
                    name='password'
                    placeholder='Enter your exsiting password'
                />
                <label className="old-password">
                    New Password
                </label>
                <input
                    className={goodPassword.goodPassword ?
                        input.length > 64 ?
                            "form-input"
                            :
                            "form-input strong-password"
                        :
                        "form-input"}
                    type='password'
                    name='password'
                    placeholder='Enter your new password'
                    onChange={handleChange}
                />
                {input !== "" &&
                    <PasswordStrengthMeter password={input} />}
                <label className="old-password">
                    Verify New Password
                </label>
                <input
                    className='form-input'
                    type='password'
                    name='password2'
                    placeholder='Reenter your password to verify'
                />
                <Button
                    buttonStyle="btn--form"
                    buttonSize="btn--large"
                >
                    Save New Password
                </Button>
            </div>
        </form>
    )
}

export default ChangePassword