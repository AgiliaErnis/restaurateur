import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../UserContext';

const useDeleteAccountForm = (callback, validate) => {
  const { setIncorrectPasswordOnDelete,
    setSuccessfullLogin } = useContext(UserContext)
  const [values, setValues] = useState({
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
    setIsSubmitting(!isSubmitting)
  };

  useEffect(
    () => {

      if (Object.keys(errors).length === 0 && isSubmitting) {
          setIncorrectPasswordOnDelete(false)
            const deleteAccountRequest = {
                method: 'DELETE',
                body: JSON.stringify(values),
                credentials: 'include',
                headers: {
                  'Content-Type': 'application/json'
                }
            }
        fetch('http://localhost:8080/auth/user', deleteAccountRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 403) {
              setIncorrectPasswordOnDelete(true)
            }
            else if (res.Status === 200) {
              callback();
              setIncorrectPasswordOnDelete(false)
              setSuccessfullLogin(false)
            }
          })
            setIncorrectPasswordOnDelete(false)
          }
    },
    [errors, isSubmitting, callback, values,
      setIncorrectPasswordOnDelete, setSuccessfullLogin]
  );

  return { handleChange, handleSubmit, values, errors, useEffect};
};

export default useDeleteAccountForm;