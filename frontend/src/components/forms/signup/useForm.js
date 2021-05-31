import { useState, useEffect } from 'react';

const useForm = (callback, validate) => {
  const [emailAlreadyUsed, setEmailAlreadyUsed] = useState(false)
  const [values, setValues] = useState({
    username: '',
    email: '',
    password: '',
    password2: ''
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
    setIsSubmitting(!isSubmitting);
  };

  useEffect(
    () => {
      if (Object.keys(errors).length === 0 &&
        isSubmitting &&
        values.password.length >= 6)
      {
        const requestValues = {
            email: values.email,
            password: values.password,
            username: values.username
        }

        const signupRequest = {
          method: 'POST',
          body: JSON.stringify(requestValues),
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          }
        }

        fetch(`${process.env.REACT_APP_PROXY}/register`, signupRequest)
          .then(response => response.json())
          .then(res => {
            if (res.status === 200) {
              callback();
            } else if (res.status === 400) {
              setIsSubmitting(false);
              setEmailAlreadyUsed(!emailAlreadyUsed)
            }
          })
      }
    },
    [errors,isSubmitting,callback,values,emailAlreadyUsed,setEmailAlreadyUsed]
  );

  return {
    handleChange, handleSubmit, values, errors,
    useEffect, emailAlreadyUsed
  };
};

export default useForm;
