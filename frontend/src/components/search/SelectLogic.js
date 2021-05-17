import { useState } from 'react'

const SelectLogic = () => {
    const [searchResult, setSearchResult] = useState(false);
    const [sortResult, setSortResult] = useState(false);

    const searchOptions = [
        { value: 'name', label: 'Search by Name' },
        { value: 'location', label: 'Search by Location' }
    ];

    const sortOptions = (sortResult !== false ?
        [
            { value: 'rating', label: 'Rating - descending' },
            { value: 'price-asc', label: 'Price - ascending' },
            { value: 'price-desc', label: 'Price - descending' },
            { value:  false, label: 'Without Sorting' }
        ]
        :
        [
            { value: 'rating', label: 'Rating - descending' },
            { value: 'price-asc', label: 'Price - ascending' },
            { value: 'price-desc', label: 'Price - descending' }
        ]);



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