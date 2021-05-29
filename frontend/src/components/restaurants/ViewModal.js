import React from 'react'
import "./ViewModal.css"

function ViewModal(props) {
    return (
        <div className="view-modal-container">
            <div className="view-modal">
                <div className="restaurant-name-cont">
                  <p>{props.name}</p>
                </div>
                <div className="view-container">
                    <p className="view">
                        <table>
                            <tr>
                                <td>Cuisines: </td>
                                <td>{props.cuisines}</td>
                            </tr>
                            <tr>
                                <td>Vegan: </td>
                                <td>{props.vegan}</td>
                            </tr>
                            <tr>
                                <td>Vegetarian: </td>
                                <td>{props.vegetarian}</td>
                            </tr>
                            <tr>
                                <td>Website: </td>
                                <td>{props.website}</td>
                            </tr>
                        </table>
                    </p>
                </div>

            </div>
        </div>
    )
}

export default ViewModal