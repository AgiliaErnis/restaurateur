import React, {useContext, useState, useEffect } from 'react'
import { Link, Redirect } from 'react-router-dom';
import { UserContext } from '../../UserContext';
import "./UserNavbar.css"

function UserNavbar() {
    const [click, setClick] = useState(false)
    const { setClickedUserMenuItem, username, logout,
        setLogout, setSuccessfullLogin } = useContext(UserContext)
    const logOut = () => {
        const logoutRequest = {
            method: 'GET',
            credentials: 'include',
            headers: {
              'Content-Type': 'application/json'
            }
        }
        fetch(`${process.env.REACT_APP_PROXY}/auth/logout`,logoutRequest)
            .then(response => response.json())
            .then(res => {
                if (res.status === 200) {
                    setLogout(true)
                    setSuccessfullLogin(false)
                }
                else {
                    setLogout(false)
                }

            }
        )
    }

    useEffect(() => {
    if(logout && !localStorage.getItem('user-logged-in'))
         <Redirect to="/" />
    }, [logout])

    return (
        <>
            <div className="user-menu-container"
                 onClick={() => setClick(!click)}>
                <div className="user-container">
                    <i class="fas fa-user"></i>
                    <span className="username">{username}</span>
                    <i class="fas fa-chevron-down"></i>
                </div>
                {click &&
                    <div className="user-menu">
                        <Link to ='/user' style={{ textDecoration: 'none' }}>
                            <div className="menu-item"
                                onClick={() => setClickedUserMenuItem("saved")}>
                                <i class="far fa-bookmark"></i>
                                <p>Saved Restaurants</p>
                            </div>
                            <div className="menu-item"
                                 onClick={() => setClickedUserMenuItem("password")}>
                                <i class="fas fa-cog"></i>
                                <p>Account Settings</p>
                            </div>
                        </Link>
                            <div className="menu-item logout" onClick={logOut}>
                                <i class="fas fa-sign-out-alt"></i>
                                <p>Log Out</p>
                            </div>

                    </div>
                }
            </div>
        </>
    )
}

export default UserNavbar
