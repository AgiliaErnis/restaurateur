export default function validatePassword(values) {
    let errors = {};

    if (!values.oldPassword) {
        errors.oldPassword = 'Password is required'
    }

    if (!values.newPassword) {
      errors.newPassword = 'Password is required';
    } else if (values.newPassword.length < 6) {
      errors.newPassword = 'Password needs to be 6 characters or more';
    }

    if (!values.newPassword2) {
      errors.newPassword2 = 'Password is required';
    } else if (values.newPassword2 !== values.newPassword) {
      errors.newPassword2 = 'Passwords do not match';
    }

    return errors;
  }