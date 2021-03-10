import React from 'react';
import '../Form.css';

const FormSuccess = () => {
  return (
    <div className='form-content '>
    <form  className='form success'>
      <h2 className='form-success'>Thank You For Signing Up!</h2>
      <img className='form-img' src='images/menu.png' alt='success-image' />
    </form>
    </div>
  );
};

export default FormSuccess;