import React from 'react';
import './Form.css'
import styled from 'styled-components'
import { MdClose } from 'react-icons/md'
import AnimatedForms from './AnimatedForms'

  const CloseModalButton = styled(MdClose)`
  cursor: pointer;
  position: absolute;
  top: 10px;
  right: 18px;
  width: 32px;
  height: 32px;
  padding: 0;
  z-index: 10;
  color: #fff;
`;


export const Modal = ({ showSignUpModal, setShowSignUpModal, showLogInModal, setShowLogInModal }) => {

  return (
    <>
      {showSignUpModal ?
        <>
          <CloseModalButton
            aria-label='Close modal'
            onClick={() => setShowSignUpModal(prev => !prev)}
          />
          <AnimatedForms signup={true} login={false} />
        </>
        :
        <>
          {showLogInModal ?
            <>
              <CloseModalButton
                aria-label='Close modal'
                onClick={() => setShowLogInModal(prev => !prev)}
              />
              <AnimatedForms signup={false} login={true} />
            </> : null}
        </>
      }
    </>
  );
};