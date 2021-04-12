import React from 'react';
import '../Form.css';
import '../AnimatedForms.css'

const FormSuccess = () => {
  return (
    <div className='form-content'>
      <form  className='login-signup-wrapper'>
        <img className='form-img' src='images/menu.png' alt='success' />
        <h2 className='form-success'>Thank You For Signing Up!</h2>
      </form>
    </div>
  );
};

export default FormSuccess;