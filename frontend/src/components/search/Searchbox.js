import React, { useEffect } from 'react';
import { Button } from '../button/Button';
import './Searchbox.css';
import Select from 'react-select'
import MobileNavbar from '../navbar/MobileNavbar'
import AdjustSearchbox from './AdjustSearchbox';
import SearchboxStyle from './SelectStyle';
import SelectLogic from './SelectLogic';

function Searchbox() {
  const { inputClassName, selectClassName,
          searchSize, searchIconClassname } = AdjustSearchbox();
  const { customStyles, customThemes } = SearchboxStyle();
  const { searchResult, searchOptions,
          setSearchResultHandler } = SelectLogic();
  const { button, showSearch } = MobileNavbar();

  useEffect(() => {
    showSearch();
  }, [showSearch]);

  window.addEventListener('resize', showSearch);

  return (
    <div className="search-box">
      <div className="main-search-container">
        {button && <i className={searchIconClassname}/> }
        <input
          type="text"
          className={inputClassName}
          placeholder={searchResult === "name" ?
            "Type restaurant, cafe... name" :
            (searchResult === "location" ?
              "Type address or nearby place of restaurant" :
              "Type restaurant, cafe... name")}
        />
      </div>
      <Select
        defaultValue={searchOptions[0]}
        options={searchOptions}
        styles={customStyles}
        theme={customThemes}
        onChange={setSearchResultHandler}
        className={selectClassName}
        placeholder="Search by"
        isSearchable
      />
      <Button
        buttonSize={searchSize}
        buttonStyle = 'btn--search'
      >
        Search
      </Button>
    </div>
  )
}

export default Searchbox;