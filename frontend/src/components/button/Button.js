import React from 'react';
import './Button.css';
import { BrowserRouter as Router } from 'react-router-dom';

const STYLES = ['btn--primary', 'btn--outline', 'btn--search', 'btn--form'];

const SIZES = ['btn--medium', 'btn--large'];

export const Button = ({
  children,
  type,
  onClick,
  buttonStyle,
  buttonSize
}) => {

  const checkButtonStyle = STYLES.includes(buttonStyle)
    ? buttonStyle
    : STYLES[0];

  const checkButtonSize = SIZES.includes(buttonSize) ? buttonSize : SIZES[0];

  return (
    <Router>
        <button
          className={`btn ${checkButtonStyle} ${checkButtonSize}`}
          onClick={onClick}
          type={type}
        >
          {children}
        </button>
    </Router>
  );
};