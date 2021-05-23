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
                    <p className="phone-num">
                        Phone Number: {props.phone}
                    </p>
                </div>

            </div>
        </div>
    )
}

export default PhoneModal