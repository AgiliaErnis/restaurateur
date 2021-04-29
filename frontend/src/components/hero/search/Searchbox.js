import React, { useState, useEffect } from 'react';
import { Button } from '../../button/Button';
import './Searchbox.css';
import Select from 'react-select'
import MobileNavbar from '../../navbar/MobileNavbar'

const searchOptions = [
  {value: 'name', label:'Search by Name'},
  {value: 'location', label:'Search by Location'}
];

const customStyles = {
  control: (base, state) => ({
    ...base,
    height: 50,
    minHeight: 50,
    borderRadius: 0,
    color: 'rgb(185, 185, 185)',
    border: state.isFocused ? `1px solid rgb(185, 185, 185)` : `1px solid rgb(185, 185, 185)`,
    boxShadow: state.isFocused ? `1px solid rgb(185, 185, 185)` : `1px solid rgb(185, 185, 185)`,
    '&:hover': {
      border: state.isFocused ? `1px solid rgb(185, 185, 185)` : `1px solid rgb(185, 185, 185)`
    }
  })
};

function customThemes (theme){
  return {
    ...theme,
    colors: {
    ...theme.colors,
    primary:` rgb(185, 185, 185)`,
    primary25:` rgb(215, 215, 215)`
    },
  }
}

function Searchbox () {
  const { button, showSearch } = MobileNavbar();
  const [result, setResult] = useState(searchOptions.value);

  const setResultHandler = e => {
    setResult(e.value);
  }

  useEffect(() => {
    showSearch();
  }, [showSearch]);

  window.addEventListener('resize', showSearch);

  return (
    <div className="search-box">
      <div className="main-search-container">
        {button && <i className="fas fa-search"/>}
        <input
          type="text"
          className="main-input"
          placeholder={result === "name" ?
            "Type restaurant, cafe... name" :
            (result === "location" ?
              "Type address or nearby place of restaurant" :
              "Type restaurant, cafe... name")}
        />
      </div>
      <Select
        defaultValue={searchOptions[0]}
        options={searchOptions}
        styles={customStyles}
        theme={customThemes}
        onChange={setResultHandler}
        className="select"
        placeholder="Search by"
        isSearchable
      />
      <Button
        buttonSize='btn--large'
        buttonStyle='btn--search'
      >
        Search
      </Button>
    </div>
  )
}

export default Searchbox;