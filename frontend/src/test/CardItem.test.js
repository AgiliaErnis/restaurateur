import * as React from "react";
import * as ReactDOM from "react-dom";
import Cards from "../components/cards/Cards";
import CardItem from "../components/cards/CardItem";
import * as enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16'
import { shallow, configure } from 'enzyme';
import { UserContext } from "../UserContext";

configure({ adapter: new Adapter() });

describe('testing CardItem', () => {

    it('should render', () => {
        shallow(<UserContext.Provider value={true}><CardItem /></UserContext.Provider>);
    })

})