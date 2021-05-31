import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../../UserContext';

const useLoginForm = (callback, validate) => {
  const { incorrectPassword, setIncorrectPassword,
    setSuccessfullLogin } = useContext(UserContext)

  const [values, setValues] = useState({
    email: '',
    password: ''
  });
  const [errors, setErrors] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleChange = e => {
    const { name, value } = e.target;
    setValues({
      ...values,
      [name]: value
    });

  };
  const handleSubmit = e => {
    e.preventDefault();

    setErrors(validate(values));
    setIsSubmitting(true)
  }

  useEffect(
    () => {
      if (Object.keys(errors).length === 0 && isSubmitting) {
        const loginRequest = {
          method: 'POST',
          body: JSON.stringify(values),
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          }
        }

        fetch(`${process.env.REACT_APP_PROXY}/login`, loginRequest)
        .then(response => response.json())
          .then(res => {
            if (res.status === 403) {
              setIncorrectPassword(true);
              setSuccessfullLogin(false);
              setIsSubmitting(false)
            } else {
              callback()
              setIncorrectPassword(false);
              setSuccessfullLogin(true)
              setIsSubmitting(true)
            };
          })
      }
    },
    [errors, callback, isSubmitting, values,
      setIncorrectPassword, setSuccessfullLogin]
  );

  return {
    handleChange, handleSubmit, values,
    errors, useEffect, incorrectPassword
  };
};

export default useLoginForm;
