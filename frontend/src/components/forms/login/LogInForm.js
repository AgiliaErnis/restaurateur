import React from 'react';
import { Button } from '../../button/Button';
import '../Form.css';

const LogInForm = () => {
  return (
    <>
      <form className='login-form'>
        <div className='form-inputs'>
          <label className='form-label'>Username</label>
          <input
            className='form-input'
            type='text'
            name='username'
            placeholder='Enter your username'
          />
        </div>
        <div className='form-inputs'>
          <label className='form-label'>Password</label>
          <input
            className='form-input'
            type='password'
            name='password'
            placeholder='Enter your password'
          />
        </div>
        <div class='form-inputs'>
          <input className='checkbox-input' type='checkbox'/>
          <span>Remember me</span>
        </div>
        <Button className='form-input-btn' buttonStyle="btn--form" buttonSize="btn--large" type='submit'>
          Log in
        </Button>
      </form>
    </>
  );
};

export default LogInForm;