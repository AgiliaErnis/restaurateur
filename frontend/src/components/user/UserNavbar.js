import React, {useContext, useState } from 'react'
import { Link } from 'react-router-dom';
import { UserContext } from '../../UserContext';
import "./UserNavbar.css"

function UserNavbar() {
    const [click, setClick] = useState(false)
    const { setClickedUserMenuItem } = useContext(UserContext)
    return (
        <>
            <div className="user-menu-container"
                 onClick={() => setClick(!click)}>
                <div className="user-container">
                    <i class="fas fa-user"></i>
                    <span className="username">Username</span>
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
                            <div className="menu-item logout">
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
