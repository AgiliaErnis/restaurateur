import React from 'react';
import validate from './signup/validateInfo';
import useForm from './useForm';
import './Form.css';

const LogInForm = ({ submitForm }) => {
  const { handleChange, handleSubmit, values, errors } = useForm(
    submitForm,
    validate
  );

  return (
    <div className='form-content'>
      <form onSubmit={handleSubmit} className='login-form' noValidate>
      <i class="fas fa-utensils" />

        <div className='form-inputs'>
          <label className='form-label'>Username</label>
          <input
            className='form-input'
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
          <input className='checkbox-input' type='checkbox' />
          <span>Remember me</span> </div>
        <button className='form-input-btn' type='submit'>
          Log in
        </button>
      </form>
    </div>
  );
};

export default LogInForm;