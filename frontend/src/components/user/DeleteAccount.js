import React, { useContext } from 'react'
import { UserContext } from '../../UserContext';
import { Button } from '../button/Button'
import validateDeleteAccount from './validateDeleteAccount';
import useDeleteAccountForm from './useDeleteAccountForm'

function DeleteAccount({ submitForm }) {
    const {incorrectPasswordOnDelete} = useContext(UserContext)
        const { handleChange, handleSubmit, values, errors } = useDeleteAccountForm (
        submitForm,
        validateDeleteAccount
    );

    return (
        <form
            onSubmit={handleSubmit}
            className='form' noValidate>
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
                <div className="form-inputs">
                <label className="old-password">
                    Password
                </label>
                <span className="input-description" style={{fontSize: "13px"}}>
                    Please provide your exsiting password in order to delete the account.
                </span>
                <input
                    className='form-input'
                    type='password'
                    name='password'
                        placeholder='Enter your password'
                        onChange={handleChange}
                        />
                        {incorrectPasswordOnDelete &&
                            !errors.password &&
                            values.password.length !== 0
                            ?
                            <p>Incorrect Password</p>
                            :
                            errors.password && <p>{errors.password}</p>}
            </div>
                <Button
                    buttonStyle="btn--form"
                    onClick={handleSubmit}
                >
                    Delete Account
                </Button>
            </div>
        </form>
    )
}

export default DeleteAccount
