import React from 'react'
import { Button } from "../button/Button"

function UpdateUsername() {
    return (
        <div className="update-username">
          <label className='old-password'>New Username</label>
          <input
            className={'form-input'}
            type='text'
            name='username'
            placeholder='Enter your new username'
            />
            <Button
                buttonStyle="btn--form"
            >
                Update Username
            </Button>
        </div>
    )
}

export default UpdateUsername
