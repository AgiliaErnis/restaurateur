import React, { useContext, useState } from 'react'
import "./UserProfile.css"
import { UserContext } from '../../UserContext';
import { UserMenuItemsData }   from './UserMenuItemsData'
import ChangePassword from './ChangePassword'
import DeleteAccount from './DeleteAccount';
import UpdateUsername from './UpdateUsername';
import { Redirect } from 'react-router';
import SavedRestaurants from './SavedRestaurants'


function UserProfile() {
    const { clickedUserMenuItem, setClickedUserMenuItem,username,savedRestaurants } = useContext(UserContext)

    const [isSubmitted, setIsSubmitted] = useState(false);
    const [newUsername, setNewUsername] = useState(false)
    const [deleteAccount, setDeleteAccount] = useState(false)

    function submitForm() {
        setIsSubmitted(!isSubmitted);
    }

    function submitNewUsername() {
        setNewUsername(!newUsername)
    }

    function submitDeleteAcount() {
        setDeleteAccount(!deleteAccount)
    }

    return (
        <div className="account-container">
        <div className="user-profile">
            <div className="profile-cover-container">
                <img src='/images/Home/hero.jpg' alt='hero-background'className="user-background"/>
                <div className="user-info">
                    <div className="user">
                        <i class="far fa-user-circle"></i>
                            <p>{username}</p>
                    </div>
                    <div className="saved-restaurants">
                            <span className="number">{savedRestaurants.length}</span>
                         <span>Saved Restaurants</span>
                    </div>
                </div>
            </div>
        </div>
            <div className="user-info-box-container">
                <div className="user-info-box">
                    <div className="info-box-menu">
                        {UserMenuItemsData.map(item =>
                         <div className="info-box-menu-item">
                                <h4 className="menu-item-header">{item.menuItem}</h4>
                                {item.options.map(option =>
                                   <div className="menu-item-options ">
                                        <span className={clickedUserMenuItem === option.keyword ?
                                            "menu-item-option clicked"
                                            :
                                            "menu-item-option"
                                            }
                                            onClick={() => setClickedUserMenuItem(option.keyword)}
                                        >
                                            {option.option}
                                        </span>
                                    </div>
                                )}
                         </div>
                        )}
                    </div>
                </div>
                <div className={`user-info-content ${clickedUserMenuItem === "saved" && "scroll"}`}>
                    <h4 className="menu-item-header content">
                        {clickedUserMenuItem === "password" ?
                            "Change Password"
                            :
                            clickedUserMenuItem === "delete"
                                ?
                                "Delete Account"
                                :
                            clickedUserMenuItem === "username"
                                ?
                                "Change Username"
                                :
                                "Saved Restaurants"}
                    </h4>
                    {clickedUserMenuItem === "password" ?
                        (!isSubmitted ? <ChangePassword submitForm={submitForm} /> :
                        <>
                            <div style={{width: "100%", textAlign: "center", marginTop:"2rem"}}>
                                <h2 style={{ margin: "2rem", marginTop: "10rem", textAlign: "center" }}>Password was successfully changed!</h2>
                            </div></>
                        )
                        :
                        clickedUserMenuItem === "delete" ?
                            (!deleteAccount ?
                                <DeleteAccount submitForm={submitDeleteAcount} />
                                :
                                <>
                                    <DeleteAccount submitForm={submitDeleteAcount} />
                                    <Redirect to='/' />
                                </>)
                            :
                            clickedUserMenuItem === "username" ?
                                (!newUsername ?
                                    <UpdateUsername submitForm={submitNewUsername} /> :
                            <>
                            <div style={{width: "100%", textAlign: "center", marginTop:"2rem"}}>
                                            <h2 style={{
                                                margin: "2rem", marginTop: "10rem",
                                                textAlign: "center"
                                            }}>
                                                Username was successfully updated!
                                            </h2>
                            </div>
                            </>)
                                :
                                <SavedRestaurants />
                    }
                </div>
            </div>
        </div>
    )
}

export default UserProfile
