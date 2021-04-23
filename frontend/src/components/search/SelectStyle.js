import SearchboxLogic from "./AdjustSearchbox";

const SearchboxStyle = () => {
    const { selectHeight } = SearchboxLogic();

    const customStyles = {
        control: (base, state) => ({
            ...base,
            borderRadius: 0,
            height: selectHeight,
            cursor: 'pointer',
            color: 'rgb(185, 185, 185)',
            border: state.isFocused ? `1px solid rgb(185, 185, 185)` :
                                      `1px solid rgb(185, 185, 185)`,
            boxShadow: state.isFocused ? `1px solid rgb(185, 185, 185)` :
                                         `1px solid rgb(185, 185, 185)`,
            '&:hover': {
                border: state.isFocused ? `1px solid rgb(185, 185, 185)` :
                                          `1px solid rgb(185, 185, 185)`
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
    return {customStyles, customThemes}
}

export default SearchboxStyle;