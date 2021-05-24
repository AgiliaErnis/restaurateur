import React from 'react'
import "./MenuModal.css"

function MenuModal(props) {
    const menu = (props.menu !== null && JSON.parse(props.menu))
    const dates = [...Object.keys(menu)]
    const content = [...Object.values(menu)]

    return (
        <div className="menu-modal-container">
            <div className={props.menu !== null ?
                "menu-modal"
                :
                "not-available-menu-modal"}>
                <div className="restaurant-name-cont">
                  <p>{props.name}</p>
                </div>
                <div className="menu-container">
                    {props.menu === null ? "Weekly Menu is Not Available" :
                        (dates.length !== 0 && dates.map(date =>
                            <div style={{ display: "flex", flexDirection: "row" }}>
                                <p className="date">{date}:</p>
                                <p style={{ textAlign: "start" }}>
                                    {content[dates.indexOf(date)]}
                                </p>
                            </div>)
                        )
                    }
                </div>
            </div>
        </div>
    )
}

export default MenuModal