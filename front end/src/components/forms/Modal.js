import React, { useState } from 'react';
import SignUp from './signup/SignUpForm';
import FormSuccess from './signup/FormSuccess';
import './Form.css'
import styled from 'styled-components'
import { MdClose } from 'react-icons/md'
import LogInForm from './LogInForm'


export const Modal = ({ showSignUpModal, setShowSignUpModal, showLogInModal, setShowLogInModal }) => {
  const [isSubmitted, setIsSubmitted] = useState(false);

  const CloseModalButton = styled(MdClose)`
  cursor: pointer;
  position: absolute;
  top: 15px;
  right: 20px;
  width: 32px;
  height: 32px;
  padding: 0;
  z-index: 10;
  color: #fff;
`;

  function submitForm() {
    setIsSubmitted(true);
  }
  return (
    <>
     {showSignUpModal ?
     <>
    <CloseModalButton
                aria-label='Close modal'
                onClick={() => setShowSignUpModal(prev => !prev)}
              />
        {!isSubmitted ? (
          <SignUp submitForm={submitForm} />
        ) : (
          <FormSuccess />
        )}</> :
        <>

       {showLogInModal ? <><CloseModalButton
                aria-label='Close modal'
                onClick={() => setShowLogInModal(prev => !prev)}
              /><LogInForm/></> : null }
        </>
       }
    </>
  );
};
