import React from 'react';
import { Button } from '../../button/Button';
import '../Form.css';
import useLoginForm from './useLoginForm'
import validateLoginInfo from './validateLoginInfo'


const LogInForm = ({ submitForm }) => {
  const { handleChange, handleSubmit, values,
    errors, incorrectPassword } = useLoginForm(
    submitForm,
    validateLoginInfo
    );

  return (
    <>
      <form className='login-form' onSubmit={handleSubmit} noValidate>
       <div className='form-inputs'>
          <label className='form-label'>Email</label>
          <input
          className='form-input blue'
            type='email'
            name='email'
            placeholder='Enter your email'
            value={values.email}
            onChange={handleChange}
          />
          {errors.email && <p>{errors.email}</p>}
      </div>
        <div className='form-inputs'>
          <label className='form-label'>Password</label>
          <input
            className='form-input blue'
            type='password'
            name='password'
            placeholder='Enter your password'
            value={values.password}
            onChange={handleChange}
          />
          {incorrectPassword === true && !errors.password ?
            <p>Incorrect Password</p>
            :
            errors.password && <p>{errors.password}</p>}
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