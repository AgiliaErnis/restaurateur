import * as React from "react";
import * as ReactDOM from "react-dom";
import Cards from "../components/cards/Cards";
import CardItem from "../components/cards/CardItem";
import * as enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16'
import { shallow, configure } from 'enzyme';

configure({ adapter: new Adapter() });

describe('testing CardItem', () => {
    
    it('should render', () => {
        shallow(<CardItem />);
    })
    

    // it('should render label', () => {
    //     const label_example = "test";
    //     const cardWrapper = enzyme.mount(<CardItem label={label_example} />);
    //     expect(cardWrapper.props().label).toEqual(label_example);
    // })

})