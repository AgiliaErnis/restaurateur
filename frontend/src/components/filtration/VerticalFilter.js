import React, { useState} from 'react'
import './VerticalFilter.css'

export const VerticalFilter = (props) => {
    const [seeMoreCuisines, setSeeMoreCuisines] = useState(false);
    const [seeMoreLocalities, setSeeMoreLocalities] = useState(false);
    const [checkedPriceRange, setCheckedPriceRange] = useState([]);
    const [checkedFeatured, setCheckedFeatured] = useState([]);

    const priceRanges = [
        {
            filterValue: "0-300",
            value: "0 - 300 Kč"
        },
        {
           filterValue: "300-600",
            value: "300 - 600 Kč"
        },
        {
           filterValue: "600-",
           value: "Over 600 Kč"
        }]

    const featuredOptions = [
        {
            filterValue: "vegetarian",
            value: "Vegetarian"
        },
        {
            filterValue: "vegan",
            value: "Vegan"
        },
        {
            filterValue: "gluten-free",
            value: "Gluten Free"
        }
    ]

    function handleClick () {
        setSeeMoreCuisines(!seeMoreCuisines)
    }

    function handleClickLocalities () {
        setSeeMoreLocalities(!seeMoreLocalities)
    }

    const handleTogglePriceRange = (value) => {
        const currentIndex = checkedPriceRange.indexOf(value)
        const newChecked = [...checkedPriceRange]

        if (currentIndex === -1) {
            newChecked.push(value);
        } else {
            newChecked.splice(currentIndex, 1)
        }

        setCheckedPriceRange(newChecked);
        props.handlePriceRangeFilters(newChecked);
    }

    const handleToggleFeatured = (value) => {
        const currentIndex = checkedFeatured.indexOf(value)
        const newcheckedFeatured = [...checkedFeatured]

        if (currentIndex === -1) {
            newcheckedFeatured.push(value);
        } else {
            newcheckedFeatured.splice(currentIndex, 1)
        }

        setCheckedFeatured(newcheckedFeatured);
        props.handleFeaturedFilters(newcheckedFeatured);
    }

    return (
        <div className="vertical-filter-container">
            <div className="filter-content">
                <p>Filters</p>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Cuisine</p>
                        <div className="filter-options">
                            <label>
                                <input
                                    className='option-input checkbox'
                                    type='checkbox'
                                />
                                <span className="option-name">American</span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Asian</span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Italian</span>
                            </label>
                            <div className={seeMoreCuisines ?
                                "cuisines_shown"
                                :
                                "cuisines_hidden"}>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Indian</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Japanese</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Vietnamese</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Spanish</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        Mediterranean
                                </span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">French</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Thai</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Mexican</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        International
                                </span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Czech</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">English</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Balkan</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Brazil</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Russian</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Chinese</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Greek</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Arabic</span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">Korean</span>
                                </label>
                            </div>
                        </div>
                        <p className="see-more" onClick={handleClick}>
                            {seeMoreCuisines ? "See less" : "See more"}
                        </p>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Price Range</p>
                        <div className="filter-options">
                            {priceRanges.map((priceRange) => (
                                <label onChange={() => handleTogglePriceRange(priceRange.filterValue)}>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        {priceRange.value}
                                    </span>
                                </label>))}
                        </div>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Services</p>
                        <div className="filter-options">
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Delivery</span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Takeaway</span>
                            </label>
                        </div>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Localities</p>
                        <div className="filter-options">
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">
                                    Prague 1
                                </span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Prague 2</span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Prague 3</span>
                            </label>
                            <div className={seeMoreLocalities ?
                                "shown"
                                :
                                "hidden"}>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        Prague 4
                                    </span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        Prague 5
                                    </span>
                                </label>
                                <label>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        Prague 6
                                    </span>
                                </label>
                            </div>
                        </div>
                        <p className="see-more" onClick={handleClickLocalities}>
                            {seeMoreLocalities ? "See less" : "See more"}
                        </p>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Suggested</p>
                        <div className="filter-options">
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">
                                    Open Now (22:00)
                                </span>
                            </label>
                            <label>
                                <input className='option-input checkbox'
                                    type='checkbox' />
                                <span className="option-name">Near me</span>
                            </label>
                        </div>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Featured</p>
                        <div className="filter-options">
                            {featuredOptions.map((featuredOption) => (
                                <label onChange={() => handleToggleFeatured(featuredOption.filterValue)}>
                                    <input className='option-input checkbox'
                                        type='checkbox' />
                                    <span className="option-name">
                                        {featuredOption.value}
                                    </span>
                                </label>
                            ))}
                        </div>
                    </div>
                </div>
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Distance</p>
                        <div className="filter-options">
                            <label>
                                <input className='radio' type='radio'
                                    name='distance-option' />
                                <div className="checkmark"></div>
                                <span className="option-name">
                                    Bird's-eye view
                                </span>
                            </label>
                            <label>
                                <input className='radio' type='radio'
                                    name='distance-option' />
                                <div className="checkmark"></div>
                                <span className="option-name">
                                    500 meters radius
                                </span>
                            </label>
                            <label>
                                <input className='radio' type='radio'
                                    name='distance-option' />
                                <div className="checkmark"></div>
                                <span className="option-name">1 km radius</span>
                            </label>
                            <label>
                                <input className='radio' type='radio'
                                    name='distance-option' />
                                <div className="checkmark"></div>
                                <span className="option-name">3 km radius</span>
                            </label>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}