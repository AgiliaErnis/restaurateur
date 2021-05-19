import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../UserContext';

const useChangePasswordForm = (callback, validate) => {
  const {incorrectOldPassword, setIncorrectOldPassword} = useContext(UserContext)
  const [values, setValues] = useState({
    oldPassword: '',
    newPassword: '',
    newPassword2: '',
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
    setIsSubmitting(!isSubmitting)
  };

  useEffect(
    () => {
      if (Object.keys(errors).length === 0 &&
        isSubmitting &&
        values.newPassword.length > 6) {
        const requestValues = {
          oldPassword: values.oldPassword,
          newPassword: values.newPassword
        }
        const changePasswordRequest = {
            method: 'PATCH',
          body: JSON.stringify(requestValues),
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          }
        }
        fetch('http://localhost:8080/user', changePasswordRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 403) {
              setIncorrectOldPassword(true);
            }
            else {
              callback();
              setIncorrectOldPassword(false)
            }
          })
        }
    },
    [errors, isSubmitting, callback, values,
      incorrectOldPassword, setIncorrectOldPassword]
  );

  return { handleChange, handleSubmit, values, errors, useEffect};
};

export default useChangePasswordForm;