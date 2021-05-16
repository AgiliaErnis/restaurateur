import React, { useContext, useState } from 'react'
import "./UserProfile.css"
import { UserContext } from '../../UserContext';
import { UserMenuItemsData }   from './UserMenuItemsData'
import ChangePassword from './ChangePassword'
import DeleteAccount from './DeleteAccount';
import SavedRestaurants from './SavedRestaurants';
import UpdateUsername from './UpdateUsername';


function UserProfile() {
    const { clickedUserMenuItem, setClickedUserMenuItem } = useContext(UserContext)

    const [isSubmitted, setIsSubmitted] = useState(false);

    function submitForm() {
        setIsSubmitted(!isSubmitted);
    }

    return (
        <div className="account-container">
        <div className="user-profile">
            <div className="profile-cover-container">
                <img src='/images/Home/hero.jpg' alt='hero-background'className="user-background"/>
                <div className="user-info">
                    <div className="user">
                        <i class="far fa-user-circle"></i>
                        <p>Username</p>
                    </div>
                    <div className="saved-restaurants">
                        <span className="number">4</span>
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
                        {clickedUserMenuItem === "password" ? "Change Password" : clickedUserMenuItem === "delete" ? "Delete Account" :
                            clickedUserMenuItem === "username"
                                ?
                                "Change Username"
                                :
                                "Saved Restaurants"}
                    </h4>
                    {clickedUserMenuItem === "password" ?
                        <ChangePassword submitForm={submitForm}/>
                        :
                        clickedUserMenuItem === "delete" ?
                            <DeleteAccount />
                            :
                            clickedUserMenuItem === "username" ?
                                <UpdateUsername />
                                :
                                <SavedRestaurants />
                    }
                </div>
            </div>
        </div>
    )
}

export default UserProfile
