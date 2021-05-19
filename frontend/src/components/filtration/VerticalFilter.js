import React, { useState, useContext } from 'react'
import { FiltersData, distanceOptions } from './FiltersData'
import { UserContext } from '../../UserContext';
import './VerticalFilter.css'

export const VerticalFilter = (props) => {
    const { clickedDistrict, setClickedDistrict,
            clickedSuggestion, setClickedSuggestion,
            checkedDistance, setCheckedDistance,
            chosenRestaurant} = useContext(UserContext)

    const [seeMoreCuisines, setSeeMoreCuisines] = useState(false);
    const [seeMoreLocalities, setSeeMoreLocalities] = useState(false);

    const [checkedFilters] = useState([
        {
            category: "cuisine",
            checkedOptions: []
        },
        {
            category: "district",
            checkedOptions: []
        },
        {
            category: "price-range",
            checkedOptions: []
        },
        {
            category: "other",
            checkedOptions: []
        }
    ]);

    function handleClickCuisines () {
        setSeeMoreCuisines(!seeMoreCuisines)
    }

    function handleClickLocalities () {
        setSeeMoreLocalities(!seeMoreLocalities)
    }

    const handleCheckboxToggle = (value, category) => {
        if (clickedDistrict !== false) {
            checkedFilters[1].checkedOptions.push(clickedDistrict);
            setClickedDistrict(false);
        }

        if (clickedSuggestion !== false) {
            if (clickedSuggestion === "vegetarian"
                ||
                clickedSuggestion === "gluten-free")
            {
                checkedFilters[3].checkedOptions.push(clickedSuggestion)
                setClickedSuggestion(false);
            } else {
                checkedFilters[0].checkedOptions.push(clickedSuggestion)
                setClickedSuggestion(false);
            }
        }

        checkedFilters.map(filter => {

            if (filter.category === category) {
                const currentIndex = filter.checkedOptions.indexOf(value)
                if (currentIndex === -1) {
                    filter.checkedOptions.push(value);
                } else {
                    filter.checkedOptions.splice(currentIndex, 1)
                }
                props.handlecheckedFilters(checkedFilters);
            }
            return null
         })
    }

    return (
        <div className="vertical-filter-container">
            <div className="filter-content">
                <p>Filters</p>
                {FiltersData.map(filter => (
                    <div className="filter-div">
                        <div className="filter-inner-div">
                            <p>{filter.filterName}</p>
                            <div className="filter-options">
                                {filter.options.map(option => {
                                    return (
                                        filter.options.indexOf(option) >= 3 ?
                                                <div
                                                    onChange={chosenRestaurant === false ? () =>
                                                        handleCheckboxToggle(option.filterValue,
                                                                                filter.category) : null}
                                                    className={FiltersData.indexOf(filter) >= 3 ?
                                                        (seeMoreLocalities ? "shown" : "hidden")
                                                        :
                                                        (seeMoreCuisines ? "cuisines_shown" : "cuisines_hidden")}>
                                                    <label>
                                                        <input
                                                            className='option-input checkbox'
                                                            type='checkbox'
                                                            checked={clickedDistrict === option.filterValue ?
                                                                    true
                                                                    :
                                                                    clickedSuggestion === option.filterValue ?
                                                                        true
                                                                        :
                                                                        chosenRestaurant
                                                                         !== false ? false : null}
                                                            onClick={clickedDistrict === option.filterValue ?
                                                                () => setClickedDistrict(false)
                                                                :
                                                                clickedSuggestion === option.filterValue ?
                                                                    () => setClickedSuggestion(false)
                                                                    : null}
                                                        />
                                                        <span className="option-name">{option.value}</span>
                                                    </label>
                                                </div>
                                                :
                                                <label
                                                     onChange={chosenRestaurant === false ? () =>
                                                        handleCheckboxToggle(option.filterValue,
                                                                                filter.category) : null}
                                                >
                                                    <input
                                                        className='option-input checkbox'
                                                        type='checkbox'
                                                        checked={clickedDistrict === option.filterValue ?
                                                            true
                                                            :
                                                            clickedSuggestion === option.filterValue ?
                                                                true
                                                                :
                                                                chosenRestaurant !== false ? false : null}
                                                        onClick={clickedDistrict === option.filterValue ?
                                                            () => setClickedDistrict(false)
                                                            :
                                                            clickedSuggestion === option.filterValue ?
                                                                () => setClickedSuggestion(false)
                                                                :
                                                                null}
                                                    />
                                                    <span className="option-name">{option.value}</span>
                                                </label>
                                            )
                                        }
                                    )
                                }
                            </div>
                        </div>
                        {(filter.options.length > 3) ?
                            <p
                                className="see-more"
                                onClick={FiltersData.indexOf(filter) === 3 ?
                                    handleClickLocalities
                                    :
                                    handleClickCuisines}
                            >
                                {FiltersData.indexOf(filter) === 3 ?
                                    (seeMoreLocalities ? "See less" : "See more")
                                    :
                                    (seeMoreCuisines ? "See less" : "See more")
                                }
                            </p>
                            :
                            null
                        }
                </div>
            ))}
                <div className="filter-div">
                    <div className="filter-inner-div">
                        <p>Distance</p>
                        <div className="filter-options">
                            {distanceOptions.map(option => {
                                return (
                                    <label onChange={() => setCheckedDistance(option.filterValue)}>
                                        <input className='radio' type='radio'
                                            name='distance-option'
                                            checked={chosenRestaurant !== false ? false :
                                                checkedDistance === option.filterValue ? true : null}
                                            />
                                        <div className="checkmark"></div>
                                        <span className="option-name">
                                            {option.value}
                                        </span>
                                    </label>
                                )
                            })}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}