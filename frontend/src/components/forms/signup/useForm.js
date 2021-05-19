import { useState, useEffect } from 'react';

const useForm = (callback, validate) => {
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
        values.password.length > 6)
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

        fetch('http://localhost:8080/register', signupRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 200) {
              callback();
            }
          })
      }
    },
    [errors,isSubmitting,callback,values]
  );

  return { handleChange, handleSubmit, values, errors, useEffect};
};

export default useForm;