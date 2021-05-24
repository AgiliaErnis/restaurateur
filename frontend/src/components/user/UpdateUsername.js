import React from 'react'
import { Button } from "../button/Button"
import useUpdateUsernameForm from './useUpdateUsernameForm'
import validateUsername from './validateUsername';

function UpdateUsername({ submitForm }) {
    const { handleChange, handleSubmit, errors } = useUpdateUsernameForm(
        submitForm,
        validateUsername
    );
    return (
        <form
            onSubmit={handleSubmit}
            className='form' noValidate>
            <div className="update-username">
                <div className="form-inputs">
                    <label className='old-password'>New Username</label>
                    <input
                        className={'form-input'}
                        type='text'
                        name='newUsername'
                        placeholder='Enter your new username'
                        onChange={handleChange}
                    />
                    {errors.newUsername && <p>{errors.newUsername}</p>}
                    <Button
                        buttonStyle="btn--form"
                        type='submit'
                        onClick={handleSubmit}
                    >
                        Update Username
                    </Button>
                </div>
            </div>
        </form>
    )
}

export default UpdateUsername
