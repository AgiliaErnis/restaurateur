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

    it('should render img src', () => {
        const src_example = 'images/Home/breakfast.webp';
        const cardWrapper = enzyme.mount(<CardItem src={src_example} path='/' />);
        expect(cardWrapper.props().src).toEqual(src_example);
    })

    it('should render text', () => {
        const text_example = 'Top breakfast places in Prague';
        const cardWrapper = enzyme.mount(<CardItem text={text_example} path='/' />);
        expect(cardWrapper.props().text).toEqual(text_example);
    })

    it('should render path', () => {
        const path_example = '/restaurants';
        const cardWrapper = enzyme.mount(<CardItem path={path_example} />);
        expect(cardWrapper.props().path).toEqual(path_example);
    })
})