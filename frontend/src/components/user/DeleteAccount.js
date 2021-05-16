import React from 'react'
import { Button } from '../button/Button'

function DeleteAccount() {
    return (<>
        <div style={{
            margin: "2rem",
            display: "flex"
        }}>
            <p
                style={{
                    marginRight: "5px",
                    color: "rgb(237, 90, 107)",
                    fontWeight: "bold"
                }}
            >
                Important:
            </p>
            <p
                style={{
                    fontSize: "14px",
                    marginTop: "1px"
                }}
            >
                Please note that this action is irreversible and
                all the data associated with your account will be
                permanently deleted.</p>
        </div>
        <div className="delete-account">

            <label className="old-password">
                Password
            </label>
            <span className="input-description">
                Please provide your exsiting password in order to delete the account.
            </span>
            <input
                className='form-input'
                type='password'
                name='password'
                placeholder='Enter your password'
            />
            <Button
                buttonStyle="btn--form"
            >
                Delete Account
            </Button>
        </div></>
    )
}

export default DeleteAccount
