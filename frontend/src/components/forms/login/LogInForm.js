import React from 'react';
import { Button } from '../../button/Button';
import '../Form.css';
import useLoginForm from './useLoginForm'
import validateLoginInfo from './validateLoginInfo'


const LogInForm = ({ submitForm }) => {
  const { handleChange, handleSubmit, values, errors } = useLoginForm(
    submitForm,
    validateLoginInfo
  );
  return (
    <>
      <form className='login-form' onSubmit={handleSubmit} noValidate>
        <div className='form-inputs'>
          <label className='form-label'>Username</label>
          <input
            className='form-input blue'
            type='text'
            name='username'
            placeholder='Enter your username'
            value={values.username}
            onChange={handleChange}
          />
          {errors.username && <p>{errors.username}</p>}
        </div>
        <div className='form-inputs'>
          <label className='form-label'>Password</label>
          <input
            className='form-input'
            type='password'
            name='password'
            placeholder='Enter your password'
            value={values.password}
            onChange={handleChange}
          />
          {errors.password && <p>{errors.password}</p>}
        </div>
        <div class='form-inputs'>
          <input className='checkbox-input' type='checkbox'/>
          <span>Remember me</span>
        </div>
        <Button
          className='form-input-btn'
          buttonStyle="btn--form"
          buttonSize="btn--large"
          type='submit'
          onClick={handleSubmit}
        >
          Log in
        </Button>
      </form>
    </>
  );
};

export default LogInForm;