const AdjustSearchbox = () => {
  let inputClassName = setInputClassName();
  let selectHeight = setSelectHeight();
  let selectClassName = setSelectClassName();
  let searchIconClassname = setSearchIconClassName();
  let searchSize = setSearchSize();

  function setInputClassName () {
    switch(window.location.pathname){
      case '/restaurants':
        return 'main-input-updated';
      default:
        return "main-input";
    }
  }

  function setSelectClassName () {
    switch(window.location.pathname){
      case '/restaurants':
        return 'select-updated';
      default:
        return "select";
    }
  }

  function setSearchSize () {
    switch(window.location.pathname){
      case '/restaurants':
        return 'btn--small';
      default:
        return "btn--large";
    }

  }

  function setSelectHeight () {
    switch(window.location.pathname){
      case '/restaurants':
        return '40px';
      default:
        return "50px";
    }
  }

  function setSearchIconClassName () {
    switch(window.location.pathname){
      case '/restaurants':
        return "hidden";
      default:
        return "fas fa-search";
    }
  }

  return {
    inputClassName, selectClassName, searchSize,
    selectHeight, searchIconClassname
  }
}

export default AdjustSearchbox;