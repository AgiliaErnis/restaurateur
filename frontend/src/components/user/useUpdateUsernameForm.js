import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../UserContext';

const useChangePasswordForm = (callback, validate) => {
    const { setNewUsername }  = useContext(UserContext)
  const [values, setValues] = useState({
    newUsername: ''
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
      if (Object.keys(errors).length === 0 && isSubmitting ) {
            const updateUsernameRequest = {
                method: 'PATCH',
                body: JSON.stringify(values),
                credentials: 'include',
                headers: {
                  'Content-Type': 'application/json'
                }
            }

    fetch(`${process.env.REACT_APP_PROXY}/auth/user`, updateUsernameRequest)
      .then(response => response.json())
      .then(res => {
        if (res.status === 200) {
          callback();
            setNewUsername(true)
        }
      })

      }
    },
    [errors,isSubmitting,callback,values,setNewUsername]
  );

  return { handleChange, handleSubmit, values, errors, useEffect};
};

export default useChangePasswordForm;
