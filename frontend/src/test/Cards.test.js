import * as React from "react";
import * as ReactDOM from "react-dom";
import Cards from "../components/cards/Cards";
import CardItem from "../components/cards/CardItem";
import * as enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16'
import { shallow, configure } from 'enzyme';

// tslint:disable-next-line:no-any
configure({ adapter: new Adapter() });

test('top card prints right string', () => {
    const root = document.createElement("div");
    ReactDOM.render(<Cards />, root);
    expect(root.querySelector("h1").textContent).toBe("Check Out the Top Suggestions! ");
})


describe('testing', () => {
    it('should render', () => {
        const root = document.createElement("div");
        ReactDOM.render(<CardItem />, root);
        ReactDOM.unmountComponentAtNode(root);
    })

    it('should render label', () => {
        const label_example = "test";
        const wrapper = shallow(<CardItem label={label_example} />);
        //expect(wrapper.props('label')).toEqual(label_example);
        expect(wrapper.props().label).toEqual(label_example);
    })
})


// test('breakfast card prints right info', () => {
//     const root = document.createElement("ul");
//     ReactDOM.render(<Cards />, root);
//     expect(root.querySelector("src").textContent).toBe('images/Home/breakfast.webp');
//     expect(root.querySelector("text").textContent).toBe('Top breakfast places in Prague');
//     expect(root.querySelector("label").textContent).toBe('Breakfast');
//     expect(root.querySelector("path").textContent).toBe('/restaurants');
// })