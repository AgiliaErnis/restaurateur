import * as React from "react";
import * as ReactDOM from "react-dom";
import CardItem from "./CardItem";
import { shallow } from 'enzyme';



describe('testing', () => {
    it('should render', () => {
        shallow(<CardItem />)
    })

    it('should render label', () => {
        const label_example = "test";
        const wrapper = shallow(<CardItem path='/restaurants 'label={label_example} />);
        expect(wrapper.props().label).toEqual(label_example);
    })

     it('renders a button in size of "small" with text in it', () => {
    const wrapper = shallow(
      <Button {...minProps} size="small" text="Join us" />
    );

    expect(wrapper.find(StyledButton).prop('size')).toBe('small');
    expect(wrapper.find(StyledButton).prop('text')).toBe('Join us');
  });
})