import React from 'react'
import "./PhoneModal.css"

function PhoneModal(props) {
    return (
        <div className="phone-modal-container">
            <div className="phone-modal">
                <div className="restaurant-name-cont">
                  <p>{props.name}</p>
                </div>
                <div className="phone-container">
                    {props.phone === null ?
                        "Phone Number is Not Available"
                        :
                        `Phone Number: ${props.phone}`}
                </div>
            </div>
        </div>
    )
}

export default PhoneModal