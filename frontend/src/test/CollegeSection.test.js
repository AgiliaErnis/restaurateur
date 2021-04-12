import * as React from "react";
import * as ReactDOM from "react-dom";
import CollegeSection from "../components/college/CollegeSection";
import * as enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16'
import { shallow, configure } from 'enzyme';

configure({ adapter: new Adapter() });

describe('testing CollegeSection', () => {

    it('should render', () => {
        shallow(<CollegeSection />)
    })

    // it('should render image true', () => {
    //     const showImg = whatTheFuck
    //     expect().ToBe("true");
    // })
})