import { useState } from 'react'

const SelectLogic = () => {
    const searchOptions = [
    {value: 'name', label:'Search by Name'},
    {value: 'location', label:'Search by Location'}
    ];

    const sortOptions = [
        { value: 'rating', label: 'Rating - ascending' },
        { value: 'price-ascending', label: 'Price - ascending' },
        { value: 'price-descending', label: 'Price - descending' }
    ];

    const [searchResult, setSearchResult] = useState(searchOptions.value);
    const [sortResult, setSortResult] = useState(sortOptions.value);

    const setSearchResultHandler = e => {
        setSearchResult(e.value);
    }

    const setSortResultHandler = e => {
        setSortResult(e.value);
    }

    return {
        sortResult,searchResult, searchOptions, sortOptions,
        setSearchResultHandler, setSortResultHandler
    }
}

export default SelectLogic;