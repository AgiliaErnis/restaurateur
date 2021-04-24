import * as React from "react";
import * as ReactDOM from "react-dom";
import Districts from "../components/districts/Districts";
import * as enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16'
import renderer from 'react-test-renderer';
import { configure } from 'enzyme';

configure({ adapter: new Adapter() });

describe('testing CollegeSection', () => {

    it('should render', () => {
        const DistrictsComponent = renderer.create(<Districts />).toJSON();
        expect(DistrictsComponent).toMatchSnapshot();
    })

})