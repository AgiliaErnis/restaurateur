import React, { useContext } from 'react';
import validate from './validateInfo';
import useForm from './useForm';
import { Button } from '../../button/Button'
import '../Form.css';
import PasswordStrengthMeter from './PasswordStrengthMeter';
import { UserContext } from "../../../UserContext"

const SignUpForm = ({ submitForm }) => {
  const { goodPassword } = useContext(UserContext)
  const { handleChange, handleSubmit, values, errors,emailAlreadyUsed } = useForm(
    submitForm,
    validate
  );

  return (
    <form
      onSubmit={handleSubmit}
      className='form' noValidate>
        <div className='form-inputs'>
          <label className='form-label'>Username</label>
          <input
          className={values.username.length < 6 ?
            'form-input'
            :
            'form-input green'}
            type='text'
            name='username'
            placeholder='Enter your username'
            value={values.username}
            onChange={handleChange}
          />
          {errors.username && <p>{errors.username}</p>}
        </div>
        <div className='form-inputs'>
          <label className='form-label'>Email</label>
          <input
          className={values.email.indexOf("@") === -1 ?
            'form-input'
            :
            'form-input green'}
            type='email'
            name='email'
            placeholder='Enter your email'
            value={values.email}
            onChange={handleChange}
          />
        {errors.email && <p>{errors.email}</p>}
        {emailAlreadyUsed && <p>Email is already used</p>}
      </div>
        <div className="form-inputs">
          <label className='form-label'>Password</label>
          <input
          className={goodPassword ?
            values.password.length > 64 ?
              "form-input"
              :
              "form-input green"
            :
            "form-input"}
            type='password'
            name='password'
            placeholder='Enter your password'
            value={values.password}
            onChange={handleChange}
        />
        {values.password !== "" &&
          <PasswordStrengthMeter password={values.password} />}
          {errors.password && <p>{errors.password}</p>}
        </div>
        <div className='form-inputs'>
          <label className='form-label'>Confirm Password</label>
          <input
          className={values.password !== values.password2 ?
            'form-input'
            :
            'form-input green'}
            type='password'
            name='password2'
            placeholder='Confirm your password'
            value={values.password2}
            onChange={handleChange}
          />
          {errors.password2 && <p>{errors.password2}</p>}
        </div>
        <Button
          buttonStyle="btn--form"
          buttonSize="btn--large"
          type='submit'
          onClick={handleSubmit}
        >
          Sign up
        </Button>
      </form>
  );
};

export default SignUpForm;